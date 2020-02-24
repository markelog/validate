package domain

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

var realLookup = lookup

func teardown() {
	lookup = realLookup
}

func TestSuccess(t *testing.T) {
	defer teardown()

	lookup = func(host string) ([]net.IP, error) {
		return nil, nil
	}

	result := Validate("markelog@gmail.com")

	assert.Equal(t, result.Valid, true)
}

func TestError(t *testing.T) {
	defer teardown()

	lookup = func(host string) ([]net.IP, error) {
		return nil, errors.New("Something went wrong")
	}

	result := Validate("markelog@gmail.com")

	assert.Equal(t, result.Valid, false)
}
