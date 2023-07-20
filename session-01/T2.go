package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}

	var userInput int
	var evenNumbers []int
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanln(&userInput)
		if userInput%2 == 0 {
			evenNumbers = append(evenNumbers, userInput)
		}
	}

	fmt.Println("Done getting inputs")

	for _, value := range evenNumbers {
		fmt.Println(value)
	}
}
