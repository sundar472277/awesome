package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
	ltk  bool
}
type Byage []person

func (a Byage) Len() int {
	return len(a)
}

func (c Byage) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (b Byage) Less(i, j int) bool {
	return b[i].age < b[j].age
}

func main() {
	p1 := person{
		name: "z",
		age:  88,
		ltk:  true,
	}
	p2 := person{
		name: "raji",
		age:  29,
		ltk:  true,
	}
	res := Byage{p1, p2}
	fmt.Println(res)
	fmt.Printf("%T\n", res)
	sort.Sort(res)
	fmt.Println(res)

}
