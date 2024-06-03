import {
  createMandatoryNumberCriterionOption,
  createMandatoryStringCriterionOption,
  createStringCriterionOption,
  createDateCriterionOption,
  createMandatoryTimestampCriterionOption,
  createDurationCriterionOption,
} from "./criteria/criterion";
import { HasMarkersCriterionOption } from "./criteria/has-markers";
import { SceneIsMissingCriterionOption } from "./criteria/is-missing";
import { MoviesCriterionOption } from "./criteria/movies";
import { GalleriesCriterionOption } from "./criteria/galleries";
import { OrganizedCriterionOption } from "./criteria/organized";
import { PerformersCriterionOption } from "./criteria/performers";
import { ResolutionCriterionOption } from "./criteria/resolution";
import { StudiosCriterionOption } from "./criteria/studios";
import { InteractiveCriterionOption } from "./criteria/interactive";
import {
  PerformerTagsCriterionOption,
  TagsCriterionOption,
} from "./criteria/tags";
import { ListFilterOptions, MediaSortByOptions } from "./filter-options";
import { DisplayMode } from "./types";
import {
  DuplicatedCriterionOption,
  PhashCriterionOption,
} from "./criteria/phash";
import { PerformerFavoriteCriterionOption } from "./criteria/favorite";
import { CaptionsCriterionOption } from "./criteria/captions";
import { StashIDCriterionOption } from "./criteria/stash-ids";
import { RatingCriterionOption } from "./criteria/rating";
import { PathCriterionOption } from "./criteria/path";
import { OrientationCriterionOption } from "./criteria/orientation";

const defaultSortBy = "date";
const sortByOptions = [
  "date",
  ...MediaSortByOptions,
]
  .map(ListFilterOptions.createSortBy)
  .concat([
    {
      messageID: "o_count",
      value: "o_counter",
    },
  ]);
const displayModeOptions = [
  DisplayMode.Grid,
];

const criterionOptions = [
  createStringCriterionOption("title"),
  createStringCriterionOption("code", "scene_code"),
  createStringCriterionOption("details"),
  createStringCriterionOption("director"),
  RatingCriterionOption,
  createMandatoryNumberCriterionOption("o_counter", "o_count"),
  ResolutionCriterionOption,
  OrientationCriterionOption,
  createMandatoryNumberCriterionOption("framerate"),
  createMandatoryNumberCriterionOption("bitrate"),
  createStringCriterionOption("video_codec"),
  createStringCriterionOption("audio_codec"),
  createDurationCriterionOption("duration"),
  HasMarkersCriterionOption,
  SceneIsMissingCriterionOption,
  TagsCriterionOption,
  PerformersCriterionOption,
  createMandatoryNumberCriterionOption("performer_count"),
  createMandatoryNumberCriterionOption("performer_age"),
  StudiosCriterionOption,
  MoviesCriterionOption,
  GalleriesCriterionOption,
  createStringCriterionOption("url"),
  InteractiveCriterionOption,
  CaptionsCriterionOption,
  createMandatoryNumberCriterionOption("interactive_speed"),
  createDateCriterionOption("date"),
];

export const SceneListFilterOptions = new ListFilterOptions(
  defaultSortBy,
  sortByOptions,
  displayModeOptions,
  criterionOptions
);
