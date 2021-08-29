package main

import (
	"fmt"
	"github.com/myRepo/Struct"
)

func main() {
	a := Struct.Person{
		Name:    "muni",
		Age:     65,
		Fav_ice: []string{"blue berry", "guava"},
	}
	c, d, e := Struct.Assignvalues(a)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	//accessing individual values of struct
	fmt.Println("accessing individual values of struct...")
	fmt.Println("c.name :", c.Name)
	fmt.Println("d.age :", d.Age)
	fmt.Println("e.age :", e.Age)
	// Ranging over one of the values of struct
	fmt.Println("ranging over one of the values of struct")
	for i, v := range c.Fav_ice {
		fmt.Println("index is :", i, "values is :", v)
	}
}
