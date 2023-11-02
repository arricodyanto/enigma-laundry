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
		updateOneService()
	case 4:
		deleteOneService()
	case 5:
		fmt.Println("<- Back to Main Menu")
	default:
		fmt.Printf("Menu '%d' does not exist\n", menu)
	}
}

func showService() {
	services := master.GetAllService()
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
		master.AddService(newService)
	} else if strings.ToLower(confirm) == "n" {
		fmt.Println("New Service Data was not saved.")
	} else {
		fmt.Println("Input is invalid!")
	}
}

func updateOneService() {
	showService()
	fmt.Println(strings.Repeat("=", 80))

	scanner := bufio.NewScanner(os.Stdin)
	defer ServiceManagement()
	defer utils.ErrorRecover()

	fmt.Print("Enter Service ID you want to update : ")
	scanner.Scan()
	id := scanner.Text()

	service, err := master.FindServiceById(id)
	if err != nil {
		panic("Service ID not found!")
	} else {
		var updatedService entity.Service
		fmt.Println("(Leave it blank if you don't want to change)")
		fmt.Print("Enter Service Name : ")
		scanner.Scan()
		updatedService.Service = scanner.Text()

		fmt.Print("Enter Service Unit (kg/buah) : ")
		scanner.Scan()
		updatedService.Unit = scanner.Text()

		fmt.Print("Enter Service Price : ")
		scanner.Scan()
		updatedService.Price, _ = strconv.Atoi(scanner.Text())

		fmt.Print("\nAre you sure want to update this service (y/n)? ")
		scanner.Scan()
		confirm := scanner.Text()

		if strings.ToLower(confirm) == "y" {
			if updatedService.Service != "" {
				service.Service = updatedService.Service
			}
			if updatedService.Unit != "" {
				service.Unit = updatedService.Unit
			}
			if updatedService.Price != 0 {
				service.Price = updatedService.Price
			}
			master.UpdateService(service)
		} else if strings.ToLower(confirm) == "n" {
			fmt.Println("Service Data was not saved.")
		} else {
			fmt.Println("Input is invalid!")
		}
	}
}

func deleteOneService() {
	showService()
	fmt.Println("\n" + strings.Repeat("=", 80))

	scanner := bufio.NewScanner(os.Stdin)
	defer ServiceManagement()
	defer utils.ErrorRecover()

	fmt.Print("Enter Service Id you want to delete : ")
	scanner.Scan()
	id := scanner.Text()

	_, err := master.FindServiceById(id)
	if err != nil {
		panic("Service ID not found!")
	} else {
		fmt.Print("\nAre you sure want to delete this service (y/n)? ")
		scanner.Scan()
		confirm := scanner.Text()

		if strings.ToLower(confirm) == "y" {
			master.DeleteService(id)
		} else if strings.ToLower(confirm) == "n" {
			fmt.Println("Service Data was not deleted.")
		} else {
			fmt.Println("Input is invalid!")
		}
	}
}
