# iicu - Intervals.icu CLI

CLI tools for intervals.icu, written in Go. Designed for AI agent accessibility.

## Project Structure

```
cmd/           # Cobra command definitions (one file per resource group)
api/           # HTTP client and Go types for intervals.icu API
config/        # Profile management and config file handling
skill/         # Claude skill for AI coaching workflows
docs/          # Documentation
main.go        # Entry point
```

## Key References

- **API Reference:** `docs/intervals-icu-api-reference.md` — comprehensive endpoint and model reference generated from the OpenAPI spec
- **intervals.icu Swagger:** https://intervals.icu/api/v1/docs/swagger-ui/index.html — live, authoritative API docs
- **MCP server (buggy, for reference only):** `../intervals-icu-mcp/` — Python MCP server covering 48 of ~110 endpoints

## Tech Stack

- **Language:** Go
- **CLI framework:** Cobra + Viper
- **Binary name:** `iicu`
- **Config file:** `~/.config/iicu/config.yaml`

## Conventions

- JSON output by default, `--human` flag for pretty-printed tables
- All dates use ISO-8601 format (YYYY-MM-DD or full datetime)
- Athlete ID "0" resolves to authenticated user in the API
- Auth: config file → env var → flag override layering

## Commands

Run `iicu --help` to see all command groups. Each group supports `--help` for subcommand details.

15 command groups: `athlete`, `activities`, `events`, `wellness`, `curves`, `workouts`, `training-plan`, `gear`, `sports`, `weather`, `routes`, `chat`, `custom-items`, `shared-events`, `config`.

## Development

```bash
make build      # Build for local platform
make test       # Run tests
make build-all  # Cross-compile all platforms
make install    # Install to $GOPATH/bin
```

## Testing

- Unit tests: `go test ./... -v`
- Integration tests require `IICU_API_KEY` and `IICU_ATHLETE_ID` env vars

## Skill

The Claude coaching skill is at `skill/intervals-icu-coaching.md`. It teaches Claude how to use `iicu` for coaching workflows.
