package validate

import (
	"net/mail"
)

// Email validates the email syntax
func Email(email string) *Result {

	if _, err := mail.ParseAddress(email); err != nil {
		return &Result{
			Valid: false,
		}
	}

	return &Result{
		Valid: true,
	}
}
