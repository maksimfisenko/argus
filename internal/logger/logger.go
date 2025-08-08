package logger

import (
	"github.com/sirupsen/logrus"
)

func Init(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Warnf("invalid log level '%s', defaulting to 'info'", logLevel)
		level = logrus.InfoLevel
	}

	logrus.SetLevel(level)

	return nil
}
