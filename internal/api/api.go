package api

import (
	"context"
	desc "github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api"
	log "github.com/rs/zerolog/log"
)

type api struct {
	desc.UnimplementedOcpCertificateApiServer
}

func NewOcpCertificateApi() desc.OcpCertificateApiServer {
	return &api{}
}

func (a *api) CreateCertificateV1(
	ctx context.Context,
	req *desc.CreateCertificateV1Request,
) (*desc.CreateCertificateV1Response, error) {
	log.Printf("Сreation of the certificate was successful")
	return &desc.CreateCertificateV1Response{}, nil
}

func (a *api) DescribeCertificateV1(
	ctx context.Context,
	req *desc.DescribeCertificateV1Request,
) (*desc.DescribeCertificateV1Response, error) {
	log.Printf("Reading of the certificate was successful")
	return &desc.DescribeCertificateV1Response{}, nil
}

func (a *api) ListCertificateV1(
	ctx context.Context,
	req *desc.ListCertificateV1Request,
) (*desc.ListCertificateV1Response, error) {
	log.Printf("Reading of the certificates was successful")
	return &desc.ListCertificateV1Response{}, nil
}

func (a *api) UpdateCertificateV1(
	ctx context.Context,
	req *desc.UpdateCertificateV1Request,
) (*desc.UpdateCertificateV1Response, error) {
	log.Printf("Update of the certificate was successful")
	return &desc.UpdateCertificateV1Response{}, nil
}

func (a *api) RemoveCertificateV1(
	ctx context.Context,
	req *desc.RemoveCertificateV1Request,
) (*desc.RemoveCertificateV1Response, error) {
	log.Printf("Removing of the certificate was successful")
	return &desc.RemoveCertificateV1Response{}, nil
}
