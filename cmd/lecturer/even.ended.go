/*
An "even ended number" is a number who's first and last digit are the same.
You mission, should you choose to accept it, is to count how
many multiplications of 4 digit numbers are "even ended"

1023 * 3423 = 3,501,729 -> not even ended
 1021 * 3423 = 3,494,883 -> is even ended

Answer: 3184963
*/
package main

import (
	"fmt"
)

func main() {
	count := 0
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			if isEvenEnded(a * b) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func isEvenEnded(n int) bool {
	s := fmt.Sprintf("%d", n)
	return s[0] == s[len(s)-1]
}
