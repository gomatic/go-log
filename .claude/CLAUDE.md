# go-log

CLI-agnostic structured-logging setup over `log/slog` (package `log`): `Level`, `Format`, `LoggerConfig`, and `NewLogger(w)`. Knows nothing about command-line frameworks. Generic — lives in `gomatic`, consumed by `template.cli` and the SkyKernel tools.

- Depends only on the stdlib (testify for tests). Must not import a CLI framework.
- Gate: gofumpt, vet, staticcheck, govulncheck, gocognit ≤ 7, 100% coverage.
