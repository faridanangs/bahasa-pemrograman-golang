package main

import (
	"fmt"
	"reflect"
)

type Type struct {
	Name string `required:"true" max:"10" min:"1"` // Tag ini memudahkan kita jika kita ingin membuat validasi
}

type People struct {
	Name, Addres, Bohi string `required:"true"`
}

func IsValid(value interface{}) bool {
	refVal := reflect.TypeOf(value)
	for i := 0; i < refVal.NumField(); i++ {
		field := refVal.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(value).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	user := Type{"farid"}
	// var sampleType = reflect.TypeOf(user) // untuk mendefinisikan bahawa ini adalah reflace
	// fmt.Println(sampleType.NumField()) // untuk melihat jumlah nilai yg ada di dalam struct
	// fmt.Println(sampleType.Field(0)) // untuk mengambil semua data yang sesuai dengan index yang di berikan di dalam struct
	// fmt.Println(sampleType.Field(0).Tag.Get("max")) // untuk mengambil nilai yang ada di dalam Tag
	// fmt.Println(sampleType.Field(0).Tag.Get("min"))
	// fmt.Println(sampleType.Field(0).Tag.Get("required"))
	// fmt.Println(reflect.ValueOf(user)) // Untuk mengambil data yang ada di dala user yang di kirim di dala struct
	user.Name = ""
	fmt.Println(IsValid(user))

	person := People{"Farid anang s", "Sesela", "Coding"}
	// persReflace := reflect.TypeOf(person)
	// fmt.Println(persReflace.Field(1).Name) // kita gunakan .Name untuk mengambil sesuai dgan namanya saja
	// person2 := People{"", "", ""} // Jika kita tidak menambah required di Tag maka ini akan tetap true
	// fmt.Println(reflect.ValueOf(person).Field(1)) // Untuk mengambil data yang ada di dalam struct sesuai dengan field() yang di kirim
	fmt.Println(IsValid(person))
	fmt.Println(reflect.TypeOf(person)) // ini dia mengembalikan main.nama structnya
	fmt.Println(reflect.TypeOf(user))   // ini dia mengembalikan main.nama structnya
}
