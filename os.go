package main

import (
	"fmt"
	"os")


func main() {
	args := os.Args
	fmt.Println(args)
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])


	name, error := os.Hostname()

	if error == nil {
		fmt.Println("Hostname: ",name)
	}else{
		fmt.Println("Error: ", error.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}