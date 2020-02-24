package tools

import (
	"errors"
	"strings"
)

// GetDomain returns the domain extract from the email address
func GetDomain(value string) (string, error) {
	email := strings.Split(value, "@")

	if len(email) != 2 {
		return "", errors.New("Couldn't parse the email")
	}

	return email[1], nil
}
