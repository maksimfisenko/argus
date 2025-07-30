package main

import (
	"os"
	"time"

	"github.com/maksimfisenko/argus/internal/config"
	"github.com/maksimfisenko/argus/internal/metrics"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.Warnf("invalid log level '%s', defaulting to 'info'", cfg.LogLevel)
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	if cfg.LogFile != "stdout" {
		file, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			logrus.Fatalf("cannot open log file: %v", err)
		}
		logrus.SetOutput(file)
		defer file.Close()
	}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	logrus.Infof("agent started with interval %s", cfg.Interval)

	for {
		<-ticker.C

		cpu, ram, err := metrics.Collect()
		if err != nil {
			logrus.WithError(err).Error("failed to collect metrics")
		}
		logrus.Infof("CPU: %.2f%%, RAM: %.2f%%", cpu, ram)
	}
}
