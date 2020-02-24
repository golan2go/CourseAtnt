package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("nums", nums)
	nums[1] = 8
	fmt.Println("nums after", nums)
	nums2 := nums
	fmt.Println("nums 2", nums2)
	nums2[0] = 9
	fmt.Println("nums after nums2 change", nums)
	nums = append(nums, 10)
	fmt.Println("nums after append", nums)
	fmt.Println("nums2 after append", nums2)
	fmt.Println("nums: len = ", len(nums), ", cap = ", cap(nums))
	appendSizes()

	// copy slice dest <- src
	copy(nums[1:], []int{-1, -2})
	fmt.Printf("nums 2 after copy = %v\n", nums)
	fmt.Println("concat", concat([]int{1, 2}, []int{3, 4, 5}))

	s := make([]string, 2, 10)
	fmt.Printf("s = %#v: len=%d, cap=%d\n", s, len(s), cap(s))
}

// concat gets two slices & return a slice a concatenation of both
// concat([]int{1,2}, []int{3,4,5}) -> []int{1,2,3,4,5}
// bonus: do it without a "for" loop
func concat(s1, s2 []int) []int {
	// TODO
	out := make([]int, len(s1)+len(s2))
	copy(out, s1)
	copy(out[len(s1):], s2)
	return out
}

func appendSizes() {
	currCap := 0
	var s []int
	for i := 0; i < 1500; i++ {
		s = append(s, i)
		if c := cap(s); c != currCap {
			fmt.Printf("%d -> %d\n", currCap, c)
			currCap = c
		}
	}
}
