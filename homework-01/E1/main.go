package main

import (
	"bufio"
	"os"
)

func main() {
	// Read the input from STDIN
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	strLength := len(input)

}
