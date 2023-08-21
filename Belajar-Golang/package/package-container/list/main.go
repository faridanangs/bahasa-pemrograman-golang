package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("farid")
	data.PushBack("anang")
	data.PushBack("samudra")
	data.PushBack("hebat")
	// data.InsertBefore("puuusu", data.Front().Next())
	// data.InsertAfter("puksksau", data.Back().Prev())

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Front().Next().Value)
	// fmt.Println(data.Front().Prev().Value)
	// fmt.Println(data.Back().Value)

	for value := data.Front(); value != nil; value = value.Next() {
		fmt.Println(value.Value)
	}
}
