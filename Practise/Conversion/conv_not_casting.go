package main

import "fmt"

type hotdog int

var x int
var y hotdog

func main() {
	x = 10
	y = 12
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	//below we cannot compare
	//fmt.Println(x==y)
	a := int(y)
	fmt.Println("value after conversion :", a)
	fmt.Printf("Type of a is %T\n", a)

}
