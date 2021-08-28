package main

import (
	"context"
	"fmt"
)

func main() {
	a := 10
	fmt.Println("value of int a :", a)
	b := 9.8
	fmt.Println("value of float b is :", b)
	c := "sundar"
	fmt.Println("value of string c is :", c)
	d := `hai hello
           boss `
	fmt.Println("value of back tip string is :", d)
	e := 0 + 1i
	fmt.Println("value of complex is e :", e)
	f := func() {
		fmt.Println("this is func")
	}
	fmt.Println("value of function f is :", f)
	g := context.Context(context.Background())
	fmt.Println("value of g context is :", g)
	h := true
	fmt.Println("value if h bool is :", h)
	i := []int{2, 3, 4, 5, 6}
	fmt.Println("value of slice i is :", i)
	j := [5]float64{3.4, 5.6, 7.8}
	fmt.Println("value of array is :", j)
	k := make(chan int)
	fmt.Println("value of channel int is :", k)
	l := struct {
		f_name string
		l_name string
		age    int
	}{
		f_name: "sundar",
		l_name: "muni",
		age:    18,
	}
	fmt.Println("value of struct :", l)

}
