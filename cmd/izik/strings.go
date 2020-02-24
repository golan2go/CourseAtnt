package main

import "fmt"

func main() {

	s := "kdfg \u2344"

	for i, c := range s {
		fmt.Printf("%d %c %d %T \n", i, c, c, c)
	}

	fmt.Println("=====")

	for c := range s {
		fmt.Printf("%d", c)
	}

	fmt.Println("=====")

	fmt.Println(s)

	fmt.Println("=====")

	json :=
		`
	{
		val: 12,
		obj: "12"
	}
	
	`

	fmt.Print(json)

}
