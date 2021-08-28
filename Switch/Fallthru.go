package main

import "fmt"

var b int

func main() {

	fmt.Println("enter the value of a")
	fmt.Scanln(&b)
	switch b {

	case 1:
		fmt.Println("its 1 buddy!!!")
	case 2:
		fmt.Println("its 2 mate!!!")
		fallthrough
	case 3:
		fmt.Println("its 3 mate!!!")
	default:
		fmt.Println("enter correct value bro!! pls !! beg you!!")
	}
}
