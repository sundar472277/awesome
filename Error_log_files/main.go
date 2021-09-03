package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	a, err := os.Create(`file-name.txt`)
	defer panichandle()
	log.SetOutput(a)
	if err != nil {
		log.Panicln("error in creating file")
	}
	_, err2 := os.OpenFile("file-name_4.txt", 2, 0644)
	if err2 != nil {
		log.Panicln("file open error")
	}
	err1 := os.WriteFile("file-name_3.txt", []byte("hellO"), 0644)
	if err1 != nil {
		fmt.Println("file not found")
	}
}

func panichandle() {
	if a := recover(); a != nil {
		fmt.Println("program recovered")
	}
}
