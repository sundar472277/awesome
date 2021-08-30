package main

import "fmt"

func main() {
	a := 13.25
	b := &a

	fmt.Println("value of a is :", a)
	fmt.Printf("type of a is : %T\n", a)
	fmt.Println("value of b is :", b)
	fmt.Printf("type of b is : %T\n", b)

	fmt.Println("---------------------------")
	fmt.Println("dereferencing ---------------")
	fmt.Println("value of address stored in b is :", *b)
	fmt.Println("---------------------------------------")
	*b = 20.0
	fmt.Println("value of a after :", a)
	fmt.Printf("type of a is : %T\n", a)

}
