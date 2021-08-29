package main

import "fmt"

func main() {
	s1 := struct {
		name    string
		age     int
		fav_ice []string
	}{
		name:    "sundar",
		age:     29,
		fav_ice: []string{"vennila", "mango"},
	}
	fmt.Println(s1)
}
