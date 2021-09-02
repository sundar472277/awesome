package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)
	go rec(even, odd, fanin)

	for i := 0; i <= 5; i++ {
		v, ok := <-fanin
		if ok {
			fmt.Println("value after reading fainin :", v)
		}
	}

	fmt.Println("main ended...")

}
func send(a chan int, b chan int) {
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			a <- i
		} else {
			b <- i
		}
	}
	close(a)
	close(b)
}

func rec(a chan int, b chan int, c chan int) {
	for i := 0; i <= 5; i++ {
		select {
		case v, ok := <-a:
			if ok {
				c <- v
			}
		case v, ok := <-b:
			if ok {
				c <- v
			}
		}
	}
	close(c)

}
