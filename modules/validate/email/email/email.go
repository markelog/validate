package email

import (
	"net/mail"

	"github.com/markelog/validate/modules/validate/result"
)

// Validate validates the email syntax
func Validate(email string) *result.Result {
	if _, err := mail.ParseAddress(email); err != nil {
		return &result.Result{
			Valid: false,
		}
	}

	return &result.Result{
		Valid: true,
	}
}
