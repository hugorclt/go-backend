package logger

import (
	"log/slog"
	"time"

	"github.com/go-chi/httplog/v2"
)

func GetLogger() *httplog.Logger {
	logger := httplog.NewLogger("backend", httplog.Options{
		//JSON:             true,
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})

	return logger
}
