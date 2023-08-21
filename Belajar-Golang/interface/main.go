package main

import "fmt"

// supaya interface ini bisa di gunakan maka kita harus membuat kntrak
// seperti siswaElit
// type interface
type Data interface {
	getName() (string, string)
}

// Type struct
type SiswaMiskin struct {
	Name, Addres string
}
type Animals struct {
	Name, Gender string
}
type Pemrograman struct {
	Name, Penemu string
}

// data siswa
func siswaElit(name string, siswa Data) {
	nameSiswa, addres := siswa.getName()
	fmt.Println("Hello boos", name, "saya ", nameSiswa, "and from", addres)
}
func (siswa SiswaMiskin) getName() (string, string) {
	return siswa.Name, siswa.Addres
}

// data hewan
func dataAnimal(animal Data) {
	nameAnimal, genderAnimal := animal.getName()
	fmt.Println("Ini hewah apa hayyoo? ", nameAnimal, "gendernya apa hayyo?", genderAnimal)
}
func (animal Animals) getName() (string, string) {
	return animal.Name, animal.Gender
}

// data bahasa pemrograman
func bahasaPemrograman(bahasa Data) {
	name, penemu := bahasa.getName()
	fmt.Println("Nama bahasa ", name, "penemu", penemu)
}
func (bahasa Pemrograman) getName() (string, string) {
	return bahasa.Name, bahasa.Penemu
}

func main() {
	// data siswa
	data := SiswaMiskin{Name: "kuspusi", Addres: "Sesela"}
	siswaElit("farid", data)

	// data hewan
	animal := Animals{Name: "Kucing", Gender: "male"}
	dataAnimal(animal)

	// data bahasa pemrograman
	bahasa := Pemrograman{Name: "Python", Penemu: "anangs"}
	bahasaPemrograman(bahasa)
}
