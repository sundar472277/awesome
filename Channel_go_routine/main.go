package main

import "fmt"

func main() {
	a := make(chan int)

	go foo(a)
	fmt.Println(<-a)
}

func foo(s chan int) {
	s <- 42
}
