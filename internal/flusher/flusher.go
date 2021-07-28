package flusher

import (
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	"github.com/ozoncp/ocp-certificate-api/internal/utils"
)

type Flusher interface {
	Flush(certificates []model.Certificate) []model.Certificate
}

func NewFlusher(
	chunkSize int,
	entityRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

type flusher struct {
	chunkSize  int
	entityRepo repo.Repo
}

func (f flusher) Flush(certificates []model.Certificate) []model.Certificate {
	splitsCertificates := utils.SplitToBulks(certificates, f.chunkSize)

	for i, splitsCertificate := range splitsCertificates {
		if err := f.entityRepo.AddCertificates(splitsCertificate); err != nil {
			return splitsCertificate[i:]
		}
	}

	return nil
}
