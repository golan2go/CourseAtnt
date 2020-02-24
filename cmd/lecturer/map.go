package main

import (
	"fmt"
)

func main() {
	var m map[string]float64
	fmt.Printf("type of m = %T\n", m)
	fmt.Println(m == nil, len(m))
	v1 := m["AT&T"]
	fmt.Println("v1 = ", v1) // zero value for float64

	v2, ok := m["AT&T"]
	fmt.Println("v2 = ", v2, " ok = ", ok)

	if v, ok := m["AT&T"]; !ok {
		fmt.Println("not found")
	} else {
		fmt.Println("got ", v)
	}

	if _, ok := m["AT&T"]; !ok {
		fmt.Println("missing")
	}

	// m["AT&T"] = 2.173 // panic
	stocks := make(map[string]float64)
	fmt.Println("stocks", stocks == nil, len(stocks))
	stocks["MSFT"] = 102.17
	fmt.Println(stocks)
	stocks["AAPL"] = 1123.47
	stocks["MSFT"] = 102.71 // will override
	fmt.Println(stocks)

	for k := range stocks {
		fmt.Println(k)
	}

	for k, v := range stocks {
		fmt.Println(k, "->", v)
	}

	delete(stocks, "MSFT")
	fmt.Println(stocks)
	delete(stocks, "MSFT") // can delete non existing

	delete(m, "AT&T")
	fmt.Println(m == nil, m)
}
