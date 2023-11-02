package utils

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func ErrorRecover() error {
	err := recover()
	if err != nil {
		fmt.Println("Terjadi panic:", err)
	}
	return nil
}

func IntegerToRupiahFormatter(number int) string {
	stringNumber := strconv.Itoa(number)
	formatted := ""

	for i, char := range stringNumber {
		formatted += string(char)
		if (len(stringNumber)-i-1)%3 == 0 && i != len(stringNumber)-1 {
			formatted += "."
		}
	}
	return "Rp " + formatted
}

func Validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction has been rolled back")
	} else {
		fmt.Println("Successfully" + message + "data!")
	}
}

func FormattedDate(time time.Time) string {
	result := time.Format("02-01-2006")
	return result
}
