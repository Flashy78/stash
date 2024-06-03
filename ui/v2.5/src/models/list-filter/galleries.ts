import {
  createMandatoryNumberCriterionOption,
  createStringCriterionOption,
  createDateCriterionOption,
  createMandatoryTimestampCriterionOption,
} from "./criteria/criterion";
import { PerformerFavoriteCriterionOption } from "./criteria/favorite";
import { GalleryIsMissingCriterionOption } from "./criteria/is-missing";
import { OrganizedCriterionOption } from "./criteria/organized";
import { HasChaptersCriterionOption } from "./criteria/has-chapters";
import { PerformersCriterionOption } from "./criteria/performers";
import { AverageResolutionCriterionOption } from "./criteria/resolution";
import { ScenesCriterionOption } from "./criteria/scenes";
import { StudiosCriterionOption } from "./criteria/studios";
import {
  PerformerTagsCriterionOption,
  TagsCriterionOption,
} from "./criteria/tags";
import { ListFilterOptions, MediaSortByOptions } from "./filter-options";
import { DisplayMode } from "./types";
import { RatingCriterionOption } from "./criteria/rating";
import { PathCriterionOption } from "./criteria/path";

const defaultSortBy = "date";

const sortByOptions = ["date", ...MediaSortByOptions]
  .map(ListFilterOptions.createSortBy);

const displayModeOptions = [
  DisplayMode.Grid,
];

const criterionOptions = [
  createStringCriterionOption("title"),
  createStringCriterionOption("code", "scene_code"),
  createStringCriterionOption("details"),
  createStringCriterionOption("photographer"),
  RatingCriterionOption,
  AverageResolutionCriterionOption,
  GalleryIsMissingCriterionOption,
  TagsCriterionOption,
  HasChaptersCriterionOption,
  createMandatoryNumberCriterionOption("tag_count"),
  PerformerTagsCriterionOption,
  PerformersCriterionOption,
  createMandatoryNumberCriterionOption("performer_count"),
  createMandatoryNumberCriterionOption("performer_age"),
  PerformerFavoriteCriterionOption,
  createMandatoryNumberCriterionOption("image_count"),
  ScenesCriterionOption,
  StudiosCriterionOption,
];

export const GalleryListFilterOptions = new ListFilterOptions(
  defaultSortBy,
  sortByOptions,
  displayModeOptions,
  criterionOptions
);
