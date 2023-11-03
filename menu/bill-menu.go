package menu

import (
	"bufio"
	"challenge-godb/db"
	"challenge-godb/db/transaction"
	"challenge-godb/entity"
	"challenge-godb/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
		if len(billCustomerDetails) == 0 {
			fmt.Print("\nTransaction Detail is empty.\n\n")
		} else {
			fmt.Println(strings.Repeat("=", 80))
			fmt.Println(strings.Repeat("=", 32), "Enigma Laundry", strings.Repeat("=", 32))
			fmt.Println(strings.Repeat("=", 80))

			fmt.Println("\n"+strings.Repeat("-", 34), strings.Repeat(" ", 10), strings.Repeat("-", 34))
			fmt.Println("| No", strings.Repeat(" ", 15-len("No")), "|", billCustomerDetails[0].BillDetail.Bill_Id, strings.Repeat(" ", 10-len(strconv.Itoa(billCustomerDetails[0].Bill_Id))), "|", strings.Repeat(" ", 10), "| Nama Cust", strings.Repeat(" ", 10-len("Nama Cust")), "|", billCustomerDetails[0].Customer.Name, strings.Repeat(" ", 15-len(billCustomerDetails[0].Customer.Name)), "|")
			fmt.Println("| Tanggal Masuk", strings.Repeat(" ", 15-len("Tanggal Masuk")), "|", utils.FormattedDate(billCustomerDetails[0].Bill.EntryDate), strings.Repeat(" ", 10-len(utils.FormattedDate(billCustomerDetails[0].Bill.EntryDate))), "|", strings.Repeat(" ", 10), "| No HP", strings.Repeat(" ", 10-len("No HP")), "|", billCustomerDetails[0].Customer.Contact, strings.Repeat(" ", 15-len(billCustomerDetails[0].Customer.Contact)), "|")
			fmt.Println("| Tanggal Selesai", strings.Repeat(" ", 15-len("Tanggal Selesai")), "|", utils.FormattedDate(billCustomerDetails[0].Bill.OutDate), strings.Repeat(" ", 10-len(utils.FormattedDate(billCustomerDetails[0].Bill.OutDate))), "|", strings.Repeat(" ", 10), strings.Repeat("-", 34))
			fmt.Println("| Diterima Oleh", strings.Repeat(" ", 15-len("Diterima Oleh")), "|", billCustomerDetails[0].Bill.RecipientName, strings.Repeat(" ", 10-len(billCustomerDetails[0].Bill.RecipientName)), "|")
			fmt.Println(strings.Repeat("-", 34))

			fmt.Println("\n\nNo", "  ", "Pelayanan", strings.Repeat(" ", 20), "Jumlah", "Satuan", "  ", "Harga", strings.Repeat(" ", 10-len("Harga")), "Total")
			fmt.Println(strings.Repeat("-", 80))

			for i, billDetail := range billCustomerDetails {
				fmt.Println(i+1, strings.Repeat(" ", 4-len(strconv.Itoa(billDetail.BillDetail.Id))), billDetail.Service.Service, strings.Repeat(" ", 31-len(billDetail.Service.Service)), billDetail.BillDetail.Amount, strings.Repeat(" ", 3-len(strconv.Itoa(billDetail.BillDetail.Amount))), billDetail.Service.Unit, strings.Repeat(" ", 8-len(billDetail.Service.Unit)), utils.IntegerToRupiahFormatter(billDetail.Service.Price), strings.Repeat(" ", 10-len(utils.IntegerToRupiahFormatter(billDetail.Service.Price))), utils.IntegerToRupiahFormatter(billDetail.BillDetail.Total))
			}

			fmt.Println("\n\n", strings.Repeat(" ", 51), "Total Harga ", utils.IntegerToRupiahFormatter(billCustomerDetails[0].Bill.TotalBill))
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

func AddNewTransaction() {
	db := db.ConnectDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	newBillId := transaction.GetMaxIdBill(tx)
	newBill := entity.Bill{}
	newBillDetail := entity.BillDetail{}

	showCustomer()
	fmt.Println()

	fmt.Print("Enter Customer ID that making the transaction : ")
	scanner.Scan()
	newBill.Customer_Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Enter Entry Date (Now? y/n) : ")
	scanner.Scan()
	choice := scanner.Text()
	if strings.ToLower(choice) == "y" {
		newBill.EntryDate = time.Now()
	} else if strings.ToLower(choice) == "n" {
		fmt.Print("Enter Entry Date (yyyy-mm-dd) : ")
		scanner.Scan()
		newBill.EntryDate, _ = time.Parse("2006-01-02", scanner.Text())
	} else {
		fmt.Println("Input is invalid! Start over..")
		AddNewTransaction()
	}

	fmt.Print("Enter Out Date (yyyy-mm-dd) : ")
	scanner.Scan()
	newBill.OutDate, _ = time.Parse("2006-01-02", scanner.Text())

	fmt.Print("Enter Recipient Name : ")
	scanner.Scan()
	newBill.RecipientName = scanner.Text()

	transaction.AddNewBill(newBill, tx)

	err = tx.Commit()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Transaction Commited!")
	}

	for {
		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}

		showService()
		fmt.Println()

		newBillDetail.Bill_Id = newBillId

		fmt.Print("Enter Service ID : ")
		scanner.Scan()
		newBillDetail.Service_Id, _ = strconv.Atoi(scanner.Text())

		fmt.Print("Enter the Amount according to the Units written (number only) : ")
		scanner.Scan()
		newBillDetail.Amount, _ = strconv.Atoi(scanner.Text())

		newBillDetail.Total = transaction.GetTotalPriceService(newBillDetail.Service_Id, newBillDetail.Amount, tx)

		transaction.AddNewBillDetail(newBillDetail, tx)
		transaction.UpdateBill(newBillId, tx)

		err = tx.Commit()

		if err != nil {
			panic(err)
		} else {
			fmt.Println("Transaction Commited!")
		}

		fmt.Print("Do you want to enter the transaction details again (y/n) ? ")
		scanner.Scan()
		repeat := scanner.Text()

		if strings.ToLower(repeat) == "n" {
			break
		}
	}
}
