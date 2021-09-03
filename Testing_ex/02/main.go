package _2

func Multiply(a ...int) int {
	res := 1
	for _, v := range a {
		res = res * v
	}
	return res
}
