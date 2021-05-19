package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

// func blacklistAdmin(name string) bool{
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool{
// 	return name == "root"
// }

func main() {
	Blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("Frans", Blacklist)
	registerUser("admin", Blacklist)

}
