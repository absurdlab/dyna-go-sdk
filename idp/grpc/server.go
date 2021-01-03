package grpc

import (
	"context"
	"encoding/json"
	"github.com/absurdlab/dyna-go-sdk/idp"
	"github.com/absurdlab/dyna-go-sdk/idp/types"
	"github.com/absurdlab/dyna-go-sdk/pb"
)

// NewServer takes in an implementation of idp.Interface and return a GRPC server for the identity provider to host.
func NewServer(idp idp.Interface) pb.IdpServer {
	return &server{idp: idp}
}

type server struct {
	idp idp.Interface
}

func (s *server) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	response, err := s.idp.User(ctx, &types.UsernameForDetailRequest{
		Username: request.GetUsername(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Meta:    toMeta(&response.RPCResponse),
		Subject: response.Subject,
		Detail:  response.Detail,
	}, nil
}

func (s *server) GetClaims(ctx context.Context, request *pb.GetClaimsRequest) (*pb.GetClaimsResponse, error) {
	response, err := s.idp.Claims(ctx, &types.ClaimRequest{
		Subject: request.GetSubject(),
		Claims:  request.GetClaims(),
	})
	if err != nil {
		return nil, err
	}

	var data []byte
	if response.OK {
		var err error
		data, err = json.Marshal(response.Data)
		if err != nil {
			return nil, err
		}
	}

	return &pb.GetClaimsResponse{
		Meta: toMeta(&response.RPCResponse),
		Data: data,
	}, nil
}

func (s *server) GetAuthPolicy(ctx context.Context, request *pb.AuthenticationPolicyRequest) (*pb.AuthenticationPolicyResponse, error) {
	response, err := s.idp.AuthenticationPolicy(ctx, &types.AuthenticationPolicyRequest{
		Subject: request.GetSubject(),
		Acr:     request.GetAcr(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.AuthenticationPolicyResponse{
		Meta: toMeta(&response.RPCResponse),
		Amr:  response.Amr,
		Acr:  response.Acr,
	}, nil
}

func (s *server) GetAuthDetails(ctx context.Context, request *pb.AuthenticationDetailsRequest) (*pb.AuthenticationDetailsResponse, error) {
	response, err := s.idp.AuthenticationDetails(ctx, &types.AuthenticationDetailsRequest{
		Subject: request.GetSubject(),
		Amr:     request.GetAmr(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.AuthenticationDetailsResponse{
		Meta:    toMeta(&response.RPCResponse),
		Amr:     response.Amr,
		Details: response.Details,
	}, nil
}

func (s *server) CheckPassword(ctx context.Context, request *pb.PasswordValidationRequest) (*pb.PasswordValidationResponse, error) {
	response, err := s.idp.CheckPassword(ctx, &types.PasswordValidationRequest{
		Subject:  request.GetSubject(),
		Password: request.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.PasswordValidationResponse{
		Meta: toMeta(&response.RPCResponse),
	}, nil
}

func toMeta(response *types.RPCResponse) *pb.Meta {
	return &pb.Meta{
		Ok: response.OK,
		Error: &pb.Error{
			Code:        response.Error,
			Description: response.ErrorDescription,
		},
	}
}
