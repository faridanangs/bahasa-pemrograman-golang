package web

import "time"

type BeritaResponse struct {
	ID        int       `json:"id"`
	Judul     string    `json:"judul"`
	Gambar    string    `json:"gambar"`
	Deskripsi string    `json:"deskripsi"`
	Penulis   string    `json:"penulis"`
	Waktu     time.Time `json:"waktu"`
}
