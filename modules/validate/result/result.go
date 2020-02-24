package result

// Result response from the validators
type Result struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}
