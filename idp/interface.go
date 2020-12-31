package idp

import (
	"context"
	"github.com/absurdlab/dyna-go-sdk/idp/pwd"
	"github.com/absurdlab/dyna-go-sdk/idp/user"
)

// Interface defines how Tiga interacts with the identity provider.
type Interface interface {
	// Discovery returns the identity provider's configuration. Dyna will call this method on startup and cache the response.
	Discovery(ctx context.Context) (*Discovery, error)
	// User exchanges username for subjects and details.
	User(ctx context.Context, request *user.UsernameForDetailRequest) (*user.UsernameForDetailResponse, error)
	// AuthenticationPolicy queries for authentication policy to be applied to the user.
	AuthenticationPolicy(ctx context.Context, request *user.AuthenticationPolicyRequest) (*user.AuthenticationPolicyResponse, error)
	// Claims retrieves claims for the user.
	Claims(ctx context.Context, request *user.ClaimRequest) (*user.ClaimResponse, error)
	// CheckPassword performs password validation.
	CheckPassword(ctx context.Context, request *pwd.ValidationRequest) (*pwd.ValidationResponse, error)
}
