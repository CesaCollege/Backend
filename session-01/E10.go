//package main
//
//import "fmt"
//
//type Address struct {
//	City  string
//	State string
//}
//
//type BankAccount struct {
//	Name    string
//	Family  string
//	Balance int
//	Address Address
//}
//
//func main() {
//	account := BankAccount{
//		Name:    "Mehdi",
//		Family:  "Mostafavi",
//		Balance: 1000,
//		Address: Address{
//			City:  "Mazandaran",
//			State: "Sari",
//		},
//	}
//
//	fmt.Println("Name:", account.Name)
//	fmt.Println("Family:", account.Family)
//	fmt.Println("Balance:", account.Balance)
//	fmt.Println("X:", account.Address.City)
//
//	account.Balance += 500
//	fmt.Println("Updated Balance:", account.Balance)
//}
