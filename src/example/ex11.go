package main

import (
	"fmt"
)

func main() {
	// Map
	var iamMap map[int]string
	iamMap = make(map[int]string)
	fmt.Println(iamMap)
	tickers := map[string]string{
		"KeyA": "ValA",
		"KeyB": "ValB",
		"KeyC": "ValC",
	}
	fmt.Println(tickers)
	fmt.Println()

	// Using Map
	var m map[int]string

	m = make(map[int]string)
	m[1] = "val1"
	m[11] = "val11"
	m[111] = "val111"
	fmt.Println(m)

	str := m[111]
	fmt.Println(str)

	delete(m, 11)
	fmt.Println(m)
	fmt.Println()

	// Map Key Check
	stocks := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "Facebook",
		"AMZN": "Amazon",
	}
	_, exists := stocks["MSFT"]
	if !exists {
		fmt.Println("No MSFT stocks")
	} else {
		fmt.Println("Has MSFT stocks")
	}
	fmt.Println()

	// using for loop Map
	for key, val := range stocks {
		fmt.Println(key, val)
	}
	fmt.Println()
}
