package utils

import (
	"fmt"
	"strconv"
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
