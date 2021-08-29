package main

import "fmt"

var a map[string]string

func main() {

	a = map[string]string{
		"first_name": "sundar",
		"last_name":  "muni",
		"hometown":   "theni",
		"address": `miranda line first 
                   street`,
	}
	fmt.Println(a)
	//map is unordered collection.ranging over to see it
	for i, v := range a {
		fmt.Println("index is :", i, "value is :", v)
	}
	b, err := a["first_name"]
	fmt.Println(b)
	fmt.Println(err)
	//replacing value to a map
	a["first_name"] = "raji"
	fmt.Println("printing after replacing value :", a)
	//adding value to a map
	a["ltk"] = "yes"
	fmt.Println("value after adding :", a)
	//deleting value from a map
	if _, err := a["ltk"]; err == true {
		delete(a, "ltk")
	}
	//printing after deleting a ltk value from map
	fmt.Println("printing after deleting value from map :", a)
	//retrieving a value not available in map
	d, err := a["new"]
	fmt.Println("value of key new is :", d)
	fmt.Println("error value :", err)

}
