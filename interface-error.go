package main

import (
	"errors"
	"fmt"
)

func main() {
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError.Error())
}
