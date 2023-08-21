package hel

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

type Data struct {
	Name, Address string
}

func (data *Data) dataKuu() {
	if data.Address == "sesela" {
		data.Address = "Kebon Lauk"
	}
	fmt.Println("name:", data.Name, "alamat:", data.Address)
}
func DataGue(name, address string) {
	data := Data{Name: name, Address: address}
	data.dataKuu()
	fmt.Println(data)
}

var Name string

// func init ini di jalankan sebelum funct main tanpa di panggil di dalam func main
func init() {
	Name = "farid anang s"
	fmt.Println("ini iit")
}

func NameKu() {
	fmt.Println(Name)
}
