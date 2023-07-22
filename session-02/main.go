package main

import (
	"fmt"

	"session2/parent/child1"
	"session2/parent/child2"
)

func main() {
	fmt.Println("I'm the entry point!")

	child1.Child1Method() // note that unexported functions are not visible
	child2.Child2Method()
}
