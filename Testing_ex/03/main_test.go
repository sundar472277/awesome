package _3

import "testing"

type nm struct {
	value []int
	res   int
}

func TestSum(t *testing.T) {
	n1 := nm{
		value: []int{2, 2, 2, 2},
		res:   8,
	}
	n2 := nm{
		value: []int{3, 3, 3, 3},
		res:   12,
	}
	tot := []nm{n1, n2}

	for _, v := range tot {
		res1 := Sum(v.value...)
		if res1 != v.res {
			t.Error("expected :", v.res, "got :", res1)
		}
	}

}
