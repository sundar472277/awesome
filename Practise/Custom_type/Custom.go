package main

import (
	"context"
	"fmt"
)

type typ1 int
type typ2 float64
type typ3 bool
type typ4 func()
type typ5 [5]int
type typ6 []int
type typ7 struct {
	name string
	age  int
}
type typ8 context.Context
type typ9 chan int
type typ10 interface {
	speak()
}
type typ11 map[string][]int
type typ12 *int
type typ13 string
type typ14 complex128
type typ15 complex64

func main() {
	fmt.Println("just a list of custom type and underlying type")

}
