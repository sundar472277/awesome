package main

import "fmt"

func main() {
	sm := callfunc(multiply, 2, 3, 4, 5)
	fmt.Println("final value is :", sm)
	sm1 := call1(multiply, 5, 6, 7, 8)
	fmt.Println("final value 1 is :", sm1)
}

func multiply(c int, d int) int {
	return c * d
}

func callfunc(a func(a int, b int) int, b ...int) int {
	sum := 0
	for _, v := range b {
		sum = sum + v
	}
	return sum
}

func call1(l func(a int, b int) int, m ...int) int {
	sm := 0
	for _, v := range m {
		sm = sm + v
	}
	return sm
}
