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
		x <- 42
	}()

	go func() {
		fmt.Println("this is second function")
		x <- 43
	}()

	go func() {
		fmt.Println("this is third function")
		x <- 44
	}()
	fmt.Println(<-x)
	fmt.Println(<-x)
	fmt.Println(<-x)

	fmt.Println("no of go rt af:", runtime.NumGoroutine())
	fmt.Println("main go routine ended...")
}
