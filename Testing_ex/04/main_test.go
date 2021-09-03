package _4

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	x := Greet("sundar")
	if x != "welcome mr. sundar" {
		t.Error("expected but got :", x)
	}
}
func ExampleGreet() {
	fmt.Println(Greet("sundar"))
}
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("muni")
	}
}
