package dmarc

import (
	"errors"
	"testing"

	"github.com/emersion/go-msgauth/dmarc"
	"github.com/stretchr/testify/assert"
)

var realLookup = lookup

func teardown() {
	lookup = realLookup
}

func TestSuccess(t *testing.T) {
	defer teardown()

	lookup = func(domain string) (*dmarc.Record, error) {
		return nil, nil
	}

	result := Validate("markelog@gmail.com")

	assert.Equal(t, result.Valid, true)
}

func TestError(t *testing.T) {
	defer teardown()

	lookup = func(domain string) (*dmarc.Record, error) {
		return nil, errors.New("Something went wrong")
	}

	result := Validate("markelog@gmail.com")

	assert.Equal(t, result.Valid, false)
}
