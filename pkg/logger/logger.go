package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// NewLogger initializes and returns a SugaredLogger instance.
func NewLogger() *zap.SugaredLogger {
	l, err := zap.NewDevelopment() // Use zap.NewDevelopment() for local development
	if err != nil {
		panic(err)
	}
	return l.Sugar()
}

// Fx Module for Logger
var Module = fx.Module("logger", fx.Provide(NewLogger))
