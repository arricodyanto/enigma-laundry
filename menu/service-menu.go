package menu

import (
	"bufio"
	"challenge-godb/db/manipulation"
	"challenge-godb/entity"
	"challenge-godb/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ServiceManagement() {
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("\nSERVICE MANAGAMENT MENU:")
	fmt.Println("1. Show All Services")
	fmt.Println("2. Add New Service")
	fmt.Println("3. Update Service")
	fmt.Println("4. Delete Service")
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
		showService()
		ServiceManagement()
	case 2:
		addNewService()
	case 3:
	case 4:
	case 5:
		fmt.Println("<- Back to Main Menu")
	default:
		fmt.Printf("Menu '%d' does not exist\n", menu)
	}
}

func showService() {
	services := manipulation.GetAllService()
	fmt.Println("\nID", strings.Repeat(" ", 4), "Service", strings.Repeat(" ", 40), "Unit", strings.Repeat(" ", 4), "Price")
	fmt.Println(strings.Repeat("-", 80))
	for _, service := range services {
		idLength := len(strconv.Itoa(service.Id))
		fmt.Println(service.Id, strings.Repeat(" ", 6-idLength), service.Service, strings.Repeat(" ", 47-len(service.Service)), service.Unit, strings.Repeat(" ", 8-len(service.Unit)), utils.IntegerToRupiahFormatter(service.Price))
	}
}

func addNewService() {
	scanner := bufio.NewScanner(os.Stdin)
	newService := entity.Service{}
	defer ServiceManagement()
	defer utils.ErrorRecover()

	fmt.Print("Enter Service Name : ")
	scanner.Scan()
	newService.Service = scanner.Text()

	fmt.Print("Enter Service Unit (kg/buah) : ")
	scanner.Scan()
	newService.Unit = scanner.Text()

	fmt.Print("Enter Service Price : ")
	scanner.Scan()
	newService.Price, _ = strconv.Atoi(scanner.Text())

	fmt.Print("\nAre you sure want to add this new customer (y/n)? ")
	scanner.Scan()
	confirm := scanner.Text()

	if strings.ToLower(confirm) == "y" {
		manipulation.AddService(newService)
	} else if strings.ToLower(confirm) == "n" {
		fmt.Println("New Service Data was not saved.")
	} else {
		fmt.Println("Input is invalid!")
	}
}
