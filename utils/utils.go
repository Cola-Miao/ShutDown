package utils

import "fmt"

func RecordError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
