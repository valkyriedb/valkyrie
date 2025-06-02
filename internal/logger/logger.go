package logger

import (
	"io"
	"log/slog"
)

func New(w io.Writer, env string) *slog.Logger {
	var h slog.Handler
	switch env {
	case "dev":
		h = slog.NewTextHandler(w, &slog.HandlerOptions{Level: slog.LevelDebug})
	case "prod":
		h = slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo})
	default:
		h = slog.NewTextHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo})
	}
	return slog.New(h)
}

func Err(err error) slog.Attr {
	return slog.String("error", err.Error())
}
