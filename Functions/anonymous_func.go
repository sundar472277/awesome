package main

import (
	"fmt"
)

func main() {
	//simple anonymous func
	func() {
		fmt.Println("this is ananymous func")
	}()
	//passing parameter in anonymous func

	func(a int) {
		fmt.Println("print parameter in anonymous func :", a)
	}(42)
	//passing multiple values in anonymous func

	func(a ...int) {
		fmt.Println("value of a is :", a)

	}(3, 4, 5, 6, 7)

}
