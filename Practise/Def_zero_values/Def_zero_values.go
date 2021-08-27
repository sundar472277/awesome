package main

import (
	"context"
	"fmt"
)

var x interface{}
var y bool
var z func()
var a *int
var b *float64
var c []int
var d map[string]int
var e context.Context
var f chan int

func main() {
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

}
