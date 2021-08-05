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
		certificateChannel: make(chan model.Certificate),
		close:              make(chan struct{}),
		done:               make(chan struct{}),
	}
}

type saver struct {
	capacity           uint
	flusher            flusher.Flusher
	ticker             time.Ticker
	certificateChannel chan model.Certificate
	close              chan struct{}
	done               chan struct{}
}

func (saver *saver) Save(certificateChannel model.Certificate) {
	saver.certificateChannel <- certificateChannel
}

func (saver *saver) Init() {
	var certificate []model.Certificate

	go func() {
		defer saver.ticker.Stop()

		for {
			select {
			case cert := <-saver.certificateChannel:
				certificate = append(certificate, cert)
			case <-saver.ticker.C:
				certificate = (saver.flusher).Flush(certificate)
			case <-saver.close:
				close(saver.close)
				_ = (saver.flusher).Flush(certificate)
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
