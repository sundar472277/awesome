package main

import "fmt"

func main() {
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan bool)

	go send(odd, even, quit)
	rec(odd, even, quit)

	fmt.Println("main ended...")
}

func send(a chan<- int, b chan<- int, c chan<- bool) {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			a <- i
		} else {
			b <- i
		}
	}
	c <- true
	close(c)
}
func rec(a <-chan int, b <-chan int, c <-chan bool) {
	for {
		select {
		case v := <-a:
			fmt.Println("even value is :", v)
		case v := <-b:
			fmt.Println("odd value is :", v)
		case v, ok := <-c:
			fmt.Println("value of ok is :", ok)
			if !ok {
				fmt.Println("value of bool channel !ok is :", v)
				return
			} else {
				fmt.Println("value of bool channel ok is :", v)
			}
		}
	}
}
