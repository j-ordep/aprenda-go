package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type porIdade []user

func (a porIdade) Len() int {
	return len(a)
}

func (a porIdade) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a porIdade) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}


type porSobreNome []user

func (n porSobreNome) Len() int {
	return len(n)
}

func (n porSobreNome) Less(i, j int) bool {
	return n[i].Last < n[j].Last
}

func (n porSobreNome) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// fmt.Println(users)

	// meu código

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(porSobreNome(users))
	sort.Sort(porIdade(users))

	for i, v := range users {
		fmt.Println(i, "- nome completo:", v.First, v.Last, "- idade:", v.Age)
		for _, c := range v.Sayings {
			fmt.Println("-", c)
		}
		fmt.Println("--------------------------------------------------------------")
	}

}