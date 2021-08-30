package main

import "fmt"

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * 3.14
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c1 := circle{
		radius: 5.6,
	}

	info(&c1)

}
