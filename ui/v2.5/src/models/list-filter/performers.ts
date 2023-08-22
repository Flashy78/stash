import {
  createNumberCriterionOption,
  createMandatoryNumberCriterionOption,
  createStringCriterionOption,
  createBooleanCriterionOption,
  createDateCriterionOption,
  createMandatoryTimestampCriterionOption,
  NumberCriterionOption,
  NullNumberCriterionOption,
} from "./criteria/criterion";
import { FavoriteCriterionOption } from "./criteria/favorite";
import { GenderCriterionOption } from "./criteria/gender";
import { CircumcisedCriterionOption } from "./criteria/circumcised";
import { PerformerIsMissingCriterionOption } from "./criteria/is-missing";
import { StashIDCriterionOption } from "./criteria/stash-ids";
import { StudiosCriterionOption } from "./criteria/studios";
import { TagsCriterionOption } from "./criteria/tags";
import { ListFilterOptions } from "./filter-options";
import { CriterionType, DisplayMode } from "./types";

const defaultSortBy = "name";
const sortByOptions = [
  "name",
  "random",
  "rating",
  "penis_length",
]
  .map(ListFilterOptions.createSortBy)
  .concat([
    {
      messageID: "scene_count",
      value: "scenes_count",
    },
    {
      messageID: "gallery_count",
      value: "galleries_count",
    },
    {
      messageID: "o_counter",
      value: "o_counter",
    },
  ]);

const displayModeOptions = [
  DisplayMode.Grid,
];

const numberCriteria: CriterionType[] = [
  "death_year",
  "age",
  "weight",
  "penis_length",
];

const stringCriteria: CriterionType[] = [
  "name",
  "disambiguation",
  "details",
  "country",
  "hair_color",
  "eye_color",
  "measurements",
  "fake_tits",
  "tattoos",
  "piercings",
  "aliases",
];

const criterionOptions = [
  GenderCriterionOption,
  CircumcisedCriterionOption,
  TagsCriterionOption,
  StudiosCriterionOption,
  createMandatoryNumberCriterionOption("scene_count"),
  createMandatoryNumberCriterionOption("gallery_count"),
  new NumberCriterionOption("height", "height_cm", "height_cm"),
  ...numberCriteria.map((c) => createNumberCriterionOption(c)),
  ...stringCriteria.map((c) => createStringCriterionOption(c)),
  createDateCriterionOption("birthdate"),
  createDateCriterionOption("death_date"),
];
export const PerformerListFilterOptions = new ListFilterOptions(
  defaultSortBy,
  sortByOptions,
  displayModeOptions,
  criterionOptions
);
