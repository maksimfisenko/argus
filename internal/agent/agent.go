package agent

import (
	"context"
	"errors"
	"time"

	"github.com/maksimfisenko/argus/internal/config"
	"github.com/maksimfisenko/argus/internal/metrics"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Agent) error {
	sender, err := NewSender(cfg.ServerAddress, cfg.ID)
	if err != nil {
		return errors.New("failed to set up new sender")
	}
	defer sender.Close()

	ticker := time.NewTicker(time.Duration(cfg.Interval) * time.Second)
	defer ticker.Stop()

	logrus.Infof("The agent is up [id='%s', server='%s', interval='%d seconds']", cfg.ID, cfg.ServerAddress, cfg.Interval)

	for {
		<-ticker.C

		snap, err := metrics.Collect()
		if err != nil {
			logrus.Error("Failed to collect metrics")
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err = sender.SendSnaphot(ctx, snap)
		cancel()

		if err != nil {
			logrus.Error("Failed to send snapshot")
			continue
		}

		logrus.Info("Successfully sent snapshot")
	}
}
