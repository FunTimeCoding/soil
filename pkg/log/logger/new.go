package logger

import (
	"context"
	"log"
	"log/slog"
	"os"
)

func New(c context.Context) *Logger {
	return &Logger{
		context: c,
		structured: slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		),
		plain: log.New(
			os.Stderr,
			"PLAIN: ",
			log.Ldate|log.Ltime|log.Lshortfile,
		),
	}
}
