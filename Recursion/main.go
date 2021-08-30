package main

import "fmt"

func main() {
	a := factorial(4)
	fmt.Println("factorial of a is :", a)
}

func factorial(s int) int {
	if s == 0 || s == 1 {
		return 1
	}
	return s * factorial(s-1)
}
