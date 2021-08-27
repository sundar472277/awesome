package main

import (
	"context"
	"fmt"
)

var x int
var y float64
var z []int
var a []string
var b complex64
var c complex128
var d string
var e struct{}
var f interface{}
var g map[string]int
var h [5]int
var i context.Context
var j [][]int
var k []struct{}
var l []map[string][]int
var m []interface{}
var n []context.Context
var o [][]context.Context
var p [][]map[string][][]int
var q [][]struct{}
var r [][]map[string][][]struct{}
var s [][]map[string][][]context.Context

func main() {
	foo()
}

func foo() {
	fmt.Println("var int x:", x)
	fmt.Println("var float y:", y)
	fmt.Println("var z slice of int", z)
	fmt.Println("var a slice of string :", a)
	fmt.Println("var b complex64 :", b)
	fmt.Println("var c complex 128 :", c)
	fmt.Println("var d string :", d)
	fmt.Println("var e struct :", e)
	fmt.Println("var f interface :", f)
	fmt.Println("var g map: ", g)
	fmt.Println("var h array of int :", h)
	fmt.Println("var i context :", i)
	fmt.Println("multi dim slice int :", j)
	fmt.Println("slice of struct :", k)
	fmt.Println("slice of map string and int :", l)
	fmt.Println("slice of interface :", m)
	fmt.Println("slice of context :", n)
	fmt.Println("multi dim context :", o)
	fmt.Println("multi dim map string and multi dim int :", p)
	fmt.Println("multi dim struct :", q)
	fmt.Println("multi dim map string and multi dim struct :", r)
	fmt.Println("multi dim map string and multi dim context :", s)
}
