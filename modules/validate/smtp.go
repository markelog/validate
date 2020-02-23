package validate

import (
	"strings"

	"github.com/smancke/mailck"
)

// SMTP validates the mailbox by trying to establish connection to it
func SMTP(email string) *Result {
	result, _ := mailck.Check("noreply@google.com", email)

	// Make messages consistent
	message := strings.Replace(result.Message, ".", "", -1)

	switch {

	case result.IsValid():
		return &Result{Valid: true}

	case result.IsError():
		// Something went wrong in the smtp communication
		// we can't say for sure if the address is valid or not
		return &Result{Valid: false, Reason: message}

	case result.IsInvalid():
		// Invalid for some reason
		// the reason is contained in result.ResultDetail
		// or we can check for different reasons:
		switch result {
		// domain is invalid
		case mailck.InvalidDomain:
			return &Result{Valid: false, Reason: "Invalid TLD"}
		case mailck.InvalidSyntax:
			return &Result{Valid: false, Reason: "Incorrect email syntax"}
		default:
			return &Result{Valid: false, Reason: message}
		}
	}

	return &Result{Valid: true}
}
