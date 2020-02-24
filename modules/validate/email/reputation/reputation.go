package reputation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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
	Suspicious bool   `json:"suspicious"`
	Status     string `json:"status"`
	Reason     string `json:"reason"`
}

// Validate validates emails reputation
func Validate(email string) *result.Result {
	var (
		data       response
		isDisabled = os.Getenv("EMAIL_REP_DISABLE")
		key        = os.Getenv("EMAIL_REP_KEY")
	)

	if isDisabled == "true" {
		return nil
	}

	url := fmt.Sprintf("%s/%s", host, email)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}

	if key != "" {
		req.Header.Add("KEY", key)
	}

	resp, err := client.Do(req)

	// In case remote service is unavailable - do not show anything
	if err != nil {
		return nil
	}

	// Means our limit is exceeded
	if resp.StatusCode == 429 {
		return nil
	}

	defer resp.Body.Close()

	// Can't parse the remote response - bail
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil
	}

	if data.Status == "fail" {
		return &result.Result{
			Valid: false,
			Reason: fmt.Sprintf(
				"%s%s",
				strings.Title(data.Reason[0:1]),
				data.Reason[1:],
			),
		}
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
