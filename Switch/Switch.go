package main

import "fmt"

var a int

func main() {

	fmt.Println("enter the value of a")
	fmt.Scanln(&a)
	switch a {

	case 1:
		fmt.Println("its 1 buddy!!!")
	case 2:
		fmt.Println("its 2 mate!!!")
	case 3:
		fmt.Println("its 3 mate!!!")
	default:
		fmt.Println("enter correct value bro!! pls !! beg you!!")
	}
}
