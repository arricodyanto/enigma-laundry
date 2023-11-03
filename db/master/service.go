package master

import (
	"challenge-godb/db"
	"challenge-godb/entity"
	"challenge-godb/utils"
	"database/sql"
	"fmt"
)

func scanService(rows *sql.Rows) []entity.Service {
	services := []entity.Service{}

	for rows.Next() {
		service := entity.Service{}
		err := rows.Scan(&service.Id, &service.Service, &service.Unit, &service.Price)
		if err != nil {
			panic(err)
		}
		services = append(services, service)
	}

	err := rows.Err()
	if err != nil {
		panic(err)
	}
	return services
}

func GetAllService() []entity.Service {
	db := db.ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_service;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	services := scanService(rows)
	return services
}

func FindServiceById(id string) (entity.Service, error) {
	db := db.ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_service WHERE id = $1;"

	service := entity.Service{}
	err := db.QueryRow(sqlStatement, id).Scan(&service.Id, &service.Service, &service.Unit, &service.Price)
	if err != nil {
		return service, err
	}
	return service, nil
}

func AddService(service entity.Service) {
	db := db.ConnectDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	newId := getMaxIdService(tx)
	insertNewService(newId, service, tx)

	err = tx.Commit()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Transaction Commited!")
	}
}

func getMaxIdService(tx *sql.Tx) int {
	sqlStatement := "SELECT MAX(id) FROM mst_service;"

	maxId := 0
	err := tx.QueryRow(sqlStatement).Scan(&maxId)

	utils.Validate(err, "Get New ID for Service", tx)
	return maxId + 1
}

func insertNewService(id int, service entity.Service, tx *sql.Tx) {
	sqlStatement := "INSERT INTO mst_service (id, service, unit, price) VALUES ($1, $2, $3, $4);"

	_, err := tx.Exec(sqlStatement, id, service.Service, service.Unit, service.Price)
	utils.Validate(err, "Saved New Service", tx)
}

func UpdateService(service entity.Service) {
	db := db.ConnectDB()
	defer db.Close()

	sqlStatement := "UPDATE mst_service SET service = $2, unit = $3, price = $4 WHERE id = $1;"

	_, err := db.Exec(sqlStatement, service.Id, service.Service, service.Unit, service.Price)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

func DeleteService(id string) {
	db := db.ConnectDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	billsId := getBillsIdByServiceId(id, tx)                   // mengambil semua bill_id yang terelasi dengan service yang dihapus
	updatedTotalBills := getUpdatedTotalBills(billsId, id, tx) // menghitung total_bill yang baru setelah service dihapus
	updateTotalBills(billsId, updatedTotalBills, tx)           // perbarui total_bill dengan data yang baru
	deleteOneService(id, tx)                                   // hapus servicenya

	err = tx.Commit()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Transaction Commited!")
	}
}

func deleteOneService(id string, tx *sql.Tx) {
	sqlStatement := "DELETE FROM mst_service WHERE id = $1;"

	_, err := tx.Exec(sqlStatement, id)
	utils.Validate(err, "Deleted One Service", tx)
}

func getBillsIdByServiceId(id string, tx *sql.Tx) []int {
	sqlStatement := "SELECT bill_id FROM trx_bill_detail WHERE service_id = $1;"

	rows, err := tx.Query(sqlStatement, id)
	utils.Validate(err, "Get Bills ID by Service ID", tx)
	defer rows.Close()

	billsId := scanNumbers(rows)
	return billsId
}

func scanNumbers(rows *sql.Rows) []int {
	numbers := []int{}

	for rows.Next() {
		number := 0
		err := rows.Scan(&number)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	err := rows.Err()
	if err != nil {
		panic(err)
	}
	return numbers
}

func getUpdatedTotalBills(billsId []int, serviceId string, tx *sql.Tx) []int {
	updatedTotalBills := []int{}
	for _, billId := range billsId {
		sqlStatement := "SELECT COALESCE(SUM(total), 0) AS total_bill FROM trx_bill_detail WHERE bill_id = $1 AND service_id != $2;"

		updatedTotalBill := 0
		err := tx.QueryRow(sqlStatement, billId, serviceId).Scan(&updatedTotalBill)
		utils.Validate(err, "Get Updated Total Bills", tx)

		updatedTotalBills = append(updatedTotalBills, updatedTotalBill)
	}
	return updatedTotalBills
}

func updateTotalBills(billsId []int, updatedTotalBills []int, tx *sql.Tx) {
	for i, billId := range billsId {
		updatedTotalBill := updatedTotalBills[i]
		sqlStatement := "UPDATE trx_bill SET total_bill = $1 WHERE id = $2;"
		_, err := tx.Exec(sqlStatement, updatedTotalBill, billId)
		utils.Validate(err, "Updated Total Bill", tx)
	}
}
