package validator

import (
	"regexp"
)

func ValidateEmail(mailAddress string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	_, err := regexp.MatchString(emailPattern, mailAddress)
	return err == nil
}

func ValidatePassword(password string) bool {
	passwordPattern := `^[A-Za-z\d]{8,}$`
	_, err := regexp.MatchString(passwordPattern, password)
	return err == nil
}

func ValidatePhone(phoneNumber string) bool {
	if len(phoneNumber) > 13 || len(phoneNumber) < 12 || phoneNumber[:1] != "0" {
		return false
	}

	for _, char := range phoneNumber {
		if char < '0' || char > '9' || char == '+' {
			return false
		}
	}
	return true
}
