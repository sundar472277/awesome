package _2

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	x := Multiply(4, 3)
	if x != 12 {
		t.Error("expected", 12, "but got :", x)

	}
}
