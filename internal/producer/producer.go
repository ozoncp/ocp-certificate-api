package producer

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"time"
)

// Producer - an interface for send event messages
type Producer interface {
	Send(topic string, value interface{}) error
	Close()
}

type producer struct {
	prod sarama.SyncProducer
}

// NewProducer - creates a new instance of Producer
func NewProducer(syncProducer sarama.SyncProducer) *producer {
	return &producer{
		prod: syncProducer,
	}
}

func (p *producer) Send(topic string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		log.Err(err).Msg("failed marshaling message to json:")
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(topic),
		Value:     sarama.StringEncoder(bytes),
		Partition: -1,
		Timestamp: time.Time{},
	}

	_, _, err = p.prod.SendMessage(msg)
	return err
}

// Close - close connection
func (p *producer) Close() {
	p.prod.Close()
}
