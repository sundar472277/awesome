package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("assigning this func to var")
	}
	a()

	b := func(x int) {
		fmt.Println("value of variable :", x)
	}
	b(5)

	c := func(y ...int) {
		fmt.Println("value of y is :", y)
	}
	c(1, 2, 3, 4)

	fmt.Printf("type of a is %T\n", a)
}
