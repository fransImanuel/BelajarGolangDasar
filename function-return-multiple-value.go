package main

import "fmt"

func getFullName() (string, string) {
	return "Frans", "Imanuel"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Print(firstName)
	fmt.Println(lastName)

	namaDepan, _ := getFullName()
	fmt.Println(namaDepan)
}
