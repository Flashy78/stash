package studio

import (
	"context"
	"fmt"

	"github.com/stashapp/stash/pkg/logger"
	"github.com/stashapp/stash/pkg/models"
	"github.com/stashapp/stash/pkg/models/json"
	"github.com/stashapp/stash/pkg/models/jsonschema"
	"github.com/stashapp/stash/pkg/utils"
)

type FinderImageAliasStashIDGetter interface {
	Finder
	GetImage(ctx context.Context, studioID int) ([]byte, error)
	models.AliasLoader
	models.StashIDLoader
}

// ToJSON converts a Studio object into its JSON equivalent.
func ToJSON(ctx context.Context, reader FinderImageAliasStashIDGetter, studio *models.Studio) (*jsonschema.Studio, error) {
	newStudioJSON := jsonschema.Studio{
		Name:          studio.Name,
		URL:           studio.URL,
		IgnoreAutoTag: studio.IgnoreAutoTag,
		CreatedAt:     json.JSONTime{Time: studio.CreatedAt},
		UpdatedAt:     json.JSONTime{Time: studio.UpdatedAt},
		Details:       studio.Details,
	}

	if studio.Rating != nil {
		newStudioJSON.Rating = *studio.Rating
	}

	if studio.ParentID != nil {
		parent, err := reader.Find(ctx, *studio.ParentID)
		if err != nil {
			return nil, fmt.Errorf("error getting parent studio: %v", err)
		}

		if parent != nil {
			newStudioJSON.ParentStudio = parent.Name
		}
	}

	if err := studio.LoadAliases(ctx, reader); err != nil {
		return nil, fmt.Errorf("loading studio aliases: %w", err)
	}
	newStudioJSON.Aliases = studio.Aliases.List()

	if err := studio.LoadStashIDs(ctx, reader); err != nil {
		return nil, fmt.Errorf("loading studio stash ids: %w", err)
	}
	newStudioJSON.StashIDs = studio.StashIDs.List()

	image, err := reader.GetImage(ctx, studio.ID)
	if err != nil {
		logger.Errorf("Error getting studio image: %v", err)
	}

	if len(image) > 0 {
		newStudioJSON.Image = utils.GetBase64StringFromData(image)
	}

	return &newStudioJSON, nil
}
