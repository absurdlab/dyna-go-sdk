package types

// RPCResponse is the common response payload.
type RPCResponse struct {
	// OK indicates whether is RPC is invoked successfully.
	OK bool `json:"ok"`
	// Error is the error code of the failure when OK is false.
	Error string `json:"error"`
	// ErrorDescription description describes the Error.
	ErrorDescription string `json:"error_description"`
}
