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
			{1.0, 1.0, now, "https://link.ru", false},
			{2.0, 2.0, now, "https://link.ru", false},
			{3.0, 3.0, now, "https://link.ru", false},
			{4.0, 4.0, now, "https://link.ru", false},
		}
	})

	AfterEach(func() {
		var err error
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	Context("Test MultiCreateCertificates", func() {
		BeforeEach(func() {
			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(1).
				AddRow(2).
				AddRow(3).
				AddRow(4)
			mock.ExpectQuery("INSERT INTO "+tableName).
				WithArgs(
					certificates[0].UserID, certificates[0].Created, certificates[0].Link, certificates[0].IsDeleted,
					certificates[1].UserID, certificates[1].Created, certificates[1].Link, certificates[1].IsDeleted,
					certificates[2].UserID, certificates[2].Created, certificates[2].Link, certificates[2].IsDeleted,
					certificates[3].UserID, certificates[3].Created, certificates[3].Link, certificates[3].IsDeleted,
				).WillReturnRows(rows)
		})

		It("Test add array certificates", func() {
			_, err := r.MultiCreateCertificates(ctx, certificates)
			Expect(err).Should(BeNil())
		})
	})

	Context("Test CreateCertificate", func() {
		BeforeEach(func() {
			rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
			mock.ExpectQuery("INSERT INTO "+tableName).
				WithArgs(
					certificates[0].UserID,
					certificates[0].Created,
					certificates[0].Link,
					certificates[0].IsDeleted).
				WillReturnRows(rows)

		})

		It("Test create certificate", func() {
			certificate := &model.Certificate{ID: 1.0, UserID: 1.0, Created: now, Link: "https://link.ru", IsDeleted: false}
			err := r.CreateCertificate(ctx, certificate)
			Expect(err).Should(BeNil())
			Expect(certificate.ID).Should(BeEquivalentTo(1))
		})
	})

	Context("Test UpdateCertificate", func() {
		BeforeEach(func() {
			mock.ExpectExec("UPDATE "+tableName+" SET").
				WithArgs(
					certificates[1].UserID,
					certificates[1].Created,
					certificates[1].Link,
					certificates[1].IsDeleted,
					certificates[1].ID).
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
			rows := sqlmock.NewRows([]string{"id", "user_id", "created", "link", "is_deleted"}).AddRow(
				certificates[2].ID,
				certificates[2].UserID,
				certificates[2].Created,
				certificates[2].Link,
				certificates[2].IsDeleted)
			mock.ExpectQuery(
				"SELECT id, user_id, created, link, is_deleted FROM " + tableName + " WHERE").
				WithArgs(certificates[2].ID).
				WillReturnRows(rows)
		})

		It("Test get certificate", func() {
			cert, err := r.GetCertificate(ctx, certificates[2].ID)
			Expect(err).Should(BeNil())
			Expect(*cert).Should(BeEquivalentTo(certificates[2]))
		})
	})

	Context("Test ListCertificates", func() {
		var limit uint64 = 4
		var offset uint64 = 0

		BeforeEach(func() {
			rows := sqlmock.NewRows([]string{"id", "user_id", "created", "link", "is_deleted"}).
				AddRow(certificates[0].ID, certificates[0].UserID, certificates[0].Created,
					certificates[0].Link, certificates[0].IsDeleted).
				AddRow(certificates[1].ID, certificates[1].UserID, certificates[1].Created,
					certificates[1].Link, certificates[1].IsDeleted).
				AddRow(certificates[2].ID, certificates[2].UserID, certificates[2].Created,
					certificates[2].Link, certificates[2].IsDeleted).
				AddRow(certificates[3].ID, certificates[3].UserID, certificates[3].Created,
					certificates[3].Link, certificates[3].IsDeleted)
			mock.ExpectQuery("SELECT id, user_id, created, link, is_deleted FROM " + tableName + " WHERE").
				WillReturnRows(rows)
		})

		It("Test get list certificates", func() {
			certificate, err := r.ListCertificates(ctx, limit, offset)
			Expect(err).Should(BeNil())
			Expect(certificate[1].ID).Should(BeEquivalentTo(certificates[1].ID))
			Expect(certificate[1].UserID).Should(BeEquivalentTo(certificates[1].UserID))
			Expect(certificate[1].Created).Should(BeEquivalentTo(certificates[1].Created))
			Expect(certificate[1].Link).Should(BeEquivalentTo(certificates[1].Link))
		})
	})
})
