package child1

import "fmt"

var (
	Child1Var = "child1"
	child1Var = "unexported child1"
)

func Child1Method() {
	fmt.Println("i'm child 1 exported method")

	child1Method()
}

func child1Method() {
	fmt.Println("i'm child 1 unexported method!")
}

func unusedMethod() {
	fmt.Println("i'm unused :/")
}
