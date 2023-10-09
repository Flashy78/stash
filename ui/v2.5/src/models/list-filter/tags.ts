import {
  createBooleanCriterionOption,
  createMandatoryNumberCriterionOption,
  createMandatoryStringCriterionOption,
  createStringCriterionOption,
  MandatoryNumberCriterionOption,
  createMandatoryTimestampCriterionOption,
} from "./criteria/criterion";
import { TagIsMissingCriterionOption } from "./criteria/is-missing";
import { ListFilterOptions } from "./filter-options";
import { DisplayMode } from "./types";
import {
  ChildTagsCriterionOption,
  ParentTagsCriterionOption,
} from "./criteria/tags";

const defaultSortBy = "name";
const sortByOptions = ["name", "random"]
  .map(ListFilterOptions.createSortBy)
  .concat([
    {
      messageID: "gallery_count",
      value: "galleries_count",
    },
    {
      messageID: "performer_count",
      value: "performers_count",
    },
    {
      messageID: "scene_count",
      value: "scenes_count",
    },
  ]);

const displayModeOptions = [DisplayMode.Grid];
const criterionOptions = [
  createStringCriterionOption("description"),
  createMandatoryNumberCriterionOption("scene_count"),
  createMandatoryNumberCriterionOption("gallery_count"),
  createMandatoryNumberCriterionOption("performer_count"),
  ParentTagsCriterionOption,
  new MandatoryNumberCriterionOption("parent_tag_count", "parent_count"),
  ChildTagsCriterionOption,
  new MandatoryNumberCriterionOption("sub_tag_count", "child_count"),
];

export const TagListFilterOptions = new ListFilterOptions(
  defaultSortBy,
  sortByOptions,
  displayModeOptions,
  criterionOptions
);
