package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodbye := getGoodBye

	result := sayGoodbye("Frans")
	fmt.Println(result)
}
