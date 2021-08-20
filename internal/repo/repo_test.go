package repo_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
	"github.com/ozoncp/ocp-certificate-api/internal/repo"
	"time"
)

var _ = Describe("Repo", func() {
	const tableName = "certificate"

	now := time.Now()

	var (
		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock

		ctx          context.Context
		r            repo.Repo
		certificates []model.Certificate
	)

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())
		sqlxDB = sqlx.NewDb(db, "sqlmock")

		ctx = context.Background()
		r = repo.NewRepo(sqlxDB)

		certificates = []model.Certificate{
			{1.0, 1.0, now, "http://link"},
			{2.0, 2.0, now, "http://link"},
			{3.0, 3.0, now, "http://link"},
			{4.0, 4.0, now, "http://link"},
		}
	})

	AfterEach(func() {
		var err error
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	Context("Test AddCertificates", func() {
		BeforeEach(func() {
			mock.ExpectExec("INSERT INTO "+tableName).
				WithArgs(
					certificates[0].UserId, certificates[0].Created, certificates[0].Link,
					certificates[1].UserId, certificates[1].Created, certificates[1].Link,
					certificates[2].UserId, certificates[2].Created, certificates[2].Link,
					certificates[3].UserId, certificates[3].Created, certificates[3].Link,
				).WillReturnResult(sqlmock.NewResult(1, 1))
		})

		It("Test add array certificates", func() {
			err := r.AddCertificates(ctx, certificates)
			Expect(err).Should(BeNil())
		})
	})

	Context("Test CreateCertificate", func() {
		BeforeEach(func() {
			rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
			mock.ExpectQuery("INSERT INTO "+tableName).
				WithArgs(
					certificates[0].UserId,
					certificates[0].Created,
					certificates[0].Link).WillReturnRows(rows)

		})

		It("Test create certificate", func() {
			certificate := &model.Certificate{Id: 1.0, UserId: 1.0, Created: now, Link: "http://link"}
			err := r.CreateCertificate(ctx, certificate)
			Expect(err).Should(BeNil())
			Expect(certificate.Id).Should(BeEquivalentTo(1))
		})
	})

	Context("Test UpdateCertificate", func() {
		BeforeEach(func() {
			mock.ExpectExec("UPDATE "+tableName+" SET").
				WithArgs(
					certificates[1].UserId,
					certificates[1].Created,
					certificates[1].Link,
					certificates[1].Id).
				WillReturnResult(sqlmock.NewResult(1, 1))
		})

		It("Test update certificate", func() {
			updated, err := r.UpdateCertificate(ctx, certificates[1])
			Expect(err).Should(BeNil())
			Expect(updated).Should(BeEquivalentTo(true))
		})
	})

	Context("Test GetCertificate", func() {
		BeforeEach(func() {
			rows := sqlmock.NewRows([]string{"id", "user_id", "created", "link"}).AddRow(
				certificates[2].Id,
				certificates[2].UserId,
				certificates[2].Created,
				certificates[2].Link)
			mock.ExpectQuery(
				"SELECT id, user_id, created, link FROM " + tableName + " WHERE").
				WithArgs(certificates[2].Id).
				WillReturnRows(rows)
		})

		It("Test get certificate", func() {
			cert, err := r.GetCertificate(ctx, certificates[2].Id)
			Expect(err).Should(BeNil())
			Expect(*cert).Should(BeEquivalentTo(certificates[2]))
		})
	})

	Context("Test CreateCertificate", func() {
		BeforeEach(func() {
			query := mock.ExpectExec("DELETE FROM " + tableName + " WHERE")
			query.WithArgs(certificates[3].Id)
			query.WillReturnResult(sqlmock.NewResult(1, 1))
		})

		It("Test remove certificate", func() {
			deleted, err := r.RemoveCertificate(ctx, certificates[3].Id)
			Expect(err).Should(BeNil())
			Expect(deleted).Should(BeEquivalentTo(true))
		})
	})
})
