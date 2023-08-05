package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read the input from STDIN
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

	var people [100]bool
	for i := 0; i < input; i++ {
		people[i] = true
	}

	// Start eliminating
	eliminated := 0
	idx := 1
	skip := false
	for eliminated != input-1 {
		// Handle the cycle format
		idx = idx % 100

		if skip && people[idx] {
			skip = false
		} else if !skip && people[idx] {
			skip = true
			people[idx] = false
			eliminated++
		}

		idx++
	}

	for i, person := range people {
		if person {
			fmt.Println(i + 1)
			return
		}
	}
}
