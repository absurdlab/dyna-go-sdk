package idp

// Discovery is the discoverable configuration of the identity provider. It is read on startup by
// Dyna and helps Dyna communicate with the identity provider. This file must be hosted at
// /.well-known/idp-configuration
type Discovery struct {
	// Version is the identity provider's API version. Dyna checks with this version with
	// itself to make sure the two can communicate smoothly together. If this version is
	// too old or too new for Dyna, Dyna self terminates.
	Version string `json:"version"`

	// LoginPage is the absolute path where the login page is hosted. Fragment is
	// not permitted. Query parameters will be retained. By design, the login page and
	// Dyna must be hosted under the same domain.
	LoginPage string `json:"login_page"`

	// SelectAccountPage is the absolute path where the account selection page is hosted.
	// Fragment is not permitted. Query parameters will be retained. By design, the account
	// selection page and Dyna must be hosted under the same domain.
	SelectAccountPage string `json:"select_account_page"`

	// ConsentPage is the absolute path where the consent page is hosted. Fragment is
	// not permitted. Query parameters will be retained. By design, the consent page and
	// Dyna must be hosted under the same domain.
	ConsentPage string `json:"consent_page"`

	// UsernameRPC is the JSON-RPC name of exchange for subject with username.
	UsernameRPC string `json:"user_subject_rpc"`

	// AuthenticationPolicyRPC is the JSON-RPC name of querying for authentication policy.
	AuthenticationPolicyRPC string `json:"authentication_policy_rpc"`

	// AuthenticationDetailsRPC is the JSON-RPC name of querying for authentication details.
	AuthenticationDetailsRPC string `json:"authentication_details_rpc"`

	// ClaimsRPC is the JSON-RPC name of the user claims query.
	ClaimsRPC string `json:"claims_rpc"`

	// PasswordValidationRPC is JSON-RPC name of the password validation procedure.
	PasswordValidationRPC string `json:"password_validation_rpc"`
}
