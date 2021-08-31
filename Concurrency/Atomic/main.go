package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func main() {

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(a int) {

			fmt.Println("iteration of go routine :", a)
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			atomic.LoadInt64(&counter)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("value of resultant counter :", counter)
	fmt.Println("....main ended....")
}
