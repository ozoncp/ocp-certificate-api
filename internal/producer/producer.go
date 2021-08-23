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
	Send(message Message)
}

type producer struct {
	prod    sarama.SyncProducer
	topic   string
	message chan Message
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
		prod:    syncProducer,
		topic:   cfg.GetConfigInstance().Kafka.Topic,
		message: make(chan Message),
	}
}

func (p *producer) Send(message Message) {
	bytes, err := json.Marshal(message)
	if err != nil {
		log.Err(err).Msg("failed marshaling message to json:")
		return
	}

	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(p.topic),
		Value:     sarama.StringEncoder(bytes),
		Partition: -1,
		Timestamp: time.Time{},
	}

	_, _, err = p.prod.SendMessage(msg)
	if err != nil {
		log.Error().Msgf("failed send message kafka: %v", err)
	}
}
