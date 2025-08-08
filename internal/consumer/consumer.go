package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/maksimfisenko/argus/internal/config"
	"github.com/maksimfisenko/argus/internal/db"
	"github.com/maksimfisenko/argus/internal/kafka"
	"github.com/maksimfisenko/argus/internal/web"
	"github.com/maksimfisenko/argus/proto"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Consumer) error {
	database, err := db.Open(cfg.DbPath)
	if err != nil {
		return errors.New("failed to open database")
	}
	defer database.Close()

	if err := kafka.Ping(cfg); err != nil {
		return errors.New("failed to connect to kafka")
	}

	go web.Start(database)

	r := kafka.NewReader(cfg.KafkaBrokers, cfg.KafkaTopic, cfg.KafkaGroupID)

	logrus.Infof("The consumer is up [kafka-brokers='%v', kafka-topic='%s', kafka-group-id='%s', dp-path='%s']", cfg.KafkaBrokers, cfg.KafkaTopic, cfg.KafkaGroupID, cfg.DbPath)

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			logrus.Error("Failed to read message from Kafka")
			continue
		}

		var snap proto.Snapshot
		if err := json.Unmarshal(msg.Value, &snap); err != nil {
			logrus.Error("Failed to unmarshal a snapshot")
			continue
		}

		err = db.Insert(database, snap.AgentId, snap.Cpu, snap.Memory, snap.DiskUsage, snap.AvgLoad, snap.Uptime)
		if err != nil {
			logrus.Error("Failed to insert record to db")
		}

		logrus.Infof("Successfully inserted record to db [agent-id='%s']", snap.AgentId)
	}
}
