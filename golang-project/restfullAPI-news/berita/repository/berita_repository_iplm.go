package repository

import (
	"context"
	"database/sql"
	"errors"
	"restfullapi_news/berita/model/domain"
	"restfullapi_news/helper"
)

type BeritaRepositoryIplm struct{}

func NewBeritaRepository() *BeritaRepositoryIplm {
	return &BeritaRepositoryIplm{}
}

func (repository *BeritaRepositoryIplm) Save(ctx context.Context, tx *sql.Tx, berita domain.BeritaDomain) domain.BeritaDomain {
	script := "insert into berita(gambar,judul,deskripsi,penulis) values(?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, berita.Gambar, berita.Judul, berita.Deskripsi, berita.Penulis)
	helper.IfLogingErr(err, "ExecContext create repository")

	id, err := result.LastInsertId()
	helper.IfLogingErr(err, "LastInsertId create repository")
	berita.ID = int(id)
	return berita

}

func (repository *BeritaRepositoryIplm) Update(ctx context.Context, tx *sql.Tx, berita domain.BeritaDomain) domain.BeritaDomain {
	script := "update berita set gambar =? ,judul =? ,deskripsi =? ,penulis =?  where id =?"
	_, err := tx.ExecContext(ctx, script, berita.Gambar, berita.Judul, berita.Deskripsi, berita.Penulis, berita.ID)
	helper.IfLogingErr(err, "ExecContext update repository")
	return berita
}

func (repository *BeritaRepositoryIplm) Delete(ctx context.Context, tx *sql.Tx, berita domain.BeritaDomain) {
	script := "delete from berita where id = ?"
	_, err := tx.ExecContext(ctx, script, berita.ID)
	helper.IfLogingErr(err, "ExecContext delete repository")
}

func (repository *BeritaRepositoryIplm) GetByID(ctx context.Context, tx *sql.Tx, beritaID int) (domain.BeritaDomain, error) {
	script := "select * from berita where id = ?"
	rows, err := tx.QueryContext(ctx, script, beritaID)
	helper.IfLogingErr(err, "QueryContext getbyid repository")
	defer rows.Close()

	berita := domain.BeritaDomain{}
	if rows.Next() {
		err := rows.Scan(&berita.ID, &berita.Gambar, &berita.Judul, &berita.Deskripsi, &berita.Penulis, &berita.Waktu)
		helper.IfLogingErr(err, "err di rows scan getbyid repository")
		return berita, nil
	} else {
		return berita, errors.New("berita tidak di temukan")
	}
}

func (repository *BeritaRepositoryIplm) GetAll(ctx context.Context, tx *sql.Tx) []domain.BeritaDomain {
	script := "select id,judul,deskripsi,gambar,penulis,waktu from berita"
	rows, err := tx.QueryContext(ctx, script)
	helper.IfLogingErr(err, "QueryContext getall repository")
	defer rows.Close()

	var beritas []domain.BeritaDomain
	for rows.Next() {
		berita := domain.BeritaDomain{}
		err := rows.Scan(&berita.ID, &berita.Gambar, &berita.Judul, &berita.Deskripsi, &berita.Penulis, &berita.Waktu)
		helper.IfLogingErr(err, "err di rows scan getall repository")
		beritas = append(beritas, berita)
	}
	return beritas
}
