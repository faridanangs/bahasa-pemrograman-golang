package repository

import (
	"context"
	"database/sql"
	"restfullapi_news/berita/model/domain"
)

type BeritaRepository interface {
	Save(ctx context.Context, tx *sql.Tx, berita domain.BeritaDomain) domain.BeritaDomain
	Update(ctx context.Context, tx *sql.Tx, berita domain.BeritaDomain) domain.BeritaDomain
	Delete(ctx context.Context, tx *sql.Tx, berita domain.BeritaDomain)
	GetByID(ctx context.Context, tx *sql.Tx, beritaID int) (domain.BeritaDomain, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.BeritaDomain
}
