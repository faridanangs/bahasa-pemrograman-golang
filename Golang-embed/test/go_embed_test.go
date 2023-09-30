package test

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"testing"
)

////// string
// untuk me embed atau memasukan nilai dari text ke dalam variable bertipe string kita bisa
// menggunakannya seperti ini //go:embed namatext.txt supaya nilai dari dalam file
// text masuk ke dalam variable maka kita taruh di atas variable di manapun yang penting di atasnya
// jika di bawahny maka dia akan error dan kita bisa memanggilnya secara berulang

//go:embed embed.txt
var embedText string

//go:embed embed.txt
var embedText2 string

func TestEmbedString(t *testing.T) {
	fmt.Println(embedText)
	fmt.Println(embedText2)
}

//go:embed Screenshot.png
var logo []byte

func TestEmbedByte(t *testing.T) {
	err := os.WriteFile("logo_new.png", logo, 0666)
	if err != nil {
		panic(err)
	}
}

func TestRemoveFile(t *testing.T) {
	os.Remove("logo_new.png")
}

func TestReadFile(t *testing.T) {
	data, err := os.ReadFile("logo_new.png")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

// multiple files jika kita gnakan ini maka kita harus menampungnya
// di sebuah variable yang bertipe embed.FS supaya bisa di panggil

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultiple(t *testing.T) {
	// cara memanggil multiple files kita gunakan namaVariable.ReadFile("nama_filenya")
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))
	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

// match multiple files sama seperti yang di atas ini kita bisa menggunakan
// multiple files hanya dengan menaruh tanda *.txt semua jenis
// dari file txt akan di ambil

//go:embed files/*.txt
var path embed.FS

func TestMachsMultipleFiles(t *testing.T) {
	// untuk membaca directori di dalam folder kita gunakan namaVarEmbed + ReadDir("nama_folder")
	dirEntry, _ := path.ReadDir("files")

	// kemudian kita iterasi data yang sudah kita dapat di readdir
	for _, entry := range dirEntry {
		if !entry.IsDir() {
			// untuk mendapatkan nama nama yang ada di dalam folder yang kita readdir kita bisa
			// menggunakan entry.Name()
			fmt.Println(entry.Name())

			// jika kita ingin membaca text yang ada di dalanya kita bisa menggunakan yang di bawah nin
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
