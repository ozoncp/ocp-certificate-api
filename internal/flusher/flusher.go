package flusher

import (
	"context"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	"github.com/ozoncp/ocp-certificate-api/internal/utils"
)

// Flusher - interface for add entity certificates for saving in repo
type Flusher interface {
	Flush(ctx context.Context, certificates []model.Certificate) []model.Certificate
}

// NewFlusher - creates a new instance of Flusher.
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

// Flush method to try to save the certificate entity in the repository if there was no error
func (f *flusher) Flush(ctx context.Context, certificates []model.Certificate) []model.Certificate {
	splitsCertificates := utils.SplitToBulks(certificates, f.chunkSize)

	for i, splitsCertificate := range splitsCertificates {
		if _, err := f.entityRepo.MultiCreateCertificates(ctx, splitsCertificate); err != nil {
			return certificates[i*f.chunkSize:]
		}
	}

	return nil
}
