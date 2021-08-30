package main

import "fmt"

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

const pi = 3.14

func (a circle) area() float64 {
	return pi * a.radius * a.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	c1 := circle{
		radius: 38.56,
	}

	info(c1)

}
