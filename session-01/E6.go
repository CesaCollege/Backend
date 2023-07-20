package main

import "fmt"

func main() {
	var numbers1 []int                // Slice
	numbers2 := []int{1, 2, 3}        // Slice
	var numbers3 [5]int               // Array
	numbers4 := [5]int{1, 2, 3, 4, 5} // Array

	n1 := []int{1}
	n2 := []int{4, 8, 5}
	n1 = append(n1, n2...)
	fmt.Println(n1)

	fmt.Println(numbers1[0])
	fmt.Println(numbers2[2])
	fmt.Println(numbers3[4])
	fmt.Println(numbers4[1])
}
