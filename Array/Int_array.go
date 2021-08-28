package main

import "fmt"

var A [5]int

func main() {
	Foo()

}
func Foo() {
	fmt.Println("before getting values in array :", A)
	A = [5]int{2, 3, 4, 5, 6}
	fmt.Println("after getting values in array :", A)
}
