package main

import "fmt"

// tanpa recursive
func menggunakanForLoop(value int) int {
	count := 1
	for i := value; i > 0; i-- {
		count *= i
	}
	return count
}

// menggunakan recursive
func menggunakanRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * menggunakanRecursive(value-1)
	}
}

func main() {
	fmt.Println(menggunakanForLoop(5))
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(menggunakanRecursive(9))

}
