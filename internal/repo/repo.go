package repo

import "github.com/ozoncp/ocp-certificate-api/internal/model"

type Repo interface {
	AddCertificates(certificates []model.Certificate) error
	ListCertificates(limit, offset uint64) ([]model.Certificate, error)
	DescribeCertificate(certificateId uint64) (*model.Certificate, error)
}
