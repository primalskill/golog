package golog

import (
	"context"
	"log/slog"
)

// discardHandler implements a slog.Handler where all data is discarded. Mostly used for testing code where log lines
// shouldn't be emitted.
type discardHandler struct{}

func (p discardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

func (p discardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (p discardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return p
}

func (p discardHandler) WithGroup(_ string) slog.Handler {
	return p
}
