package main

import (
	"fmt"
	"time")


func main(){
	now := time.Now()
	
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	fmt.Println(now.UnixNano())
	
	utc := time.Date(2020,10,10,10,10,10,10,time.UTC)
	fmt.Println(utc)
	
	layout := "RFC3399"
	parse, _ := time.Parse(layout,"2020-10-01")
	fmt.Println(parse)
}