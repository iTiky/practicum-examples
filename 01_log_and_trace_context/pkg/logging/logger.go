package logging

import (
	"os"

	"github.com/rs/zerolog"
)

// LoggerOption defines logger customization option.
type LoggerOption func(logger zerolog.Logger) zerolog.Logger

// WithLogLevel sets log level.
func WithLogLevel(level zerolog.Level) LoggerOption {
	return func(logger zerolog.Logger) zerolog.Logger {
		return logger.Level(level)
	}
}

// NewLogger creates a new customizable logger.
func NewLogger(opts ...LoggerOption) zerolog.Logger {
	logger := zerolog.New(os.Stdout).
		Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Logger()

	for _, opt := range opts {
		logger = opt(logger)
	}

	return logger
}
