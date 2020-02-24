package domain

import (
	"net"

	"github.com/markelog/validate/modules/validate/result"
	"github.com/markelog/validate/modules/validate/tools"
)

var lookup = net.LookupIP

// Validate validates the domain of the email
func Validate(email string) *result.Result {
	domain, err := tools.GetDomain(email)
	if err != nil {
		return &result.Result{
			Valid:  false,
			Reason: err.Error(),
		}
	}

	_, err = lookup(domain)
	if err != nil {
		return &result.Result{
			Valid:  false,
			Reason: err.Error(),
		}
	}

	return &result.Result{
		Valid: true,
	}
}
