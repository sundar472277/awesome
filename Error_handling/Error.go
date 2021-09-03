package main

import (
	"fmt"
	"log"
)

var b string
var a string
var c string

func main() {
	a := "sundar"
	b := "sundar"
	fmt.Println("main started...")
	check(a, b, c)
	fmt.Println("main ended....")
}
func check(fname string, address string, lname string) {

	defer panichandle()

	if fname == "" {
		panic("fname cannot be nill")
	}
	if address == "" {
		log.Panicln("address cannot be null")
	}
	if lname == "" {
		log.Fatalln("lname cannot be null")
	}
	fmt.Println("function ended...")
}

func panichandle() {
	if a := recover(); a != nil {
		fmt.Println("program recovered")
	}
}
