package Saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	x := Greeting("sundar")
	if x != "welcome my dear :sundar" {
		t.Error("Expected :", "welcome my dear :sundar", "but got", x)
	}
}

func ExampleGreet() {
	fmt.Println(Greeting("James"))
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greeting("James")
	}
}
