# Intervals.icu API Reference

> Generated from OpenAPI spec v1.0.0

---

## Overview

| Item | Value |
|------|-------|
| **Base URL** | `https://intervals.icu` |
| **API prefix** | `/api/v1` |
| **Date format** | ISO-8601 local dates (`2024-01-15`) and datetimes (`2024-01-15T16:18:49`) |
| **Spec version** | v1.0.0 |

### Authentication

Two authentication methods are supported (either may be used):

| Method | Type | Details |
|--------|------|---------|
| **API Key** (Basic Auth) | HTTP Basic | Username: `API_KEY`, Password: your API key from `/settings` |
| **OAuth Bearer Token** | HTTP Bearer | Use bearer token obtained via OAuth flow for an athlete |

### Conventions

- **Athlete ID `"0"`** in path parameters resolves to the authenticated user's own athlete ID.
- **All dates** use ISO-8601 local format (no timezone offset): `2024-01-15` for dates, `2024-01-15T16:18:49` for datetimes.
- **CSV output**: Many endpoints support `.csv` extension in the path (e.g., `/wellness.csv`, `/activities.csv`).
- **Activity uploads** accept FIT, TCX, GPX files (or `.zip`/`.gz` archives of the same).
- **Wellness data** uses metric units only. The `locked` flag on a wellness record prevents external overwrites.
- **Webhooks** are available for events: `ACTIVITY_UPLOADED`, `ACTIVITY_ANALYZED`, `CALENDAR_UPDATED`, and others.
- **Activity types** include: `Ride`, `Run`, `Swim`, `WeightTraining`, `Hike`, `Walk`, `VirtualRide`, `VirtualRun`, `MountainBikeRide`, `GravelRide`, `TrailRun`, `Rowing`, and 50+ others.

---

## Endpoints by Resource

### Athlete

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}` | Get athlete with sportSettings and custom_items | `id`\* (path) | `WithSportSettings` |
| PUT | `/athlete/{id}` | Update an athlete | `id`\* (path); body: `AthleteUpdateDTO` | `Athlete` |
| GET | `/athlete/{id}/profile` | Get athlete profile info | `id`\* (path) | `AthleteProfile` |
| GET | `/athlete/{id}/settings/{deviceClass}` | Get settings for phone/tablet/desktop | `id`\* (path), `deviceClass`\* (path) | `object` |
| GET | `/athlete/{id}/training-plan` | Get the athlete's training plan | `id`\* (path) | `AthleteTrainingPlan` |
| PUT | `/athlete/{id}/training-plan` | Change the athlete's training plan | `id`\* (path); body: `AthleteTrainingPlanUpdate` | `AthleteTrainingPlan` |
| PUT | `/athlete-plans` | Change training plans for a list of athletes | body: `AthleteTrainingPlanUpdate[]` | `object` |
| GET | `/athlete/{id}/athlete-summary{ext}` | Summary info for followed athletes | `id`\* (path), `ext`\* (path), `start` (query), `end` (query), `tags` (query) | `SummaryWithCats[]` |

### Activities

#### Core CRUD

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/activity/{id}` | Get an activity | `id`\* (path), `intervals` (query, bool) | `Activity` or `ActivityWithIntervals` |
| PUT | `/activity/{id}` | Update activity | `id`\* (path); body: `Activity` | `Activity` |
| DELETE | `/activity/{id}` | Delete an activity | `id`\* (path) | `ActivityId` |
| GET | `/athlete/{id}/activities` | List activities for date range (desc order) | `id`\* (path), `oldest`\* (query, ISO date), `newest` (query), `route_id` (query, int64), `limit` (query, int), `fields` (query, array) | `Activity[]` |
| GET | `/athlete/{athleteId}/activities/{ids}` | Fetch multiple activities by id | `athleteId`\* (path), `ids`\* (path, array), `intervals` (query, bool) | `Activity[]` |
| POST | `/athlete/{id}/activities` | Upload activity file (FIT/TCX/GPX/ZIP/GZ) | `id`\* (path), `name` (query), `description` (query), `external_id` (query), `paired_event_id` (query, int); body: multipart file | `UploadResponse` |
| POST | `/athlete/{id}/activities/manual` | Create a manual activity | `id`\* (path); body: `Activity` | `Activity` |
| GET | `/athlete/{id}/activities.csv` | Download activities as CSV | `id`\* (path) | CSV file |

#### Search

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}/activities/search` | Search by name or tag (summary) | `id`\* (path), `q`\* (query), `limit` (query) | `ActivitySearchResult[]` |
| GET | `/athlete/{id}/activities/search-full` | Search by name or tag (full objects) | `id`\* (path), `q`\* (query), `limit` (query) | `Activity[]` |
| GET | `/athlete/{id}/activities/interval-search` | Find activities with matching intervals | `id`\* (path), `minSecs`\*, `maxSecs`\*, `minIntensity`\*, `maxIntensity`\* (all query, int), `type` (query, enum), `minReps` (default 1), `maxReps` (default 999999), `limit` (default 30) | `Activity[]` |
| GET | `/athlete/{id}/activities-around` | Activities before/after another activity | `id`\* (path), `activity_id`\* (query), `route_id` (query, int64), `limit` (query, default 30) | `Activity[]` |
| GET | `/athlete/{id}/activity-tags` | List all activity tags | `id`\* (path) | `string[]` |

#### Intervals

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/activity/{id}/intervals` | Get activity intervals | `id`\* (path) | `IntervalsDTO` |
| PUT | `/activity/{id}/intervals` | Update/replace intervals | `id`\* (path), `all` (query, bool, default true); body: `Interval[]` | `IntervalsDTO` |
| PUT | `/activity/{id}/intervals/{intervalId}` | Update/create a single interval | `id`\* (path), `intervalId`\* (path, int); body: `Interval` | `IntervalsDTO` |
| PUT | `/activity/{id}/delete-intervals` | Delete intervals | `id`\* (path); body: `Interval[]` | `IntervalsDTO` |
| PUT | `/activity/{id}/split-interval` | Split an interval | `id`\* (path), `splitAt`\* (query, int) | `IntervalsDTO` |
| GET | `/activity/{id}/interval-stats` | Stats for part of an activity | `id`\* (path), `start_index`\* (query, int), `end_index`\* (query, int) | `Interval` |
| GET | `/activity/{id}/best-efforts` | Find best efforts in activity | `id`\* (path), `stream`\* (query), `duration` (query, int, secs), `distance` (query, float, meters), `count` (default 8), `minValue`, `excludeIntervals`, `startIndex`, `endIndex` | `BestEfforts` |

