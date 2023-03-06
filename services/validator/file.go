package validator

import "fmt"

func ValidateFile(file []byte) error {
	var err error

	fmt.Sprintf("%T\n", file)

	if err != nil {
		return err
	}
	return nil
}
