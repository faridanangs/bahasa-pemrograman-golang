package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data ke-" + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	siswa := ring.New(5)

	for i := 0; i < siswa.Len(); i++ {
		siswa.Value = "siswa ke-" + strconv.FormatInt(int64(i), 10)
		siswa = siswa.Next()
	}

	dataSiswa := data.Link(siswa) // kita gunakan untuk menghubungkan data dan siswa

	dataSiswa.Do(func(value any) {
		fmt.Println(value)
	})

}
