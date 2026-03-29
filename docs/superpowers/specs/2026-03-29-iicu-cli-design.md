# iicu CLI — Design Spec

CLI tools for intervals.icu written in Go, designed for AI agent accessibility and wrapped in a Claude skill for coaching workflows.

## Goals

- Full CRUD access to the intervals.icu API via a single cross-platform binary
- JSON-first output optimized for AI agents, with optional human-readable formatting
- Multi-profile config supporting coaches managing multiple athletes
- Claude skill enabling AI-driven coaching workflows (training plans, performance analysis, recovery checks)
- Open-source, distributed as pre-built binaries for macOS, Linux, and Windows

## Non-Goals

- OAuth authentication (API key covers CLI use cases)
- Caching or offline mode (intervals.icu is the source of truth)
- GUI or TUI interface
- Package manager distribution (homebrew, apt — future work)

---

## Architecture

### Project Structure

```
iicu/
├── cmd/
│   ├── root.go              # Root command, global flags, config loading
│   ├── activities.go         # iicu activities [list|get|search|update|delete|streams|intervals|efforts]
│   ├── athlete.go            # iicu athlete [profile|fitness]
│   ├── chat.go               # iicu chat [list|get|messages|send|delete-message|activity-comments|comment]
│   ├── config_cmd.go         # iicu config [init|profiles|use|add|remove]
│   ├── curves.go             # iicu curves [power|hr|pace]
│   ├── custom_items.go       # iicu custom-items [list|get|create|update|delete|reorder]
│   ├── events.go             # iicu events [list|get|create|update|delete|bulk-create|bulk-delete|duplicate]
│   ├── gear.go               # iicu gear [list|create|update|delete|add-reminder|update-reminder]
│   ├── routes.go             # iicu routes [list|get|update|compare]
│   ├── shared_events.go      # iicu shared-events [get|create|update|delete]
│   ├── sports.go             # iicu sports [list|create|update|delete|apply]
│   ├── training_plan.go      # iicu training-plan [get|set]
│   ├── weather.go            # iicu weather [config|update-config|forecast]
│   ├── wellness.go           # iicu wellness [get|update]
│   └── workouts.go           # iicu workouts [folders|list|create|update|delete|bulk-create|duplicate|tags|apply-plan]
├── api/
│   ├── client.go             # HTTP client, auth, base URL, error handling
│   └── types.go              # Go structs matching API responses
├── config/
│   └── config.go             # Profile management, config file R/W
├── skill/
│   └── intervals-icu-coaching.md  # Claude skill for agentskills.io marketplace
├── docs/
│   └── intervals-icu-api-reference.md  # Full API reference (110 endpoints)
├── main.go                   # Entry point
├── Makefile                  # Build targets
├── go.mod
└── go.sum
```

### Dependencies

- `github.com/spf13/cobra` — CLI framework
- `github.com/spf13/viper` — config management
- Go standard library for everything else (`net/http`, `encoding/json`, `fmt`, `os`)

### Module Path

`github.com/elricho/iicu`

---

## Command Groups

### `iicu athlete`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `profile` | `GET /athlete/{id}` | Get athlete profile (zones, FTP, weight, etc.) |
| `fitness` | `GET /athlete/{id}/athlete-summary` | Get fitness summary (CTL, ATL, TSB) |

### `iicu activities`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `list` | `GET /athlete/{id}/activities` | List recent activities with date range filters |
| `get <id>` | `GET /activity/{id}` | Get full activity details |
| `search <query>` | `GET /athlete/{id}/activities/search` | Search activities by name/tag |
| `update <id>` | `PUT /activity/{id}` | Update activity fields |
| `delete <id>` | `DELETE /activity/{id}` | Delete an activity |
| `streams <id>` | `GET /activity/{id}/streams` | Get activity data streams (power, HR, etc.) |
| `intervals <id>` | `GET /activity/{id}/intervals` | Get activity intervals/laps |
| `efforts <id>` | `GET /activity/{id}/best-efforts` | Get best efforts for an activity |

### `iicu events`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `list` | `GET /athlete/{id}/events` | List calendar events (date range) |
| `get <id>` | `GET /athlete/{id}/events/{eventId}` | Get single event |
| `create` | `POST /athlete/{id}/events` | Create event/workout on calendar |
| `update <id>` | `PUT /athlete/{id}/events/{eventId}` | Update an event |
| `delete <id>` | `DELETE /athlete/{id}/events/{eventId}` | Delete an event |
| `bulk-create` | `POST /athlete/{id}/events/bulk` | Create multiple events at once |
| `bulk-delete` | `PUT /athlete/{id}/events/bulk-delete` | Delete events by ID |
| `duplicate <id>` | `POST /athlete/{id}/duplicate-events` | Duplicate events to other dates |

