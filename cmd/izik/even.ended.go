package main

import "fmt"
import "math"

func main() {

	var count = 0
	for a := 1000; a < 10000; a++ {
		for b := a; b < 10000; b++ {
			if isEventEnded(a * b) {
				count++
			}
		}
	}

	fmt.Println(count)

}

func isEventEnded(num int) bool {
	last := num % 10
	len := int(math.Log10(float64(num)))
	pow := math.Pow(float64(len), 10)
	c := num / int(pow)
	return last == c

}
