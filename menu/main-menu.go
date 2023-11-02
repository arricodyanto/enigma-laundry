package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowMenu() {
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println(strings.Repeat("=", 27), "Enigma Laundy Management", strings.Repeat("=", 27))
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("\n1. Customer Management")
	fmt.Println("2. Service Management")
	fmt.Println("3. Show All Transaction")
	fmt.Println("4. Insert New Transaction")
	fmt.Println("5. Exit")
	fmt.Println("\n" + strings.Repeat("=", 80))

	var menu int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Select the menu you want to use : ")
	scanner.Scan()
	menu, _ = strconv.Atoi(scanner.Text())

	switch menu {
	case 1:
		CustomerManagement()
	case 5:
		exitProgram()
	default:
		fmt.Printf("Menu '%d' does not exist\n", menu)
	}
}

func exitProgram() {
	fmt.Println("Closing the program..")
	os.Exit(0)
}
