package helper

import (
	"restfullapi_news/berita/model/domain"
	"restfullapi_news/berita/model/web"
)

func ToBeritaResponse(berita domain.BeritaDomain) web.BeritaResponse {
	return web.BeritaResponse{
		ID:        berita.ID,
		Gambar:    berita.Gambar,
		Judul:     berita.Judul,
		Deskripsi: berita.Deskripsi,
		Penulis:   berita.Penulis,
		Waktu:     berita.Waktu,
	}
}

func ToBeritaResponses(berita []domain.BeritaDomain) []web.BeritaResponse {
	var responses []web.BeritaResponse
	for _, value := range berita {
		responses = append(responses, ToBeritaResponse(value))
	}
	return responses
}
