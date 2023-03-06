package validator

import (
	"fmt"
	"strconv"
)

func ValidateAirac(airac string) error {
	var err error

	_, err = strconv.ParseInt(airac, 10, 16)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
