package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// InitializeLogger creates and returns a new logrus instance
func InitializeLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetLevel(logrus.ErrorLevel)

	return logger
}
