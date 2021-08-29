package main

import "fmt"

type secret_agent struct {
	f_name string
	l_name string
	ltk    bool
}
type CIA struct {
	secret_agent
	afg bool
}

func main() {
	s1 := CIA{
		secret_agent: secret_agent{
			f_name: "jeorge",
			l_name: "kutty",
			ltk:    true,
		},
		afg: true,
	}

	s2 := secret_agent{
		f_name: "meera ",
		l_name: "geroge",
		ltk:    false,
	}
	fmt.Println("value of S1 is :", s1)
	fmt.Println("value of s2 is :", s2)
}
