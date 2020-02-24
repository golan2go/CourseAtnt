/* What is the most common word in road.txt? (ignoring case)

- Use strings.Fields to split to words
- Use strings.ToLower to convert to lower case
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, word := range strings.Fields(poem) {
		counts[strings.ToLower(word)]++
	}

	maxWord, maxCount := "", 0
	for word, count := range counts {
		if count > maxCount {
			maxWord, maxCount = word, count
		}
	}
	fmt.Println(maxWord)
}

var poem = `
The Road goes ever on and on
Down from the door where it began
Now far ahead the Road has gone
And I must follow if I can
Pursuing it with eager feet
Until it joins some larger way
Where many paths and errands meet
And whither then I cannot say