#### Streams and Analysis

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/activity/{id}/streams{ext}` | List streams for the activity | `id`\* (path), `ext`\* (path), `types` (query, array) | `ActivityStream[]` |
| GET | `/activity/{id}/map` | Get activity map data | `id`\* (path), `bounds` (query, array), `boundsOnly` (default false), `weather` (default false) | `MapData` |
| GET | `/activity/{id}/segments` | Get activity segments | `id`\* (path) | `IcuSegment[]` |
| GET | `/activity/{id}/weather-summary` | Get weather summary | `id`\* (path), `start_index` (query, default 0), `end_index` (query, default 0) | `ActivityWeatherSummary` |
| GET | `/activity/{id}/time-at-hr` | Time at heart rate data | `id`\* (path) | `Plot` |

#### Power Curves (Activity-Level)

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/activity/{id}/power-curve{ext}` | Activity power curve (JSON/CSV) | `id`\* (path), `ext`\* (path), `fatigue` (query, `kj0`/`kj1`) | `PowerCurve` |
| GET | `/activity/{id}/power-curves{ext}` | Multi-stream power curves | `id`\* (path), `ext`\* (path), `types` (query, array), `fatigue` (query, array) | `PowerCurve[]` |
| GET | `/activity/{id}/power-histogram` | Power histogram | `id`\* (path), `bucketSize` (query, default 25) | `Bucket[]` |
| GET | `/activity/{id}/power-spike-model` | Power spike detection model | `id`\* (path) | `PowerModel` |
| GET | `/activity/{id}/power-vs-hr{ext}` | Power vs heart rate data | `id`\* (path), `ext`\* (path) | `PowerVsHRPlot` |

#### HR Curves (Activity-Level)

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/activity/{id}/hr-curve{ext}` | Activity HR curve (JSON/CSV) | `id`\* (path), `ext`\* (path) | `HRCurve` |
| GET | `/activity/{id}/hr-histogram` | HR histogram | `id`\* (path), `bucketSize` (query, default 5) | `Bucket[]` |
| GET | `/activity/{id}/hr-load-model` | HR training load model | `id`\* (path) | `HRLoadModel` |

#### Pace Curves (Activity-Level)

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/activity/{id}/pace-curve{ext}` | Activity pace curve (JSON/CSV) | `id`\* (path), `ext`\* (path), `gap` (query, bool, default false) | `PaceCurve` |
| GET | `/activity/{id}/pace-histogram` | Pace histogram | `id`\* (path) | `Bucket[]` |
| GET | `/activity/{id}/gap-histogram` | Gradient adjusted pace histogram | `id`\* (path) | `Bucket[]` |

#### File Downloads

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/activity/{id}/file` | Download original activity file | `id`\* (path) | file |
| GET | `/activity/{id}/fit-file` | Download generated FIT file | `id`\* (path), `power` (default true), `hr` (default true) | file |
| GET | `/activity/{id}/gpx-file` | Download generated GPX file | `id`\* (path), `power` (default true), `hr` (default true) | file |
| POST | `/athlete/{id}/download-fit-files` | Download zip of generated FIT files | `id`\* (path), `ids`\* (query, array), `power` (default true), `hr` (default true) | ZIP file |

### Performance Curves (Athlete-Level)

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}/power-curves{ext}` | Best power curves for athlete | `id`\*, `ext`\* (path), `type`\* (query, sport enum), `newest`, `curves` (query, array), `includeRanks`, `subMaxEfforts`, `now`, `pmType` (enum: `MS_2P`, `MORTON_3P`, `FFT_CURVES`, `ECP`), `filters`, `f1`\*, `f2`\*, `f3`\* | `DataCurveSetPowerCurve` |
| GET | `/athlete/{id}/hr-curves{ext}` | Best HR curves for athlete | `id`\*, `ext`\* (path), `type` (query, sport enum), `newest`, `curves`, `subMaxEfforts`, `now`, `filters`, `f1`\*, `f2`\*, `f3`\* | `DataCurveSetHRCurve` |
| GET | `/athlete/{id}/pace-curves{ext}` | Best pace curves for athlete | `id`\*, `ext`\* (path), `type` (query, sport enum), `newest`, `curves`, `includeRanks`, `subMaxEfforts`, `now`, `gap`, `pmType` (enum: `CS`), `filters`, `f1`\*, `f2`\*, `f3`\* | `DataCurveSetPaceCurve` |
| GET | `/athlete/{id}/activity-power-curves{ext}` | Best power for durations across activities | `id`\*, `ext`\* (path), `oldest`\*, `newest`\* (query, ISO date), `type` (query, sport enum), `secs` (query, array), `fatigue`, `filters` | `ActivityPowerCurvePayload` |
| GET | `/athlete/{id}/activity-hr-curves{ext}` | Best HR for durations across activities | `id`\*, `ext`\* (path), `oldest`\*, `newest`\* (query, ISO date), `type` (query, sport enum), `secs` (query, array), `filters` | `ActivityHRCurvePayload` |
| GET | `/athlete/{id}/activity-pace-curves{ext}` | Best pace for distances across activities | `id`\*, `ext`\* (path), `oldest`\*, `newest`\* (query, ISO date), `type` (query, sport enum), `distances` (query, array, meters), `gap` (bool), `filters` | (no body documented) |
| GET | `/athlete/{id}/power-hr-curve` | Power vs HR curve for date range | `id`\* (path), `start`\* (query, ISO date), `end`\* (query, ISO date) | `PowerHRCurve` |
| GET | `/athlete/{id}/mmp-model` | Power model for %MMP workout steps | `id`\* (path), `type`\* (query, sport enum) | `PowerModel` |

