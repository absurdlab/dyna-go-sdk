package user

import (
	"encoding/json"
	"github.com/absurdlab/dyna-go-sdk/idp"
)

// UsernameForDetailRequest is the payload for username subject query. It serves for
// both Dyna-Identity communication and UI-Dyna communication
type UsernameForDetailRequest struct {
	Username string `json:"username"`
}

// UsernameForDetailResponse is the response payload for UsernameForSubjectRequest. It serves for
// both Dyna-Identity communication and UI-Dyna communication.
type UsernameForDetailResponse struct {
	idp.RPCResponse
	// Subject is the identifier for the user.
	Subject string `json:"subject,omitempty"`
	// Detail is the UI related data that will be relayed to the UI.
	Detail json.RawMessage `json:"detail,omitempty"`
}

// AuthenticationPolicyRequest is the RPC payload for querying the authentication strategy
// for the current session.
//
// As Dyna progresses, additional data may be passed in for a policy decision. For example,
// ip address, user agent, time, etc.
type AuthenticationPolicyRequest struct {
	// Subject is the resolved user identifier
	Subject string `json:"subject"`
	// Acr is the intended authentication context reference class
	// as part of the OIDC request. This value is optional. It
	// represents the ideal level of authentication hoped by the
	// client. Identity providers may use best effort to find a
	// strategy to fulfill this level, or it may ignore it altogether.
	Acr string `json:"acr,omitempty"`
}

// AuthenticationPolicyResponse is the RPC payload to respond to AuthenticationPolicyRequest.
type AuthenticationPolicyResponse struct {
	idp.RPCResponse
	// Amr is the resolved authentication methods to be carried out
	// by Dyna in order. The returned value must be recognizable by
	// Dyna, otherwise the authentication flow is broken. Dyna controls
	// the flow of these steps and proceed to the next only on success.
	Amr []string `json:"amr"`
	// Acr is the authentication context reference value achieved when
	// all Amr is completed. It may be different with the acr passed
	// in AuthenticationPolicyRequest.
	Acr string `json:"acr"`
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
	idp.RPCResponse
	// Data contains the claim response data. It may or may not
	// contain all the requested claims.
	Data map[string]interface{} `json:"data"`
}
