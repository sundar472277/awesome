package main

import "fmt"

var a string

func main() {
	a = "sundar muni"
	for i, v := range a {
		fmt.Println("index is :", i, "value is :", v)
	}
	fmt.Println("-------------------------------------------")

	b := []byte(a)
	fmt.Println("value of slice of bytes is :", b)

	fmt.Println("----------------------------------------------")

	for _, v := range a {
		fmt.Printf("binary is %b\n", v)
		fmt.Printf("hexa-decimal is %#X\n :", v)
		fmt.Printf("UTF is %#U\n :", v)
	}

}
