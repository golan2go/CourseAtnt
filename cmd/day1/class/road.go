package main

import (
	"fmt"
	"strings"
)

var road = `
The Road goes ever on and on
Down from the door where it began
Now far ahead the Road has gone
And I must follow if I can
Pursuing it with eager feet
Until it joins some larger way
Where many paths and errands meet
And whither then I cannot say
`

func main() {

	var winner string
	max := 0

	words := make(map[string]int)

	terms := strings.Fields(road)

	for _, term := range terms {
		term = strings.ToLower(term)
		words[term]++
		if words[term] > max {
			max = words[term]
			winner = term
		}
	}

	fmt.Println("winner [", winner, "] max=[", max, "]")

	fmt.Println(words)

}
