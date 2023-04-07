package input

import (
	"context"
	"logbeat/internal/config"
	"logbeat/internal/filter"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func InputKafka(input config.Input) {
	if input.CommitInterval == 0 {
		input.CommitInterval = 10
	}
	for _, topic := range input.Topic {
		c := kafka.ReaderConfig{
			Brokers:        input.Broker,
			Topic:          topic,
			GroupID:        input.GroupId,
			MinBytes:       10240,
			MaxBytes:       10240000,
			CommitInterval: time.Second * time.Duration(input.CommitInterval),
		}
		go InputKafkaWorker(c, input.Codec)
	}
}

func InputKafkaWorker(c kafka.ReaderConfig, codec string) {
	r := kafka.NewReader(c)
	fields := logrus.Fields{
		"broker":    c.Brokers,
		"topic":     c.Topic,
		"partition": r.Stats().Partition,
		"group_id":  c.GroupID,
	}
	logrus.WithFields(fields).Info("input kafka client")
	defer r.Close()
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			logrus.WithFields(fields).Error(err)
			time.Sleep(time.Second)
		}
		filter.FilterChan <- Codec(m.Value, codec)
	}
}
