package manipulation

import (
	"challenge-godb/db"
	"challenge-godb/entity"
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
	var err error

	sqlStatement := "INSERT INTO mst_customer (name, contact) VALUES($1, $2);"

	_, err = db.Exec(sqlStatement, customer.Name, customer.Contact)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Saved New Customer Data!")
	}
}

func UpdateCustomer(customer entity.Customer) {
	db := db.ConnectDB()
	defer db.Close()

	sqlStatement := "UPDATE mst_customer SET name = $2, contact = $3 WHERE id = $1;"

	_, err := db.Exec(sqlStatement, customer.Id, customer.Name, customer.Contact)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfull Update Data!")
	}
}

func DeleteCustomer(id string) {
	db := db.ConnectDB()
	defer db.Close()

	sqlStatement := "DELETE FROM mst_customer WHERE id = $1;"

	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Deleted One Customer!")
	}
}
