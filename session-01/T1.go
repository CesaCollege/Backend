package main

import "fmt"

func main() {
	var inputInteger int
	lastIntegerInput := 0
	_, _ = fmt.Scanln(&inputInteger)

	if inputInteger > lastIntegerInput {
		fmt.Println(inputInteger)
	}
	lastIntegerInput = inputInteger
}
