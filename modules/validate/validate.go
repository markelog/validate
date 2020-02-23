package validate

// Validator is a type for validators functions
type Validator func(string) *Result

// Validators is the list of all available validators
var Validators = map[string]map[string]Validator{
	"email": map[string]Validator{
		"smtp":       SMTP,
		"domain":     EmailDomain,
		"regexp":     Email,
		"reputation": Reputation,
		"dmarc":      DMARC,
	},
}

// Result response from the validators
type Result struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
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
func (validate Validate) Validate(field string) (bool, map[string]*Result) {
	result := map[string]*Result{}
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
