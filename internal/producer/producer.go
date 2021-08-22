package producer

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	cfg "github.com/ozoncp/ocp-certificate-api/internal/config"
	"github.com/rs/zerolog/log"
	"time"
)

// Producer - an interface for send event messages
type Producer interface {
	Send(message Message)
	Init(ctx context.Context)
	Close()
}

type producer struct {
	prod    sarama.AsyncProducer
	topic   string
	message chan Message
	close   chan struct{}
}

// NewProducer - creates a new instance of Producer
func NewProducer() *producer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	asyncProducer, err := sarama.NewAsyncProducer(cfg.GetConfigInstance().Kafka.Brokers, config)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create Sarama new async producer")
	}

	return &producer{
		prod:    asyncProducer,
		topic:   cfg.GetConfigInstance().Kafka.Topic,
		message: make(chan Message),
		close:   make(chan struct{}),
	}
}

// Init - starting loop processing incoming event messages
func (p *producer) Init(ctx context.Context) {
	go func() {
		defer p.prod.Close()

		for {
			select {
			case message := <-p.message:
				msg, err := p.PrepareMessage(message)
				if err != nil {
					log.Err(err).Msg("failed marshaling message to json:")
					return
				}
				p.prod.Input() <- msg
			case <-ctx.Done():
				p.close <- struct{}{}
				return
			}
		}
	}()
}

func (p *producer) PrepareMessage(message Message) (*sarama.ProducerMessage, error) {
	bytes, err := json.Marshal(message)
	if err != nil {
		log.Err(err).Msg("failed marshaling message to json:")
		return &sarama.ProducerMessage{}, err
	}

	return &sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(p.topic),
		Value:     sarama.StringEncoder(bytes),
		Partition: -1,
		Timestamp: time.Time{},
	}, nil
}

// Send - send message into kafka
func (p *producer) Send(message Message) {
	p.message <- message
}

// Close - send signal closing for get messages
func (p *producer) Close() {
	<-p.close
}
