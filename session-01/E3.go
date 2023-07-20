package main

import "fmt"

func main() {
	age := 61

	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 60 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

	if age := 25; age >= 18 {
		fmt.Println("You are an adult.")
	}
}