### Events / Calendar

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}/events{format}` | List events on calendar (.csv for CSV) | `id`\*, `format`\* (path), `oldest` (query, ISO date), `newest` (query), `category` (query, array), `limit` (query), `calendar_id` (query), `ext` (query, workout format), `powerRange`, `hrRange`, `paceRange`, `locale`, `resolve` (bool) | `Event[]` |
| POST | `/athlete/{id}/events` | Create an event | `id`\* (path); body: `EventEx` | `Event` |
| PUT | `/athlete/{id}/events` | Update all events for date range | `id`\* (path), `oldest`\*, `newest`\* (query, ISO date); body: `Event` | `Event[]` |
| DELETE | `/athlete/{id}/events` | Delete events for date range | `id`\* (path), `oldest`\* (query), `newest` (query), `createdById` (query), `category`\* (query, array) | (no body) |
| POST | `/athlete/{id}/events/bulk` | Create multiple events | `id`\* (path), `upsert` (query, bool, default false); body: `EventEx[]` | `Event[]` |
| PUT | `/athlete/{id}/events/bulk-delete` | Delete events by id/external_id | `id`\* (path); body: `DoomedEvent[]` | `DeleteEventsResponse` |
| GET | `/athlete/{id}/events/{eventId}` | Get an event | `id`\*, `eventId`\* (path) | `Event` |
| PUT | `/athlete/{id}/events/{eventId}` | Update an event | `id`\*, `eventId`\* (path); body: `EventEx` | `Event` |
| DELETE | `/athlete/{id}/events/{eventId}` | Delete an event | `id`\*, `eventId`\* (path), `others` (query, bool), `notBefore` (query, ISO date) | `object` |
| POST | `/athlete/{id}/events/{eventId}/mark-done` | Create manual activity from planned workout | `id`\*, `eventId`\* (path) | `Activity` |
| GET | `/athlete/{id}/events/{eventId}/download{ext}` | Download planned workout (zwo/mrc/erg/fit) | `id`\*, `eventId`\*, `ext`\* (path) | file |
| POST | `/athlete/{id}/duplicate-events` | Duplicate events on calendar | `id`\* (path); body: `DuplicateEventsDTO` | `Event[]` |
| GET | `/athlete/{id}/event-tags` | List all event tags | `id`\* (path) | `string[]` |
| GET | `/athlete/{id}/fitness-model-events` | Events influencing fitness calculation | `id`\* (path) | `Event[]` |
| GET | `/athlete/{id}/workouts.zip` | Download workouts as zip (zwo/mrc/erg/fit) | `id`\* (path), `ext`\* (query), `oldest`\*, `newest`\* (query, ISO date), `powerRange`, `hrRange`, `paceRange`, `locale` | ZIP file |

**Event categories:** `WORKOUT`, `RACE_A`, `RACE_B`, `RACE_C`, `NOTE`, `PLAN`, `HOLIDAY`, `SICK`, `INJURED`, `SET_EFTP`, `FITNESS_DAYS`, `SEASON_START`, `TARGET`, `SET_FITNESS`

### Wellness

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}/wellness{ext}` | List wellness records for date range | `id`\*, `ext`\* (path), `oldest` (query, ISO date), `newest` (query), `cols` (query, array), `fields` (query, array) | `Wellness[]` |
| GET | `/athlete/{id}/wellness/{date}` | Get wellness record for date | `id`\*, `date`\* (path, ISO date) | `Wellness` |
| PUT | `/athlete/{id}/wellness/{date}` | Update wellness record for date | `id`\*, `date`\* (path); body: `Wellness` | `Wellness` |
| PUT | `/athlete/{id}/wellness` | Update a wellness record (id = day) | `id`\* (path); body: `Wellness` | `Wellness` |
| POST | `/athlete/{id}/wellness` | Upload wellness CSV | `id`\* (path), `ignoreMissingFields` (query, bool, default false); body: multipart file | `object` |
| PUT | `/athlete/{id}/wellness-bulk` | Update multiple wellness records | `id`\* (path); body: `Wellness[]` | (no body) |

### Workouts / Library

