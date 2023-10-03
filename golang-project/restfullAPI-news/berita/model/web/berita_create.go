package web

type BeritaCreateRequest struct {
	Judul     string `validate:"required,max=300,min=10" json:"judul"`
	Gambar    string `validate:"required,max=335" json:"gambar"`
	Deskripsi string `validate:"required" json:"deskripsi"`
	Penulis   string `validate:"required,max=40,min=4" json:"penulis"`
}
