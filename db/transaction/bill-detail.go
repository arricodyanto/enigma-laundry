package transaction

import (
	"challenge-godb/db"
	"challenge-godb/entity"
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

		err := rows.Scan(&billDetail.Id, &bill.Id, &customer.Id, &customer.Name, &customer.Contact, &service.Service, &billDetail.Amount, &service.Unit, &service.Price, &billDetail.Total, &bill.TotalBill, &bill.RecipientName, &bill.EntryDate, &bill.OutDate)
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

	sqlStatement := "SELECT d.id, b.id, c.id, c.name, c.contact, s.service, d.amount, s.unit, s.price, d.total, b.total_bill, b.recipient_name, b.entry_date, b.out_date FROM trx_bill_detail AS d JOIN trx_bill AS b ON b.id=d.bill_id JOIN mst_service AS s ON d.service_id=s.id JOIN mst_customer AS c ON b.customer_id=c.id WHERE c.id=$1;"

	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	billCustomerDetails := scanBillCustomerDetail(rows)
	return billCustomerDetails
}

// func GetBillCustomerById(id string) ([]BillCustomerDetail, error) {
// 	db := db.ConnectDB()
// 	defer db.Close()

// 	sqlStatement := "SELECT d.id, b.id, c.name, c.contact, s.service, d.amount, s.unit, d.total, b.total_bill, b.recipient_name, b.entry_date, b.out_date FROM trx_bill_detail AS d JOIN trx_bill AS b ON b.id=d.bill_id JOIN mst_service AS s ON d.service_id=s.id JOIN mst_customer AS c ON b.customer_id=c.id WHERE d.id = $1;"

// 	// rows, err := db.Query(sqlStatement, id)
// 	billDetail := entity.BillDetail{}
// 	bill := entity.Bill{}
// 	service := entity.Service{}
// 	customer := entity.Customer{}

// 	err := db.QueryRow(sqlStatement, id).Scan(&billDetail.Id, &bill.Id, &customer.Name, &customer.Contact, &service.Service, &billDetail.Amount, &service.Unit, &billDetail.Total, &bill.TotalBill, &bill.RecipientName, &bill.EntryDate, &bill.OutDate)

// 	billCustomerDetails := []BillCustomerDetail{}

// 	billCustomerDetail := BillCustomerDetail{
// 		BillDetail: billDetail,
// 		Bill:       bill,
// 		Service:    service,
// 		Customer:   customer,
// 	}
// 	billCustomerDetails = append(billCustomerDetails, billCustomerDetail)

// 	if err != nil {
// 		return billCustomerDetails, err
// 	}
// 	return billCustomerDetails, nil
// }
