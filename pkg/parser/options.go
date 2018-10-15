package parser

import (
	"go.uber.org/zap"
)

// Options is a configuration for parser container.
type Options struct {
	Logger *zap.Logger
}

// DefaultLogger returns basic logger instance.
func DefaultLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
