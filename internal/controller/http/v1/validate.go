package v1

import (
	"fmt"
	"log"
	"net/mail"
	"unicode/utf8"
)

func validateLogin(email, password string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Printf("invalid email format: %s", err.Error())
		return fmt.Errorf("invalid email format: %w", err)
	}
	if utf8.RuneCountInString(password) < 4 {
		log.Println("invalid password")
		return fmt.Errorf("invalid passowrd")
	}
	return nil
}
