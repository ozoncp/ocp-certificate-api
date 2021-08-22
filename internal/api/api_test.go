package api_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-certificate-api/internal/api"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"github.com/ozoncp/ocp-certificate-api/internal/producer"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	desc "github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var _ = Describe("Api", func() {
	const tableName = "certificate"
	const batchSize = 2

	now := time.Now()

	var (
		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock

		ctx          context.Context
		r            repo.Repo
		grpc         desc.OcpCertificateApiServer
		certificates []model.Certificate
		p            producer.Producer

		err error
	)

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())
		sqlxDB = sqlx.NewDb(db, "sqlmock")

		ctx = context.Background()
		r = repo.NewRepo(sqlxDB)
		p = producer.NewProducer()
		p.Init(ctx)
		grpc = api.NewOcpCertificateApi(r, p, batchSize)

		certificates = []model.Certificate{
			{1.0, 1.0, now, "http://link"},
			{2.0, 2.0, now, "http://link"},
			{3.0, 3.0, now, "http://link"},
			{4.0, 4.0, now, "http://link"},
		}
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		p.Close()
		Expect(err).Should(BeNil())
	})

	Context("Test MultiCreateCertificatesV1", func() {
		var req *desc.MultiCreateCertificatesV1Request

		BeforeEach(func() {
			multiCertificates := make([]*desc.Certificate, 0, len(certificates))
			for _, certificate := range certificates {
				multiCertificates = append(multiCertificates, &desc.Certificate{
					Id:      certificate.Id,
					UserId:  certificate.UserId,
					Created: timestamppb.New(certificate.Created),
					Link:    certificate.Link,
				})
			}

			req = &desc.MultiCreateCertificatesV1Request{
				Certificates: multiCertificates,
			}

			rows1 := sqlmock.NewRows([]string{"id"}).
				AddRow(1).
				AddRow(2)
			mock.ExpectQuery("INSERT INTO "+tableName).
				WithArgs(
					multiCertificates[0].UserId, multiCertificates[0].Created.AsTime(), multiCertificates[0].Link,
					multiCertificates[1].UserId, multiCertificates[1].Created.AsTime(), multiCertificates[1].Link).
				WillReturnRows(rows1)

			rows2 := sqlmock.NewRows([]string{"id"}).
				AddRow(3).
				AddRow(4)
			mock.ExpectQuery("INSERT INTO "+tableName).
				WithArgs(
					multiCertificates[2].UserId, multiCertificates[2].Created.AsTime(), multiCertificates[2].Link,
					multiCertificates[3].UserId, multiCertificates[3].Created.AsTime(), multiCertificates[3].Link).
				WillReturnRows(rows2)

		})

		It("Test create certificate", func() {
			grpc = api.NewOcpCertificateApi(r, p, batchSize)
			Expect(grpc).ShouldNot(BeNil())
			response, err := grpc.MultiCreateCertificatesV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(len(response.CertificateIds)).Should(BeEquivalentTo(len(certificates)))
		})
	})

	Context("Test CreateCertificateV1", func() {
		var req *desc.CreateCertificateV1Request

		BeforeEach(func() {
			req = &desc.CreateCertificateV1Request{
				Certificate: &desc.Certificate{
					Id:      certificates[0].Id,
					UserId:  certificates[0].UserId,
					Created: timestamppb.New(certificates[0].Created),
					Link:    certificates[0].Link,
				},
			}

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(1)
			mock.ExpectQuery("INSERT INTO "+tableName).
				WithArgs(req.Certificate.UserId, req.Certificate.Created.AsTime(), req.Certificate.Link).
				WillReturnRows(rows)

		})

		It("Test create certificate", func() {
			grpc = api.NewOcpCertificateApi(r, p, batchSize)
			Expect(grpc).ShouldNot(BeNil())
			response, err := grpc.CreateCertificateV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response.CertificateId).Should(BeEquivalentTo(1))
		})
	})

	Context("Test GetCertificateV1", func() {
		var req *desc.GetCertificateV1Request

		BeforeEach(func() {
			req = &desc.GetCertificateV1Request{
				CertificateId: certificates[1].Id,
			}

			rows := sqlmock.NewRows([]string{"id", "user_id", "created", "link"}).
				AddRow(
					certificates[1].Id,
					certificates[1].UserId,
					certificates[1].Created,
					certificates[1].Link)
			mock.ExpectQuery("SELECT id, user_id, created, link FROM " + tableName + " WHERE").
				WithArgs(req.CertificateId).
				WillReturnRows(rows)

		})

		It("Test Get certificate", func() {
			grpc = api.NewOcpCertificateApi(r, p, batchSize)
			Expect(grpc).ShouldNot(BeNil())
			Expect(err).Should(BeNil())

			response, err := grpc.GetCertificateV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response.Certificate.Id).Should(BeEquivalentTo(certificates[1].Id))
			Expect(response.Certificate.UserId).Should(BeEquivalentTo(certificates[1].UserId))
			Expect(response.Certificate.Created.AsTime().Unix()).Should(BeEquivalentTo(certificates[1].Created.Unix()))
			Expect(response.Certificate.Link).Should(BeEquivalentTo(certificates[1].Link))
		})
	})

	Context("Test UpdateCertificateV1Request", func() {
		var req *desc.UpdateCertificateV1Request

		BeforeEach(func() {
			req = &desc.UpdateCertificateV1Request{
				Certificate: &desc.Certificate{
					Id:      certificates[3].Id,
					UserId:  certificates[3].UserId,
					Created: timestamppb.New(certificates[3].Created),
					Link:    certificates[3].Link,
				},
			}

			mock.ExpectExec("UPDATE "+tableName+" SET").
				WithArgs(
					req.Certificate.UserId,
					req.Certificate.Created.AsTime(),
					req.Certificate.Link,
					req.Certificate.Id,
				).
				WillReturnResult(sqlmock.NewResult(1, 1))

		})

		It("Test update certificate", func() {
			grpc = api.NewOcpCertificateApi(r, p, batchSize)
			Expect(grpc).ShouldNot(BeNil())
			Expect(err).Should(BeNil())

			response, err := grpc.UpdateCertificateV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response.Updated).Should(BeEquivalentTo(true))
		})
	})

	Context("Test ListCertificateV1Request", func() {
		var req *desc.ListCertificateV1Request

		BeforeEach(func() {
			req = &desc.ListCertificateV1Request{
				Limit:  3,
				Offset: 0,
			}

			rows := sqlmock.NewRows([]string{"id", "user_id", "created", "link"}).
				AddRow(certificates[0].Id, certificates[0].UserId, certificates[0].Created, certificates[0].Link).
				AddRow(certificates[1].Id, certificates[1].UserId, certificates[1].Created, certificates[1].Link).
				AddRow(certificates[2].Id, certificates[2].UserId, certificates[2].Created, certificates[2].Link).
				AddRow(certificates[3].Id, certificates[3].UserId, certificates[3].Created, certificates[3].Link)
			mock.ExpectQuery(
				"SELECT id, user_id, created, link FROM " + tableName + " LIMIT 3 OFFSET 0").
				WillReturnRows(rows)

		})

		It("Test get list certificates", func() {
			grpc = api.NewOcpCertificateApi(r, p, batchSize)
			Expect(grpc).ShouldNot(BeNil())
			Expect(err).Should(BeNil())

			response, err := grpc.ListCertificateV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response.Certificates[0].Id).Should(BeEquivalentTo(certificates[0].Id))
			Expect(response.Certificates[1].Id).Should(BeEquivalentTo(certificates[1].Id))
			Expect(response.Certificates[2].Id).Should(BeEquivalentTo(certificates[2].Id))
			Expect(response.Certificates[3].Id).Should(BeEquivalentTo(certificates[3].Id))
			Expect(response.Certificates[0].Link).Should(BeEquivalentTo(certificates[0].Link))
			Expect(response.Certificates[1].Link).Should(BeEquivalentTo(certificates[1].Link))
			Expect(response.Certificates[2].Link).Should(BeEquivalentTo(certificates[2].Link))
			Expect(response.Certificates[3].Link).Should(BeEquivalentTo(certificates[3].Link))
		})
	})

	Context("Test RemoveCertificateV1Request", func() {
		var req *desc.RemoveCertificateV1Request

		BeforeEach(func() {
			req = &desc.RemoveCertificateV1Request{
				CertificateId: certificates[1].Id,
			}

			mock.ExpectExec("DELETE FROM " + tableName).
				WillReturnResult(sqlmock.NewResult(1, 1))

		})

		It("Test remove certificate", func() {
			grpc = api.NewOcpCertificateApi(r, p, batchSize)
			Expect(grpc).ShouldNot(BeNil())
			Expect(err).Should(BeNil())

			response, err := grpc.RemoveCertificateV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response.Removed).Should(BeEquivalentTo(true))
		})
	})
})
