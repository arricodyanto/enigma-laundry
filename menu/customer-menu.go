package menu

import (
	"bufio"
	"challenge-godb/db/master"
	"challenge-godb/entity"
	"challenge-godb/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CustomerManagement() {
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("\nCUSTOMER MANAGAMENT MENU:")
	fmt.Println("1. Show All Customers")
	fmt.Println("2. Add New Customer")
	fmt.Println("3. Update Customer")
	fmt.Println("4. Delete Customer")
	fmt.Println("5. Back")
	fmt.Println("\n" + strings.Repeat("=", 80))

	var menu int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Select Customer Menu : ")
	scanner.Scan()
	menu, _ = strconv.Atoi(scanner.Text())

	defer utils.ErrorRecover()

	switch menu {
	case 1:
		showCustomer()
		CustomerManagement()
	case 2:
		addNewCustomer()
	case 3:
		updateOneCustomer()
	case 4:
		deleteOneCustomer()
	case 5:
		fmt.Println("<<- Back to Main Menu")
	default:
		fmt.Printf("Menu '%d' does not exist\n", menu)
	}
}

func showCustomer() {
	customers := master.GetAllCustomer()
	fmt.Println("\nID", strings.Repeat(" ", 4), "Name", strings.Repeat(" ", 40), "No. HP", strings.Repeat(" ", 15))
	fmt.Println(strings.Repeat("-", 80))
	for _, customer := range customers {
		idLength := len(strconv.Itoa(customer.Id))
		nameLength := len(customer.Name)
		fmt.Println(customer.Id, strings.Repeat(" ", 6-idLength), customer.Name, strings.Repeat(" ", 44-nameLength), customer.Contact)
	}
}

func addNewCustomer() {
	scanner := bufio.NewScanner(os.Stdin)
	newCustomer := entity.Customer{}
	defer CustomerManagement()
	defer utils.ErrorRecover()

	fmt.Print("Enter Customer Name : ")
	scanner.Scan()
	newCustomer.Name = scanner.Text()

	fmt.Print("Enter Customer Contact : ")
	scanner.Scan()
	newCustomer.Contact = scanner.Text()

	fmt.Print("\nAre you sure want to add this new customer (y/n)? ")
	scanner.Scan()
	confirm := scanner.Text()

	if strings.ToLower(confirm) == "y" {
		master.AddCustomer(newCustomer)
	} else if strings.ToLower(confirm) == "n" {
		fmt.Println("New Customer Data was not saved.")
	} else {
		fmt.Println("Input is invalid!")
	}
}

func updateOneCustomer() {
	showCustomer()
	fmt.Println("\n" + strings.Repeat("=", 80))

	scanner := bufio.NewScanner(os.Stdin)
	defer CustomerManagement()
	defer utils.ErrorRecover()

	fmt.Print("Enter Customer ID you want to update : ")
	scanner.Scan()
	id := scanner.Text()

	customer, err := master.FindCustomerById(id)
	if err != nil {
		panic("Customer ID not found!")
	} else {
		var updatedCustomer entity.Customer
		fmt.Println("(Leave it blank if you don't want to change)")
		fmt.Print("Enter Customer Name : ")
		scanner.Scan()
		updatedCustomer.Name = scanner.Text()

		fmt.Print("Enter Customer Phone : ")
		scanner.Scan()
		updatedCustomer.Contact = scanner.Text()

		fmt.Print("\nAre you sure want to update this customer (y/n)? ")
		scanner.Scan()
		confirm := scanner.Text()

		if strings.ToLower(confirm) == "y" {
			if updatedCustomer.Name != "" {
				customer.Name = updatedCustomer.Name
			}
			if updatedCustomer.Contact != "" {
				customer.Contact = updatedCustomer.Contact
			}
			master.UpdateCustomer(customer)
		} else if strings.ToLower(confirm) == "n" {
			fmt.Println("Customer Data was not saved.")
		} else {
			fmt.Println("Input is invalid!")
		}
	}
}

func deleteOneCustomer() {
	showCustomer()
	fmt.Println("\n" + strings.Repeat("=", 80))

	scanner := bufio.NewScanner(os.Stdin)
	defer CustomerManagement()
	defer utils.ErrorRecover()

	fmt.Print("Enter Customer ID you want to delete : ")
	scanner.Scan()
	id := scanner.Text()

	_, err := master.FindCustomerById(id)
	if err != nil {
		panic("Customer ID not found!")
	} else {
		fmt.Print("\nAre you sure want to delete this customer (y/n)? ")
		scanner.Scan()
		confirm := scanner.Text()

		if strings.ToLower(confirm) == "y" {
			master.DeleteCustomer(id)
		} else if strings.ToLower(confirm) == "n" {
			fmt.Println("Customer Data was not deleted.")
		} else {
			fmt.Println("Input is invalid!")
		}
	}

}
