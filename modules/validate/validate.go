package validate

import (
	"github.com/markelog/validate/modules/validate/email/dmarc"
	"github.com/markelog/validate/modules/validate/email/domain"
	"github.com/markelog/validate/modules/validate/email/email"
	"github.com/markelog/validate/modules/validate/email/reputation"
	"github.com/markelog/validate/modules/validate/email/smtp"
	"github.com/markelog/validate/modules/validate/result"
)

// Validator is a type for validators functions
type Validator func(string) *result.Result

// Validators is the list of all available validators
var Validators = map[string]map[string]Validator{
	"email": {
		"smtp":       smtp.Validate,
		"domain":     domain.Validate,
		"regexp":     email.Validate,
		"reputation": reputation.Validate,
		"dmarc":      dmarc.Validate,
	},
}

// Validate is the main struct representing the validator module
type Validate struct {
	value string
}

// New returns the new Validate struct
func New(value string) *Validate {
	return &Validate{
		value: value,
	}
}

// Validate runs all the validations
func (validate Validate) Validate(field string) (bool, map[string]*result.Result) {
	result := map[string]*result.Result{}
	valid := true

	for name, validator := range Validators[field] {
		res := validator(validate.value)

		if valid {
			valid = res.Valid
		}

		if res == nil {
			continue
		}

		result[name] = res
	}

	return valid, result
}
