# go-log

A thin [`log/slog`](https://pkg.go.dev/log/slog)-based configuration layer: textual `Level` and `Format` value types bindable from flags, producing a `*slog.Logger`.

## Install

```sh
go get github.com/gomatic/go-log
```

## Usage

```go
package main

import (
	"os"

	log "github.com/gomatic/go-log"
)

func main() {
	logger := log.LoggerConfig{Level: "info", Format: log.FormatJSON}.NewLogger(os.Stderr)
	logger.Info("ready", "addr", ":8080")
}
```

`Level` accepts `debug`, `info`, `warn`, or `error` (defaulting to `info` when empty or invalid); `Format` accepts `FormatText` or `FormatJSON` (defaulting to text when unknown). The package knows nothing about command-line frameworks — binding these types to flags lives in the consumer.
