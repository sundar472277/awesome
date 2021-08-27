package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("num of go routine :", runtime.NumGoroutine())
	fmt.Println("check for error :", ctx.Err())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done with cancel")
				return
			default:
				n++
				time.Sleep(time.Second * 1)
				fmt.Println("working", n)

			}
		}

	}()
	time.Sleep(time.Second * 6)
	fmt.Println("error check 2", ctx.Err())
	fmt.Println("context is about to get cancelled")
	fmt.Println("number of go-routine after :", runtime.NumGoroutine())
	cancel()
}
