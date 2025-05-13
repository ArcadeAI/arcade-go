// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadeengine

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/arcade-engine-go/internal/apijson"
	"github.com/stainless-sdks/arcade-engine-go/internal/apiquery"
	"github.com/stainless-sdks/arcade-engine-go/internal/requestconfig"
	"github.com/stainless-sdks/arcade-engine-go/option"
	"github.com/stainless-sdks/arcade-engine-go/packages/param"
	"github.com/stainless-sdks/arcade-engine-go/shared"
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
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
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

// The properties AuthRequirement, UserID are required.
type AuthRequestParam struct {
	AuthRequirement AuthRequestAuthRequirementParam `json:"auth_requirement,omitzero,required"`
	UserID          string                          `json:"user_id,required"`
	paramObj
}

func (r AuthRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AuthRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRequestAuthRequirementParam struct {
	// one of ID or ProviderID must be set
	ID param.Opt[string] `json:"id,omitzero"`
	// one of ID or ProviderID must be set
	ProviderID   param.Opt[string]                     `json:"provider_id,omitzero"`
	ProviderType param.Opt[string]                     `json:"provider_type,omitzero"`
	Oauth2       AuthRequestAuthRequirementOauth2Param `json:"oauth2,omitzero"`
	paramObj
}

func (r AuthRequestAuthRequirementParam) MarshalJSON() (data []byte, err error) {
	type shadow AuthRequestAuthRequirementParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthRequestAuthRequirementParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRequestAuthRequirementOauth2Param struct {
	Scopes []string `json:"scopes,omitzero"`
	paramObj
}

func (r AuthRequestAuthRequirementOauth2Param) MarshalJSON() (data []byte, err error) {
	type shadow AuthRequestAuthRequirementOauth2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthRequestAuthRequirementOauth2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthAuthorizeParams struct {
	AuthRequest AuthRequestParam
	paramObj
}

func (r AuthAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.AuthRequest)
}
func (r *AuthAuthorizeParams) UnmarshalJSON(data []byte) error {
	return r.AuthRequest.UnmarshalJSON(data)
}

type AuthStatusParams struct {
	// Authorization ID
	ID string `query:"id,required" json:"-"`
	// Timeout in seconds (max 59)
	Wait param.Opt[int64] `query:"wait,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuthStatusParams]'s query parameters as `url.Values`.
func (r AuthStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
