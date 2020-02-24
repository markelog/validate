package smtp

import (
	"strings"

	"github.com/markelog/validate/modules/validate/result"
	"github.com/smancke/mailck"
)

var checker = mailck.Check

// Validate validates the mailbox by trying to establish connection to it
func Validate(email string) *result.Result {
	check, _ := checker("noreply@google.com", email)

	// Make messages consistent
	message := strings.Replace(check.Message, ".", "", -1)

	switch {

	case check.IsValid():
		return &result.Result{Valid: true}

	case check.IsError():
		// Something went wrong in the smtp communication
		// we can't say for sure if the address is valid or not
		return &result.Result{Valid: false, Reason: message}

	case check.IsInvalid():
		// Invalid for some reason
		// the reason is contained in result.result.ResultDetail
		// or we can check for different reasons:
		switch check {
		// domain is invalid
		case mailck.InvalidDomain:
			return &result.Result{Valid: false, Reason: "Invalid TLD"}
		case mailck.InvalidSyntax:
			return &result.Result{Valid: false, Reason: "Incorrect email syntax"}
		default:
			return &result.Result{Valid: false, Reason: message}
		}
	}

	return &result.Result{Valid: true}
}
