package main

import "fmt"

func main() {
	a := make(chan int)
	c := make(chan bool)

	go send(a, c)
	for i := 0; i <= 5; i++ {
		select {
		case v := <-a:
			fmt.Println("received value from a :", v)
		case v, ok := <-c:
			if !ok {
				return
			} else {
				fmt.Println("value recieved from c :", v)
			}
		}
	}

}
func send(a chan int, c chan bool) {
	for i := 0; i <= 10; i++ {
		a <- i
	}
	close(a)
	c <- true
}
