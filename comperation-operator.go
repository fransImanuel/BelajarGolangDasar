package main

func main() {
	var (
		name1 = "Frans"
		name2 = "Frans"
	)

	var result bool = name1 == name2
	println(result)

	var (
		value1 = 100
		value2 = 200
	)
	println(value1 > value2)
	println(value1 < value2)
	println(value1 == value2)
	println(value1 != value2)

}
