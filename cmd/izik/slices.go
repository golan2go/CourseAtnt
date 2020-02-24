package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	b := []int{3, 4}
	res := concat(a, b)

	pr(a)
	pr(b)
	pr(res)

	c := 3
	print(c * 3)

}

func pr(a []int) {
	fmt.Printf("  cap=[%v]  arr=%v \n", cap(a), a)
}

func concat(s1, s2 []int) []int {
	res := make([]int, len(s1)+len(s2))
	copy(res[:len(s1)], s1)
	copy(res[len(s1):], s2)
	return res
}
