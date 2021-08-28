package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("main go routine started")
	fmt.Println("no of go rt bf:", runtime.NumGoroutine())

	x := make(chan int)

	go func() {
		fmt.Println("this is first function")
		for i := 0; i <= 5; i++ {
			x <- i
		}
		close(x)
	}()

	for v := range x {
		fmt.Println("value retrieved from channel x :", v)
	}

	fmt.Println("no of go rt af:", runtime.NumGoroutine())
	fmt.Println("main go routine ended...")
}
