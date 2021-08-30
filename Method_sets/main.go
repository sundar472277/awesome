package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	age   int
	ltk   bool
}

type human interface {
	speak()
}

func (a *person) speak() {
	fmt.Printf("type of a inside function : %T\n", a)
	fmt.Println("name of person is :", a.fname, " ", a.lname)
	(*a).fname = "ajith"
	fmt.Println("value of person struct after updating :", *a)
}

func call(b human) {

	b.speak()
}

func main() {

	p1 := person{
		fname: "sundar",
		lname: "muni",
		age:   28,
		ltk:   true,
	}
	p2 := person{
		fname: "vijay",
		lname: "joseph",
		age:   56,
		ltk:   false,
	}
	fmt.Println("printing value of struct before calling :", p1)
	(&p1).speak()
	fmt.Println("printing person struct after calling :", p1)

	fmt.Println("value of p2 before calling function :", p2)
	call(&p2)
	fmt.Println("value of p2 after calling function :", p2)

}
