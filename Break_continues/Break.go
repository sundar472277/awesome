package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		if i == 2 {
			continue
		} else if i == 4 {
			break
		}
		fmt.Println("iteration is :", i)
	}
}
