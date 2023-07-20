package main

import "fmt"

func main() {
	numbers := []int{5, 2, 3, 7, 0}
	for key, value := range numbers {
		fmt.Println(key, value)
	}
}
