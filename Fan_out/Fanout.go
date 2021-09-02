package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	a := make(chan int)
	go send(a)
	rec(a)
	wg.Wait()
	fmt.Println("main ended...")

}
func send(a chan int) {
	for i := 0; i <= 100; i++ {
		a <- i
	}
	close(a)
}
func rec(a chan int) {

	for v := range a {
		wg.Add(1)
		go func(a int) {
			counter(a)
			wg.Done()
		}(v)
	}
}
func counter(a int) {
	fmt.Println("Counter value is :", a)
}
