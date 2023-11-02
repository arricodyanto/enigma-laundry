package entity

import (
	"time"
)

type Bill struct {
	Id            int
	Customer_Id   int
	EntryDate     time.Time
	OutDate       time.Time
	RecipientName string
	TotalBill     int
}
