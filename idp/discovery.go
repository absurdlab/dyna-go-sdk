package idp

// Discovery is the discoverable configuration of the identity provider. It is read on startup by
// Dyna and helps Dyna communicate with the identity provider. This file must be hosted at
// /.well-known/idp-configuration
type Discovery struct {
	// Version is the identity provider's API version. Dyna checks with this version with
	// itself to make sure the two can communicate smoothly together. If this version is
	// too old or too new for Dyna, Dyna self terminates.
	Version string `json:"version"`

	// LoginPagePath is the relative path where the login page is hosted. Fragment is
	// not permitted. Query parameters will be retained. By design, the login page and
	// Dyna must be hosted under the same domain. Dyna will use its own base URL and
	// the LoginPagePath to render the absolute URL.
	LoginPagePath string `json:"login_page_path"`

	// SelectAccountPagePath is the relative path where the account selection page is hosted.
	// Fragment is not permitted. Query parameters will be retained. By design, the account
	// selection page and Dyna must be hosted under the same domain. Dyna will use its own
	// base URL and the SelectAccountPagePath to render the absolute URL.
	SelectAccountPagePath string `json:"select_account_page_path"`

	// ConsentPagePath is the relative path where the consent page is hosted. Fragment is
	// not permitted. Query parameters will be retained. By design, the consent page and
	// Dyna must be hosted under the same domain. Dyna will use its own base URL and
	// the ConsentPagePath to render the absolute URL.
	ConsentPagePath string `json:"consent_page_path"`

	// UsernameRPC is the JSON-RPC name of exchange for subject with username.
	// user.UsernameForDetailRequest and user.UsernameForDetailResponse is used for communication.
	UsernameRPC string `json:"user_subject_rpc"`

	// AuthenticationPolicyRPC is the JSON-RPC name of querying for authentication policy.
	// user.AuthenticationPolicyRequest and user.AuthenticationPolicyResponse is used for communication.
	AuthenticationPolicyRPC string `json:"authentication_policy_rpc"`

	// ClaimsRPC is the JSON-RPC name of the user claims query. user.ClaimRequest and user.ClaimResponse
	// is used for communication.
	ClaimsRPC string `json:"claims_rpc"`

	// PasswordValidationRPC is JSON-RPC name of the password validation procedure.
	// pwd.ValidationRequest and pwd.ValidationResponse is used for communication.
	PasswordValidationRPC string `json:"password_validation_rpc"`
}
