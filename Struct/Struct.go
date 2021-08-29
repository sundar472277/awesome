package Struct

type Person struct {
	Name    string
	Age     int
	Fav_ice []string
}

func Assignvalues(S Person) (Person, Person, Person) {
	P1 := Person{
		Name:    "sundar",
		Age:     29,
		Fav_ice: []string{"vennila", "milk_shake"},
	}
	P2 := Person{
		Name:    "raji",
		Age:     29,
		Fav_ice: []string{"straw", "orange"},
	}
	return P1, P2, S
}
