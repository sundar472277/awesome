package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.GOMAXPROCS(2))
	fmt.Println(runtime.Gosched)

}
