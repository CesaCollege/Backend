package main

import "fmt"

type Account struct {
	Name    string
	Family  string
	Balance int
}

func main() {
	account := Account{
		Name:    "Mehdi",
		Family:  "Mostafavi",
		Balance: 1000,
	}

	fmt.Println("Name:", account.Name)
	fmt.Println("Family:", account.Family)
	fmt.Println("Balance:", account.Balance)

	account.Balance += 500
	fmt.Println("Updated Balance:", account.Balance)
}
