package main

import (
	"fmt"
	"runtime"
	"time"
)

func lm(a string) {
	for i := 0; true; i++ {
		fmt.Println(i, a)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	fmt.Println("main started...")
	fmt.Println("before no of go :", runtime.NumGoroutine())
	go lm("first")

	lm("second")
	fmt.Println("after no of go :", runtime.NumGoroutine())
	fmt.Println("main ended...")

}
