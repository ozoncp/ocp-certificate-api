package api

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	cfg "github.com/ozoncp/ocp-certificate-api/internal/config"
	"github.com/ozoncp/ocp-certificate-api/internal/metrics"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"github.com/ozoncp/ocp-certificate-api/internal/producer"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	"github.com/ozoncp/ocp-certificate-api/internal/utils"
	desc "github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"time"
)

type api struct {
	desc.UnimplementedOcpCertificateApiServer
	repo repo.Repo
	prod producer.Producer
}

// NewOcpCertificateApi constructor
func NewOcpCertificateApi(repo repo.Repo, prod producer.Producer) desc.OcpCertificateApiServer {
	return &api{
		repo: repo,
		prod: prod,
	}
}

// MultiCreateCertificatesV1 request for create new certificates
func (a *api) MultiCreateCertificatesV1(
	ctx context.Context,
	req *desc.MultiCreateCertificatesV1Request,
) (*desc.MultiCreateCertificatesV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Invalid arguments was received when creating a multi certificates")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateCertificatesV1")
	defer span.Finish()

	var certificates []model.Certificate
	for _, certificate := range req.Certificates {
		certificates = append(certificates, model.Certificate{
			Id:      certificate.Id,
			UserId:  certificate.UserId,
			Created: certificate.Created.AsTime(),
			Link:    certificate.Link,
		})
	}

	batchSize := cfg.GetConfigInstance().BatchSize
	certBulks := utils.SplitToBulks(certificates, batchSize)
	response := &desc.MultiCreateCertificatesV1Response{}
	for i := 0; i < len(certBulks); i++ {
		certIds, err := a.repo.MultiCreateCertificates(ctx, certBulks[i])
		if err != nil {
			childSpan := tracer.StartSpan("Size of data 0 bytes", opentracing.ChildOf(span.Context()))
			childSpan.Finish()
			log.Error().Err(err).Msgf("error when try multi create certificates with butchSize %d", batchSize)
			return response, status.Error(codes.Internal, err.Error())
		}

		childSpan := tracer.StartSpan(
			fmt.Sprintf("Size of data %d bytes",
				uintptr(len(certBulks[i]))*reflect.TypeOf(certBulks[i]).Elem().Size()),
			opentracing.ChildOf(span.Context()),
		)
		childSpan.Finish()

		response.CertificateIds = append(response.CertificateIds, certIds...)
	}

	log.Info().Msg("multi creating of the certificates was successful")

	return response, nil
}

// CreateCertificateV1 request for create new certificate
func (a *api) CreateCertificateV1(
	ctx context.Context,
	req *desc.CreateCertificateV1Request,
) (*desc.CreateCertificateV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Invalid arguments was received when creating a certificate")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("CreateCertificateV1")
	defer span.Finish()

	certificate := &model.Certificate{
		UserId:  req.Certificate.UserId,
		Created: req.Certificate.Created.AsTime(),
		Link:    req.Certificate.Link,
	}

	err := a.repo.CreateCertificate(ctx, certificate)

	if err != nil {
		log.Error().Err(err).Msg("error create certificate")
		return nil, status.Error(codes.Internal, err.Error())
	}

	span.SetTag("id", certificate.Id)
	metrics.CreateCounterInc()
	a.prod.Send(producer.CreateMessage(producer.Create, certificate.Id, time.Now()))
	response := &desc.CreateCertificateV1Response{
		CertificateId: certificate.Id,
	}

	log.Info().Msg("creation of the certificate was successful")

	return response, nil
}

// GetCertificateV1 request for get single certificate
func (a *api) GetCertificateV1(
	ctx context.Context,
	req *desc.GetCertificateV1Request,
) (*desc.GetCertificateV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Invalid arguments was received when getting a certificate")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("GetCertificateV1Request")
	defer span.Finish()

	certificate, err := a.repo.GetCertificate(ctx, req.CertificateId)
	if err != nil {
		log.Error().Err(err).Msg("error get the certificate")
		return nil, status.Error(codes.Internal, err.Error())
	}

	span.SetTag("id", certificate.Id)
	response := &desc.GetCertificateV1Response{
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
		log.Error().Err(err).Msg("failed get list certificates")
		return nil, status.Error(codes.Internal, "failed get list certificates")
	}

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("ListCertificateV1Request")
	defer span.Finish()

	log.Info().Msgf("found count certificates: %d", len(listCertificates))

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

	childSpan := tracer.StartSpan(
		fmt.Sprintf("Found count cerificates: %d", len(listCertificates)),
		opentracing.ChildOf(span.Context()),
	)
	defer childSpan.Finish()

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
		log.Error().Err(err).Msg("Invalid arguments was received when updating a certificate")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("UpdateCertificateV1Request")
	defer span.Finish()

	certificate := model.Certificate{
		Id:      req.Certificate.Id,
		UserId:  req.Certificate.UserId,
		Created: req.Certificate.Created.AsTime(),
		Link:    req.Certificate.Link,
	}

	updated, err := a.repo.UpdateCertificate(ctx, certificate)

	if err != nil {
		log.Error().Err(err).Msg("failed when try update certificate")
		return nil, status.Error(codes.Internal, err.Error())
	}

	span.SetTag("id", certificate.Id)
	metrics.UpdateCounterInc()
	a.prod.Send(producer.CreateMessage(producer.Update, certificate.Id, time.Now()))
	response := &desc.UpdateCertificateV1Response{
		Updated: updated,
	}

	log.Info().Msg("update of the certificate was successful")

	return response, nil
}

// RemoveCertificateV1 request for remove certificate
func (a *api) RemoveCertificateV1(
	ctx context.Context,
	req *desc.RemoveCertificateV1Request,
) (*desc.RemoveCertificateV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Invalid arguments was received when removing a certificate")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("RemoveCertificateV1Request")
	defer span.Finish()

	removed, err := a.repo.RemoveCertificate(ctx, req.CertificateId)

	if err != nil {
		log.Error().Err(err).Msg("failed when try remove certificate")
		return nil, status.Error(codes.Internal, err.Error())
	}

	span.SetTag("id", req.CertificateId)
	metrics.RemoveCounterInc()
	a.prod.Send(producer.CreateMessage(producer.Remove, req.CertificateId, time.Now()))

	response := &desc.RemoveCertificateV1Response{
		Removed: removed,
	}

	log.Info().Msg("removing of the certificate was successful")

	return response, nil
}
