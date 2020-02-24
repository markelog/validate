package reputation

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var realHost = host

func teardown() {
	host = realHost
}

func TestSuccss(t *testing.T) {
	defer teardown()
	resp, _ := json.Marshal(response{false})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}))
	defer server.Close()

	host = server.URL

	result := Validate("markelog@gmail.com")

	assert.Equal(t, result.Valid, true)
}

func TestError(t *testing.T) {
	defer teardown()
	resp, _ := json.Marshal(response{true})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}))
	defer server.Close()

	host = server.URL

	result := Validate("markelog@gmail.com")
	assert.Equal(t, result.Valid, false)
}

func TestIncorrectJSONResponse(t *testing.T) {
	defer teardown()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("nope"))
	}))
	defer server.Close()

	host = server.URL

	result := Validate("markelog@gmail.com")
	assert.Nil(t, result)
}

func TestBadURL(t *testing.T) {
	defer teardown()

	host = ""

	result := Validate("markelog@gmail.com")
	assert.Nil(t, result)
}
