package main

import "fmt"

func main() {
	b := cut(3)
	fmt.Println("value after returning :", b())
}

func cut(a int) func() int {
	return func() int {
		return a
	}
}
