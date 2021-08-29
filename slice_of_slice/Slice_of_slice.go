package main

import "fmt"

func main() {
	a := []int{3, 4, 5, 6, 7, 8, 9}

	fmt.Println("before slicing :", a)

	a = a[2:]
	fmt.Println("after first slicing :", a)
	a = a[:2]
	fmt.Println("after second slice :", a)
}
