package Paying

import "strings"

var c []string

func ConcatString(s []string) string {
	a := ""
	for _, v := range s {
		a = a + v
	}
	return a
}

func SplitString(b string) []string {
	c := strings.Split(b, " ")
	return c

}
