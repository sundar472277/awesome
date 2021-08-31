package main

import (
	"fmt"
	"time"
)

func incrementer(c int) int {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("iteration :", i)
			c++
		}()
	}
	time.Sleep(time.Second * 6)
	return c
}

func main() {
	//shared variable
	counter := 0
	g := incrementer(counter)

	time.Sleep(time.Second * 6)
	fmt.Println("value after increment :", g)

}
