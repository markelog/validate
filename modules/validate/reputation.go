package validate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Third-party service we use to determine reputation
// Be careful, amount of requests is limited
const host = "https://emailrep.io"

var client = &http.Client{Timeout: 10 * time.Second}

type response struct {
	Suspicious bool `json:"suspicious"`
}

// Reputation validates emails reputation
func Reputation(email string) *Result {
	var data response
	resp, err := client.Get(fmt.Sprintf("%s/%s", host, email))
	// In case remote service is unavailable - do not show anything
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	// Can't parse the remote response - bail
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil
	}

	if data.Suspicious {
		return &Result{
			Valid:  false,
			Reason: "Suspicious email",
		}
	}

	return &Result{
		Valid: true,
	}
}
