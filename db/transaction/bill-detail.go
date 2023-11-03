package transaction

import (
	"challenge-godb/db"
	"challenge-godb/entity"
	"challenge-godb/utils"
	"database/sql"
)

type BillCustomerDetail struct {
	entity.BillDetail
	entity.Bill
	entity.Service
	entity.Customer
}

func scanBillCustomerDetail(rows *sql.Rows) []BillCustomerDetail {
	billCustomerDetails := []BillCustomerDetail{}

	for rows.Next() {
		billDetail := entity.BillDetail{}
		bill := entity.Bill{}
		service := entity.Service{}
		customer := entity.Customer{}

		err := rows.Scan(&billDetail.Id, &billDetail.Bill_Id, &customer.Id, &customer.Name, &customer.Contact, &service.Service, &billDetail.Amount, &service.Unit, &service.Price, &billDetail.Total, &bill.TotalBill, &bill.RecipientName, &bill.EntryDate, &bill.OutDate)
		if err != nil {
			panic(err)
		}
		billCustomerDetail := BillCustomerDetail{
			BillDetail: billDetail,
			Bill:       bill,
			Service:    service,
			Customer:   customer,
		}
		billCustomerDetails = append(billCustomerDetails, billCustomerDetail)
	}

	err := rows.Err()
	if err != nil {
		panic(err)
	}
	return billCustomerDetails
}

func GetBillDetailsByCustomerId(id string) []BillCustomerDetail {
	db := db.ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT d.id, d.bill_id, c.id, c.name, c.contact, s.service, d.amount, s.unit, s.price, d.total, b.total_bill, b.recipient_name, b.entry_date, b.out_date FROM trx_bill_detail AS d JOIN trx_bill AS b ON b.id=d.bill_id JOIN mst_service AS s ON d.service_id=s.id JOIN mst_customer AS c ON b.customer_id=c.id WHERE d.bill_id=$1;"

	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	billCustomerDetails := scanBillCustomerDetail(rows)
	return billCustomerDetails
}

func AddNewBillDetail(billDetail entity.BillDetail, tx *sql.Tx) {
	newId := getMaxIdBillDetail(tx)
	insertNewBillDetail(newId, billDetail, tx)
}

func getMaxIdBillDetail(tx *sql.Tx) int {
	sqlStatement := "SELECT MAX(id) FROM trx_bill_detail;"

	maxId := 0
	err := tx.QueryRow(sqlStatement).Scan(&maxId)

	utils.Validate(err, "Get New ID for Bill Detail", tx)
	return maxId + 1
}

func insertNewBillDetail(id int, billDetail entity.BillDetail, tx *sql.Tx) {
	sqlStatement := "INSERT INTO trx_bill_detail (id, bill_id, service_id, amount, total) VALUES ($1, $2, $3, $4, $5);"

	_, err := tx.Exec(sqlStatement, id, billDetail.Bill_Id, billDetail.Service_Id, billDetail.Amount, billDetail.Total)
	utils.Validate(err, "Saved New Bill Detail", tx)
}

func GetTotalPriceService(id int, amount int, tx *sql.Tx) int {
	sqlStatement := "SELECT price FROM mst_service WHERE id = $1;"

	price := 0
	err := tx.QueryRow(sqlStatement, id).Scan(&price)
	utils.Validate(err, "Get The Price Service", tx)

	totalPrice := amount * price
	return totalPrice
}
