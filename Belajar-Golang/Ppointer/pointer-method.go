package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) sayHello() {
	// supaya data yang ada di dalam struct ikut berubah kita gunakan pointer
	// method * di awal struct untuk memberi tahu bahwa ini adalah pointer
	man.Name = "Hello Mr." + man.Name
}

func main() {
	man := Man{"Farid"} // kita tidak perlu menambah & di dalam pointer method
	man.sayHello()
	fmt.Println(man)
}
