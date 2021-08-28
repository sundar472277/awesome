package main

import "fmt"

var x string
var a string

func main() {
	y := "sundar"
	fmt.Println("value of y is :", y)
	a = y
	fmt.Println("value of a after assignment :", a)
	a = "muni"
	fmt.Println("value of reassigned a is :", a)
	x = a
	fmt.Println("value of x is :", a)
	fmt.Println("value of x[1] :", x[1])
	//strings can be re assigned but cannot be mutable
	x[1] = 120

	for _, v := range x {
		fmt.Println(v)
	}

}
