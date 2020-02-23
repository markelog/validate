package validate

import (
	"github.com/emersion/go-msgauth/dmarc"
)

// DMARC checks presence of DMARC DNS TXT records
func DMARC(email string) *Result {
	domain, err := getDomain(email)
	if err != nil {
		return &Result{
			Valid:  false,
			Reason: err.Error(),
		}
	}

	_, err = dmarc.Lookup(domain)
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
