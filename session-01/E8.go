package main

import "fmt"

func main() {
	ages := map[string]int{
		"John":  25,
		"Alice": 30,
		"Bob":   50,
	}

	fmt.Println("John's age:", ages["John"])

	ages["Eve"] = 28

	ages["Alice"] = 31

	delete(ages, "Bob")

	for name, age := range ages {
		fmt.Printf("%s is %v years old\n", name, age)
	}
}
