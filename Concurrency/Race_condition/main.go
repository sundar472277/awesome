package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started...")
	fmt.Printf("type of wg :%T\n", wg)
	counter := 0
	wg.Add(11)
	for i := 0; i <= 10; i++ {

		go func(a int) {
			fmt.Println("routine count :", a)
			counter++
			wg.Done()
		}(i)
	}
	//time.Sleep(time.Millisecond)
	//instead of time we can use wait group
	wg.Wait()
	fmt.Println("counter value after cal :", counter)
	fmt.Println("number of goroutines after :", runtime.NumGoroutine())
	fmt.Println("main ended...")

}
