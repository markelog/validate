package smtp

import (
	"testing"

	"github.com/smancke/mailck"
	"github.com/stretchr/testify/assert"
)

var realCheck = checker

func teardown() {
	checker = realCheck
}

func TestSuccess(t *testing.T) {
	defer teardown()
	checker = func(a, b string) (mailck.Result, error) {
		return mailck.Result{
			Result: mailck.ValidState,
		}, nil
	}

	res := Validate("markelog@gmail.com")

	assert.True(t, res.Valid)
}

func TestBadSyntax(t *testing.T) {
	res := Validate("nope")

	assert.False(t, res.Valid)
	assert.Equal(t, "Incorrect email syntax", res.Reason)
}

func TestError(t *testing.T) {
	defer teardown()
	checker = func(a, b string) (mailck.Result, error) {
		return mailck.Result{
			Result:  mailck.InvalidState,
			Message: "sup.",
		}, nil
	}

	res := Validate("markelog@gmail.com")

	assert.False(t, res.Valid)
	assert.Equal(t, "sup", res.Reason)
}

func TestInvalidDomain(t *testing.T) {
	defer teardown()
	checker = func(a, b string) (mailck.Result, error) {
		return mailck.InvalidDomain, nil
	}

	res := Validate("markelog@gmail.com")

	assert.False(t, res.Valid)
	assert.Equal(t, "Invalid TLD", res.Reason)
}
