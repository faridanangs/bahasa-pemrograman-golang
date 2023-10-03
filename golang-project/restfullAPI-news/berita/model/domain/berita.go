package domain

import "time"

type BeritaDomain struct {
	ID        int
	Gambar    string
	Judul     string
	Deskripsi string
	Penulis   string
	Waktu     time.Time
}
