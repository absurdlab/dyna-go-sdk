package types

import "encoding/json"

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
	RPCResponse
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

// AuthenticationDetailsRequest requests the IDP to return detail information to assist the
// frontend app to start a certain amr based authentication flow.
type AuthenticationDetailsRequest struct {
	// Subject is the user identifier applying this method.
	Subject string `json:"subject"`
	// Amr is the next authentication method Dyna is prepared to execute
	Amr string `json:"amr"`
}

// AuthenticationDetailsResponse responds to AuthenticationDetailsRequest.
type AuthenticationDetailsResponse struct {
	RPCResponse
	// Amr is the requested amr.
	Amr string `json:"amr"`
	// Details is the arbitrary data Dyna will pass onto the frontend.
	Details json.RawMessage `json:"details"`
}
