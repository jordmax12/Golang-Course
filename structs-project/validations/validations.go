package validations

import (
	"fmt"
)

func ValidateNonEmpty(value string) error {
	if value == "" {
		return fmt.Errorf("value cannot be empty")
	}
	return nil
}
