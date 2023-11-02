package menu

import (
	"bufio"
	"challenge-godb/db/transaction"
	"challenge-godb/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowAllTransaction() {
	showBillCustomer()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(strings.Repeat("=", 105))
	fmt.Print("Show Transaction Details (Input ID) / Enter 0 to Back : ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	if id == 0 {
		fmt.Println("<- Back to Main Menu")
	} else {

	}
}

func showBillCustomer() {
	billCustomers := transaction.GetAllBillCustomer()
	fmt.Println("\nID", strings.Repeat(" ", 4), "Customer", strings.Repeat(" ", 15), "Phone", strings.Repeat(" ", 10), "Total Price", strings.Repeat(" ", 3), "Entry Date", " ", "Out Date", "  ", "Recipient Name")
	fmt.Println(strings.Repeat("-", 105))
	for _, billCustomer := range billCustomers {
		fmt.Println(billCustomer.Bill.Id, strings.Repeat(" ", 6-len(strconv.Itoa(billCustomer.Bill.Id))), billCustomer.Customer.Name, strings.Repeat(" ", 23-len(billCustomer.Customer.Name)), billCustomer.Customer.Contact, strings.Repeat(" ", 15-len(billCustomer.Customer.Contact)), utils.IntegerToRupiahFormatter(billCustomer.Bill.TotalBill), strings.Repeat(" ", 14-len(utils.IntegerToRupiahFormatter(billCustomer.Bill.TotalBill))), utils.FormattedDate(billCustomer.Bill.EntryDate), strings.Repeat(" ", 11-len(utils.FormattedDate(billCustomer.Bill.EntryDate))), utils.FormattedDate(billCustomer.Bill.OutDate), strings.Repeat(" ", 10-len(utils.FormattedDate(billCustomer.Bill.OutDate))), billCustomer.Bill.RecipientName)
	}
}
