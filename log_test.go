package log_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	xlog "github.com/gomatic/go-log"
)

func TestNewLoggerFormats(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		format xlog.Format
		isJSON bool
	}{
		{"text", xlog.FormatText, false},
		{"json", xlog.FormatJSON, true},
		{"unknown defaults to text", "xml", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := assert.New(t)
			var buf bytes.Buffer
			cfg := xlog.LoggerConfig{LogLevel: "info", LogFormat: tt.format}
			cfg.NewLogger(&buf).Info("hello", "k", "v")
			line := buf.String()
			want.Contains(line, "hello")
			if tt.isJSON {
				want.True(json.Valid(bytes.TrimSpace(buf.Bytes())), "expected JSON output")
			} else {
				want.False(json.Valid(bytes.TrimSpace(buf.Bytes())), "expected text output")
			}
		})
	}
}

func TestNewLoggerLevels(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		level     xlog.Level
		wantDebug bool
	}{
		{"debug enables debug", "debug", true},
		{"warn suppresses info", "warn", false},
		{"empty defaults to info", "", false},
		{"invalid defaults to info", "loud", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := assert.New(t)
			var buf bytes.Buffer
			cfg := xlog.LoggerConfig{LogLevel: tt.level, LogFormat: xlog.FormatText}
			log := cfg.NewLogger(&buf)
			log.Debug("dbg")
			log.Info("inf")
			out := buf.String()
			want.Equal(tt.wantDebug, strings.Contains(out, "dbg"))
			// info is emitted at info/debug/warn-default but suppressed only above warn
			want.Equal(tt.level != "warn", strings.Contains(out, "inf"))
		})
	}
}

func TestNewLoggerReturnsLogger(t *testing.T) {
	t.Parallel()
	cfg := xlog.LoggerConfig{LogLevel: "info", LogFormat: xlog.FormatText}
	assert.New(t).IsType(&slog.Logger{}, cfg.NewLogger(&bytes.Buffer{}))
}
