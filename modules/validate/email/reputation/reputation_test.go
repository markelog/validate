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
	resp, _ := json.Marshal(response{Suspicious: false})

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
	resp, _ := json.Marshal(response{Suspicious: true})

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

func TestExceedLimit(t *testing.T) {
	defer teardown()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("nope"))
	}))
	defer server.Close()

	host = server.URL

	result := Validate("markelog@gmail.com")
	assert.Nil(t, result)
}

func TestIncorrectFailedResponse(t *testing.T) {
	defer teardown()

	resp, _ := json.Marshal(response{Status: "fail", Reason: "test"})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}))
	defer server.Close()

	host = server.URL

	result := Validate("markelog@gmail.com")
	assert.Equal(t, result.Valid, false)
	assert.Equal(t, result.Reason, "Test")
}
