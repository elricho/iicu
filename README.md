# iicu — intervals.icu CLI

Command-line tools for the [intervals.icu](https://intervals.icu) training platform, written in Go and designed to be accessible to AI agents as well as humans.

JSON output by default; pass `--human` for pretty-printed tables.

## Install

### macOS / Linux

**Homebrew:**

```sh
brew install elricho/iicu/iicu
```

**Install script:**

```sh
curl -fsSL https://raw.githubusercontent.com/elricho/iicu/main/install.sh | sh
```

### Windows

**Scoop:**

```powershell
scoop bucket add iicu https://github.com/elricho/scoop-iicu
scoop install iicu
```

**PowerShell script:**

```powershell
irm https://raw.githubusercontent.com/elricho/iicu/main/install.ps1 | iex
```

### Other options

- **Go:** `go install github.com/elricho/iicu@latest`
- **Manual:** download a binary from the [latest release](https://github.com/elricho/iicu/releases/latest) and put it on your `PATH`.
- **From source:** `git clone` this repo, then `make install`.

The install scripts honour `IICU_VERSION` (pin a specific tag) and `IICU_INSTALL_DIR` (change the install location).

## Setup

Get your API key from intervals.icu under **Settings → Developer**, then run:

```sh
iicu config init
```

This creates `~/.config/iicu/config.yaml`. Credentials resolve in this order: config file → environment variables (`IICU_API_KEY`, `IICU_ATHLETE_ID`) → command-line flags (`--api-key`, `--athlete-id`), with later sources overriding earlier ones.

Multiple athletes are supported via profiles:

```sh
iicu config add coach        # add a named profile
iicu config use coach        # switch the default profile
iicu config profiles         # list configured profiles
iicu --profile coach ...     # use a profile for one command
```

Athlete ID `0` always resolves to the authenticated user.

## Usage

```sh
iicu athlete                          # your profile and fitness data
iicu activities list                  # recent activities (JSON)
iicu activities list --human          # ...as a table
iicu wellness --athlete-id 0           # wellness / health data
iicu events list                      # calendar events and planned workouts
iicu curves --help                    # performance curves (power, HR, pace)
```

15 command groups are available: `athlete`, `activities`, `events`, `wellness`,
`curves`, `workouts`, `training-plan`, `gear`, `sports`, `weather`, `routes`,
`chat`, `custom-items`, `shared-events`, `config`.

Run `iicu --help` for the full list, or `iicu <group> --help` for a group's
subcommands.

## AI coaching skill

`skill/intervals-icu-coaching.md` is an example Claude skill that teaches an agent how to
use `iicu` for coaching workflows.

## Development

```sh
make build      # build for the local platform
make test       # run tests
make build-all  # cross-compile all platforms
make install    # install to $GOPATH/bin
```

Integration tests require `IICU_API_KEY` and `IICU_ATHLETE_ID` to be set.

Releases are cut by GoReleaser when a `v*` tag is pushed — see `.goreleaser.yaml`
and `.github/workflows/release.yml`.

## License

MIT
