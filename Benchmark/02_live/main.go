package main

import (
	"fmt"
	"github.com/myRepo/Benchmark/02_live/Paying"
)

func main() {
	m := `Coverage in programming is how much of our code is covered by tests. We can use the “-cover” flag to run coverage analysis on our code. We can use the flag and required file name “-coverprofile <some file name>” to write our coverage analysis to a file.`

	n := Paying.SplitString(m)
	fmt.Println("after splitting :", n)

	o := Paying.ConcatString(n)
	fmt.Println("after concatination :", o)

}