#### Folders and Plans

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}/folders` | List all folders, plans, and workouts | `id`\* (path) | `Folder[]` |
| POST | `/athlete/{id}/folders` | Create a folder or plan | `id`\* (path); body: `CreateFolderDTO` | `Folder` |
| PUT | `/athlete/{id}/folders/{folderId}` | Update a folder or plan | `id`\*, `folderId`\* (path); body: `Folder` | `Folder` |
| DELETE | `/athlete/{id}/folders/{folderId}` | Delete folder/plan and all workouts | `id`\*, `folderId`\* (path) | `object` |
| GET | `/athlete/{id}/folders/{folderId}/shared-with` | List shared athletes for folder | `id`\*, `folderId`\* (path) | `SharedWith[]` |
| PUT | `/athlete/{id}/folders/{folderId}/shared-with` | Add/remove shared athletes | `id`\*, `folderId`\* (path); body: `SharedWith[]` | `SharedWith[]` |
| PUT | `/athlete/{id}/folders/{folderId}/workouts` | Update workouts on a plan (hide_from_athlete) | `id`\*, `folderId`\* (path), `oldest`\*, `newest`\* (query, int); body: `Workout` | `Workout[]` |
| POST | `/athlete/{id}/folders/{folderId}/import-workout` | Import workout from file (zwo/mrc/erg/fit) | `id`\*, `folderId`\* (path), `type`\* (query, sport enum); body: multipart file | `Workout` |

#### Workouts

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}/workouts` | List all workouts in library | `id`\* (path) | `Workout[]` |
| POST | `/athlete/{id}/workouts` | Create a workout in library | `id`\* (path); body: `WorkoutEx` | `Workout` |
| POST | `/athlete/{id}/workouts/bulk` | Create multiple workouts | `id`\* (path); body: `WorkoutEx[]` | `Workout[]` |
| GET | `/athlete/{id}/workouts/{workoutId}` | Get a workout | `id`\*, `workoutId`\* (path) | `Workout` |
| PUT | `/athlete/{id}/workouts/{workoutId}` | Update a workout | `id`\*, `workoutId`\* (path); body: `WorkoutEx` | `Workout` |
| DELETE | `/athlete/{id}/workouts/{workoutId}` | Delete a workout | `id`\*, `workoutId`\* (path), `others` (query, bool) | `int[]` |
| POST | `/athlete/{id}/duplicate-workouts` | Duplicate workouts on a plan | `id`\* (path); body: `DuplicateWorkoutsDTO` | `Workout[]` |
| GET | `/athlete/{id}/workout-tags` | List all workout tags | `id`\* (path) | `string[]` |
| PUT | `/athlete/{id}/apply-plan-changes` | Apply plan changes to calendar | `id`\* (path) | `object` |

#### Workout Conversion

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| POST | `/athlete/{id}/download-workout{ext}` | Convert workout to zwo/mrc/erg/fit | `id`\*, `ext`\* (path); body: `Workout` | file |
| POST | `/download-workout{ext}` | Convert workout (no athlete context) | `ext`\* (path); body: `Workout` | file |

### Gear

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}/gear{ext}` | List athlete gear (.csv for CSV) | `id`\*, `ext`\* (path) | `Gear[]` |
| POST | `/athlete/{id}/gear` | Create gear or component | `id`\* (path); body: `Gear` | `Gear` |
| PUT | `/athlete/{id}/gear/{gearId}` | Update gear or component | `id`\*, `gearId`\* (path); body: `Gear` | `Gear` |
| DELETE | `/athlete/{id}/gear/{gearId}` | Delete gear or component | `id`\*, `gearId`\* (path) | (no body) |
| GET | `/athlete/{id}/gear/{gearId}/calc` | Recalculate gear stats | `id`\*, `gearId`\* (path) | `GearStats` |
| POST | `/athlete/{id}/gear/{gearId}/replace` | Retire component and replace with copy | `id`\*, `gearId`\* (path); body: `Gear` | `Gear[]` |
| POST | `/athlete/{id}/gear/{gearId}/reminder` | Create a gear reminder | `id`\*, `gearId`\* (path); body: `GearReminder` | `Gear` |
| PUT | `/athlete/{id}/gear/{gearId}/reminder/{reminderId}` | Update a gear reminder | `id`\*, `gearId`\*, `reminderId`\* (path), `reset`\* (query, bool), `snoozeDays`\* (query, int); body: `GearReminder` | `Gear` |
| DELETE | `/athlete/{id}/gear/{gearId}/reminder/{reminderId}` | Delete a gear reminder | `id`\*, `gearId`\*, `reminderId`\* (path) | `Gear` |

**Gear types:** `Bike`, `Shoes`, `Wetsuit`, `RowingMachine`, `Skis`, `Snowboard`, `Equipment`, `Trainer`, `Tyre`, `Wheel`, `Chain`, `Cassette`, `PowerMeter`, and 30+ component types.

### Sport Settings

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{athleteId}/sport-settings` | List sport settings | `athleteId`\* (path) | `SportSettings[]` |
| POST | `/athlete/{athleteId}/sport-settings` | Create sport settings with defaults | `athleteId`\* (path); body: `SportSettings` | `SportSettings` |
| PUT | `/athlete/{athleteId}/sport-settings` | Update multiple sport settings | `athleteId`\* (path), `recalcHrZones`\* (query, bool); body: `SportSettings[]` | `SportSettings[]` |
| GET | `/athlete/{athleteId}/sport-settings/{id}` | Get sport settings by id or type (e.g. Run) | `athleteId`\*, `id`\* (path) | `SportSettings` |
| PUT | `/athlete/{athleteId}/sport-settings/{id}` | Update sport settings by id or type | `athleteId`\*, `id`\* (path), `recalcHrZones`\* (query, bool); body: `SportSettings` | `SportSettings` |
| DELETE | `/athlete/{athleteId}/sport-settings/{id}` | Delete sport settings | `athleteId`\*, `id`\* (path) | `object` |
| PUT | `/athlete/{athleteId}/sport-settings/{id}/apply` | Apply settings to matching activities (async) | `athleteId`\*, `id`\* (path) | `object` |
| GET | `/athlete/{athleteId}/sport-settings/{id}/matching-activities` | List activities matching settings | `athleteId`\*, `id`\* (path) | `ActivityMini[]` |
| GET | `/athlete/{athleteId}/sport-settings/{id}/pace_distances` | Pace curve distances and best effort defaults | `athleteId`\*, `id`\* (path) | `PaceDistancesDTO` |
| GET | `/pace_distances` | List pace curve distances (global) | (none) | `PaceDistancesDTO` |

