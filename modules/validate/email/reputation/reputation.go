package reputation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/markelog/validate/modules/validate/result"
)

var (
	// Third-party service we use to determine reputation
	// Be careful, amount of requests is limited
	host = "https://emailrep.io"

	client = &http.Client{Timeout: 10 * time.Second}
)

type response struct {
	Suspicious bool `json:"suspicious"`
}

// Validate validates emails reputation
func Validate(email string) *result.Result {
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
		return &result.Result{
			Valid:  false,
			Reason: "Suspicious email",
		}
	}

	return &result.Result{
		Valid: true,
	}
}
