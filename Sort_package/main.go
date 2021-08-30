package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{99, 21, 35, 1, 9}
	fmt.Println("before sorting int slice:", x)
	sort.Ints(x)
	fmt.Println("after sorting int slice :", x)

	y := []float64{4.5, 1.8, 99.67}
	fmt.Println("before sorting float slice :", y)
	sort.Float64s(y)
	fmt.Println("after sorting float slice :", y)

	z := []string{"hi", "hello", "how", "are"}
	fmt.Println("before sorting string slice :", z)
	sort.Strings(z)
	fmt.Println("after sorting string slice", z)
}
