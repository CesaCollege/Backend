package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read the input from STDIN
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_, _ = strconv.Atoi(scanner.Text())

	// Read the numbers from STDIN
	scanner.Scan()
	strNumbers := strings.Split(scanner.Text(), " ")
	var numbers []int
	for _, number := range strNumbers {
		n, _ := strconv.Atoi(number)
		numbers = append(numbers, n)
	}

	// Find the maximum in these numbers
	maximum := math.MinInt
	for _, n := range numbers {
		if n > maximum {
			maximum = n
		}
	}

	fmt.Println(maximum)
}
