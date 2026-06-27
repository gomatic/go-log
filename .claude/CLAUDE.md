# go-log

CLI-agnostic structured-logging setup over `log/slog` (package `slogx`): `LogLevel`, `LogFormat`, `LoggerConfig`, and `NewLogger(w)`. It builds a `*slog.Logger` over any writer and knows nothing about command-line frameworks — binding these types to flags lives in a consumer (`go-app`).

- Depends only on the stdlib (testify for tests). Must not import urfave/cli or any CLI glue.
- Quality gate: gofumpt, `go vet`, staticcheck, govulncheck, gocognit ≤ 7, **100% coverage**.
