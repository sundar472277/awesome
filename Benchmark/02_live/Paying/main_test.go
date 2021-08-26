package Paying

import "testing"

func BenchmarkConcatString(b *testing.B) {
	c = []string{"hi", "hello", "how", "are", "you"}
	for i := 0; i <= b.N; i++ {
		ConcatString(c)
	}
}

func BenchmarkSplitString(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		SplitString("hhhjjhjhjh jjhjhhjhjgdd  kkkkk  dddwfggwe")
	}
}
