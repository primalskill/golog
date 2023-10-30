package log

import (
	"log/slog"
	"os"
)

func NewDevelopment() *slog.Logger {
	return slog.New(newDevHandler())
}

func NewDiscard() *slog.Logger {
	return slog.New(discardHandler{})
}

func NewProduction() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			AddSource:   true,
			Level:       slog.LevelInfo,
			ReplaceAttr: prodReplacer,
		}),
	)
}
