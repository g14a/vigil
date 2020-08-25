package utils

import (
	"fmt"
)

func CheckError(err error) {
	if err != nil {
		fmt.Printf("%v", err)
	}
}
