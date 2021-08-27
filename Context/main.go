package main

import (
	"fmt"
	"time"
)

func main() {

	a, _ := search("delhi", "kabul")
	time.Sleep(5)
	fmt.Println("main ended....and res :", a)
}

func search(from string, to string) ([]string, error) {
	time.Sleep(time.Second * 2)
	a := []string{"sundar", "raji", "muni", from, to}
	return a, nil

}
