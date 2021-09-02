package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("value of context :", ctx)
	fmt.Printf("type of context : %T\n", ctx)
	fmt.Println("context error :", ctx.Err())

}
