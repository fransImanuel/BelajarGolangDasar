package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Selesai")
	message := recover()
	fmt.Println("Error dengan message:", message)

}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("APLIKASI BERJALAN")
}

func main() {
	runApp(true)
}
