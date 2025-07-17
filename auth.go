// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadego

import (
	"context"
	"net/http"
	"net/url"

	"github.com/ArcadeAI/arcade-go/internal/apijson"
	"github.com/ArcadeAI/arcade-go/internal/apiquery"
	"github.com/ArcadeAI/arcade-go/internal/param"
	"github.com/ArcadeAI/arcade-go/internal/requestconfig"
	"github.com/ArcadeAI/arcade-go/option"
	"github.com/ArcadeAI/arcade-go/shared"
)

// AuthService contains methods and other services that help with interacting with
// the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r *AuthService) {
	r = &AuthService{}
	r.Options = opts
	return
}

// Starts the authorization process for given authorization requirements
func (r *AuthService) Authorize(ctx context.Context, body AuthAuthorizeParams, opts ...option.RequestOption) (res *shared.AuthorizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/auth/authorize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Checks the status of an ongoing authorization process for a specific tool. If
// 'wait' param is present, does not respond until either the auth status becomes
// completed or the timeout is reached.
func (r *AuthService) Status(ctx context.Context, query AuthStatusParams, opts ...option.RequestOption) (res *shared.AuthorizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/auth/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AuthRequestParam struct {
	AuthRequirement param.Field[AuthRequestAuthRequirementParam] `json:"auth_requirement,required"`
	UserID          param.Field[string]                          `json:"user_id,required"`
	// Optional: if provided, the user will be redirected to this URI after
	// authorization
	NextUri param.Field[string] `json:"next_uri"`
}

func (r AuthRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRequestAuthRequirementParam struct {
	// one of ID or ProviderID must be set
	ID     param.Field[string]                                `json:"id"`
	Oauth2 param.Field[AuthRequestAuthRequirementOauth2Param] `json:"oauth2"`
	// one of ID or ProviderID must be set
	ProviderID   param.Field[string] `json:"provider_id"`
	ProviderType param.Field[string] `json:"provider_type"`
}

func (r AuthRequestAuthRequirementParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRequestAuthRequirementOauth2Param struct {
	Scopes param.Field[[]string] `json:"scopes"`
}

func (r AuthRequestAuthRequirementOauth2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthAuthorizeParams struct {
	AuthRequest AuthRequestParam `json:"auth_request,required"`
}

func (r AuthAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AuthRequest)
}

type AuthStatusParams struct {
	// Authorization ID
	ID param.Field[string] `query:"id,required"`
	// Timeout in seconds (max 59)
	Wait param.Field[int64] `query:"wait"`
}

// URLQuery serializes [AuthStatusParams]'s query parameters as `url.Values`.
func (r AuthStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
