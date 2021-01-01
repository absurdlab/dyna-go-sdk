package types

import "encoding/json"

// UsernameForDetailRequest is the payload for username subject query. It serves for
// both Dyna-Identity communication and UI-Dyna communication
type UsernameForDetailRequest struct {
	Username string `json:"username"`
}

// UsernameForDetailResponse is the response payload for UsernameForSubjectRequest. It serves for
// both Dyna-Identity communication and UI-Dyna communication.
type UsernameForDetailResponse struct {
	RPCResponse
	// Subject is the identifier for the user.
	Subject string `json:"subject,omitempty"`
	// Detail is the UI related data that will be relayed to the UI.
	Detail json.RawMessage `json:"detail,omitempty"`
}

// ClaimRequest is the payload for requesting claims. It serves for both Dyna-Identity and
// Tiga-Dyna communication.
type ClaimRequest struct {
	// Subject is the resolved user identifier
	Subject string `json:"subject"`
	// Claims is the list of claim names.
	Claims []string `json:"claims"`
}

// ClaimResponse responds to ClaimRequest. It serves for both Dyna-Identity and Tiga-Dyna communication.
type ClaimResponse struct {
	RPCResponse
	// Data contains the claim response data. It may or may not
	// contain all the requested claims.
	Data map[string]interface{} `json:"data"`
}
