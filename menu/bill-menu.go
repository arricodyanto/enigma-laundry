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
	showBillDetail()
}

func showBillCustomer() {
	billCustomers := transaction.GetAllBillCustomer()
	fmt.Println("\nID", strings.Repeat(" ", 4), "Customer", strings.Repeat(" ", 15), "Phone", strings.Repeat(" ", 10), "Total Price", strings.Repeat(" ", 3), "Entry Date", " ", "Out Date", "  ", "Recipient Name")
	fmt.Println(strings.Repeat("-", 105))
	for _, billCustomer := range billCustomers {
		fmt.Println(billCustomer.Bill.Id, strings.Repeat(" ", 6-len(strconv.Itoa(billCustomer.Bill.Id))), billCustomer.Customer.Name, strings.Repeat(" ", 23-len(billCustomer.Customer.Name)), billCustomer.Customer.Contact, strings.Repeat(" ", 15-len(billCustomer.Customer.Contact)), utils.IntegerToRupiahFormatter(billCustomer.Bill.TotalBill), strings.Repeat(" ", 14-len(utils.IntegerToRupiahFormatter(billCustomer.Bill.TotalBill))), utils.FormattedDate(billCustomer.Bill.EntryDate), strings.Repeat(" ", 11-len(utils.FormattedDate(billCustomer.Bill.EntryDate))), utils.FormattedDate(billCustomer.Bill.OutDate), strings.Repeat(" ", 10-len(utils.FormattedDate(billCustomer.Bill.OutDate))), billCustomer.Bill.RecipientName)
	}
}

func showBillDetail() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(strings.Repeat("=", 105))
	fmt.Print("Show Transaction Details (Input ID) / Enter 0 to Back : ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	if id == 0 {
		fmt.Println("<<- Back to Main Menu")
		ShowMenu()
	} else {
		billCustomerDetails := transaction.GetBillDetailsByCustomerId(strconv.Itoa(id))

		for _, billCustomerDetail := range billCustomerDetails {
			fmt.Println(billCustomerDetail)
			fmt.Println(strings.Repeat("=", 80))
			fmt.Println(strings.Repeat("=", 32), "Enigma Laundry", strings.Repeat("=", 32))
			fmt.Println(strings.Repeat("=", 80))

			fmt.Println("\n"+strings.Repeat("-", 34), strings.Repeat(" ", 10), strings.Repeat("-", 34))
			fmt.Println("| No", strings.Repeat(" ", 15-len("No")), "|", billCustomerDetail.Bill.Id, strings.Repeat(" ", 10-len(strconv.Itoa(billCustomerDetail.Bill_Id))), "|", strings.Repeat(" ", 10), "| Nama Cust", strings.Repeat(" ", 10-len("Nama Cust")), "|", billCustomerDetail.Customer.Name, strings.Repeat(" ", 15-len(billCustomerDetail.Customer.Name)), "|")
			fmt.Println("| Tanggal Masuk", strings.Repeat(" ", 15-len("Tanggal Masuk")), "|", utils.FormattedDate(billCustomerDetail.Bill.EntryDate), strings.Repeat(" ", 10-len(utils.FormattedDate(billCustomerDetail.Bill.EntryDate))), "|", strings.Repeat(" ", 10), "| No HP", strings.Repeat(" ", 10-len("No HP")), "|", billCustomerDetail.Customer.Contact, strings.Repeat(" ", 15-len(billCustomerDetail.Customer.Contact)), "|")
			fmt.Println("| Tanggal Selesai", strings.Repeat(" ", 15-len("Tanggal Selesai")), "|", utils.FormattedDate(billCustomerDetail.Bill.OutDate), strings.Repeat(" ", 10-len(utils.FormattedDate(billCustomerDetail.Bill.OutDate))), "|", strings.Repeat(" ", 10), strings.Repeat("-", 34))
			fmt.Println("| Diterima Oleh", strings.Repeat(" ", 15-len("Diterima Oleh")), "|", billCustomerDetail.Bill.RecipientName, strings.Repeat(" ", 10-len(billCustomerDetail.Bill.RecipientName)), "|")
			fmt.Println(strings.Repeat("-", 34))

			fmt.Println("\n\nNo", "  ", "Pelayanan", strings.Repeat(" ", 20), "Jumlah", "Satuan", "  ", "Harga", strings.Repeat(" ", 10-len("Harga")), "Total")
			fmt.Println(strings.Repeat("-", 80))

			for _, billDetail := range billCustomerDetails {
				fmt.Println(billDetail.BillDetail.Id, strings.Repeat(" ", 2-len(strconv.Itoa(billDetail.BillDetail.Id))), billDetail.Service.Service, strings.Repeat(" ", 33-len(billDetail.Service.Service)), billDetail.BillDetail.Amount, strings.Repeat(" ", 3-len(strconv.Itoa(billDetail.BillDetail.Amount))), billDetail.Service.Unit, strings.Repeat(" ", 8-len(billDetail.Service.Unit)), utils.IntegerToRupiahFormatter(billDetail.Service.Price), strings.Repeat(" ", 10-len(utils.IntegerToRupiahFormatter(billDetail.Service.Price))), utils.IntegerToRupiahFormatter(billDetail.BillDetail.Total))
			}

			fmt.Println("\n\n", strings.Repeat(" ", 51), "Total Harga ", utils.IntegerToRupiahFormatter(billCustomerDetail.Bill.TotalBill))
			fmt.Print(strings.Repeat("-", 80), "\n", strings.Repeat("=", 80), "\n\n\n")
		}
	}

	fmt.Print("Enter '0' for the previous menu or enter anything else to return to the main menu : ")
	scanner.Scan()
	choice := scanner.Text()

	if choice == "0" {
		fmt.Println("<- Back to All Transaction Menu")
		ShowAllTransaction()
	} else {
		fmt.Println("<<- Back to Main Menu")
	}
}
