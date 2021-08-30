package main

import (
	"fmt"
)

func main() {

	{
		x := 13
		fmt.Println("value inside the block :", x)
	}
	//below statement throw error because its used outside block
	//fmt.Println("value outside the block :",x)
}
