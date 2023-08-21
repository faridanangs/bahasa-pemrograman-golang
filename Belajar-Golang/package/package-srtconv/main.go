package main

import (
	"fmt"
	"strconv"
)

func main() {

	// strconv.Parse kita gunakan untuk mengubah string sesuai dengan convigure
	// strconv.Format kita gunakan untuk mengubah nilai convigure menjadi string

	//String to Boolean
	parsBool, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(parsBool)
	} else {
		fmt.Println(err.Error())
	}

	//String to Intjer
	parseInt, err := strconv.ParseInt("36576832", 10, 32)
	if err == nil {
		intVal := int32(parseInt)
		fmt.Println(intVal)
	} else {
		fmt.Println(err.Error())
	}

	// String to Float
	parseFloat, _ := strconv.ParseFloat("10.2001010", 10)
	fmt.Println(parseFloat)

	// Boot to String
	formatBool := strconv.FormatBool(true)
	fmt.Println(formatBool)

	// Int to String
	formatInt := strconv.FormatInt(12012, 2)
	fmt.Println(formatInt)

	fmt.Println(" ")
	// Float to Sring
	// e
	formatFloatE := strconv.FormatFloat(10.212121, 'e', 3, 64)
	fmt.Printf("%+v\n", formatFloatE)
	// g
	formatFloatG := strconv.FormatFloat(10.212121, 'g', 5, 64)
	fmt.Printf("%+v\n", formatFloatG)
	// x
	formatFloatX := strconv.FormatFloat(10.212121, 'x', 2, 64)
	fmt.Printf("%+v\n", formatFloatX)
	// f
	formatFloatF := strconv.FormatFloat(10.212121, 'f', 4, 64)
	fmt.Printf("%+v\n", formatFloatF)
}
