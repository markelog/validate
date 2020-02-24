package dmarc

import (
	"github.com/emersion/go-msgauth/dmarc"
	"github.com/markelog/validate/modules/validate/result"
	"github.com/markelog/validate/modules/validate/tools"
)

var lookup = dmarc.Lookup

// Validate checks presence of DMARC DNS TXT records
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
