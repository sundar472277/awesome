package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	//git clone https://github.com/golang/crypto.git
	//create a folder golang.org/x/crypto/ under GOPATH/src
	//use this git clone under GOPATH/src/golang.org/x/crypto/
	s := `password123`
	fmt.Println()
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 2)
	fmt.Println("bs value is :", bs)
	fmt.Println("error is 1 :", err)
	if err != nil {
		fmt.Println("error is :", err)
	}

	err1 := bcrypt.CompareHashAndPassword(bs, []byte(s))
	if err != nil {
		fmt.Println("login failed", err1)
	} else {
		fmt.Println("login success")
	}

}
