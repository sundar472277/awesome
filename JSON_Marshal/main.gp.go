package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string
	Lname string
	Age   int
	Ltk   bool
}

func main() {
	p1 := person{
		Fname: "sundar",
		Lname: "muni",
		Age:   28,
		Ltk:   true,
	}

	a, ok := json.Marshal(p1)
	fmt.Println("after marshalling : ", string(a))
	if ok != nil {
		fmt.Println("error in marshalling the data")
	}

}
