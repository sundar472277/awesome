package main

import "fmt"

func main() {
	x := 10
	y := fmt.Sprintf("value of int x is %d\n", x)
	fmt.Println(y)

	a := 13.40
	b := fmt.Sprintf("value of float a is %f\n", a)
	fmt.Println(b)

	c := true
	d := fmt.Sprintf("value of bool c is %t\n", c)
	fmt.Println(d)

	e := "string"
	f := fmt.Sprintf("value of string e is %s\t :", e)
	fmt.Println(f)

}
