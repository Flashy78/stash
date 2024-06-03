import {
  createBooleanCriterionOption,
  createMandatoryNumberCriterionOption,
  createMandatoryStringCriterionOption,
  createStringCriterionOption,
  createMandatoryTimestampCriterionOption,
} from "./criteria/criterion";
import { FavoriteStudioCriterionOption } from "./criteria/favorite";
import { StudioIsMissingCriterionOption } from "./criteria/is-missing";
import { RatingCriterionOption } from "./criteria/rating";
import { StashIDCriterionOption } from "./criteria/stash-ids";
import { ParentStudiosCriterionOption } from "./criteria/studios";
import { ListFilterOptions } from "./filter-options";
import { DisplayMode } from "./types";

const defaultSortBy = "name";
const sortByOptions = ["name", "random"]
  .map(ListFilterOptions.createSortBy)
  .concat([
    {
      messageID: "gallery_count",
      value: "galleries_count",
    },
    {
      messageID: "scene_count",
      value: "scenes_count",
    },
    {
      messageID: "subsidiary_studio_count",
      value: "child_count",
    },
  ]);

const displayModeOptions = [DisplayMode.Grid];
const criterionOptions = [
  FavoriteStudioCriterionOption,
  createMandatoryStringCriterionOption("name"),
  createStringCriterionOption("details"),
  ParentStudiosCriterionOption,
  createMandatoryNumberCriterionOption("scene_count"),
  createMandatoryNumberCriterionOption("gallery_count"),
  createStringCriterionOption("url"),
  createStringCriterionOption("aliases"),
  createMandatoryNumberCriterionOption(
    "child_count",
    "subsidiary_studio_count"
  ),
];

export const StudioListFilterOptions = new ListFilterOptions(
  defaultSortBy,
  sortByOptions,
  displayModeOptions,
  criterionOptions
);
