package saver

import (
	"github.com/ozoncp/ocp-certificate-api/internal/flusher"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"time"
)

// Saver - an interface for saving certificate entities.
type Saver interface {
	Save(certificate model.Certificate)
	Init()
	Close()
}

// NewSaver - creates a new instance of Saver.
func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
	ticker time.Ticker,
) Saver {
	return &saver{
		flusher:          flusher,
		ticker:           ticker,
		certificateModel: make([]model.Certificate, 0, capacity),
		cert:             make(chan model.Certificate),
		close:            make(chan struct{}),
	}
}

type saver struct {
	flusher          flusher.Flusher
	ticker           time.Ticker
	certificateModel []model.Certificate
	cert             chan model.Certificate
	close            chan struct{}
}

// Save - saving certificates entities into the repo
func (saver *saver) Save(certificateChannel model.Certificate) {
	saver.cert <- certificateChannel
}

// Init - starting loop processing incoming events
func (saver *saver) Init() {
	go func() {
		defer saver.ticker.Stop()

		for {
			select {
			case cert := <-saver.cert:
				saver.certificateModel = append(saver.certificateModel, cert)
			case <-saver.ticker.C:
				saver.certificateModel = saver.flusher.Flush(saver.certificateModel)
			case <-saver.close:
				close(saver.cert)
				close(saver.close)
				return
			}
		}
	}()
}

// Close - send signal closing the saver
func (saver *saver) Close() {
	saver.close <- struct{}{}
}
