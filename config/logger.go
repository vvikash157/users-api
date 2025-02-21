package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// InitializeLogger creates and returns a new logrus instance
func InitializeLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)                     // Log to console
	logger.SetFormatter(&logrus.JSONFormatter{})    // JSON formatting
	logger.SetLevel(logrus.DebugLevel)              // Set log level

	return logger
}
