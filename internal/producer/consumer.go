package producer

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/opentracing/opentracing-go"
	cfg "github.com/ozoncp/ocp-certificate-api/internal/config"
	"github.com/ozoncp/ocp-certificate-api/internal/metrics"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	"github.com/rs/zerolog/log"
)

type Consumer interface {
	Subscribe(topic string, actionType ActionType) error
	Close()
}

// NewConsumer - creates a new instance of Consumer
// Batch asynchronous entry in the database. When you pull the addition of in the `repo`,
// there is an entry in KAFKA, and the consumer reads the packages and writes to the database.
func NewConsumer(r repo.Repo, metr metrics.Metrics) *consumer {
	return &consumer{
		r:    r,
		metr: metr,
	}
}

type consumer struct {
	r    repo.Repo
	scar sarama.Consumer
	metr metrics.Metrics
}

// Subscribe - subscribe topic
func (c *consumer) Subscribe(topic string, actionType ActionType) error {
	cons, err := sarama.NewConsumer(cfg.GetConfigInstance().Kafka.Brokers, nil)
	if err != nil {
		return err
	}
	c.scar = cons

	partitions, err := c.scar.Partitions(topic)
	if err != nil {
		return err
	}

	offsetOldest := sarama.OffsetOldest
	for _, partition := range partitions {
		pc, err := c.scar.ConsumePartition(topic, partition, offsetOldest)
		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				c.handleMessage(message, actionType)
			}
		}(pc)
	}

	return nil
}

// Close - close connection
func (c *consumer) Close() {
	c.scar.Close()
}

// handleMessage - handle message from topic
func (c *consumer) handleMessage(message *sarama.ConsumerMessage, actionType ActionType) {
	log.Info().Msgf("Receive message: %v", string(message.Value))
	c.metr.MultiCreateCounterInc()
	tcr := opentracing.GlobalTracer()

	var messages Messages
	if err := json.Unmarshal(message.Value, &messages); err != nil {
		log.Error().Err(err).Msg("Failed parse messages")
		return
	}

	switch actionType {
	case MultiCreate:
		span := tcr.StartSpan("MultiCreateCertificatesV1")
		defer span.Finish()

		_, err := c.r.MultiCreateCertificates(context.Background(), messages.Body)
		if err != nil {
			childSpan := tcr.StartSpan("Size of data 0 bytes", opentracing.ChildOf(span.Context()))
			childSpan.Finish()
			log.Error().Err(err).Msgf("error when save multi create certificates in repo")
			return
		}
	}

	log.Info().Msg("Consumer success")
}
