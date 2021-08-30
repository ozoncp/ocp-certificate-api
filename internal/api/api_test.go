package api_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Shopify/sarama/mocks"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-certificate-api/internal/api"
	"github.com/ozoncp/ocp-certificate-api/internal/broker"
	mockMetr "github.com/ozoncp/ocp-certificate-api/internal/mocks"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	desc "github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var _ = Describe("Api", func() {
	const tableName = "certificate"

	now := time.Now()

	var (
		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock

		controller   *gomock.Controller
		m            *mockMetr.MockMetrics
		ctx          context.Context
		r            repo.Repo
		grpc         desc.OcpCertificateApiServer
		certificates []model.Certificate
		synProdMock  *mocks.SyncProducer

		p broker.Producer
		s broker.Consumer
	)

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())
		sqlxDB = sqlx.NewDb(db, "sqlmock")

		ctx = context.Background()
		r = repo.NewRepo(sqlxDB)

		controller = gomock.NewController(GinkgoT())
		m = mockMetr.NewMockMetrics(controller)
		synProdMock = mocks.NewSyncProducer(controller.T, nil)

		p = broker.NewProducer(synProdMock)
		s = broker.NewConsumer(r, m)
		grpc = api.NewOcpCertificateAPI(r, m, p, s)
		link := "https://link.ru"

		certificates = []model.Certificate{
			{1.0, 1.0, now, link, false},
			{2.0, 2.0, now, link, false},
			{3.0, 3.0, now, link, false},
			{4.0, 4.0, now, link, false},
		}
	})

	AfterEach(func() {
		var err error
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	Context("Test MultiCreateCertificatesV1", func() {
		var req *desc.MultiCreateCertificatesV1Request
		var err error

		BeforeEach(func() {
			synProdMock.ExpectSendMessageAndSucceed()
			m.EXPECT().MultiCreateCounterInc()
			multiCertificates := make([]*desc.NewCertificate, 0, len(certificates))
			for _, certificate := range certificates {
				multiCertificates = append(multiCertificates, &desc.NewCertificate{
					UserId:    certificate.UserID,
					Created:   timestamppb.New(certificate.Created),
					Link:      certificate.Link,
					IsDeleted: certificate.IsDeleted,
				})
			}

			req = &desc.MultiCreateCertificatesV1Request{
				Certificates: multiCertificates,
			}
		})

		It("Test create certificate", func() {
			grpc = api.NewOcpCertificateAPI(r, m, p, s)
			Expect(grpc).ShouldNot(BeNil())
			_, err = grpc.MultiCreateCertificatesV1(ctx, req)
			Expect(err).Should(BeNil())
		})
	})

	Context("Test CreateCertificateV1", func() {
		var req *desc.CreateCertificateV1Request

		BeforeEach(func() {
			synProdMock.ExpectSendMessageAndSucceed()
			m.EXPECT().CreateCounterInc()
			req = &desc.CreateCertificateV1Request{
				Certificate: &desc.NewCertificate{
					UserId:    certificates[0].UserID,
					Created:   timestamppb.New(certificates[0].Created),
					Link:      certificates[0].Link,
					IsDeleted: certificates[0].IsDeleted,
				},
			}

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(1)
			mock.ExpectQuery("INSERT INTO "+tableName).
				WithArgs(req.Certificate.UserId, req.Certificate.Created.AsTime(), req.Certificate.Link, req.Certificate.IsDeleted).
				WillReturnRows(rows)

		})

		It("Test create certificate", func() {
			grpc = api.NewOcpCertificateAPI(r, m, p, s)
			Expect(grpc).ShouldNot(BeNil())
			response, err := grpc.CreateCertificateV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response.CertificateId).Should(BeEquivalentTo(1))
		})
	})

	Context("Test GetCertificateV1", func() {
		var req *desc.GetCertificateV1Request
		var err error

		BeforeEach(func() {
			req = &desc.GetCertificateV1Request{
				CertificateId: certificates[1].ID,
			}

			rows := sqlmock.NewRows([]string{"id", "user_id", "created", "link", "is_deleted"}).
				AddRow(
					certificates[1].ID,
					certificates[1].UserID,
					certificates[1].Created,
					certificates[1].Link,
					certificates[1].IsDeleted)
			mock.ExpectQuery("SELECT id, user_id, created, link, is_deleted FROM " + tableName + " WHERE").
				WithArgs(req.CertificateId).
				WillReturnRows(rows)

		})

		It("Test Get certificate", func() {
			grpc = api.NewOcpCertificateAPI(r, m, p, s)
			Expect(grpc).ShouldNot(BeNil())
			Expect(err).Should(BeNil())

			response, err := grpc.GetCertificateV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response.Certificate.Id).Should(BeEquivalentTo(certificates[1].ID))
			Expect(response.Certificate.UserId).Should(BeEquivalentTo(certificates[1].UserID))
			Expect(response.Certificate.Created.AsTime().Unix()).Should(BeEquivalentTo(certificates[1].Created.Unix()))
			Expect(response.Certificate.Link).Should(BeEquivalentTo(certificates[1].Link))
		})
	})

	Context("Test UpdateCertificateV1Request", func() {
		var req *desc.UpdateCertificateV1Request
		var err error

		BeforeEach(func() {
			synProdMock.ExpectSendMessageAndSucceed()
			m.EXPECT().UpdateCounterInc()
			req = &desc.UpdateCertificateV1Request{
				Certificate: &desc.Certificate{
					Id:      certificates[3].ID,
					UserId:  certificates[3].UserID,
					Created: timestamppb.New(certificates[3].Created),
					Link:    certificates[3].Link,
				},
			}

			mock.ExpectExec("UPDATE "+tableName+" SET").
				WithArgs(
					req.Certificate.UserId,
					req.Certificate.Created.AsTime(),
					req.Certificate.Link,
					req.Certificate.IsDeleted,
					req.Certificate.Id,
				).
				WillReturnResult(sqlmock.NewResult(1, 1))

		})

		It("Test update certificate", func() {
			grpc = api.NewOcpCertificateAPI(r, m, p, s)
			Expect(grpc).ShouldNot(BeNil())
			Expect(err).Should(BeNil())

			response, err := grpc.UpdateCertificateV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response.Updated).Should(BeEquivalentTo(true))
		})
	})

	Context("Test ListCertificateV1Request", func() {
		var req *desc.ListCertificateV1Request
		var err error

		BeforeEach(func() {
			req = &desc.ListCertificateV1Request{
				Limit:  3,
				Offset: 0,
			}

			rows := sqlmock.NewRows([]string{"id", "user_id", "created", "link", "is_deleted"}).
				AddRow(certificates[0].ID, certificates[0].UserID, certificates[0].Created,
					certificates[0].Link, certificates[0].IsDeleted).
				AddRow(certificates[1].ID, certificates[1].UserID, certificates[1].Created,
					certificates[1].Link, certificates[1].IsDeleted).
				AddRow(certificates[2].ID, certificates[2].UserID, certificates[2].Created,
					certificates[2].Link, certificates[2].IsDeleted).
				AddRow(certificates[3].ID, certificates[3].UserID, certificates[3].Created,
					certificates[3].Link, certificates[3].IsDeleted)
			mock.ExpectQuery(
				"SELECT id, user_id, created, link, is_deleted FROM " + tableName + " WHERE").
				WillReturnRows(rows)

		})

		It("Test get list certificates", func() {
			grpc = api.NewOcpCertificateAPI(r, m, p, s)
			Expect(grpc).ShouldNot(BeNil())
			Expect(err).Should(BeNil())

			response, err := grpc.ListCertificateV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response.Certificates[0].Id).Should(BeEquivalentTo(certificates[0].ID))
			Expect(response.Certificates[1].Id).Should(BeEquivalentTo(certificates[1].ID))
			Expect(response.Certificates[2].Id).Should(BeEquivalentTo(certificates[2].ID))
			Expect(response.Certificates[3].Id).Should(BeEquivalentTo(certificates[3].ID))
			Expect(response.Certificates[0].Link).Should(BeEquivalentTo(certificates[0].Link))
			Expect(response.Certificates[1].Link).Should(BeEquivalentTo(certificates[1].Link))
			Expect(response.Certificates[2].Link).Should(BeEquivalentTo(certificates[2].Link))
			Expect(response.Certificates[3].Link).Should(BeEquivalentTo(certificates[3].Link))
		})
	})

	Context("Test RemoveCertificateV1Request", func() {
		var req *desc.RemoveCertificateV1Request
		var err error

		BeforeEach(func() {
			synProdMock.ExpectSendMessageAndSucceed()
			m.EXPECT().RemoveCounterInc()
			req = &desc.RemoveCertificateV1Request{
				CertificateId: certificates[1].ID,
			}

			mock.ExpectExec("DELETE FROM " + tableName).
				WillReturnResult(sqlmock.NewResult(1, 1))

		})

		It("Test remove certificate", func() {
			grpc = api.NewOcpCertificateAPI(r, m, p, s)
			Expect(grpc).ShouldNot(BeNil())
			Expect(err).Should(BeNil())

			response, err := grpc.RemoveCertificateV1(ctx, req)
			Expect(err).Should(BeNil())
			Expect(response.Removed).Should(BeEquivalentTo(true))
		})
	})
})
