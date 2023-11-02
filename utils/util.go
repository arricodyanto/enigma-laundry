package utils

import "fmt"

func ErrorRecover() error {
	err := recover()
	if err != nil {
		fmt.Println("Terjadi panic:", err)
	}
	return nil
}
