package pwd

import "github.com/absurdlab/dyna-go-sdk/idp"

// ValidationRequest is the payload for validating a password. It serves for both
// Dyna-Identity communication and UI-Dyna communication.
type ValidationRequest struct {
	Subject  string `json:"subject"`
	Password string `json:"password"`
}

// ValidationResponse is the payload for responding to ValidationRequest. It serves for both
// Dyna-Identity communication and UI-Dyna communication.
type ValidationResponse struct {
	idp.RPCResponse
}
