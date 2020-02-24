package tools_test

import (
	"testing"

	"github.com/markelog/validate/modules/validate/tools"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	domain, err := tools.GetDomain("markelog@gmail.com")

	assert.Equal(t, domain, "gmail.com")
	assert.NoError(t, err)
}

func TestError(t *testing.T) {
	domain, err := tools.GetDomain("nope")

	assert.Equal(t, domain, "")
	assert.Error(t, err, "Couldn't parse the email")
}
