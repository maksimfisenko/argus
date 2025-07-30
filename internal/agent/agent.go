package agent

import (
	"context"
	"time"

	"github.com/maksimfisenko/argus/internal/config"
	"github.com/maksimfisenko/argus/internal/metrics"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config) error {
	sender, err := NewSender(cfg.ServerAddress)
	if err != nil {
		return err
	}
	defer sender.Close()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	logrus.Infof("agent started with interval %s, sending metrics to '%s'", cfg.Interval, cfg.ServerAddress)

	for {
		<-ticker.C

		snap, err := metrics.Collect()
		if err != nil {
			logrus.WithError(err).Error("failed to collect metrics")
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err = sender.SendSnaphot(ctx, snap)
		cancel()

		if err != nil {
			logrus.WithError(err).Error("failed to send snapshot")
			continue
		}
	}
}
