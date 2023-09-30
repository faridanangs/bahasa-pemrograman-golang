package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

func goroutine(syncMap *sync.Map, wg *sync.WaitGroup, i int) {
	wg.Done()
	syncMap.Store(i, i)
}

func TestMap(t *testing.T) {
	syncMap := &sync.Map{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go goroutine(syncMap, wg, i)
	}

	for j := 0; j < 100; j++ {
		syncMap.Delete(j)
	}
	syncMap.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})

	wg.Wait()
	fmt.Println("selesai")
}
