package producer

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	cfg "github.com/ozoncp/ocp-certificate-api/internal/config"
	"github.com/rs/zerolog/log"
	"time"
)

// Producer - an interface for send event messages
type Producer interface {
	Send(message Message) error
}

type producer struct {
	prod  sarama.SyncProducer
	topic string
}

// NewProducer - creates a new instance of Producer
func NewProducer() *producer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(cfg.GetConfigInstance().Kafka.Brokers, config)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create Sarama new sync producer")
	}

	return &producer{
		prod:  syncProducer,
		topic: cfg.GetConfigInstance().Kafka.Topic,
	}
}

func (p *producer) Send(message Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		log.Err(err).Msg("failed marshaling message to json:")
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(p.topic),
		Value:     sarama.StringEncoder(bytes),
		Partition: -1,
		Timestamp: time.Time{},
	}

	_, _, err = p.prod.SendMessage(msg)
	return err
}
