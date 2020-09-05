package main

import (
	"go.uber.org/zap"
	"os"
)

var logger *zap.Logger

func initLogging() {
	if os.Getenv("DEBUG") != "" {
		cfg := zap.NewDevelopmentConfig()
		cfg.Level.SetLevel(zap.DebugLevel)
		logger, _ = cfg.Build()
		logger.Info("Starting in debug mode")
	} else {
		logger, _ = zap.NewProduction()
	}
}
