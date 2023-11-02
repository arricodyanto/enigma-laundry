package entity

import "time"

type Transaction struct {
	Id            int
	Customer_Id   int
	EntryDate     time.Time
	OutDate       time.Time
	RecipientName string
	TotalBill     int
}
