package grpc

import (
	"context"
	"encoding/json"
	"github.com/absurdlab/dyna-go-sdk/idp"
	"github.com/absurdlab/dyna-go-sdk/idp/types"
	"github.com/absurdlab/dyna-go-sdk/pb"
)

// NewClient returns an idp.Interface implementation that adapts the GRPC stub client. This implementation
// does not attempt to convert or log any error returned by the stub, it returns them as is.
func NewClient(stub pb.IdpClient) idp.Interface {
	return &client{stub: stub}
}

type client struct {
	stub pb.IdpClient
}

func (c *client) User(ctx context.Context, request *types.UsernameForDetailRequest) (*types.UsernameForDetailResponse, error) {
	response, err := c.stub.GetUser(ctx, &pb.GetUserRequest{
		Username: request.Username,
	})
	if err != nil {
		return nil, err
	}

	return &types.UsernameForDetailResponse{
		RPCResponse: *fromMeta(response.GetMeta()),
		Subject:     response.GetSubject(),
		Detail:      response.GetDetail(),
	}, nil
}

func (c *client) AuthenticationPolicy(ctx context.Context, request *types.AuthenticationPolicyRequest) (*types.AuthenticationPolicyResponse, error) {
	response, err := c.stub.GetAuthPolicy(ctx, &pb.AuthenticationPolicyRequest{
		Subject: request.Subject,
		Acr:     request.Acr,
	})
	if err != nil {
		return nil, err
	}

	return &types.AuthenticationPolicyResponse{
		RPCResponse: *fromMeta(response.GetMeta()),
		Amr:         response.GetAmr(),
		Acr:         response.GetAcr(),
	}, nil
}

func (c *client) AuthenticationDetails(ctx context.Context, request *types.AuthenticationDetailsRequest) (*types.AuthenticationDetailsResponse, error) {
	response, err := c.stub.GetAuthDetails(ctx, &pb.AuthenticationDetailsRequest{
		Subject: request.Subject,
		Amr:     request.Amr,
	})
	if err != nil {
		return nil, err
	}

	return &types.AuthenticationDetailsResponse{
		RPCResponse: *fromMeta(response.GetMeta()),
		Amr:         response.GetAmr(),
		Details:     response.GetDetails(),
	}, nil
}

func (c *client) Claims(ctx context.Context, request *types.ClaimRequest) (*types.ClaimResponse, error) {
	response, err := c.stub.GetClaims(ctx, &pb.GetClaimsRequest{
		Subject: request.Subject,
		Claims:  request.Claims,
	})
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	if err := json.Unmarshal(response.GetData(), &data); err != nil {
		return nil, err
	}

	return &types.ClaimResponse{
		RPCResponse: *fromMeta(response.GetMeta()),
		Data:        data,
	}, nil

}

func (c *client) CheckPassword(ctx context.Context, request *types.PasswordValidationRequest) (*types.PasswordValidationResponse, error) {
	response, err := c.stub.CheckPassword(ctx, &pb.PasswordValidationRequest{
		Subject:  request.Subject,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.PasswordValidationResponse{
		RPCResponse: *fromMeta(response.GetMeta()),
	}, nil
}

func fromMeta(meta *pb.Meta) *types.RPCResponse {
	if meta.GetOk() {
		return &types.RPCResponse{OK: true}
	}
	return &types.RPCResponse{
		OK:               false,
		Error:            meta.GetError().GetCode(),
		ErrorDescription: meta.GetError().GetDescription(),
	}
}
