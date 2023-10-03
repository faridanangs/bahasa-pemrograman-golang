package service

import (
	"context"
	"database/sql"
	"restfullapi_news/berita/model/domain"
	"restfullapi_news/berita/model/web"
	"restfullapi_news/berita/repository"
	"restfullapi_news/helper"

	"github.com/go-playground/validator"
)

type BeritaServiceIplm struct {
	BeritaRepository repository.BeritaRepository
	Db               *sql.DB
	Validate         *validator.Validate
}

func NewBeritaService(beritaRepository repository.BeritaRepository, db *sql.DB, validate *validator.Validate) *BeritaServiceIplm {
	return &BeritaServiceIplm{
		Db:               db,
		Validate:         validate,
		BeritaRepository: beritaRepository,
	}
}

func (service *BeritaServiceIplm) Create(ctx context.Context, request web.BeritaCreateRequest) web.BeritaResponse {
	err := service.Validate.Struct(request)
	helper.IfLogingErr(err, "error di validate.struct create service")

	tx, err := service.Db.Begin()
	helper.IfLogingErr(err, "error di db.begin create service")
	defer helper.CommitOrRollback(tx)

	berita := domain.BeritaDomain{
		Gambar:    request.Gambar,
		Judul:     request.Judul,
		Deskripsi: request.Deskripsi,
		Penulis:   request.Penulis,
	}

	berita = service.BeritaRepository.Save(ctx, tx, berita)

	return helper.ToBeritaResponse(berita)

}
func (service *BeritaServiceIplm) Update(ctx context.Context, request web.BeritaUpdateRequest) web.BeritaResponse {
	err := service.Validate.Struct(request)
	helper.IfLogingErr(err, "error di validate.struct update service")

	tx, err := service.Db.Begin()
	helper.IfLogingErr(err, "error di Db.Begin update service")
	defer helper.CommitOrRollback(tx)

	berita, err := service.BeritaRepository.GetByID(ctx, tx, request.ID)

	berita = domain.BeritaDomain{
		ID:        request.ID,
		Gambar:    request.Gambar,
		Judul:     request.Judul,
		Deskripsi: request.Deskripsi,
		Penulis:   request.Penulis,
	}

	berita = service.BeritaRepository.Update(ctx, tx, berita)

	return helper.ToBeritaResponse(berita)
}
func (service *BeritaServiceIplm) Delete(ctx context.Context, beritaID int) {
	tx, err := service.Db.Begin()
	helper.IfLogingErr(err, "error di Db.Begin delete service")
	defer helper.CommitOrRollback(tx)

	berita, err := service.BeritaRepository.GetByID(ctx, tx, beritaID)
	service.BeritaRepository.Delete(ctx, tx, berita)

}
func (service *BeritaServiceIplm) GetByID(ctx context.Context, beritaID int) web.BeritaResponse {
	tx, err := service.Db.Begin()
	helper.IfLogingErr(err, "error di Db.Begin getbyid service")
	defer helper.CommitOrRollback(tx)

	berita, err := service.BeritaRepository.GetByID(ctx, tx, beritaID)
	return helper.ToBeritaResponse(berita)

}
func (service *BeritaServiceIplm) GetAll(ctx context.Context) []web.BeritaResponse {
	tx, err := service.Db.Begin()
	helper.IfLogingErr(err, "error di Db.Begin getall service")
	defer helper.CommitOrRollback(tx)

	beritas := service.BeritaRepository.GetAll(ctx, tx)
	return helper.ToBeritaResponses(beritas)
}
