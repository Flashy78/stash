import {
  createMandatoryNumberCriterionOption,
  createMandatoryStringCriterionOption,
  createStringCriterionOption,
  NullNumberCriterionOption,
  createMandatoryTimestampCriterionOption,
  createDateCriterionOption,
} from "./criteria/criterion";
import { PerformerFavoriteCriterionOption } from "./criteria/favorite";
import { ImageIsMissingCriterionOption } from "./criteria/is-missing";
import { OrganizedCriterionOption } from "./criteria/organized";
import { PerformersCriterionOption } from "./criteria/performers";
import { ResolutionCriterionOption } from "./criteria/resolution";
import { StudiosCriterionOption } from "./criteria/studios";
import {
  PerformerTagsCriterionOption,
  TagsCriterionOption,
} from "./criteria/tags";
import { ListFilterOptions, MediaSortByOptions } from "./filter-options";
import { DisplayMode } from "./types";

const defaultSortBy = "title";

const sortByOptions = [
  "o_counter",
  "date",
  ...MediaSortByOptions,
].map(ListFilterOptions.createSortBy);

const displayModeOptions = [DisplayMode.Wall];
const criterionOptions = [
  createStringCriterionOption("title"),
  createMandatoryNumberCriterionOption("o_counter"),
  ResolutionCriterionOption,
  TagsCriterionOption,
  PerformersCriterionOption,
  createMandatoryNumberCriterionOption("performer_count"),
  createMandatoryNumberCriterionOption("performer_age"),
  StudiosCriterionOption,
  createDateCriterionOption("date"),
];
export const ImageListFilterOptions = new ListFilterOptions(
  defaultSortBy,
  sortByOptions,
  displayModeOptions,
  criterionOptions
);
