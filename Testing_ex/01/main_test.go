package _1

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	x := Sum(3, 4, 5)
	if x != 12 {
		fmt.Println("expected", 12, "but got :", x)
	}
}
