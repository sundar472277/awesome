package main

import "fmt"

func main() {
	fmt.Println("main started...")
	a := make(chan int)

	go foo(a)

	for v := range a {
		fmt.Println(v)
	}
	fmt.Println("main ended...")
}

func foo(a chan<- int) {
	for i := 0; i < 5; i++ {
		//fmt.Println("iteration :",i)
		a <- i
	}
	close(a)
}
