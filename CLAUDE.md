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
