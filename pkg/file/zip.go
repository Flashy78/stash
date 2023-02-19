package file

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"

	"github.com/stashapp/stash/pkg/logger"
	"github.com/xWTF/chardet"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

var (
	errNotReaderAt  = errors.New("not a ReaderAt")
	errZipFSOpenZip = errors.New("cannot open zip file inside zip file")
)

// ZipFS is a file system backed by a zip file.
type ZipFS struct {
	*zip.Reader
	zipFileCloser io.Closer
	zipInfo       fs.FileInfo
	zipPath       string
}

func newZipFS(fs FS, path string, info fs.FileInfo) (*ZipFS, error) {
	reader, err := fs.Open(path)
	if err != nil {
		return nil, err
	}

	asReaderAt, _ := reader.(io.ReaderAt)
	if asReaderAt == nil {
		reader.Close()
		return nil, errNotReaderAt
	}

	zipReader, err := zip.NewReader(asReaderAt, info.Size())
	if err != nil {
		reader.Close()
		return nil, err
	}

	// Concat all Name and Comment for better detection result
	var buffer bytes.Buffer
	for _, f := range zipReader.File {
		buffer.WriteString(f.Name)
		buffer.WriteString(f.Comment)
	}
	buffer.WriteString(zipReader.Comment)

	// Detect encoding
	d, err := chardet.NewTextDetector().DetectBest(buffer.Bytes())
	if err != nil {
		reader.Close()
		return nil, fmt.Errorf("unable to detect decoding: %w", err)
	}

	// If the charset is not UTF8, decode'em
	if d.Charset != "UTF-8" {
		logger.Debugf("Detected non-utf8 zip charset %s (%s): %s", d.Charset, d.Language, path)

		e, _ := charset.Lookup(d.Charset)
		if e == nil {
			reader.Close()
			return nil, fmt.Errorf("failed to lookup charset %s, language %s", d.Charset, d.Language)
		}

		decoder := e.NewDecoder()
		for _, f := range zipReader.File {
			f.Name, _, err = transform.String(decoder, f.Name)
			if err != nil {
				reader.Close()
				return nil, fmt.Errorf("failed to decode %v: %w", []byte(f.Name), err)
			}
			// Comments are not decoded cuz stash doesn't use that
		}
	}

	return &ZipFS{
		Reader:        zipReader,
		zipFileCloser: reader,
		zipInfo:       info,
		zipPath:       path,
	}, nil
}

func (f *ZipFS) rel(name string) (string, error) {
	if f.zipPath == name {
		return ".", nil
	}

	relName, err := filepath.Rel(f.zipPath, name)
	if err != nil {
		return "", fmt.Errorf("internal error getting relative path: %w", err)
	}

	// convert relName to use slash, since zip files do so regardless
	// of os
	relName = filepath.ToSlash(relName)

	return relName, nil
}

func (f *ZipFS) Stat(name string) (fs.FileInfo, error) {
	reader, err := f.Open(name)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return reader.Stat()
}

func (f *ZipFS) Lstat(name string) (fs.FileInfo, error) {
	return f.Stat(name)
}

func (f *ZipFS) OpenZip(name string) (*ZipFS, error) {
	return nil, errZipFSOpenZip
}

func (f *ZipFS) IsPathCaseSensitive(path string) (bool, error) {
	return true, nil
}

type zipReadDirFile struct {
	fs.File
}

func (f *zipReadDirFile) ReadDir(n int) ([]fs.DirEntry, error) {
	asReadDirFile, _ := f.File.(fs.ReadDirFile)
	if asReadDirFile == nil {
		return nil, fmt.Errorf("internal error: not a ReadDirFile")
	}

	return asReadDirFile.ReadDir(n)
}

func (f *ZipFS) Open(name string) (fs.ReadDirFile, error) {
	relName, err := f.rel(name)
	if err != nil {
		return nil, err
	}

	r, err := f.Reader.Open(relName)
	if err != nil {
		return nil, err
	}

	return &zipReadDirFile{
		File: r,
	}, nil
}

func (f *ZipFS) Close() error {
	return f.zipFileCloser.Close()
}

// openOnly returns a ReadCloser where calling Close will close the zip fs as well.
func (f *ZipFS) OpenOnly(name string) (io.ReadCloser, error) {
	r, err := f.Open(name)
	if err != nil {
		return nil, err
	}

	return &wrappedReadCloser{
		ReadCloser: r,
		outer:      f,
	}, nil
}

type wrappedReadCloser struct {
	io.ReadCloser
	outer io.Closer
}

func (f *wrappedReadCloser) Close() error {
	_ = f.ReadCloser.Close()
	return f.outer.Close()
}
