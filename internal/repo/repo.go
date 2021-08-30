package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
)

const tableName = "certificate"

var ErrorCertificateNotFound = errors.New("certificate not found")

// Repo - repository interface for entity certificate
type Repo interface {
	MultiCreateCertificates(ctx context.Context, certificates []model.Certificate) ([]uint64, error)
	CreateCertificate(ctx context.Context, certificate *model.Certificate) error
	UpdateCertificate(ctx context.Context, certificate model.Certificate) (bool, error)
	ListCertificates(ctx context.Context, limit, offset uint64) ([]model.Certificate, error)
	GetCertificate(ctx context.Context, certificateID uint64) (*model.Certificate, error)
}

type repo struct {
	db *sqlx.DB
}

// NewRepo creates new instance
func NewRepo(db *sqlx.DB) Repo {
	return &repo{
		db: db,
	}
}

// MultiCreateCertificates - creating array certificates in database
func (r *repo) MultiCreateCertificates(ctx context.Context, certificates []model.Certificate) ([]uint64, error) {
	query := squirrel.
		Insert(tableName).
		Columns("user_id", "created", "link", "is_deleted").
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	for _, certificate := range certificates {
		query = query.Values(certificate.UserID, certificate.Created, certificate.Link, certificate.IsDeleted)
	}

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	certIds := make([]uint64, 0)
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		certIds = append(certIds, id)
	}
	return certIds, nil
}

// CreateCertificate - creating single certificate in database
func (r *repo) CreateCertificate(ctx context.Context, certificate *model.Certificate) error {
	query := squirrel.Insert(tableName).
		Columns("user_id", "created", "link", "is_deleted").
		Values(certificate.UserID, certificate.Created, certificate.Link, certificate.IsDeleted).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	err := query.QueryRowContext(ctx).Scan(&certificate.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCertificate - update certificate in database
func (r *repo) UpdateCertificate(ctx context.Context, certificate model.Certificate) (bool, error) {
	query := squirrel.Update(tableName).
		Set("user_id", certificate.UserID).
		Set("created", certificate.Created).
		Set("link", certificate.Link).
		Set("is_deleted", certificate.IsDeleted).
		Where(squirrel.Eq{"id": certificate.ID}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	exec, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	rowsAffected, err := exec.RowsAffected()
	if err == sql.ErrNoRows {
		return false, ErrorCertificateNotFound
	}

	if err != nil {
		return false, err
	}

	if rowsAffected <= 0 {
		return false, ErrorCertificateNotFound
	}

	return true, nil
}

// ListCertificates - get list certificate from database
func (r *repo) ListCertificates(ctx context.Context, limit, offset uint64) ([]model.Certificate, error) {
	query := squirrel.Select("id", "user_id", "created", "link", "is_deleted").
		From(tableName).
		Where(squirrel.Eq{"is_deleted": false}).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var certificates []model.Certificate
	for rows.Next() {
		var certificate model.Certificate
		if err = rows.Scan(
			&certificate.ID,
			&certificate.UserID,
			&certificate.Created,
			&certificate.Link,
			&certificate.IsDeleted); err != nil {
			return nil, err
		}
		certificates = append(certificates, certificate)
	}

	return certificates, nil
}

// GetCertificate - get single certificate from database
func (r *repo) GetCertificate(ctx context.Context, certificateID uint64) (*model.Certificate, error) {
	query := squirrel.Select("id", "user_id", "created", "link", "is_deleted").
		From(tableName).
		Where(squirrel.Eq{"id": certificateID}, squirrel.Eq{"is_deleted": false}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	var certificate model.Certificate

	if err := query.QueryRowContext(ctx).
		Scan(&certificate.ID,
			&certificate.UserID,
			&certificate.Created,
			&certificate.Link,
			&certificate.IsDeleted); err != nil {
		return nil, err
	}

	return &certificate, nil
}
