package main

import (
	"fmt"
)

func main() {
	// Create the long number
	longNumber := ""
	for i := 0; i < 5000; i++ {
		longNumber += fmt.Sprintf("%d", i+1)
	}

}
