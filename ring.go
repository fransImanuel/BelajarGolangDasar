package main

import (
	"fmt"
	"container/ring"	
	"strconv"
)


func main() {
	var data *ring.Ring = ring.New(5)

	// data.Value = "Eko"

	// var data2 = data.Next()
	// data2.Value = "Kurniawan"

	for i := 0; i < data.Len(); i++{
		data.Value = "Data"+ strconv.FormatInt(int64(i),10)
		data = data.Next()
	}

	// fmt.Println(*data)
	data.Do(func(value interface{}){
		fmt.Println(value)
	})
}