### Weather

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}/weather-config` | Get weather forecast configuration | `id`\* (path) | `WeatherConfig` |
| PUT | `/athlete/{id}/weather-config` | Update weather forecast configuration | `id`\* (path); body: `WeatherConfig` | `WeatherConfig` |
| GET | `/athlete/{id}/weather-forecast` | Get weather forecast | `id`\* (path) | `WeatherDTO` |

### Routes

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}/routes` | List routes with activity counts | `id`\* (path) | `WithCount[]` |
| GET | `/athlete/{id}/routes/{route_id}` | Get a route | `id`\*, `route_id`\* (path, int64), `includePath` (query, bool, default false) | `AthleteRoute` |
| PUT | `/athlete/{id}/routes/{route_id}` | Update a route | `id`\*, `route_id`\* (path); body: `AthleteRoute` | `AthleteRoute` |
| GET | `/athlete/{id}/routes/{route_id}/similarity/{other_id}` | Route similarity comparison | `id`\*, `route_id`\*, `other_id`\* (path, all int64) | `RouteSimilarity` |

### Chats / Messages

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}/chats` | List chats (most recent first) | `id`\* (path) | `Chat[]` |
| GET | `/chats/{id}` | Get a chat by id | `id`\* (path, int) | `Chat` |
| GET | `/chats/{id}/messages` | List messages (most recent first) | `id`\* (path, int), `beforeId` (query, int64), `limit` (query, int, default 30, max 100) | `Message[]` |
| POST | `/chats/send-message` | Send a message | body: `NewMessage` | `SendResponse` |
| DELETE | `/chats/{id}/messages/{msgId}` | Delete a message | `id`\* (path, int), `msgId`\* (path, int64) | `object` |
| PUT | `/chats/{id}/messages/{msgId}/seen` | Mark message as last seen | `id`\* (path, int), `msgId`\* (path, int64) | `object` |
| GET | `/activity/{id}/messages` | List activity comments | `id`\* (path), `sinceId` (query, int64), `limit` (query, int, default 100) | `Message[]` |
| POST | `/activity/{id}/messages` | Add comment to activity | `id`\* (path); body: `NewActivityMsg` | `NewMsg` |

### Custom Items

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| GET | `/athlete/{id}/custom-item` | List custom items | `id`\* (path) | `CustomItem[]` |
| POST | `/athlete/{id}/custom-item` | Create a custom item | `id`\* (path); body: `CustomItem` | `NewCustomItem` |
| GET | `/athlete/{id}/custom-item/{itemId}` | Get a custom item | `id`\*, `itemId`\* (path) | `CustomItem` |
| PUT | `/athlete/{id}/custom-item/{itemId}` | Update a custom item | `id`\*, `itemId`\* (path); body: `CustomItem` | `CustomItem` |
| DELETE | `/athlete/{id}/custom-item/{itemId}` | Delete a custom item | `id`\*, `itemId`\* (path) | (no body) |
| PUT | `/athlete/{id}/custom-item-indexes` | Re-order custom items | `id`\* (path); body: `CustomItem[]` | (no body) |
| POST | `/athlete/{id}/custom-item/{itemId}/image` | Upload image for custom item | `id`\*, `itemId`\* (path); body: multipart file | `CustomItem` |

**Custom item types:** `FITNESS_CHART`, `TRACE_CHART`, `INPUT_FIELD`, `ACTIVITY_FIELD`, `INTERVAL_FIELD`, `ACTIVITY_STREAM`, `ACTIVITY_CHART`, `ACTIVITY_HISTOGRAM`, `ACTIVITY_HEATMAP`, `ACTIVITY_MAP`, `ACTIVITY_PANEL`, `ZONES`

### Shared Events

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| POST | `/shared-event` | Create a shared event (e.g. race) | `linkToEventId` (query, int); body: `SharedEvent` | `SharedEvent` |
| GET | `/shared-event/{id}` | Get a shared event | `id`\* (path, int) | `SharedEvent` |
| PUT | `/shared-event/{id}` | Update a shared event | `id`\* (path, int); body: `SharedEvent` | `SharedEvent` |
| DELETE | `/shared-event/{id}` | Delete a shared event | `id`\* (path, int) | (no body) |

### OAuth / App Management

| Method | Path | Summary | Key Parameters | Response |
|--------|------|---------|---------------|----------|
| DELETE | `/disconnect-app` | Disconnect athlete from OAuth app | (bearer token identifies app) | (no body) |

---

## Key Data Models

### Activity

The core activity object with 150+ fields. Key fields:

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Activity ID |
| `name` | string | Activity name |
| `type` | string | Sport type (Ride, Run, Swim, etc.) |
| `start_date_local` | string | Local start date/time |
| `start_date` | string | UTC start date/time |
| `distance` | float | Distance in meters |
| `moving_time` | int | Moving time in seconds |
| `elapsed_time` | int | Elapsed time in seconds |
| `total_elevation_gain` | float | Elevation gain in meters |
| `average_speed` | float | Average speed (m/s) |
| `max_speed` | float | Max speed (m/s) |
| `average_watts` | int | via `icu_average_watts` |
| `icu_weighted_avg_watts` | int | Normalized power |
| `icu_ftp` | int | FTP at time of activity |
| `icu_training_load` | int | Training load (TSS) |
| `icu_atl` | float | Acute training load (fatigue) |
| `icu_ctl` | float | Chronic training load (fitness) |
| `icu_intensity` | float | Intensity factor |
| `icu_efficiency_factor` | float | Efficiency factor |
| `icu_variability_index` | float | Variability index |
| `average_heartrate` | int | Avg HR (bpm) |
| `max_heartrate` | int | Max HR (bpm) |
| `average_cadence` | float | Avg cadence |
| `calories` | int | Calories burned |
| `device_watts` | bool | Power from device (not estimated) |
| `trainer` | bool | Indoor trainer ride |
| `source` | enum | `STRAVA`, `UPLOAD`, `MANUAL`, `GARMIN_CONNECT`, `POLAR`, `SUUNTO`, `COROS`, `WAHOO`, `ZWIFT`, etc. |
| `sub_type` | enum | `NONE`, `COMMUTE`, `WARMUP`, `COOLDOWN`, `RACE` |
| `perceived_exertion` | float | RPE |
| `icu_rpe` | int | Session RPE |
| `feel` | int | Feel rating |
| `compliance` | float | Workout compliance |
| `tags` | string[] | Activity tags |
| `stream_types` | string[] | Available data streams |
| `paired_event_id` | int | Linked calendar event |
| `icu_color` | string | Display color |
| `gear` | StravaGear | Gear used |
| `icu_zone_times` | ZoneTime[] | Time in each power zone |
| `icu_hr_zone_times` | int[] | Time in each HR zone |
| `pace_zone_times` | int[] | Time in each pace zone |
| `decoupling` | float | Aerobic decoupling % |
| `strava_id` | string | Strava activity ID |
| `external_id` | string | External source ID |
| `weather fields` | various | `average_weather_temp`, `average_wind_speed`, `headwind_percent`, etc. |

### Athlete

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Athlete ID |
| `name` | string | Display name |
| `firstname` | string | First name |
| `lastname` | string | Last name |
| `email` | string | Email address |
| `sex` | string | Gender |
| `weight` | float | Weight (kg) |
| `icu_weight` | float | Current weight |
| `icu_resting_hr` | int | Resting HR |
| `timezone` | string | Timezone |
| `locale` | string | Locale |
| `city`, `state`, `country` | string | Location |
| `visibility` | enum | `PRIVATE`, `PUBLIC`, `HIDDEN` |
| `status` | enum | `ACTIVE`, `DORMANT`, `ARCHIVED` |
| `plan` | enum | `FREE`, `PREMIUM`, `SUPPORTER`, `WHITELABEL` |
| `measurement_preference` | string | Units preference |
| `icu_coach` | bool | Is a coach |
| `icu_api_key` | string | API key |
| `bikes` | StravaGear[] | Bikes from Strava |
| `shoes` | StravaGear[] | Shoes from Strava |
| `training_plan_id` | int | Active training plan |
| `height` | float | Height |
| `icu_date_of_birth` | string | Date of birth |
| `icu_form_as_percent` | bool | Show form as percentage |
| `icu_mmp_days` | int | Days for MMP calculation |

### Event

Calendar events including planned workouts, notes, races.

| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Event ID |
| `start_date_local` | string | Start date (ISO-8601) |
| `end_date_local` | string | End date (multi-day events) |
| `category` | enum | `WORKOUT`, `RACE_A`, `RACE_B`, `RACE_C`, `NOTE`, `PLAN`, `HOLIDAY`, `SICK`, `INJURED`, `SET_EFTP`, `FITNESS_DAYS`, `SEASON_START`, `TARGET`, `SET_FITNESS` |
| `name` | string | Event name |
| `description` | string | Description/notes |
| `type` | string | Activity type |
| `color` | string | Display color |
| `indoor` | bool | Indoor workout |
| `moving_time` | int | Planned duration (seconds) |
| `distance` | float | Planned distance |
| `icu_ftp` | int | FTP used for targets |
| `icu_training_load` | int | Planned load |
| `icu_atl` | float | ATL after event |
| `icu_ctl` | float | CTL after event |
| `icu_intensity` | float | Intensity factor |
| `target` | enum | `AUTO`, `POWER`, `HR`, `PACE` |
| `workout_doc` | object | Structured workout definition |
| `joules` | int | Planned work (joules) |
| `load_target` | int | Load target |
| `time_target` | int | Time target |
| `distance_target` | float | Distance target |
| `tags` | string[] | Tags |
| `external_id` | string | External ID |
| `athlete_cannot_edit` | bool | Lock for coached athletes |
| `hide_from_athlete` | bool | Hidden from athlete |
| `shared_event_id` | int | Linked shared event |
| `sub_type` | enum | `NONE`, `COMMUTE`, `WARMUP`, `COOLDOWN`, `RACE` |

### Wellness

Daily wellness/health metrics.

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Date (ISO-8601 day) |
| `weight` | float | Body weight (kg) |
| `restingHR` | int | Resting heart rate (bpm) |
| `hrv` | float | HRV (rMSSD) |
| `hrvSDNN` | float | HRV (SDNN) |
| `sleepSecs` | int | Sleep duration (seconds) |
| `sleepScore` | float | Sleep score |
| `sleepQuality` | int | Sleep quality (1-5) |
| `avgSleepingHR` | float | Average sleeping HR |
| `soreness` | int | Soreness (1-5) |
| `fatigue` | int | Fatigue (1-5) |
| `stress` | int | Stress (1-5) |
| `mood` | int | Mood (1-5) |
| `motivation` | int | Motivation (1-5) |
| `injury` | int | Injury (1-5) |
| `spO2` | float | Blood oxygen % |
| `systolic` | int | Blood pressure systolic |
| `diastolic` | int | Blood pressure diastolic |
| `hydration` | int | Hydration level |
| `hydrationVolume` | float | Hydration volume |
| `readiness` | float | Readiness score |
| `bodyFat` | float | Body fat % |
| `vo2max` | float | VO2max |
| `kcalConsumed` | int | Calories consumed |
| `steps` | int | Daily steps |
| `respiration` | float | Respiration rate |
| `bloodGlucose` | float | Blood glucose |
| `lactate` | float | Blood lactate |
| `baevskySI` | float | Baevsky stress index |
| `menstrualPhase` | enum | `PERIOD`, `FOLLICULAR`, `OVULATING`, `LUTEAL`, `NONE` |
| `menstrualPhasePredicted` | enum | Same as above (predicted) |
| `comments` | string | Notes |
| `locked` | bool | Prevents external overwrites |
| `tempWeight` | bool | Temporary weight (not used for trends) |
| `tempRestingHR` | bool | Temporary resting HR |
| `ctl` | float | Chronic training load (fitness) |
| `atl` | float | Acute training load (fatigue) |
| `rampRate` | float | CTL ramp rate |
| `sportInfo` | SportInfo[] | Per-sport load info |
| `updated` | datetime | Last update timestamp |

### Workout

Workout definition in the library.

| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Workout ID |
| `athlete_id` | string | Owner athlete |
| `folder_id` | int | Parent folder ID |
| `name` | string | Workout name |
| `description` | string | Description |
| `type` | string | Activity type (Ride, Run, etc.) |
| `indoor` | bool | Indoor workout |
| `color` | string | Display color |
| `moving_time` | int | Duration (seconds) |
| `icu_training_load` | int | Estimated training load |
| `icu_intensity` | float | Estimated intensity |
| `joules` | int | Estimated work (joules) |
| `workout_doc` | object | Structured workout definition |
| `target` | enum | `AUTO`, `POWER`, `HR`, `PACE` |
| `targets` | string[] | Available targets |
| `day` | int | Day within plan |
| `days` | int | Number of days (multi-day) |
| `tags` | string[] | Tags |
| `sub_type` | enum | `NONE`, `COMMUTE`, `WARMUP`, `COOLDOWN`, `RACE` |
| `distance` | float | Target distance |
| `for_week` | bool | Applies to whole week |
| `carbs_per_hour` | int | Nutrition target |
| `updated` | datetime | Last modified |

### SportSettings

Per-sport configuration for zones, thresholds, and display.

| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Settings ID |
| `athlete_id` | string | Athlete ID |
| `types` | string[] | Activity types this applies to |
| `ftp` | int | FTP (watts) |
| `indoor_ftp` | int | Indoor FTP |
| `w_prime` | int | W' (joules) |
| `p_max` | int | P-max (watts) |
| `lthr` | int | Lactate threshold HR (bpm) |
| `max_hr` | int | Maximum HR (bpm) |
| `threshold_pace` | float | Threshold pace |
| `power_zones` | int[] | Power zone boundaries |
| `power_zone_names` | string[] | Power zone names |
| `hr_zones` | int[] | HR zone boundaries |
| `hr_zone_names` | string[] | HR zone names |
| `pace_zones` | float[] | Pace zone boundaries |
| `pace_zone_names` | string[] | Pace zone names |
| `sweet_spot_min` | int | Sweet spot lower bound |
| `sweet_spot_max` | int | Sweet spot upper bound |
| `power_spike_threshold` | int | Spike detection threshold |
| `warmup_time` | int | Default warmup (seconds) |
| `cooldown_time` | int | Default cooldown (seconds) |
| `hr_load_type` | enum | `AVG_HR`, `HR_ZONES`, `HRSS` |
| `pace_load_type` | enum | `SWIM`, `RUN` |
| `pace_units` | enum | `SECS_100M`, `SECS_100Y`, `MINS_KM`, `MINS_MILE`, `SECS_500M` |
| `gap_model` | enum | `NONE`, `STRAVA_RUN` |
| `elevation_correction` | enum | `NO`, `AUTO`, `YES` |
| `load_order` | enum | `POWER_HR_PACE`, `POWER_PACE_HR`, etc. |
| `default_gear_id` | string | Default gear |
| `default_indoor_gear_id` | string | Default indoor gear |
| `extract_workouts` | bool | Auto-extract workout intervals |
| `best_effort_distances` | float[] | Distances for best efforts |
| `power_field` | string | Power field to use |
| `mmp_model` | PowerModel | Power model for %MMP |

### Gear

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Gear ID |
| `athlete_id` | string | Owner athlete |
| `type` | enum | `Bike`, `Shoes`, `Wetsuit`, `Chain`, `Tyre`, `Wheel`, etc. (40+ types) |
| `name` | string | Gear name |
| `purchased` | string | Purchase date |
| `notes` | string | Notes |
| `distance` | float | Total distance (meters) |
| `time` | float | Total time (seconds) |
| `activities` | int | Activity count |
| `retired` | string | Retirement date |
| `component` | bool | Is a component (not top-level gear) |
| `component_ids` | string[] | Child component IDs |
| `reminders` | GearReminder[] | Maintenance reminders |
| `activity_filters` | ActivityFilter[] | Auto-assignment filters |
| `use_elapsed_time` | bool | Use elapsed vs moving time |

### GearReminder

| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Reminder ID |
| `gear_id` | string | Parent gear |
| `name` | string | Reminder name |
| `distance` | float | Distance threshold (meters) |
| `time` | float | Time threshold (seconds) |
| `activities` | int | Activity count threshold |
| `days` | int | Day count threshold |
| `last_reset` | datetime | Last reset time |
| `percent_used` | float | % of threshold used |
| `snoozed_until` | datetime | Snooze until date |

### Interval

| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Interval ID |
| `type` | enum | `RECOVERY`, `WORK` |
| `label` | string | Interval label |
| `start_index` | int | Start stream index |
| `end_index` | int | End stream index |
| `start_time` | int | Start time (seconds) |
| `end_time` | int | End time (seconds) |
| `distance` | float | Distance (meters) |
| `moving_time` | int | Moving time (seconds) |
| `elapsed_time` | int | Elapsed time (seconds) |
| `average_watts` | int | Avg power |
| `weighted_average_watts` | int | Normalized power |
| `max_watts` | int | Max power |
| `average_watts_kg` | float | Watts/kg |
| `intensity` | int | Intensity % |
| `training_load` | float | Interval load |
| `joules` | int | Work done |
| `average_heartrate` | int | Avg HR |
| `max_heartrate` | int | Max HR |
| `average_cadence` | float | Avg cadence |
| `average_speed` | float | Avg speed |
| `total_elevation_gain` | float | Elevation gain |
| `average_gradient` | float | Avg gradient % |
| `zone` | int | Power zone |
| `group_id` | string | Interval group |
| `gap` | float | Gradient adjusted pace |
| `decoupling` | float | Aerobic decoupling |
| `wbal_start` | int | W'bal at start |
| `wbal_end` | int | W'bal at end |
| `strain_score` | float | Strain score |

### Folder

| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Folder ID |
| `athlete_id` | string | Owner athlete |
| `type` | enum | `FOLDER`, `PLAN` |
| `name` | string | Folder/plan name |
| `description` | string | Description |
| `visibility` | enum | `PRIVATE`, `PUBLIC` |
| `children` | Workout[] | Workouts in folder |
| `start_date_local` | string | Plan start date |
| `rollout_weeks` | int | Auto-rollout weeks |
| `num_workouts` | int | Workout count |
| `duration_weeks` | int | Plan duration in weeks |
| `hours_per_week_min` | int | Min hours/week |
| `hours_per_week_max` | int | Max hours/week |
| `activity_types` | string[] | Sports covered |
| `workout_targets` | string[] | Workout targets available |
| `shareToken` | string | Sharing token |

### Chat

| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Chat ID |
| `type` | enum | `PRIVATE`, `GROUP`, `ACTIVITY` |
| `name` | string | Chat name |
| `members` | ChatMember[] | Chat members |
| `join_policy` | enum | `OPEN`, `ASK`, `INVITE_ONLY` |
| `new_message_count` | int | Unread messages |
| `role` | enum | `MEMBER`, `FOLLOWER`, `COACH`, `ADMIN` |

### Message

| Field | Type | Description |
|-------|------|-------------|
| `id` | int64 | Message ID |
| `athlete_id` | string | Author |
| `content` | string | Message content |
| `type` | enum | `TEXT`, `FOLLOW_REQ`, `COACH_REQ`, `COACH_ME_REQ`, `ACTIVITY`, `NOTE`, `JOIN_REQ`, `ACCEPT_COACHING_GROUP` |
| `created` | datetime | Timestamp |
| `activity_id` | string | Linked activity |
| `start_index` | int | Stream range start |
| `end_index` | int | Stream range end |

### SharedEvent

| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Shared event ID |
| `name` | string | Event name |
| `category` | enum | `RACE`, `WORKOUT` |
| `start_date_local` | string | Event date |
| `visibility` | enum | `PUBLIC`, `GROUP` |
| `types` | string[] | Activity types |
| `description` | string | Description |
| `website` | string | Event website |
| `location` | string | Location name |
| `lat`, `lon` | float | Coordinates |
| `country`, `region` | string | Location details |
| `polyline` | string | Route polyline |
| `external_id` | string | External ID |

### CustomItem

| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Item ID |
| `type` | enum | `FITNESS_CHART`, `TRACE_CHART`, `INPUT_FIELD`, `ACTIVITY_FIELD`, `INTERVAL_FIELD`, `ACTIVITY_STREAM`, `ACTIVITY_CHART`, `ACTIVITY_HISTOGRAM`, `ACTIVITY_HEATMAP`, `ACTIVITY_MAP`, `ACTIVITY_PANEL`, `ZONES` |
| `name` | string | Item name |
| `description` | string | Description |
| `visibility` | enum | `PRIVATE`, `FOLLOWERS`, `PUBLIC` |
| `content` | object | Item configuration |
| `index` | int | Display order |

### WeatherConfig

| Field | Type | Description |
|-------|------|-------------|
| `forecasts` | Forecast[] | Configured forecast locations |

### AthleteRoute

| Field | Type | Description |
|-------|------|-------------|
| Route fields include name, distance, elevation, GPS path, and activity linkage. |

### PowerModel

| Field | Type | Description |
|-------|------|-------------|
| Used for %MMP workout resolution and power curve modeling. Model types: `MS_2P`, `MORTON_3P`, `FFT_CURVES`, `ECP`. |

---

## Appendix: All 101 Schema Objects

Activity, ActivityCharts, ActivityFilter, ActivityHRCurve, ActivityHRCurvePayload, ActivityId, ActivityMini, ActivityPowerCurve, ActivityPowerCurvePayload, ActivitySearchResult, ActivityStream, ActivityWeather, ActivityWeatherSummary, ActivityWithIntervals, Anomaly, Athlete, AthleteProfile, AthleteRoute, AthleteSearchResult, AthleteTrainingPlan, AthleteTrainingPlanUpdate, AthleteUpdateDTO, Attachment, BestEfforts, Bucket, CategorySummary, Chat, ChatMember, Closest, CoachTick, CreateFolderDTO, Curve, CustomItem, DataCurve, DataCurvePt, DataCurveSetHRCurve, DataCurveSetPaceCurve, DataCurveSetPowerCurve, DeleteEventsResponse, Display, DoomedEvent, DuplicateEventsDTO, DuplicateWorkoutsDTO, Effort, Event, EventEx, Folder, Forecast, Gear, GearReminder, GearStats, HRCurve, HRLoadModel, HRRecovery, Hidden, IcuAchievement, IcuSegment, Ignore, Interval, IntervalGroup, IntervalsDTO, MapData, Message, NewActivityMsg, NewCustomItem, NewMessage, NewMsg, PaceCurve, PaceDistancesDTO, PaceModel, Plot, Pos, PowerCurve, PowerHRCurve, PowerModel, PowerVsHRPlot, PushError, Rank, RouteSimilarity, SendResponse, Settings, SharedEvent, SharedWith, SportInfo, SportSettings, StravaGear, SummaryWithCats, Time, UploadResponse, WeatherConfig, WeatherDTO, WeatherPoint, Wellness, WindRose, WithCount, WithSportSettings, Workout, WorkoutEx, ZoneInfo, ZoneSet, ZoneTime
