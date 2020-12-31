package errors

import "errors"

var (
	BadPassword      = errors.New("bad_password")
	ExpiredPassword  = errors.New("expired_password")
	UsernameNotFound = errors.New("username_not_found")
	Internal         = errors.New("internal")
)
