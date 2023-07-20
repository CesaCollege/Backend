package main

import "fmt"

func sum1(x int, y int) int {
	return x + y
}

func sum2(x, y int) int {
	return x + y
}

func sum3(x, y int) (sum int) {
	sum = x + y
	return
}

func calculate(x, y int) (sum int, divide int) {
	sum = x + y
	divide = x / y
	return
}

func main() {
	fmt.Println(sum1(1, 2))
	fmt.Println(sum2(1, 2))
	fmt.Println(sum3(1, 2))
	fmt.Println(calculate(1, 2))
}
