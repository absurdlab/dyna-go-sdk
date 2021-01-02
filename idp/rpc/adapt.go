package rpc

import (
	"context"
	"github.com/absurdlab/dyna-go-sdk/idp"
	"github.com/absurdlab/dyna-go-sdk/idp/types"
)

type Server struct {
	Idp idp.Interface
}

func (s *Server) User(request *types.UsernameForDetailRequest, response *types.UsernameForDetailResponse) (err error) {
	response, err = s.Idp.User(context.Background(), request)
	return
}

func (s *Server) AuthenticationPolicy(request *types.AuthenticationPolicyRequest, response *types.AuthenticationPolicyResponse) (err error) {
	response, err = s.Idp.AuthenticationPolicy(context.Background(), request)
	return
}

func (s *Server) AuthenticationDetails(request *types.AuthenticationDetailsRequest, response *types.AuthenticationDetailsResponse) (err error) {
	response, err = s.Idp.AuthenticationDetails(context.Background(), request)
	return
}

func (s *Server) Claims(request *types.ClaimRequest, response *types.ClaimResponse) (err error) {
	response, err = s.Idp.Claims(context.Background(), request)
	return
}

func (s *Server) CheckPassword(request *types.PasswordValidationRequest, response *types.PasswordValidationResponse) (err error) {
	response, err = s.Idp.CheckPassword(context.Background(), request)
	return
}

func (s *Server) Health(request *types.HealthRequest, response *types.HealthResponse) (err error) {
	response, err = s.Idp.Health(context.Background(), request)
	return
}
