package master

import (
	"challenge-godb/db"
	"challenge-godb/entity"
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

	sqlStatement := "INSERT INTO mst_service (service, unit, price) VALUES ($1, $2, $3);"

	_, err := db.Exec(sqlStatement, service.Service, service.Unit, service.Price)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Saved New Service!")
	}
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

	sqlStatement := "DELETE FROM mst_service WHERE id = $1;"

	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Deleted One Service Data!")
	}
}
