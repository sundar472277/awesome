package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	a, b := io.WriteString(os.Stdout, "hello-world")
	fmt.Println("a :", a)
	fmt.Println("b :", b)
}
