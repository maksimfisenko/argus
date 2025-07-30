package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init(logLevel, logFile string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Warnf("invalid log level '%s', defaulting to 'info'", logLevel)
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	if logFile != "stdout" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			logrus.Fatalf("cannot open log file: %v", err)
		}
		logrus.SetOutput(file)
		defer file.Close()
	}

	return nil
}
