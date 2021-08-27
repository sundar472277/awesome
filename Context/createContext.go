package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	cxt := context.Background()
	fmt.Println("value of context :", cxt)
	fmt.Println("if error:", cxt.Err())
	fmt.Printf("type of context is %T\n:", cxt)
	fmt.Println("------------------------------------")
	cxt1, _ := context.WithCancel(cxt)
	fmt.Println("value of context :", cxt1)
	fmt.Println("if error:", cxt1.Err())
	fmt.Printf("type of context is %T\n:", cxt1)
	fmt.Println("------------------------------------")
	cxt2, _ := context.WithDeadline(cxt1, time.Time{})
	fmt.Println("value of context :", cxt2)
	fmt.Println("if error:", cxt2.Err())
	fmt.Printf("type of context is %T\n:", cxt2)
	fmt.Println("------------------------------------")
	cxt3, _ := context.WithTimeout(cxt, time.Second*4)
	fmt.Println("value of context :", cxt3)
	fmt.Println("if error:", cxt3.Err())
	fmt.Printf("type of context is %T\n:", cxt3)
	fmt.Println("------------------------------------")

}
