package main

import "fmt"

func main() {
	a := []int{3, 4, 5, 6, 7, 8}
	//variadic parameters should be at last
	afoo("sundar", a...)
}
func afoo(a string, s ...int) {
	fmt.Println(a, s)
}
