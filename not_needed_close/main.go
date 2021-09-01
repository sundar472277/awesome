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
	//close(a)
	//close(b)
	//close(c)
}
func rec(a <-chan int, b <-chan int, c <-chan bool) {
	for {
		select {
		case <-a:
			fmt.Println("even recieved")
		case <-b:
			fmt.Println("odd recieved")
		case <-c:
			fmt.Println("time to close recieveing")
			return
		}
	}
}
