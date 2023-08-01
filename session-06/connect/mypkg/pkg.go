package mypkg

import "fmt"

func init() {
	fmt.Println("init from mypkg")
}

func ExportedSomething() {
	fmt.Println("i'm doing something from mypkg")
}
