package utils

import "fmt"

func GetKey() int {
	var key int

	_, err := fmt.Scanln(&key)
	RecordError(err)

	return key
}
