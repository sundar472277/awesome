package main

import "fmt"

func main() {
	a := 11
	fmt.Printf("binary of a before bit shifting : %b\n", a)
	b := a << 1
	fmt.Printf("value of a after bit shifting : %b\n", b)
}
