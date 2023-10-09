import {
  createStringCriterionOption,
  createDateCriterionOption,
  createMandatoryTimestampCriterionOption,
  createDurationCriterionOption,
} from "./criteria/criterion";
import { MovieIsMissingCriterionOption } from "./criteria/is-missing";
import { StudiosCriterionOption } from "./criteria/studios";
import { PerformersCriterionOption } from "./criteria/performers";
import { ListFilterOptions } from "./filter-options";
import { DisplayMode } from "./types";
import { RatingCriterionOption } from "./criteria/rating";

const defaultSortBy = "name";

const sortByOptions = ["name", "date", "rating"]
  .map(ListFilterOptions.createSortBy);
const displayModeOptions = [DisplayMode.Grid];
const criterionOptions = [
  StudiosCriterionOption,
  createStringCriterionOption("director"),
  createStringCriterionOption("synopsis"),
  PerformersCriterionOption,
  createDateCriterionOption("date"),
];

export const MovieListFilterOptions = new ListFilterOptions(
  defaultSortBy,
  sortByOptions,
  displayModeOptions,
  criterionOptions
);
