package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// WithLogger can be used to compose a struct with a logger.
//
// Example:
//
//	type MyStruct struct {
//		*utils.WithLogger
//	}
type WithLogger struct {
	Logger *zerolog.Logger
}

func NewWithLogger(logger *zerolog.Logger) *WithLogger {
	return &WithLogger{Logger: logger}
}

func NewTestWithLogger() *WithLogger {
	logger := NewLogger(&DefaultConfig{IsDebug: true})
	return &WithLogger{Logger: logger}
}

// ProvideLogger returns a new logger instance for dependency injection
func ProvideLogger(config *DefaultConfig) *zerolog.Logger {
	return NewLogger(config)
}

// NewLogger returns a standard zerolog logger.
func NewLogger(config *DefaultConfig) *zerolog.Logger {
	var logger zerolog.Logger

	// Configure zerolog based on environment
	if config.IsDebug {
		// Development mode: pretty console output with colors and readable format
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		logger = zerolog.New(output).With().Timestamp().Caller().Logger()
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		// Production mode: JSON output for structured logging
		logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return &logger
}
