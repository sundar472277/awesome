package main

import (
	"fmt"
	"time"
)

//infinite for loop creation
func infin(s string) {
	for i := 0; true; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 900)
	}
}

func main() {
	fmt.Println("main started....")
	// calling first function
	infin("first")
	//calling again
	infin("second")

	fmt.Println("main ended....")

}
