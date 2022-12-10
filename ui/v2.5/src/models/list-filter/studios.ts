import {
  createBooleanCriterionOption,
  createMandatoryNumberCriterionOption,
  createMandatoryStringCriterionOption,
  createStringCriterionOption,
  NullNumberCriterionOption,
  createMandatoryTimestampCriterionOption,
} from "./criteria/criterion";
import { StudioIsMissingCriterionOption } from "./criteria/is-missing";
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
  ]);

const displayModeOptions = [DisplayMode.Grid];
const criterionOptions = [
  createMandatoryNumberCriterionOption("scene_count"),
  createMandatoryNumberCriterionOption("gallery_count"),
];

export const StudioListFilterOptions = new ListFilterOptions(
  defaultSortBy,
  sortByOptions,
  displayModeOptions,
  criterionOptions
);
