package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []string{"sundar", "muni", "indra"}
	c := [5]string{"one", "two", "three"}
	d := map[string]string{
		"name": "sundar",
		"last": "muni",
	}
	e := struct {
		name string
		age  int
	}{
		name: "ind",
		age:  29,
	}
	fmt.Println("map is :", d)
	fmt.Println("struct is :", e)

	//we cannot range over struct

	for i, v := range d {
		fmt.Println("index is :", i, "value is :", v)
	}

	for i, v := range a {
		fmt.Println("index is :", i, "value is :", v)
	}

	for i, v := range b {
		fmt.Println("index is :", i, "value is :", v)
	}

	for i, v := range c {
		fmt.Println("index is :", i, "value is :", v)
	}
}
