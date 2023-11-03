package transaction

import (
	"challenge-godb/db"
	"challenge-godb/entity"
	"challenge-godb/utils"
	"database/sql"
)

type BillCustomer struct {
	entity.Bill
	entity.Customer
}

func scanBillCustomer(rows *sql.Rows) []BillCustomer {
	billCustomers := []BillCustomer{}

	for rows.Next() {
		bill := entity.Bill{}
		customer := entity.Customer{}
		err := rows.Scan(&bill.Id, &customer.Name, &customer.Contact, &bill.TotalBill, &bill.EntryDate, &bill.OutDate, &bill.RecipientName)
		if err != nil {
			panic(err)
		}
		billCustomer := BillCustomer{
			Bill:     bill,
			Customer: customer,
		}
		billCustomers = append(billCustomers, billCustomer)
	}
	err := rows.Err()
	if err != nil {
		panic(err)
	}
	return billCustomers
}

func GetAllBillCustomer() []BillCustomer {
	db := db.ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT b.id, c.name, c.contact, b.total_bill, b.entry_date, b.out_date, b.recipient_name FROM trx_bill AS b JOIN mst_customer AS c ON b.customer_id=c.id;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	billCustomers := scanBillCustomer(rows)
	return billCustomers
}

func AddNewBill(bill entity.Bill, tx *sql.Tx) {
	newId := GetMaxIdBill(tx)
	insertNewBill(newId, bill, tx)
}

func GetMaxIdBill(tx *sql.Tx) int {
	sqlStatement := "SELECT MAX(id) FROM trx_bill;"

	maxId := 0
	err := tx.QueryRow(sqlStatement).Scan(&maxId)

	utils.Validate(err, "Get New ID for Bill", tx)
	return maxId + 1
}

func insertNewBill(id int, bill entity.Bill, tx *sql.Tx) {
	sqlStatement := "INSERT INTO trx_bill (id, customer_id, entry_date, out_date, recipient_name) VALUES ($1, $2, $3, $4, $5);"

	_, err := tx.Exec(sqlStatement, id, bill.Customer_Id, bill.EntryDate, bill.OutDate, bill.RecipientName)
	utils.Validate(err, "Saved New Bill", tx)
}

func UpdateBill(id int, tx *sql.Tx) {
	totalBill := getTotalBill(id, tx)
	updateTotalBill(id, totalBill, tx)
}

func getTotalBill(bill_id int, tx *sql.Tx) int {
	sqlStatement := "SELECT SUM(total) FROM trx_bill_detail WHERE bill_id = $1;"

	totalBill := 0
	err := tx.QueryRow(sqlStatement, bill_id).Scan(&totalBill)
	utils.Validate(err, "Get Total Bill From Customer", tx)
	return totalBill
}

func updateTotalBill(id int, totalBill int, tx *sql.Tx) {
	sqlStatement := "UPDATE trx_bill SET total_bill = $2 WHERE id = $1;"

	_, err := tx.Exec(sqlStatement, id, totalBill)
	utils.Validate(err, "Updated Total Bill From Customer", tx)
}
