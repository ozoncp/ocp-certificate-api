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
		capacity:           capacity,
		flusher:            flusher,
		ticker:             ticker,
		certificateModel:   make([]model.Certificate, 0, capacity),
		certificateChannel: make(chan model.Certificate),
		close:              make(chan struct{}),
		done:               make(chan struct{}),
	}
}

type saver struct {
	capacity           uint
	flusher            flusher.Flusher
	ticker             time.Ticker
	certificateModel   []model.Certificate
	certificateChannel chan model.Certificate
	close              chan struct{}
	done               chan struct{}
}

func (saver *saver) Save(certificateChannel model.Certificate) {
	saver.certificateChannel <- certificateChannel
}

func (saver *saver) Init() {
	go func() {
		defer saver.ticker.Stop()

		for {
			select {
			case cert := <-saver.certificateChannel:
				saver.certificateModel = append(saver.certificateModel, cert)
			case <-saver.ticker.C:
				saver.certificateModel = saver.flusher.Flush(saver.certificateModel)
			case <-saver.close:
				close(saver.close)
				_ = saver.flusher.Flush(saver.certificateModel)
				saver.done <- struct{}{}
				return
			}
		}
	}()
}

func (saver *saver) Close() {
	saver.close <- struct{}{}
	<-saver.done
}
