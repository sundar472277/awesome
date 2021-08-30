package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string
	Lname string
	Age   int
	Ltk   bool
}

var p1 person

func main() {
	A := `{"Fname":"sundar","Lname":"muni","Age":28,"Ltk":true}`
	fmt.Println(A)
	json.Unmarshal([]byte(A), &p1)

	fmt.Println("after unmarshalling :", p1)
}
