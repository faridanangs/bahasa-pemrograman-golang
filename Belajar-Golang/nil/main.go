package main

import "fmt"

// nil dapat di gunakan di interface{}, []slice, map[], pointer, channel
func data(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	var result map[string]string = data("farrrrr")
	if result == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(result)
	}
}
