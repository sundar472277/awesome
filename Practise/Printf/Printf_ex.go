package main

import "fmt"

func main() {
	a := 10
	b := 12.30
	c := true
	d := "string"

	fmt.Printf("Type of a is %T\n", a)
	fmt.Printf("type of b is %T\n", b)
	fmt.Printf("Type of c is %T\n", c)
	fmt.Printf("Type of c is %T\n", d)
	fmt.Printf("Binary of a is %b\n", a)
	fmt.Printf("hexadecimal of a is %#x\n", a)
	fmt.Printf("another form of Hexadecimal is %#x\n", a)
	fmt.Printf("utf of a is %#U\n", a)
	fmt.Println("--------------------------")
	fmt.Printf("value of a is %v\n", a)
	z := fmt.Sprintf("value of a again %v", a)
	fmt.Println(z)
}
