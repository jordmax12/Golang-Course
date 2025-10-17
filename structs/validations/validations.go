package validations

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func ValidateNonEmpty(value string) error {
	if value == "" {
		return fmt.Errorf("value cannot be empty")
	}
	return nil
}

func ValidateBirthdate(birthdate string) error {
	if birthdate == "" {
		return fmt.Errorf("birthdate cannot be empty")
	}
	
	_, err := time.Parse("01/02/2006", birthdate)
	if err != nil {
		return fmt.Errorf("invalid date format, use MM/DD/YYYY")
	}
	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	if !strings.Contains(email, "@") {
		return fmt.Errorf("email must contain @")
	}

	email = strings.TrimSpace(email)
	// rudimentary regex for basic email validation
	// Accepts most correct addresses but rejects most obviously incorrect ones
	re := `^[a-zA-Z0-9._%%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	matched := false
	if matched, _ = regexp.MatchString(re, email); !matched {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	hasSpecial := false
	specialChars := "!@#$%^&*()-_=+[]{}|;:',.<>/?`~"
	for _, c := range password {
		if strings.ContainsRune(specialChars, c) {
			hasSpecial = true
			break
		}
	}
	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}

	hasUpper := false
	for _, c := range password {
		if c >= 'A' && c <= 'Z' {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		return fmt.Errorf("password must contain at least one upper case letter")
	}

	hasNumber := false
	for _, c := range password {
		if c >= '0' && c <= '9' {
			hasNumber = true
			break
		}
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}
	return nil
}