package types

// PasswordValidationRequest is the payload for validating a password. It serves for both
// Dyna-Identity communication and UI-Dyna communication.
type PasswordValidationRequest struct {
	Subject  string `json:"subject"`
	Password string `json:"password"`
}

// PasswordValidationResponse is the payload for responding to PasswordValidationRequest. It serves for both
// Dyna-Identity communication and UI-Dyna communication.
type PasswordValidationResponse struct {
	RPCResponse
}
