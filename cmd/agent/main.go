package main

import (
	"time"

	"github.com/maksimfisenko/argus/internal/metrics"
	"github.com/sirupsen/logrus"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C

		cpu, ram, err := metrics.Collect()
		if err != nil {
			logrus.WithError(err).Error("failed to collect metrics")
		}
		logrus.Infof("CPU: %.2f%%, RAM: %.2f%%", cpu, ram)
	}
}
