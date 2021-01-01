package idp

import (
	"context"
	"github.com/absurdlab/dyna-go-sdk/idp/types"
)

// Interface defines how Tiga interacts with the identity provider.
type Interface interface {
	// User exchanges username for subjects and details.
	User(ctx context.Context, request *types.UsernameForDetailRequest) (*types.UsernameForDetailResponse, error)
	// AuthenticationPolicy queries for authentication policy to be applied to the user.
	AuthenticationPolicy(ctx context.Context, request *types.AuthenticationPolicyRequest) (*types.AuthenticationPolicyResponse, error)
	// AuthenticationDetails queries for the details data to start an authentication scheme.
	AuthenticationDetails(ctx context.Context, request *types.AuthenticationDetailsRequest) (*types.AuthenticationDetailsResponse, error)
	// Claims retrieves claims for the user.
	Claims(ctx context.Context, request *types.ClaimRequest) (*types.ClaimResponse, error)
	// CheckPassword performs password validation.
	CheckPassword(ctx context.Context, request *types.PasswordValidationRequest) (*types.PasswordValidationResponse, error)
}
