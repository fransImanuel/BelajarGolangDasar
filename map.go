package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{
	// }

	person := map[string]string{
		"name":    "frans",
		"address": "kutabumi",
	}

	person["title"] = "Programmer"

	book := make(map[string]string)
	book["title"] = "Belajar golang"
	book["author"] = "Eko"
	book["ups"] = "salah"
	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
