package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	// Read the input from STDIN
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

}
