package validator

import (
	"strconv"
)

func ValidateAirac(airac string) error {
	var err error

	_, err = strconv.ParseInt(airac, 10, 16)
	if err != nil {
		return err
	}
	return nil
}
