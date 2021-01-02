package rpc

import (
	"context"
	"github.com/absurdlab/dyna-go-sdk/idp"
	"github.com/absurdlab/dyna-go-sdk/idp/types"
)

func User(idp idp.Interface) func(request types.UsernameForDetailRequest, response *types.UsernameForDetailResponse) error {
	return func(request types.UsernameForDetailRequest, response *types.UsernameForDetailResponse) (err error) {
		response, err = idp.User(context.Background(), &request)
		return
	}
}

func AuthenticationPolicy(idp idp.Interface) func(request types.AuthenticationPolicyRequest, response *types.AuthenticationPolicyResponse) error {
	return func(request types.AuthenticationPolicyRequest, response *types.AuthenticationPolicyResponse) (err error) {
		response, err = idp.AuthenticationPolicy(context.Background(), &request)
		return
	}
}

func AuthenticationDetails(idp idp.Interface) func(request types.AuthenticationDetailsRequest, response *types.AuthenticationDetailsResponse) error {
	return func(request types.AuthenticationDetailsRequest, response *types.AuthenticationDetailsResponse) (err error) {
		response, err = idp.AuthenticationDetails(context.Background(), &request)
		return
	}
}

func Claims(idp idp.Interface) func(request types.ClaimRequest, response *types.ClaimResponse) error {
	return func(request types.ClaimRequest, response *types.ClaimResponse) (err error) {
		response, err = idp.Claims(context.Background(), &request)
		return
	}
}

func CheckPassword(idp idp.Interface) func(request types.PasswordValidationRequest, response *types.PasswordValidationResponse) error {
	return func(request types.PasswordValidationRequest, response *types.PasswordValidationResponse) (err error) {
		response, err = idp.CheckPassword(context.Background(), &request)
		return
	}
}

func Health(idp idp.Interface) func(request types.HealthRequest, response *types.HealthResponse) error {
	return func(request types.HealthRequest, response *types.HealthResponse) (err error) {
		response, err = idp.Health(context.Background(), &request)
		return
	}
}
