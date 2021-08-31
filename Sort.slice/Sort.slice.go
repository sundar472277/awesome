package main

import (
	"fmt"
	"sort"
)

type person struct {
	fname string
	lname string
	age   int
}

type Byage []person

func main() {

	p1 := person{
		fname: "sundar",
		lname: "muni",
		age:   28,
	}
	p2 := person{
		fname: "raji",
		lname: "muni",
		age:   29,
	}
	sli := Byage{p1, p2}
	sort.Slice(sli, func(i, j int) bool {
		return sli[i].age < sli[j].age
	})
	fmt.Println(sli)
}
