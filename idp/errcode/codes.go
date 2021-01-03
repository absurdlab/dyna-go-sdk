package errcode

// Standard error codes returned in the rpc response.
const (
	BadPassword      = "bad_password"
	ExpiredPassword  = "expired_password"
	UsernameNotFound = "username_not_found"
	SubjectNotFound  = "subject_not_found"
	AmrNotSupported  = "amr_not_supported"
	Internal         = "internal"
)
