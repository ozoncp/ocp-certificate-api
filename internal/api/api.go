package api

import (
	"context"
	"errors"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	desc "github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

type api struct {
	desc.UnimplementedOcpCertificateApiServer
	repo repo.Repo
}

// NewOcpCertificateApi constructor
func NewOcpCertificateApi(repo repo.Repo) desc.OcpCertificateApiServer {
	return &api{
		repo: repo,
	}
}

// CreateCertificateV1 request for create new certificate
func (a *api) CreateCertificateV1(
	ctx context.Context,
	req *desc.CreateCertificateV1Request,
) (*desc.CreateCertificateV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid arguments")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	certificate := model.Certificate{
		UserId:  req.Certificate.UserId,
		Created: req.Certificate.Created.AsTime(),
		Link:    req.Certificate.Link,
	}

	certificateId, err := a.repo.CreateCertificate(ctx, certificate)

	if err != nil {
		log.Error().Err(err).Msg("error create certificate")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &desc.CreateCertificateV1Response{
		CertificateId: certificateId,
	}

	log.Info().Msg("creation of the certificate was successful")

	return response, nil
}

// DescribeCertificateV1 request for get single certificate
func (a *api) DescribeCertificateV1(
	ctx context.Context,
	req *desc.DescribeCertificateV1Request,
) (*desc.DescribeCertificateV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	certificate, err := a.repo.DescribeCertificate(ctx, req.CertificateId)
	if err != nil {
		if err == repo.ErrorCertificateNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("Failed to describe the data")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &desc.DescribeCertificateV1Response{
		Certificate: &desc.Certificate{
			Id:      certificate.Id,
			UserId:  certificate.UserId,
			Created: timestamppb.New(certificate.Created),
			Link:    certificate.Link,
		},
	}

	log.Info().Msg("reading of the certificate was successful")

	return response, nil
}

// ListCertificateV1 request for get list certificates
func (a *api) ListCertificateV1(
	ctx context.Context,
	req *desc.ListCertificateV1Request,
) (*desc.ListCertificateV1Response, error) {
	listCertificates, err := a.repo.ListCertificates(ctx, req.Limit, req.Offset)
	if err != nil {
		if errors.Is(err, repo.ErrorCertificateNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("failed get list certificates")
		return nil, status.Error(codes.Internal, "failed get list certificates")
	}

	log.Info().Msg("found count certificates: " + strconv.Itoa(len(listCertificates)))

	certificates := make([]*desc.Certificate, 0, len(listCertificates))
	for _, certificate := range listCertificates {
		cert := &desc.Certificate{
			Id:      certificate.Id,
			UserId:  certificate.UserId,
			Created: timestamppb.New(certificate.Created),
			Link:    certificate.Link,
		}

		certificates = append(certificates, cert)
	}

	response := &desc.ListCertificateV1Response{
		Certificates: certificates,
	}

	log.Info().Msg("reading of the certificates was successful")

	return response, nil
}

// UpdateCertificateV1 request for update certificate
func (a *api) UpdateCertificateV1(
	ctx context.Context,
	req *desc.UpdateCertificateV1Request,
) (*desc.UpdateCertificateV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("failed when try update certificate")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	certificate := model.Certificate{
		Id:      req.Certificate.Id,
		UserId:  req.Certificate.UserId,
		Created: req.Certificate.Created.AsTime(),
		Link:    req.Certificate.Link,
	}

	updated, err := a.repo.UpdateCertificate(ctx, certificate)

	if err != nil {
		if err == repo.ErrorCertificateNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("failed when try update certificate")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &desc.UpdateCertificateV1Response{
		Updated: updated,
	}

	log.Printf("update of the certificate was successful")

	return response, nil
}

// RemoveCertificateV1 request for remove certificate
func (a *api) RemoveCertificateV1(
	ctx context.Context,
	req *desc.RemoveCertificateV1Request,
) (*desc.RemoveCertificateV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	removed, err := a.repo.RemoveCertificate(ctx, req.CertificateId)

	if err != nil {
		if errors.Is(err, repo.ErrorCertificateNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("failed when try remove certificate")
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &desc.RemoveCertificateV1Response{
		Removed: removed,
	}

	log.Printf("removing of the certificate was successful")

	return response, nil
}
