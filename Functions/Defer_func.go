package main

import "fmt"

func main() {

	hi()
	bar()
	// last in first out for defer functions
	defer zoo()
	defer how()
}
func hi() {
	fmt.Println("inside hi func")
}
func bar() {
	fmt.Println("inside bar func")
}
func zoo() {
	fmt.Println("inside zoo func")
}
func how() {
	fmt.Println("insie how function")
}
