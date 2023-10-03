package service

import (
	"context"
	"restfullapi_news/berita/model/web"
)

type BeritaService interface {
	Create(ctx context.Context, request web.BeritaCreateRequest) web.BeritaResponse
	Update(ctx context.Context, request web.BeritaUpdateRequest) web.BeritaResponse
	Delete(ctx context.Context, beritaID int)
	GetByID(ctx context.Context, beritaID int) web.BeritaResponse
	GetAll(ctx context.Context) []web.BeritaResponse
}
