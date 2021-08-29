package main

import "fmt"

type inter interface {
	speak()
}

type person struct {
	name string
	age  int
	ltk  bool
}

type hotdog int

var x hotdog

func (a person) speak() {
	fmt.Println("inside speak method :", a)
}

func (b hotdog) speak() {
	fmt.Println("inside hotdog :", b)
}

func call(a inter) {
	a.speak()
}

func main() {

	p1 := person{
		name: "sundar",
		age:  28,
		ltk:  false,
	}

	p1.speak()
	call(p1)

	call(x)

}
