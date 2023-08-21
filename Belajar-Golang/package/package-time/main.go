package main

import (
	"fmt"
	"time"
)

func main() {
	// Time Now
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Local().Year())
	fmt.Println(now.Local().Month())
	fmt.Println(now.Local().Day())
	fmt.Println(now.Local().Hour())
	fmt.Println(now.Local().Minute())
	fmt.Println(now.Local().Second())
	fmt.Println(now.Local().Nanosecond())
	fmt.Println(now.UTC().Location())
	fmt.Println(now.AddDate(2023, 02, 29))

	// kita atur sendiri datenya
	utc := time.Date(2023, 20, 22, 12, 50, 20, 100, time.UTC)
	fmt.Println(utc)

	//
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2023-10-20")
	fmt.Println(parse)

	//
	time.Sleep(5000)
}
