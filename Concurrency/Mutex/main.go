package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	counter := 0

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(a int) {
			mu.Lock()
			fmt.Println("Go routine iteration :", a)

			counter++
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("counter result :", counter)
	fmt.Println("main ended....")
}