**Event categories:** `WORKOUT`, `RACE_A`, `RACE_B`, `RACE_C`, `NOTE`, `PLAN`, `HOLIDAY`, `SICK`, `INJURED`, `SET_EFTP`, `FITNESS_DAYS`, `SEASON_START`, `TARGET`, `SET_FITNESS`

### `iicu wellness`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `get` | `GET /athlete/{id}/wellness` or `GET /athlete/{id}/wellness/{date}` | Get wellness data (date or range) |
| `update` | `PUT /athlete/{id}/wellness/{date}` | Update wellness entry |

### `iicu curves`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `power` | `GET /athlete/{id}/power-curves` | Power duration curves |
| `hr` | `GET /athlete/{id}/hr-curves` | Heart rate curves |
| `pace` | `GET /athlete/{id}/pace-curves` | Pace curves |

### `iicu workouts`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `folders` | `GET /athlete/{id}/folders` | List workout library folders |
| `list` | `GET /athlete/{id}/workouts` | List all workouts or workouts in a folder |
| `create` | `POST /athlete/{id}/workouts` | Create a workout |
| `update <id>` | `PUT /athlete/{id}/workouts/{workoutId}` | Update a workout |
| `delete <id>` | `DELETE /athlete/{id}/workouts/{workoutId}` | Delete a workout |
| `bulk-create` | `POST /athlete/{id}/workouts/bulk` | Create multiple workouts |
| `duplicate` | `POST /athlete/{id}/duplicate-workouts` | Duplicate workouts on a plan |
| `tags` | `GET /athlete/{id}/workout-tags` | List all workout tags |
| `apply-plan` | `PUT /athlete/{id}/apply-plan-changes` | Apply plan changes to calendar |

### `iicu gear`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `list` | `GET /athlete/{id}/gear` | List all gear |
| `create` | `POST /athlete/{id}/gear` | Add new gear |
| `update <id>` | `PUT /athlete/{id}/gear/{gearId}` | Update gear details |
| `delete <id>` | `DELETE /athlete/{id}/gear/{gearId}` | Delete gear |
| `add-reminder <id>` | `POST /athlete/{id}/gear/{gearId}/reminder` | Add maintenance reminder |
| `update-reminder <id> <rid>` | `PUT /athlete/{id}/gear/{gearId}/reminder/{reminderId}` | Update a reminder |

### `iicu sports`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `list` | `GET /athlete/{athleteId}/sport-settings` | List sport settings |
| `create` | `POST /athlete/{athleteId}/sport-settings` | Create sport config |
| `update <id>` | `PUT /athlete/{athleteId}/sport-settings/{id}` | Update sport settings |
| `delete <id>` | `DELETE /athlete/{athleteId}/sport-settings/{id}` | Delete sport settings |
| `apply <id>` | `PUT /athlete/{athleteId}/sport-settings/{id}/apply` | Apply settings to activities |

### `iicu weather`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `config` | `GET /athlete/{id}/weather-config` | Get weather forecast config |
| `update-config` | `PUT /athlete/{id}/weather-config` | Update forecast locations |
| `forecast` | `GET /athlete/{id}/weather-forecast` | Get weather forecast |

### `iicu routes`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `list` | `GET /athlete/{id}/routes` | List routes with activity counts |
| `get <id>` | `GET /athlete/{id}/routes/{route_id}` | Get a route |
| `update <id>` | `PUT /athlete/{id}/routes/{route_id}` | Update a route |
| `compare <id> <other-id>` | `GET /athlete/{id}/routes/{route_id}/similarity/{other_id}` | Route similarity comparison |

### `iicu chat`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `list` | `GET /athlete/{id}/chats` | List chats |
| `get <id>` | `GET /chats/{id}` | Get a chat |
| `messages <id>` | `GET /chats/{id}/messages` | List messages in a chat |
| `send` | `POST /chats/send-message` | Send a message |
| `delete-message <chat-id> <msg-id>` | `DELETE /chats/{id}/messages/{msgId}` | Delete a message |
| `activity-comments <activity-id>` | `GET /activity/{id}/messages` | List activity comments |
| `comment <activity-id>` | `POST /activity/{id}/messages` | Add comment to activity |

### `iicu custom-items`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `list` | `GET /athlete/{id}/custom-item` | List custom items |
| `get <id>` | `GET /athlete/{id}/custom-item/{itemId}` | Get a custom item |
| `create` | `POST /athlete/{id}/custom-item` | Create a custom item |
| `update <id>` | `PUT /athlete/{id}/custom-item/{itemId}` | Update a custom item |
| `delete <id>` | `DELETE /athlete/{id}/custom-item/{itemId}` | Delete a custom item |
| `reorder` | `PUT /athlete/{id}/custom-item-indexes` | Re-order custom items |

