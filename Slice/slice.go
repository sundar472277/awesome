package Slice

import "fmt"

func RangeSlice(s []int) {
	for _, v := range s {
		fmt.Println("each value in slice :", v)
	}
}
