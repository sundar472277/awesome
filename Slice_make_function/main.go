package main

import "fmt"

func main() {
	a := make([]int, 10)
	fmt.Println("value of slice a :", a)
	fmt.Printf("type of slice a is : %T\n", a)
	fmt.Println("length of slice a is :", len(a))
	fmt.Println("max capacity if slice is :", cap(a))
	//Ranging over slice and assigning value to slice
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println("value after filling slice  :", a)
	//appending slice using for range loop
	for i := 10; i <= 20; i++ {
		a = append(a, i)
	}
	fmt.Println("value of slice after appending :", a)
	fmt.Println("length after appending :", len(a))
	fmt.Println("capacity after appending :", cap(a))

}
