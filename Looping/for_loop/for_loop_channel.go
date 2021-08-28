package main

import "fmt"

func main() {
	a := make(chan int, 5)

	for i := 0; i <= 4; i++ {
		a <- i
	}
	close(a)
	for v := range a {
		fmt.Println("pulling values from channel is :", v)
	}
}
