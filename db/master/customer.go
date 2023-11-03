package master

import (
	"challenge-godb/db"
	"challenge-godb/entity"
	"challenge-godb/utils"
	"database/sql"
	"fmt"
)

func scanCustomer(rows *sql.Rows) []entity.Customer {
	customers := []entity.Customer{}
	var err error

	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Contact)
		if err != nil {
			panic(err)
		}
		customers = append(customers, customer)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return customers
}

func FindCustomerById(id string) (entity.Customer, error) {
	db := db.ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_customer WHERE id = $1;"

	customer := entity.Customer{}
	err := db.QueryRow(sqlStatement, id).Scan(&customer.Id, &customer.Name, &customer.Contact)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func GetAllCustomer() []entity.Customer {
	db := db.ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_customer;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := scanCustomer(rows)
	return customers
}

func AddCustomer(customer entity.Customer) {
	db := db.ConnectDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	newId := getMaxIdCustomer(tx)
	insertNewCustomer(newId, customer, tx)

	err = tx.Commit()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Transaction Commited!")
	}
}

func getMaxIdCustomer(tx *sql.Tx) int {
	sqlStatement := "SELECT MAX(id) FROM mst_customer;"

	maxId := 0
	err := tx.QueryRow(sqlStatement).Scan(&maxId)

	utils.Validate(err, "Get New ID for Customer", tx)
	return maxId + 1
}

func insertNewCustomer(id int, customer entity.Customer, tx *sql.Tx) {
	sqlStatement := "INSERT INTO mst_customer (id, name, contact) VALUES($1, $2, $3);"

	_, err := tx.Exec(sqlStatement, id, customer.Name, customer.Contact)
	utils.Validate(err, "Saved New Customer", tx)
}

func UpdateCustomer(customer entity.Customer) {
	db := db.ConnectDB()
	defer db.Close()

	sqlStatement := "UPDATE mst_customer SET name = $2, contact = $3 WHERE id = $1;"

	_, err := db.Exec(sqlStatement, customer.Id, customer.Name, customer.Contact)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

func DeleteCustomer(id string) {
	db := db.ConnectDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	deleteOneCustomer(id, tx)

	err = tx.Commit()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Transaction Commited!")
	}

}

func deleteOneCustomer(id string, tx *sql.Tx) {
	sqlStatement := "DELETE FROM mst_customer WHERE id = $1;"

	_, err := tx.Exec(sqlStatement, id)
	utils.Validate(err, "Deleted One Customer", tx)
}
