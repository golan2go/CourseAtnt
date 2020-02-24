package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello ☺"
	fmt.Printf("s = %q, len=%d\n", s, len(s))
	fmt.Printf("nums of runes = %d\n", utf8.RuneCountInString(s))
	for i, c := range s {
		fmt.Printf("%d: %c (%T)\n", i, c, c)
	}
	fmt.Printf("s[0] = %c (%T)", s[0], s[0])

	// raw string
	poem := `
The road goes ever on and on
Down from the door where it began.
`
	fmt.Println(poem)

	s2, s3 := "a\tb", `a\tb`
	fmt.Printf("s2 = %s, s3 = %s\n", s2, s3)

	s4 := fmt.Sprintf("the answer is %d", 42)
	fmt.Println(s4)
	// see also bytes.Buffer, strings.Builder, text/template & html/template
}
