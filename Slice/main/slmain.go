package main

import (
	"fmt"
	"github.com/myRepo/Slice"
)

var b int

func main() {
	A := []int{4, 4, 4, 4}
	Slice.RangeSlice(A)

	fmt.Println("enter the values to get appended")
	fmt.Scanln(&b)
	Slice.Append(A, b)
	fmt.Println("value of slice after appending :", A)
	fmt.Println("its wont fetch value from Append function")

}
