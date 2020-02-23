package validate

import (
	"net"
)

// EmailDomain validates the domain of the email
func EmailDomain(email string) *Result {
	domain, err := getDomain(email)
	if err != nil {
		return &Result{
			Valid:  false,
			Reason: err.Error(),
		}
	}

	_, err = net.LookupIP(domain)
	if err != nil {
		return &Result{
			Valid:  false,
			Reason: err.Error(),
		}
	}

	return &Result{
		Valid: true,
	}
}
