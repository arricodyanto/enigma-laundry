package main

import (
	"challenge-godb/menu"
)

func main() {
	// customers := manipulation.GetAllCustomer()
	// for _, customer := range customers {
	// 	fmt.Println(customer.Id, customer.Name, customer.Contact)
	// }
	// defer errorRecover()

	for {
		menu.ShowMenu()
	}
}

// func printError(err error) {
// 	if err != nil {
// 		fmt.Println("Error occured:", err)
// 	}
// }