### `iicu shared-events`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `get <id>` | `GET /shared-event/{id}` | Get a shared event |
| `create` | `POST /shared-event` | Create a shared event |
| `update <id>` | `PUT /shared-event/{id}` | Update a shared event |
| `delete <id>` | `DELETE /shared-event/{id}` | Delete a shared event |

### `iicu training-plan`

| Command | API Endpoint | Description |
|---------|-------------|-------------|
| `get` | `GET /athlete/{id}/training-plan` | Get athlete's training plan |
| `set` | `PUT /athlete/{id}/training-plan` | Change athlete's training plan |

### `iicu config`

| Command | Description |
|---------|-------------|
| `init` | Interactive setup — API key, athlete ID, creates config file |
| `profiles` | List configured profiles |
| `use <name>` | Switch default profile |
| `add <name>` | Add a new profile |
| `remove <name>` | Remove a profile |

### Global Flags

| Flag | Description |
|------|-------------|
| `--profile <name>` | Use a specific profile |
| `--athlete <id>` | Override athlete ID |
| `--api-key <key>` | Override API key |
| `--json` | Force JSON output (default) |
| `--human` | Pretty-printed table output |
| `--help` | Help for any command |

---

## Authentication & Config

### Config File

Location: `~/.config/iicu/config.yaml`

```yaml
default_profile: me

profiles:
  me:
    api_key: "abc123..."
    athlete_id: "i12345"

  jane:
    api_key: "def456..."
    athlete_id: "i67890"

  bob:
    api_key: "ghi789..."
    athlete_id: "i11111"
```

### Override Layering (highest wins)

1. `--api-key` / `--athlete-id` command flags
2. `IICU_API_KEY` / `IICU_ATHLETE_ID` environment variables
3. `--profile <name>` flag (selects from config file)
4. `default_profile` in config file

---

## API Client

### `api/client.go`

```go
type Client struct {
    BaseURL    string       // https://intervals.icu/api/v1
    APIKey     string
    AthleteID  string
    HTTPClient *http.Client
}
```

### Behaviors

- **Auth:** HTTP Basic Auth on every request (username: `API_KEY`, password: the actual key)
- **Athlete ID:** Always sends the real athlete ID from config, never the "0" shortcut
- **Errors:** Parses API error responses into structured `APIError` with status code, message, body
- **Date handling:** Helper accepting flexible input (`2024-01-15`, `today`, `yesterday`, `-7d`) and always outputting ISO-8601
- **Types:** Go structs in `api/types.go` matching API models — add as needed, not all 101 upfront
- **Output:** Commands marshal API responses directly to stdout as JSON

### Not Included (intentionally)

- No retry logic or rate limiting (not documented by intervals.icu)
- No caching (API is source of truth)
- No OAuth (API key covers CLI use cases)

---

## Skill

File: `skill/intervals-icu-coaching.md`

Published to the agentskills.io marketplace so Claude can use `iicu` as a coaching tool.

### Frontmatter

```yaml
---
name: intervals-icu-coaching
description: Use when managing training plans, analyzing athlete performance, creating workouts, or interacting with intervals.icu data for cycling, running, or triathlon coaching
---
```

### Contents

- **Overview** — what `iicu` is, how it connects to intervals.icu
- **Setup** — how to verify `iicu` is available and configured (`iicu athlete profile`)
- **Command reference** — all 14 resource groups with common examples
- **Coaching workflows:**
  - Build a training block (read profile → check fitness → create workouts → schedule events)
  - Weekly training review (recent activities → wellness trends → planned vs actual load)
  - Race preparation (power curves → targets → taper plan)
  - Recovery check (wellness → ATL/CTL/TSB → adjust upcoming workouts)
- **Data interpretation** — CTL/ATL/TSB, power curves, zone definitions, compliance scores
- **Date handling** — format examples
- **Output format** — JSON structure, key fields per response type

The skill focuses on **when and why** to use commands. Agents can run `iicu <command> --help` for flag-level syntax.

---

## Build & Distribution

### Makefile Targets

| Target | Description |
|--------|-------------|
| `build` | Build for local platform |
| `build-all` | Cross-compile: macOS (arm64, amd64), Linux (arm64, amd64), Windows (amd64) |
| `test` | Run unit tests |
| `install` | Build and copy to `$GOPATH/bin` |

### Distribution

- GitHub releases with pre-built binaries for all platforms
- Users download binary, put on PATH
- No package manager integration initially

### Testing

- **Unit tests** for `api/` — mock HTTP responses, verify request construction and response parsing
- **Unit tests** for `config/` — profile CRUD, layering logic
- **Integration tests** (skipped by default, need real API key) — smoke-test actual API calls
