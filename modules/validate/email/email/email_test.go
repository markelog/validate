package email_test

import (
	"testing"

	"github.com/markelog/validate/modules/validate/email/email"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	result := email.Validate("markelog@gmail.com")

	assert.Equal(t, result.Valid, true)
}

func TestError(t *testing.T) {
	result := email.Validate("nope")

	assert.Equal(t, result.Valid, false)
}
