package repo

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-certificate-api/internal/model"
)

const tableName = "certificate"

var ErrorCertificateNotFound = errors.New("certificate not found")

// Repo - repository interface for entity certificate
type Repo interface {
	AddCertificates(ctx context.Context, certificates []model.Certificate) error
	CreateCertificate(ctx context.Context, certificate *model.Certificate) error
	UpdateCertificate(ctx context.Context, certificate model.Certificate) (bool, error)
	ListCertificates(ctx context.Context, limit, offset uint64) ([]model.Certificate, error)
	GetCertificate(ctx context.Context, certificateId uint64) (*model.Certificate, error)
	RemoveCertificate(ctx context.Context, certificateId uint64) (bool, error)
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

// AddCertificates - creating array certificates in database
func (r *repo) AddCertificates(ctx context.Context, certificates []model.Certificate) error {
	query := squirrel.
		Insert(tableName).
		Columns("user_id", "created", "link").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	for _, certificate := range certificates {
		query = query.Values(certificate.UserId, certificate.Created, certificate.Link)
	}

	exec, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return ErrorCertificateNotFound
	}

	return nil
}

// CreateCertificate - creating single certificate in database
func (r *repo) CreateCertificate(ctx context.Context, certificate *model.Certificate) error {
	query := squirrel.Insert(tableName).
		Columns("user_id", "created", "link").
		Values(certificate.UserId, certificate.Created, certificate.Link).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	err := query.QueryRowContext(ctx).Scan(&certificate.Id)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCertificate - update certificate in database
func (r *repo) UpdateCertificate(ctx context.Context, certificate model.Certificate) (bool, error) {
	query := squirrel.Update(tableName).
		Set("user_id", certificate.UserId).
		Set("created", certificate.Created).
		Set("link", certificate.Link).
		Where(squirrel.Eq{"id": certificate.Id}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	exec, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	rowsAffected, err := exec.RowsAffected()
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
	query := squirrel.Select("id", "user_id", "created", "link").
		From(tableName).
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
		if err := rows.Scan(
			&certificate.Id,
			&certificate.UserId,
			&certificate.Created,
			&certificate.Link); err != nil {
			return nil, err
		}
		certificates = append(certificates, certificate)
	}

	return certificates, nil
}

// GetCertificate - get single certificate from database
func (r *repo) GetCertificate(ctx context.Context, certificateId uint64) (*model.Certificate, error) {
	query := squirrel.Select("id", "user_id", "created", "link").
		From(tableName).
		Where(squirrel.Eq{"id": certificateId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	var certificate model.Certificate

	if err := query.QueryRowContext(ctx).
		Scan(&certificate.Id,
			&certificate.UserId,
			&certificate.Created,
			&certificate.Link); err != nil {
		return nil, err
	}

	return &certificate, nil
}

// RemoveCertificate - remove single certificate in database
func (r *repo) RemoveCertificate(ctx context.Context, certificateId uint64) (bool, error) {
	query := squirrel.Delete(tableName).
		Where(squirrel.Eq{"id": certificateId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected <= 0 {
		return false, ErrorCertificateNotFound
	}

	return true, err
}
