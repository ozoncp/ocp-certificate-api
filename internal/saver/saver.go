package saver

import (
	"github.com/ozoncp/ocp-certificate-api/internal/flusher"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"time"
)

type Saver interface {
	Save(certificate model.Certificate)
	Init()
	Close()
}

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

func (saver *saver) Save(certificateChannel model.Certificate) {
	saver.cert <- certificateChannel
}

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

func (saver *saver) Close() {
	saver.close <- struct{}{}
}
