package main

func main() {
	var (
		nilai32 int32 = 129
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai32)
	)

	println(nilai32)
	println(nilai64)
	println(nilai8)

	var name = "Frans"
	var e byte = name[0]
	var eString = string(e)

	println(name)
	println(eString)
}
