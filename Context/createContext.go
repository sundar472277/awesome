package main

import (
	"context"
	"fmt"
)

func main() {
	cxt := context.Background()
	fmt.Println("value of context :", cxt)
	fmt.Println("if error:", cxt.Err())
	fmt.Printf("type of context is %T\n:", cxt)

}
