package main

import "fmt"

func main() {
	f := make(chan int, 2)
	f <- 13
	v, ok := <-f
	fmt.Println("value of v is :", v)
	fmt.Println("value ok is :", ok)

}
