package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	for i := 0; i <= 10; i++ {
		select {
		case v := <-even:
			fmt.Println("value after recieveing even :", v)
		case v := <-odd:
			fmt.Println("value after recieving odd :", v)
		case _, ok := <-quit:
			if !ok {
				return
			}
		}
	}
}
func send(a chan<- int, b chan<- int, c chan<- bool) {
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			a <- i
		} else {
			b <- i
		}
	}
	close(a)
	close(b)
	c <- true
}
