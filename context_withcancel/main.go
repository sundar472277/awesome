package main

import (
	"context"
	"fmt"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("context value is :", ctx)
	fmt.Printf("context type is %T\n", ctx)
	fmt.Println("context error is :", ctx.Err())
	cancel()
	fmt.Println("context value after cancel :", ctx)
	fmt.Println("context error after cancel :", ctx.Err())
}
