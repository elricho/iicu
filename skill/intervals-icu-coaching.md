---
name: intervals-icu-coaching
description: Use when managing training plans, analyzing athlete performance, creating workouts, checking wellness and fitness data, or interacting with intervals.icu for cycling, running, or triathlon coaching
---

# Intervals.icu Coaching via iicu CLI

## Overview

`iicu` is a CLI tool for the intervals.icu training platform. It provides full API access for managing an athlete's training data — activities, workouts, events, wellness, performance curves, and more.

## Setup Check

Before using, verify the CLI is available and configured:

```bash
iicu athlete profile
```

If this fails with a credentials error, the user needs to run `iicu config init` with their API key from https://intervals.icu/settings.

## Command Groups

| Group | Purpose | Key commands |
|-------|---------|-------------|
| `athlete` | Profile and fitness | `profile`, `fitness` |
| `activities` | Training history | `list`, `get`, `search`, `streams`, `intervals`, `efforts` |
| `events` | Calendar management | `list`, `create`, `update`, `delete`, `bulk-create` |
| `wellness` | Health metrics | `get`, `update` |
| `curves` | Performance curves | `power`, `hr`, `pace` |
| `workouts` | Workout library | `folders`, `list`, `create`, `bulk-create`, `apply-plan` |
| `training-plan` | Plan assignment | `get`, `set` |
| `gear` | Equipment tracking | `list`, `create`, `update` |
| `sports` | Zone/threshold config | `list`, `update`, `apply` |
| `weather` | Forecasts | `forecast` |
| `routes` | Saved routes | `list`, `get`, `compare` |
| `chat` | Messaging | `list`, `messages`, `send`, `activity-comments` |

Run `iicu <group> --help` for full flag details on any command.

## Output

All commands output JSON by default. Use `--human` for pretty-printed output. Use `--profile <name>` to switch between athlete profiles.

## Date Inputs

Date flags accept flexible formats: `2024-06-15`, `today`, `yesterday`, `-7d`, `-30d`.

## Write Operations

Commands that create or update data (`create`, `update`, `bulk-create`) read JSON from stdin:

```bash
echo '{"name":"Easy Ride","category":"WORKOUT","start_date_local":"2024-07-01","type":"Ride","moving_time":3600}' | iicu events create
```

## Coaching Workflows

### Assess Current Fitness

1. `iicu athlete profile` — get FTP, weight, zones
2. `iicu wellness get --date today` — check today's readiness, sleep, HRV
3. `iicu activities list --oldest -7d` — review last week's training
4. `iicu curves power --type Ride` — check power curve trends

### Build a Training Block

1. Check current load: `iicu wellness get --oldest -7d`
2. Review recent training: `iicu activities list --oldest -14d`
3. Check power curves for current fitness: `iicu curves power --type Ride`
4. Create workouts in library: `echo '<workout-json>' | iicu workouts create`
5. Schedule on calendar: `echo '<events-json>' | iicu events bulk-create`

### Weekly Review

1. `iicu activities list --oldest -7d` — what was done
2. `iicu events list --oldest -7d --category WORKOUT` — what was planned
3. Compare planned vs actual load, compliance, and intensity
4. `iicu wellness get --oldest -7d` — track fatigue, sleep, HRV trends
5. Adjust upcoming week based on findings

### Recovery Check

1. `iicu wellness get --date today` — check fatigue, soreness, sleep, HRV
2. `iicu activities list --oldest -3d` — recent training load
3. If fatigued: reduce upcoming workout intensity or add rest day
4. Update calendar: `echo '<updated-event>' | iicu events update <id>`

## Key Metrics

- **CTL** (Chronic Training Load): Fitness — rolling ~42-day average of daily training load
- **ATL** (Acute Training Load): Fatigue — rolling ~7-day average
- **TSB** (Training Stress Balance): Form — CTL minus ATL. Positive = fresh, negative = fatigued
- **IF** (Intensity Factor): Ratio of normalized power to FTP
- **TSS** (Training Stress Score): Training load for a session
- **Compliance**: How closely an executed workout matched the plan (0-100%)
- **Decoupling**: Aerobic efficiency drift during a workout (lower = better)

## Event Categories

Use these values for the `category` field when creating events:

- `WORKOUT` — planned training session
- `RACE_A` / `RACE_B` / `RACE_C` — races by priority
- `NOTE` — calendar note
- `HOLIDAY` / `SICK` / `INJURED` — time off
- `TARGET` — performance target
