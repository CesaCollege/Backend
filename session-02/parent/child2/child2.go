package child2

import "fmt"

var (
	Child2Var = "child2"
	child2Var = "unexported child2"
)

func Child2Method() {
	fmt.Println("i'm child 2 exported method")

	child2Method()
}

func child2Method() {
	fmt.Println("i'm child 2 unexported method!")
}

func unusedMethod() {
	fmt.Println("i'm unused :/")
}
