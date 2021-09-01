package main

import (
	"fmt"
	"time"
)

func main() {
	s := make(chan int)

	go func() {
		s <- 42
	}()

	fmt.Println(<-s)
	time.Sleep(time.Second * 3)
}
