package main

import "fmt"

func main() {
	name := "farid"
	count := 0

	closure := func() {
		var name string
		var count int = 2

		// dia akan mengubah data yang ada di atasnya jika kita tidak
		// membuat vaiable di dalam function closue
		name = "usin"

		// closure bersarang
		closureCount := func() {
			count++
		}
		closureCount()
		closureCount()
		closureCount()
		closureCount()

		fmt.Println("di dalam closure", name)
		fmt.Println("di dalam closure", count)
	}

	closure()

	fmt.Println(name)
	fmt.Println(count)
}
