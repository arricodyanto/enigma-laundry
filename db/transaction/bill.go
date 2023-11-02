package transaction

import (
	"challenge-godb/db"
	"challenge-godb/entity"
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
