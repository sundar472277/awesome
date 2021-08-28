package Slice

import "fmt"

func Append(S []int, A int) {
	S = append(S, A)
	fmt.Println("inside append function :", S)
}
