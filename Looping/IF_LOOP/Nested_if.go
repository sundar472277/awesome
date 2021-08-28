package main

import "fmt"

var a int

func main() {
	fmt.Println("enter the number between 3 and 5 :")
	fmt.Scanln(&a)
	if a == 3 {
		fmt.Println("its 3 bro!!!")
	} else if a == 4 {
		fmt.Println("its 4 buddy!!!")
	} else if a == 5 {
		fmt.Println("its 5 mate!!")
	} else {
		fmt.Println("are you out of your mind ? huh ?")
	}
}
