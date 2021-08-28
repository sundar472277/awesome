package main

import "fmt"

func main() {
	i := 0
	for {
		if i <= 3 {
			fmt.Println("printing the iteration :", i)
		} else {
			break
		}
		i++
	}
}
