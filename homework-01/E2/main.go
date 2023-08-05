package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Create the long number
	longNumber := ""
	for i := 0; i < 5000; i++ {
		longNumber += fmt.Sprintf("%d", i+1)
	}

	// Read the input from STDIN
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

	fmt.Println(string(longNumber[input-1]))
}
