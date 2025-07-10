package main

import (
	"fmt"
	"sort"
)

func main() {

	ss := []string{"abobora", "laranja", "banana", "beringela", "maça"}
	fmt.Println(ss)
	
	sort.Strings(ss) // ordenou por ordem alfabetica
	fmt.Println(ss)

	nums := []int{3, 5, 4, 6, 8, 1, 7}
	fmt.Println(nums)
	
	sort.Ints(nums)
	fmt.Println(nums)
}