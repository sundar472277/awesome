package main

import "fmt"

var b [5]int

func main() {
	for i, v := range b {
		fmt.Println("index value is :", i, "value after ranging :", v)
	}
}
