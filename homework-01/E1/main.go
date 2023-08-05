package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read the input from STDIN
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	strLength := len(input)

	for i := 0; i < strLength/2; i++ {
		if input[i] != input[strLength-i-1] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
