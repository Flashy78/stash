import React from "react";
import { Button, ButtonGroup } from "react-bootstrap";
import { FormattedMessage } from "react-intl";
import cx from "classnames";

import { Icon } from "src/components/Shared/Icon";
import { OperationButton } from "src/components/Shared/OperationButton";
import { StudioSelect, SelectObject } from "src/components/Shared/Select";
import * as GQL from "src/core/generated-graphql";

import { OptionalField } from "../IncludeButton";
import { faSave } from "@fortawesome/free-solid-svg-icons";

interface IStudioResultProps {
  studio: GQL.ScrapedStudio;
  selectedID: string | undefined;
  setSelectedID: (id: string | undefined) => void;
  onCreate: () => void;
  onLink?: () => Promise<void>;
  endpoint?: string;
  stashBoxBaseURL?: string;
}

const StudioResult: React.FC<IStudioResultProps> = ({
  studio,
  selectedID,
  setSelectedID,
  onCreate,
  onLink,
  endpoint,
  stashBoxBaseURL,
}) => {
  const { data: studioData, loading: stashLoading } = GQL.useFindStudioQuery({
    variables: { id: studio.stored_id ?? "" },
    skip: !studio.stored_id,
  });

  const matchedStudio = studioData?.findStudio;
  const matchedStashID = matchedStudio?.stash_ids.some(
    (stashID) => stashID.endpoint === endpoint && stashID.stash_id
  );

  const handleSelect = (studios: SelectObject[]) => {
    if (studios.length) {
      setSelectedID(studios[0].id);
    } else {
      setSelectedID(undefined);
    }
  };

  const handleSkip = () => {
    setSelectedID(undefined);
  };

  if (stashLoading) return <div>Loading studio</div>;

  function renderStudioName(
    name: string,
    url: string | undefined,
    id: string | undefined | null
  ) {
    return url && id ? (
      <a href={url + "studios/" + id} target="_blank" rel="noreferrer">
        {name}
      </a>
    ) : (
      name
    );
  }

  if (matchedStudio && matchedStashID) {
    return (
      <div className="row no-gutters my-2">
        <div className="entity-name">
          <FormattedMessage id="countables.studios" values={{ count: 1 }} />:
          <b className="ml-2">
            {renderStudioName(
              studio.name,
              stashBoxBaseURL,
              studio.remote_site_id
            )}
          </b>
        </div>
        <span className="ml-auto">
          <OptionalField
            exclude={selectedID === undefined}
            setExclude={(v) =>
              v ? handleSkip() : setSelectedID(matchedStudio.id)
            }
          >
            <div>
              <span className="mr-2">
                <FormattedMessage id="component_tagger.verb_matched" />:
              </span>
              <b className="col-3 text-right">
                {renderStudioName(matchedStudio.name, "/", matchedStudio.id)}
              </b>
            </div>
          </OptionalField>
        </span>
      </div>
    );
  }

  function maybeRenderLinkButton() {
    if (endpoint && onLink) {
      return (
        <OperationButton
          variant="secondary"
          disabled={selectedID === undefined}
          operation={onLink}
          hideChildrenWhenLoading
        >
          <Icon icon={faSave} />
        </OperationButton>
      );
    }
  }

  const selectedSource = !selectedID ? "skip" : "existing";

  return (
    <div className="row no-gutters align-items-center mt-2">
      <div className="entity-name">
        <FormattedMessage id="countables.studios" values={{ count: 1 }} />:
        <b className="ml-2">
          {renderStudioName(
            studio.name,
            stashBoxBaseURL,
            studio.remote_site_id
          )}
        </b>
      </div>
      <ButtonGroup>
        <Button variant="secondary" onClick={() => onCreate()}>
          <FormattedMessage id="actions.create" />
        </Button>
        <Button
          variant={selectedSource === "skip" ? "primary" : "secondary"}
          onClick={() => handleSkip()}
        >
          <FormattedMessage id="actions.skip" />
        </Button>
        <StudioSelect
          ids={selectedID ? [selectedID] : []}
          onSelect={handleSelect}
          className={cx("studio-select", {
            "studio-select-active": selectedSource === "existing",
          })}
          isClearable={false}
        />
        {maybeRenderLinkButton()}
      </ButtonGroup>
    </div>
  );
};

export default StudioResult;
