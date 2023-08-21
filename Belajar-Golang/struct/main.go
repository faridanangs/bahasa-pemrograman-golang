package main

import "fmt"

type Siswa struct {
	Name, Addres string
	Age          int
}

func (siswa Siswa) saySiswa(name string) {
	fmt.Println("hello ", name, "my name is", siswa.Name)
}

func (siswa Siswa) siswaSultan(name string) {
	fmt.Println("Amppunn boss ", name, ", nama gue", siswa.Name)
}

func dataSiswa(name string, addres string, age int) {
	// cara pertama
	// var data Siswa
	// data.Name = name
	// data.Addres = addres
	// data.Age = age
	// fmt.Println(data)

	// cara kedua
	// Lebih baik pake cara ke 2
	data := Siswa{
		Name:   name,
		Addres: addres,
		Age:    age,
	}

	// cara ketiga
	// ini rentant error jika kita ubah urutan data yang ada di dalam
	// struct ini akan eror
	// data := Siswa{name, addres, age}
	data.saySiswa("diana")
	data.siswaSultan("anangs")
	fmt.Println(data)
}

func main() {
	dataSiswa("udin", "sesela", 17)

}
