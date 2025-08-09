// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadego

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ArcadeAI/arcade-go/internal/apijson"
	"github.com/ArcadeAI/arcade-go/internal/param"
	"github.com/ArcadeAI/arcade-go/internal/requestconfig"
	"github.com/ArcadeAI/arcade-go/option"
)

// AdminAuthProviderService contains methods and other services that help with
// interacting with the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminAuthProviderService] method instead.
type AdminAuthProviderService struct {
	Options []option.RequestOption
}

// NewAdminAuthProviderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminAuthProviderService(opts ...option.RequestOption) (r *AdminAuthProviderService) {
	r = &AdminAuthProviderService{}
	r.Options = opts
	return
}

// Create a new auth provider
func (r *AdminAuthProviderService) New(ctx context.Context, body AdminAuthProviderNewParams, opts ...option.RequestOption) (res *AuthProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/admin/auth_providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List a page of auth providers that are available to the caller
func (r *AdminAuthProviderService) List(ctx context.Context, opts ...option.RequestOption) (res *AdminAuthProviderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/admin/auth_providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a specific auth provider
func (r *AdminAuthProviderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AuthProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/admin/auth_providers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the details of a specific auth provider
func (r *AdminAuthProviderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AuthProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/admin/auth_providers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch an existing auth provider
func (r *AdminAuthProviderService) Patch(ctx context.Context, id string, body AdminAuthProviderPatchParams, opts ...option.RequestOption) (res *AuthProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/admin/auth_providers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AuthProviderCreateRequestParam struct {
	ID          param.Field[string] `json:"id,required"`
	Description param.Field[string] `json:"description"`
	// The unique external ID for the auth provider
	ExternalID param.Field[string]                               `json:"external_id"`
	Oauth2     param.Field[AuthProviderCreateRequestOauth2Param] `json:"oauth2"`
	ProviderID param.Field[string]                               `json:"provider_id"`
	Status     param.Field[string]                               `json:"status"`
	Type       param.Field[string]                               `json:"type"`
}

func (r AuthProviderCreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderCreateRequestOauth2Param struct {
	ClientID                  param.Field[string]                                                        `json:"client_id,required"`
	AuthorizeRequest          param.Field[AuthProviderCreateRequestOauth2AuthorizeRequestParam]          `json:"authorize_request"`
	ClientSecret              param.Field[string]                                                        `json:"client_secret"`
	Pkce                      param.Field[AuthProviderCreateRequestOauth2PkceParam]                      `json:"pkce"`
	RefreshRequest            param.Field[AuthProviderCreateRequestOauth2RefreshRequestParam]            `json:"refresh_request"`
	ScopeDelimiter            param.Field[AuthProviderCreateRequestOauth2ScopeDelimiter]                 `json:"scope_delimiter"`
	TokenIntrospectionRequest param.Field[AuthProviderCreateRequestOauth2TokenIntrospectionRequestParam] `json:"token_introspection_request"`
	TokenRequest              param.Field[AuthProviderCreateRequestOauth2TokenRequestParam]              `json:"token_request"`
	UserInfoRequest           param.Field[AuthProviderCreateRequestOauth2UserInfoRequestParam]           `json:"user_info_request"`
}

func (r AuthProviderCreateRequestOauth2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderCreateRequestOauth2AuthorizeRequestParam struct {
	Endpoint            param.Field[string]                                                             `json:"endpoint,required"`
	AuthMethod          param.Field[string]                                                             `json:"auth_method"`
	Method              param.Field[string]                                                             `json:"method"`
	Params              param.Field[map[string]string]                                                  `json:"params"`
	RequestContentType  param.Field[AuthProviderCreateRequestOauth2AuthorizeRequestRequestContentType]  `json:"request_content_type"`
	ResponseContentType param.Field[AuthProviderCreateRequestOauth2AuthorizeRequestResponseContentType] `json:"response_content_type"`
	ResponseMap         param.Field[map[string]string]                                                  `json:"response_map"`
}

func (r AuthProviderCreateRequestOauth2AuthorizeRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderCreateRequestOauth2AuthorizeRequestRequestContentType string

const (
	AuthProviderCreateRequestOauth2AuthorizeRequestRequestContentTypeApplicationXWwwFormUrlencoded AuthProviderCreateRequestOauth2AuthorizeRequestRequestContentType = "application/x-www-form-urlencoded"
	AuthProviderCreateRequestOauth2AuthorizeRequestRequestContentTypeApplicationJson               AuthProviderCreateRequestOauth2AuthorizeRequestRequestContentType = "application/json"
)

func (r AuthProviderCreateRequestOauth2AuthorizeRequestRequestContentType) IsKnown() bool {
	switch r {
	case AuthProviderCreateRequestOauth2AuthorizeRequestRequestContentTypeApplicationXWwwFormUrlencoded, AuthProviderCreateRequestOauth2AuthorizeRequestRequestContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderCreateRequestOauth2AuthorizeRequestResponseContentType string

const (
	AuthProviderCreateRequestOauth2AuthorizeRequestResponseContentTypeApplicationXWwwFormUrlencoded AuthProviderCreateRequestOauth2AuthorizeRequestResponseContentType = "application/x-www-form-urlencoded"
	AuthProviderCreateRequestOauth2AuthorizeRequestResponseContentTypeApplicationJson               AuthProviderCreateRequestOauth2AuthorizeRequestResponseContentType = "application/json"
)

func (r AuthProviderCreateRequestOauth2AuthorizeRequestResponseContentType) IsKnown() bool {
	switch r {
	case AuthProviderCreateRequestOauth2AuthorizeRequestResponseContentTypeApplicationXWwwFormUrlencoded, AuthProviderCreateRequestOauth2AuthorizeRequestResponseContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderCreateRequestOauth2PkceParam struct {
	CodeChallengeMethod param.Field[string] `json:"code_challenge_method"`
	Enabled             param.Field[bool]   `json:"enabled"`
}

func (r AuthProviderCreateRequestOauth2PkceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderCreateRequestOauth2RefreshRequestParam struct {
	Endpoint            param.Field[string]                                                           `json:"endpoint,required"`
	AuthMethod          param.Field[string]                                                           `json:"auth_method"`
	Method              param.Field[string]                                                           `json:"method"`
	Params              param.Field[map[string]string]                                                `json:"params"`
	RequestContentType  param.Field[AuthProviderCreateRequestOauth2RefreshRequestRequestContentType]  `json:"request_content_type"`
	ResponseContentType param.Field[AuthProviderCreateRequestOauth2RefreshRequestResponseContentType] `json:"response_content_type"`
	ResponseMap         param.Field[map[string]string]                                                `json:"response_map"`
}

func (r AuthProviderCreateRequestOauth2RefreshRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderCreateRequestOauth2RefreshRequestRequestContentType string

const (
	AuthProviderCreateRequestOauth2RefreshRequestRequestContentTypeApplicationXWwwFormUrlencoded AuthProviderCreateRequestOauth2RefreshRequestRequestContentType = "application/x-www-form-urlencoded"
	AuthProviderCreateRequestOauth2RefreshRequestRequestContentTypeApplicationJson               AuthProviderCreateRequestOauth2RefreshRequestRequestContentType = "application/json"
)

func (r AuthProviderCreateRequestOauth2RefreshRequestRequestContentType) IsKnown() bool {
	switch r {
	case AuthProviderCreateRequestOauth2RefreshRequestRequestContentTypeApplicationXWwwFormUrlencoded, AuthProviderCreateRequestOauth2RefreshRequestRequestContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderCreateRequestOauth2RefreshRequestResponseContentType string

const (
	AuthProviderCreateRequestOauth2RefreshRequestResponseContentTypeApplicationXWwwFormUrlencoded AuthProviderCreateRequestOauth2RefreshRequestResponseContentType = "application/x-www-form-urlencoded"
	AuthProviderCreateRequestOauth2RefreshRequestResponseContentTypeApplicationJson               AuthProviderCreateRequestOauth2RefreshRequestResponseContentType = "application/json"
)

func (r AuthProviderCreateRequestOauth2RefreshRequestResponseContentType) IsKnown() bool {
	switch r {
	case AuthProviderCreateRequestOauth2RefreshRequestResponseContentTypeApplicationXWwwFormUrlencoded, AuthProviderCreateRequestOauth2RefreshRequestResponseContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderCreateRequestOauth2ScopeDelimiter string

const (
	AuthProviderCreateRequestOauth2ScopeDelimiterUnknown0  AuthProviderCreateRequestOauth2ScopeDelimiter = ","
	AuthProviderCreateRequestOauth2ScopeDelimiterLowercase AuthProviderCreateRequestOauth2ScopeDelimiter = " "
)

func (r AuthProviderCreateRequestOauth2ScopeDelimiter) IsKnown() bool {
	switch r {
	case AuthProviderCreateRequestOauth2ScopeDelimiterUnknown0, AuthProviderCreateRequestOauth2ScopeDelimiterLowercase:
		return true
	}
	return false
}

type AuthProviderCreateRequestOauth2TokenIntrospectionRequestParam struct {
	Endpoint            param.Field[string]                                                                      `json:"endpoint,required"`
	Triggers            param.Field[AuthProviderCreateRequestOauth2TokenIntrospectionRequestTriggersParam]       `json:"triggers,required"`
	AuthMethod          param.Field[string]                                                                      `json:"auth_method"`
	Method              param.Field[string]                                                                      `json:"method"`
	Params              param.Field[map[string]string]                                                           `json:"params"`
	RequestContentType  param.Field[AuthProviderCreateRequestOauth2TokenIntrospectionRequestRequestContentType]  `json:"request_content_type"`
	ResponseContentType param.Field[AuthProviderCreateRequestOauth2TokenIntrospectionRequestResponseContentType] `json:"response_content_type"`
	ResponseMap         param.Field[map[string]string]                                                           `json:"response_map"`
}

func (r AuthProviderCreateRequestOauth2TokenIntrospectionRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderCreateRequestOauth2TokenIntrospectionRequestTriggersParam struct {
	OnTokenGrant   param.Field[bool] `json:"on_token_grant"`
	OnTokenRefresh param.Field[bool] `json:"on_token_refresh"`
}

func (r AuthProviderCreateRequestOauth2TokenIntrospectionRequestTriggersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderCreateRequestOauth2TokenIntrospectionRequestRequestContentType string

const (
	AuthProviderCreateRequestOauth2TokenIntrospectionRequestRequestContentTypeApplicationXWwwFormUrlencoded AuthProviderCreateRequestOauth2TokenIntrospectionRequestRequestContentType = "application/x-www-form-urlencoded"
	AuthProviderCreateRequestOauth2TokenIntrospectionRequestRequestContentTypeApplicationJson               AuthProviderCreateRequestOauth2TokenIntrospectionRequestRequestContentType = "application/json"
)

func (r AuthProviderCreateRequestOauth2TokenIntrospectionRequestRequestContentType) IsKnown() bool {
	switch r {
	case AuthProviderCreateRequestOauth2TokenIntrospectionRequestRequestContentTypeApplicationXWwwFormUrlencoded, AuthProviderCreateRequestOauth2TokenIntrospectionRequestRequestContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderCreateRequestOauth2TokenIntrospectionRequestResponseContentType string

const (
	AuthProviderCreateRequestOauth2TokenIntrospectionRequestResponseContentTypeApplicationXWwwFormUrlencoded AuthProviderCreateRequestOauth2TokenIntrospectionRequestResponseContentType = "application/x-www-form-urlencoded"
	AuthProviderCreateRequestOauth2TokenIntrospectionRequestResponseContentTypeApplicationJson               AuthProviderCreateRequestOauth2TokenIntrospectionRequestResponseContentType = "application/json"
)

func (r AuthProviderCreateRequestOauth2TokenIntrospectionRequestResponseContentType) IsKnown() bool {
	switch r {
	case AuthProviderCreateRequestOauth2TokenIntrospectionRequestResponseContentTypeApplicationXWwwFormUrlencoded, AuthProviderCreateRequestOauth2TokenIntrospectionRequestResponseContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderCreateRequestOauth2TokenRequestParam struct {
	Endpoint            param.Field[string]                                                         `json:"endpoint,required"`
	AuthMethod          param.Field[string]                                                         `json:"auth_method"`
	Method              param.Field[string]                                                         `json:"method"`
	Params              param.Field[map[string]string]                                              `json:"params"`
	RequestContentType  param.Field[AuthProviderCreateRequestOauth2TokenRequestRequestContentType]  `json:"request_content_type"`
	ResponseContentType param.Field[AuthProviderCreateRequestOauth2TokenRequestResponseContentType] `json:"response_content_type"`
	ResponseMap         param.Field[map[string]string]                                              `json:"response_map"`
}

func (r AuthProviderCreateRequestOauth2TokenRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderCreateRequestOauth2TokenRequestRequestContentType string

const (
	AuthProviderCreateRequestOauth2TokenRequestRequestContentTypeApplicationXWwwFormUrlencoded AuthProviderCreateRequestOauth2TokenRequestRequestContentType = "application/x-www-form-urlencoded"
	AuthProviderCreateRequestOauth2TokenRequestRequestContentTypeApplicationJson               AuthProviderCreateRequestOauth2TokenRequestRequestContentType = "application/json"
)

func (r AuthProviderCreateRequestOauth2TokenRequestRequestContentType) IsKnown() bool {
	switch r {
	case AuthProviderCreateRequestOauth2TokenRequestRequestContentTypeApplicationXWwwFormUrlencoded, AuthProviderCreateRequestOauth2TokenRequestRequestContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderCreateRequestOauth2TokenRequestResponseContentType string

const (
	AuthProviderCreateRequestOauth2TokenRequestResponseContentTypeApplicationXWwwFormUrlencoded AuthProviderCreateRequestOauth2TokenRequestResponseContentType = "application/x-www-form-urlencoded"
	AuthProviderCreateRequestOauth2TokenRequestResponseContentTypeApplicationJson               AuthProviderCreateRequestOauth2TokenRequestResponseContentType = "application/json"
)

func (r AuthProviderCreateRequestOauth2TokenRequestResponseContentType) IsKnown() bool {
	switch r {
	case AuthProviderCreateRequestOauth2TokenRequestResponseContentTypeApplicationXWwwFormUrlencoded, AuthProviderCreateRequestOauth2TokenRequestResponseContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderCreateRequestOauth2UserInfoRequestParam struct {
	Endpoint            param.Field[string]                                                            `json:"endpoint,required"`
	Triggers            param.Field[AuthProviderCreateRequestOauth2UserInfoRequestTriggersParam]       `json:"triggers,required"`
	AuthMethod          param.Field[string]                                                            `json:"auth_method"`
	Method              param.Field[string]                                                            `json:"method"`
	Params              param.Field[map[string]string]                                                 `json:"params"`
	RequestContentType  param.Field[AuthProviderCreateRequestOauth2UserInfoRequestRequestContentType]  `json:"request_content_type"`
	ResponseContentType param.Field[AuthProviderCreateRequestOauth2UserInfoRequestResponseContentType] `json:"response_content_type"`
	ResponseMap         param.Field[map[string]string]                                                 `json:"response_map"`
}

func (r AuthProviderCreateRequestOauth2UserInfoRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderCreateRequestOauth2UserInfoRequestTriggersParam struct {
	OnTokenGrant   param.Field[bool] `json:"on_token_grant"`
	OnTokenRefresh param.Field[bool] `json:"on_token_refresh"`
}

func (r AuthProviderCreateRequestOauth2UserInfoRequestTriggersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderCreateRequestOauth2UserInfoRequestRequestContentType string

const (
	AuthProviderCreateRequestOauth2UserInfoRequestRequestContentTypeApplicationXWwwFormUrlencoded AuthProviderCreateRequestOauth2UserInfoRequestRequestContentType = "application/x-www-form-urlencoded"
	AuthProviderCreateRequestOauth2UserInfoRequestRequestContentTypeApplicationJson               AuthProviderCreateRequestOauth2UserInfoRequestRequestContentType = "application/json"
)

func (r AuthProviderCreateRequestOauth2UserInfoRequestRequestContentType) IsKnown() bool {
	switch r {
	case AuthProviderCreateRequestOauth2UserInfoRequestRequestContentTypeApplicationXWwwFormUrlencoded, AuthProviderCreateRequestOauth2UserInfoRequestRequestContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderCreateRequestOauth2UserInfoRequestResponseContentType string

const (
	AuthProviderCreateRequestOauth2UserInfoRequestResponseContentTypeApplicationXWwwFormUrlencoded AuthProviderCreateRequestOauth2UserInfoRequestResponseContentType = "application/x-www-form-urlencoded"
	AuthProviderCreateRequestOauth2UserInfoRequestResponseContentTypeApplicationJson               AuthProviderCreateRequestOauth2UserInfoRequestResponseContentType = "application/json"
)

func (r AuthProviderCreateRequestOauth2UserInfoRequestResponseContentType) IsKnown() bool {
	switch r {
	case AuthProviderCreateRequestOauth2UserInfoRequestResponseContentTypeApplicationXWwwFormUrlencoded, AuthProviderCreateRequestOauth2UserInfoRequestResponseContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderResponse struct {
	ID          string                      `json:"id"`
	Binding     AuthProviderResponseBinding `json:"binding"`
	CreatedAt   string                      `json:"created_at"`
	Description string                      `json:"description"`
	Oauth2      AuthProviderResponseOauth2  `json:"oauth2"`
	ProviderID  string                      `json:"provider_id"`
	Status      string                      `json:"status"`
	Type        string                      `json:"type"`
	UpdatedAt   string                      `json:"updated_at"`
	JSON        authProviderResponseJSON    `json:"-"`
}

// authProviderResponseJSON contains the JSON metadata for the struct
// [AuthProviderResponse]
type authProviderResponseJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Oauth2      apijson.Field
	ProviderID  apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthProviderResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseJSON) RawJSON() string {
	return r.raw
}

type AuthProviderResponseBinding struct {
	ID   string                          `json:"id"`
	Type AuthProviderResponseBindingType `json:"type"`
	JSON authProviderResponseBindingJSON `json:"-"`
}

// authProviderResponseBindingJSON contains the JSON metadata for the struct
// [AuthProviderResponseBinding]
type authProviderResponseBindingJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthProviderResponseBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseBindingJSON) RawJSON() string {
	return r.raw
}

type AuthProviderResponseBindingType string

const (
	AuthProviderResponseBindingTypeStatic  AuthProviderResponseBindingType = "static"
	AuthProviderResponseBindingTypeTenant  AuthProviderResponseBindingType = "tenant"
	AuthProviderResponseBindingTypeProject AuthProviderResponseBindingType = "project"
	AuthProviderResponseBindingTypeAccount AuthProviderResponseBindingType = "account"
)

func (r AuthProviderResponseBindingType) IsKnown() bool {
	switch r {
	case AuthProviderResponseBindingTypeStatic, AuthProviderResponseBindingTypeTenant, AuthProviderResponseBindingTypeProject, AuthProviderResponseBindingTypeAccount:
		return true
	}
	return false
}

type AuthProviderResponseOauth2 struct {
	AuthorizeRequest AuthProviderResponseOauth2AuthorizeRequest `json:"authorize_request"`
	ClientID         string                                     `json:"client_id"`
	ClientSecret     AuthProviderResponseOauth2ClientSecret     `json:"client_secret"`
	Pkce             AuthProviderResponseOauth2Pkce             `json:"pkce"`
	// The redirect URI required for this provider.
	RedirectUri               string                                              `json:"redirect_uri"`
	RefreshRequest            AuthProviderResponseOauth2RefreshRequest            `json:"refresh_request"`
	ScopeDelimiter            string                                              `json:"scope_delimiter"`
	TokenIntrospectionRequest AuthProviderResponseOauth2TokenIntrospectionRequest `json:"token_introspection_request"`
	TokenRequest              AuthProviderResponseOauth2TokenRequest              `json:"token_request"`
	UserInfoRequest           AuthProviderResponseOauth2UserInfoRequest           `json:"user_info_request"`
	JSON                      authProviderResponseOauth2JSON                      `json:"-"`
}

// authProviderResponseOauth2JSON contains the JSON metadata for the struct
// [AuthProviderResponseOauth2]
type authProviderResponseOauth2JSON struct {
	AuthorizeRequest          apijson.Field
	ClientID                  apijson.Field
	ClientSecret              apijson.Field
	Pkce                      apijson.Field
	RedirectUri               apijson.Field
	RefreshRequest            apijson.Field
	ScopeDelimiter            apijson.Field
	TokenIntrospectionRequest apijson.Field
	TokenRequest              apijson.Field
	UserInfoRequest           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AuthProviderResponseOauth2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseOauth2JSON) RawJSON() string {
	return r.raw
}

type AuthProviderResponseOauth2AuthorizeRequest struct {
	AuthMethod          string                                         `json:"auth_method"`
	Endpoint            string                                         `json:"endpoint"`
	ExpirationFormat    string                                         `json:"expiration_format"`
	Method              string                                         `json:"method"`
	Params              map[string]string                              `json:"params"`
	RequestContentType  string                                         `json:"request_content_type"`
	ResponseContentType string                                         `json:"response_content_type"`
	ResponseMap         map[string]string                              `json:"response_map"`
	JSON                authProviderResponseOauth2AuthorizeRequestJSON `json:"-"`
}

// authProviderResponseOauth2AuthorizeRequestJSON contains the JSON metadata for
// the struct [AuthProviderResponseOauth2AuthorizeRequest]
type authProviderResponseOauth2AuthorizeRequestJSON struct {
	AuthMethod          apijson.Field
	Endpoint            apijson.Field
	ExpirationFormat    apijson.Field
	Method              apijson.Field
	Params              apijson.Field
	RequestContentType  apijson.Field
	ResponseContentType apijson.Field
	ResponseMap         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AuthProviderResponseOauth2AuthorizeRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseOauth2AuthorizeRequestJSON) RawJSON() string {
	return r.raw
}

type AuthProviderResponseOauth2ClientSecret struct {
	Binding  AuthProviderResponseOauth2ClientSecretBinding `json:"binding"`
	Editable bool                                          `json:"editable"`
	Exists   bool                                          `json:"exists"`
	Hint     string                                        `json:"hint"`
	Value    string                                        `json:"value"`
	JSON     authProviderResponseOauth2ClientSecretJSON    `json:"-"`
}

// authProviderResponseOauth2ClientSecretJSON contains the JSON metadata for the
// struct [AuthProviderResponseOauth2ClientSecret]
type authProviderResponseOauth2ClientSecretJSON struct {
	Binding     apijson.Field
	Editable    apijson.Field
	Exists      apijson.Field
	Hint        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthProviderResponseOauth2ClientSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseOauth2ClientSecretJSON) RawJSON() string {
	return r.raw
}

type AuthProviderResponseOauth2ClientSecretBinding string

const (
	AuthProviderResponseOauth2ClientSecretBindingStatic  AuthProviderResponseOauth2ClientSecretBinding = "static"
	AuthProviderResponseOauth2ClientSecretBindingTenant  AuthProviderResponseOauth2ClientSecretBinding = "tenant"
	AuthProviderResponseOauth2ClientSecretBindingProject AuthProviderResponseOauth2ClientSecretBinding = "project"
	AuthProviderResponseOauth2ClientSecretBindingAccount AuthProviderResponseOauth2ClientSecretBinding = "account"
)

func (r AuthProviderResponseOauth2ClientSecretBinding) IsKnown() bool {
	switch r {
	case AuthProviderResponseOauth2ClientSecretBindingStatic, AuthProviderResponseOauth2ClientSecretBindingTenant, AuthProviderResponseOauth2ClientSecretBindingProject, AuthProviderResponseOauth2ClientSecretBindingAccount:
		return true
	}
	return false
}

type AuthProviderResponseOauth2Pkce struct {
	CodeChallengeMethod string                             `json:"code_challenge_method"`
	Enabled             bool                               `json:"enabled"`
	JSON                authProviderResponseOauth2PkceJSON `json:"-"`
}

// authProviderResponseOauth2PkceJSON contains the JSON metadata for the struct
// [AuthProviderResponseOauth2Pkce]
type authProviderResponseOauth2PkceJSON struct {
	CodeChallengeMethod apijson.Field
	Enabled             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AuthProviderResponseOauth2Pkce) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseOauth2PkceJSON) RawJSON() string {
	return r.raw
}

type AuthProviderResponseOauth2RefreshRequest struct {
	AuthMethod          string                                       `json:"auth_method"`
	Endpoint            string                                       `json:"endpoint"`
	ExpirationFormat    string                                       `json:"expiration_format"`
	Method              string                                       `json:"method"`
	Params              map[string]string                            `json:"params"`
	RequestContentType  string                                       `json:"request_content_type"`
	ResponseContentType string                                       `json:"response_content_type"`
	ResponseMap         map[string]string                            `json:"response_map"`
	JSON                authProviderResponseOauth2RefreshRequestJSON `json:"-"`
}

// authProviderResponseOauth2RefreshRequestJSON contains the JSON metadata for the
// struct [AuthProviderResponseOauth2RefreshRequest]
type authProviderResponseOauth2RefreshRequestJSON struct {
	AuthMethod          apijson.Field
	Endpoint            apijson.Field
	ExpirationFormat    apijson.Field
	Method              apijson.Field
	Params              apijson.Field
	RequestContentType  apijson.Field
	ResponseContentType apijson.Field
	ResponseMap         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AuthProviderResponseOauth2RefreshRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseOauth2RefreshRequestJSON) RawJSON() string {
	return r.raw
}

type AuthProviderResponseOauth2TokenIntrospectionRequest struct {
	AuthMethod          string                                                      `json:"auth_method"`
	Enabled             bool                                                        `json:"enabled"`
	Endpoint            string                                                      `json:"endpoint"`
	ExpirationFormat    string                                                      `json:"expiration_format"`
	Method              string                                                      `json:"method"`
	Params              map[string]string                                           `json:"params"`
	RequestContentType  string                                                      `json:"request_content_type"`
	ResponseContentType string                                                      `json:"response_content_type"`
	ResponseMap         map[string]string                                           `json:"response_map"`
	Triggers            AuthProviderResponseOauth2TokenIntrospectionRequestTriggers `json:"triggers"`
	JSON                authProviderResponseOauth2TokenIntrospectionRequestJSON     `json:"-"`
}

// authProviderResponseOauth2TokenIntrospectionRequestJSON contains the JSON
// metadata for the struct [AuthProviderResponseOauth2TokenIntrospectionRequest]
type authProviderResponseOauth2TokenIntrospectionRequestJSON struct {
	AuthMethod          apijson.Field
	Enabled             apijson.Field
	Endpoint            apijson.Field
	ExpirationFormat    apijson.Field
	Method              apijson.Field
	Params              apijson.Field
	RequestContentType  apijson.Field
	ResponseContentType apijson.Field
	ResponseMap         apijson.Field
	Triggers            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AuthProviderResponseOauth2TokenIntrospectionRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseOauth2TokenIntrospectionRequestJSON) RawJSON() string {
	return r.raw
}

type AuthProviderResponseOauth2TokenIntrospectionRequestTriggers struct {
	OnTokenGrant   bool                                                            `json:"on_token_grant"`
	OnTokenRefresh bool                                                            `json:"on_token_refresh"`
	JSON           authProviderResponseOauth2TokenIntrospectionRequestTriggersJSON `json:"-"`
}

// authProviderResponseOauth2TokenIntrospectionRequestTriggersJSON contains the
// JSON metadata for the struct
// [AuthProviderResponseOauth2TokenIntrospectionRequestTriggers]
type authProviderResponseOauth2TokenIntrospectionRequestTriggersJSON struct {
	OnTokenGrant   apijson.Field
	OnTokenRefresh apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AuthProviderResponseOauth2TokenIntrospectionRequestTriggers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseOauth2TokenIntrospectionRequestTriggersJSON) RawJSON() string {
	return r.raw
}

type AuthProviderResponseOauth2TokenRequest struct {
	AuthMethod          string                                     `json:"auth_method"`
	Endpoint            string                                     `json:"endpoint"`
	ExpirationFormat    string                                     `json:"expiration_format"`
	Method              string                                     `json:"method"`
	Params              map[string]string                          `json:"params"`
	RequestContentType  string                                     `json:"request_content_type"`
	ResponseContentType string                                     `json:"response_content_type"`
	ResponseMap         map[string]string                          `json:"response_map"`
	JSON                authProviderResponseOauth2TokenRequestJSON `json:"-"`
}

// authProviderResponseOauth2TokenRequestJSON contains the JSON metadata for the
// struct [AuthProviderResponseOauth2TokenRequest]
type authProviderResponseOauth2TokenRequestJSON struct {
	AuthMethod          apijson.Field
	Endpoint            apijson.Field
	ExpirationFormat    apijson.Field
	Method              apijson.Field
	Params              apijson.Field
	RequestContentType  apijson.Field
	ResponseContentType apijson.Field
	ResponseMap         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AuthProviderResponseOauth2TokenRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseOauth2TokenRequestJSON) RawJSON() string {
	return r.raw
}

type AuthProviderResponseOauth2UserInfoRequest struct {
	AuthMethod          string                                            `json:"auth_method"`
	Endpoint            string                                            `json:"endpoint"`
	ExpirationFormat    string                                            `json:"expiration_format"`
	Method              string                                            `json:"method"`
	Params              map[string]string                                 `json:"params"`
	RequestContentType  string                                            `json:"request_content_type"`
	ResponseContentType string                                            `json:"response_content_type"`
	ResponseMap         map[string]string                                 `json:"response_map"`
	Triggers            AuthProviderResponseOauth2UserInfoRequestTriggers `json:"triggers"`
	JSON                authProviderResponseOauth2UserInfoRequestJSON     `json:"-"`
}

// authProviderResponseOauth2UserInfoRequestJSON contains the JSON metadata for the
// struct [AuthProviderResponseOauth2UserInfoRequest]
type authProviderResponseOauth2UserInfoRequestJSON struct {
	AuthMethod          apijson.Field
	Endpoint            apijson.Field
	ExpirationFormat    apijson.Field
	Method              apijson.Field
	Params              apijson.Field
	RequestContentType  apijson.Field
	ResponseContentType apijson.Field
	ResponseMap         apijson.Field
	Triggers            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AuthProviderResponseOauth2UserInfoRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseOauth2UserInfoRequestJSON) RawJSON() string {
	return r.raw
}

type AuthProviderResponseOauth2UserInfoRequestTriggers struct {
	OnTokenGrant   bool                                                  `json:"on_token_grant"`
	OnTokenRefresh bool                                                  `json:"on_token_refresh"`
	JSON           authProviderResponseOauth2UserInfoRequestTriggersJSON `json:"-"`
}

// authProviderResponseOauth2UserInfoRequestTriggersJSON contains the JSON metadata
// for the struct [AuthProviderResponseOauth2UserInfoRequestTriggers]
type authProviderResponseOauth2UserInfoRequestTriggersJSON struct {
	OnTokenGrant   apijson.Field
	OnTokenRefresh apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AuthProviderResponseOauth2UserInfoRequestTriggers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authProviderResponseOauth2UserInfoRequestTriggersJSON) RawJSON() string {
	return r.raw
}

type AuthProviderUpdateRequestParam struct {
	ID          param.Field[string]                               `json:"id"`
	Description param.Field[string]                               `json:"description"`
	Oauth2      param.Field[AuthProviderUpdateRequestOauth2Param] `json:"oauth2"`
	ProviderID  param.Field[string]                               `json:"provider_id"`
	Status      param.Field[string]                               `json:"status"`
	Type        param.Field[string]                               `json:"type"`
}

func (r AuthProviderUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderUpdateRequestOauth2Param struct {
	AuthorizeRequest param.Field[AuthProviderUpdateRequestOauth2AuthorizeRequestParam] `json:"authorize_request"`
	ClientID         param.Field[string]                                               `json:"client_id"`
	ClientSecret     param.Field[string]                                               `json:"client_secret"`
	Pkce             param.Field[AuthProviderUpdateRequestOauth2PkceParam]             `json:"pkce"`
	RefreshRequest   param.Field[AuthProviderUpdateRequestOauth2RefreshRequestParam]   `json:"refresh_request"`
	ScopeDelimiter   param.Field[AuthProviderUpdateRequestOauth2ScopeDelimiter]        `json:"scope_delimiter"`
	TokenRequest     param.Field[AuthProviderUpdateRequestOauth2TokenRequestParam]     `json:"token_request"`
	UserInfoRequest  param.Field[AuthProviderUpdateRequestOauth2UserInfoRequestParam]  `json:"user_info_request"`
}

func (r AuthProviderUpdateRequestOauth2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderUpdateRequestOauth2AuthorizeRequestParam struct {
	AuthMethod          param.Field[string]                                                             `json:"auth_method"`
	Endpoint            param.Field[string]                                                             `json:"endpoint"`
	Method              param.Field[string]                                                             `json:"method"`
	Params              param.Field[map[string]string]                                                  `json:"params"`
	RequestContentType  param.Field[AuthProviderUpdateRequestOauth2AuthorizeRequestRequestContentType]  `json:"request_content_type"`
	ResponseContentType param.Field[AuthProviderUpdateRequestOauth2AuthorizeRequestResponseContentType] `json:"response_content_type"`
	ResponseMap         param.Field[map[string]string]                                                  `json:"response_map"`
}

func (r AuthProviderUpdateRequestOauth2AuthorizeRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderUpdateRequestOauth2AuthorizeRequestRequestContentType string

const (
	AuthProviderUpdateRequestOauth2AuthorizeRequestRequestContentTypeApplicationXWwwFormUrlencoded AuthProviderUpdateRequestOauth2AuthorizeRequestRequestContentType = "application/x-www-form-urlencoded"
	AuthProviderUpdateRequestOauth2AuthorizeRequestRequestContentTypeApplicationJson               AuthProviderUpdateRequestOauth2AuthorizeRequestRequestContentType = "application/json"
)

func (r AuthProviderUpdateRequestOauth2AuthorizeRequestRequestContentType) IsKnown() bool {
	switch r {
	case AuthProviderUpdateRequestOauth2AuthorizeRequestRequestContentTypeApplicationXWwwFormUrlencoded, AuthProviderUpdateRequestOauth2AuthorizeRequestRequestContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderUpdateRequestOauth2AuthorizeRequestResponseContentType string

const (
	AuthProviderUpdateRequestOauth2AuthorizeRequestResponseContentTypeApplicationXWwwFormUrlencoded AuthProviderUpdateRequestOauth2AuthorizeRequestResponseContentType = "application/x-www-form-urlencoded"
	AuthProviderUpdateRequestOauth2AuthorizeRequestResponseContentTypeApplicationJson               AuthProviderUpdateRequestOauth2AuthorizeRequestResponseContentType = "application/json"
)

func (r AuthProviderUpdateRequestOauth2AuthorizeRequestResponseContentType) IsKnown() bool {
	switch r {
	case AuthProviderUpdateRequestOauth2AuthorizeRequestResponseContentTypeApplicationXWwwFormUrlencoded, AuthProviderUpdateRequestOauth2AuthorizeRequestResponseContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderUpdateRequestOauth2PkceParam struct {
	CodeChallengeMethod param.Field[string] `json:"code_challenge_method"`
	Enabled             param.Field[bool]   `json:"enabled"`
}

func (r AuthProviderUpdateRequestOauth2PkceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderUpdateRequestOauth2RefreshRequestParam struct {
	AuthMethod          param.Field[string]                                                           `json:"auth_method"`
	Endpoint            param.Field[string]                                                           `json:"endpoint"`
	Method              param.Field[string]                                                           `json:"method"`
	Params              param.Field[map[string]string]                                                `json:"params"`
	RequestContentType  param.Field[AuthProviderUpdateRequestOauth2RefreshRequestRequestContentType]  `json:"request_content_type"`
	ResponseContentType param.Field[AuthProviderUpdateRequestOauth2RefreshRequestResponseContentType] `json:"response_content_type"`
	ResponseMap         param.Field[map[string]string]                                                `json:"response_map"`
}

func (r AuthProviderUpdateRequestOauth2RefreshRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderUpdateRequestOauth2RefreshRequestRequestContentType string

const (
	AuthProviderUpdateRequestOauth2RefreshRequestRequestContentTypeApplicationXWwwFormUrlencoded AuthProviderUpdateRequestOauth2RefreshRequestRequestContentType = "application/x-www-form-urlencoded"
	AuthProviderUpdateRequestOauth2RefreshRequestRequestContentTypeApplicationJson               AuthProviderUpdateRequestOauth2RefreshRequestRequestContentType = "application/json"
)

func (r AuthProviderUpdateRequestOauth2RefreshRequestRequestContentType) IsKnown() bool {
	switch r {
	case AuthProviderUpdateRequestOauth2RefreshRequestRequestContentTypeApplicationXWwwFormUrlencoded, AuthProviderUpdateRequestOauth2RefreshRequestRequestContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderUpdateRequestOauth2RefreshRequestResponseContentType string

const (
	AuthProviderUpdateRequestOauth2RefreshRequestResponseContentTypeApplicationXWwwFormUrlencoded AuthProviderUpdateRequestOauth2RefreshRequestResponseContentType = "application/x-www-form-urlencoded"
	AuthProviderUpdateRequestOauth2RefreshRequestResponseContentTypeApplicationJson               AuthProviderUpdateRequestOauth2RefreshRequestResponseContentType = "application/json"
)

func (r AuthProviderUpdateRequestOauth2RefreshRequestResponseContentType) IsKnown() bool {
	switch r {
	case AuthProviderUpdateRequestOauth2RefreshRequestResponseContentTypeApplicationXWwwFormUrlencoded, AuthProviderUpdateRequestOauth2RefreshRequestResponseContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderUpdateRequestOauth2ScopeDelimiter string

const (
	AuthProviderUpdateRequestOauth2ScopeDelimiterUnknown1  AuthProviderUpdateRequestOauth2ScopeDelimiter = ","
	AuthProviderUpdateRequestOauth2ScopeDelimiterLowercase AuthProviderUpdateRequestOauth2ScopeDelimiter = " "
)

func (r AuthProviderUpdateRequestOauth2ScopeDelimiter) IsKnown() bool {
	switch r {
	case AuthProviderUpdateRequestOauth2ScopeDelimiterUnknown1, AuthProviderUpdateRequestOauth2ScopeDelimiterLowercase:
		return true
	}
	return false
}

type AuthProviderUpdateRequestOauth2TokenRequestParam struct {
	AuthMethod          param.Field[string]                                                         `json:"auth_method"`
	Endpoint            param.Field[string]                                                         `json:"endpoint"`
	Method              param.Field[string]                                                         `json:"method"`
	Params              param.Field[map[string]string]                                              `json:"params"`
	RequestContentType  param.Field[AuthProviderUpdateRequestOauth2TokenRequestRequestContentType]  `json:"request_content_type"`
	ResponseContentType param.Field[AuthProviderUpdateRequestOauth2TokenRequestResponseContentType] `json:"response_content_type"`
	ResponseMap         param.Field[map[string]string]                                              `json:"response_map"`
}

func (r AuthProviderUpdateRequestOauth2TokenRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderUpdateRequestOauth2TokenRequestRequestContentType string

const (
	AuthProviderUpdateRequestOauth2TokenRequestRequestContentTypeApplicationXWwwFormUrlencoded AuthProviderUpdateRequestOauth2TokenRequestRequestContentType = "application/x-www-form-urlencoded"
	AuthProviderUpdateRequestOauth2TokenRequestRequestContentTypeApplicationJson               AuthProviderUpdateRequestOauth2TokenRequestRequestContentType = "application/json"
)

func (r AuthProviderUpdateRequestOauth2TokenRequestRequestContentType) IsKnown() bool {
	switch r {
	case AuthProviderUpdateRequestOauth2TokenRequestRequestContentTypeApplicationXWwwFormUrlencoded, AuthProviderUpdateRequestOauth2TokenRequestRequestContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderUpdateRequestOauth2TokenRequestResponseContentType string

const (
	AuthProviderUpdateRequestOauth2TokenRequestResponseContentTypeApplicationXWwwFormUrlencoded AuthProviderUpdateRequestOauth2TokenRequestResponseContentType = "application/x-www-form-urlencoded"
	AuthProviderUpdateRequestOauth2TokenRequestResponseContentTypeApplicationJson               AuthProviderUpdateRequestOauth2TokenRequestResponseContentType = "application/json"
)

func (r AuthProviderUpdateRequestOauth2TokenRequestResponseContentType) IsKnown() bool {
	switch r {
	case AuthProviderUpdateRequestOauth2TokenRequestResponseContentTypeApplicationXWwwFormUrlencoded, AuthProviderUpdateRequestOauth2TokenRequestResponseContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderUpdateRequestOauth2UserInfoRequestParam struct {
	AuthMethod          param.Field[string]                                                            `json:"auth_method"`
	Endpoint            param.Field[string]                                                            `json:"endpoint"`
	Method              param.Field[string]                                                            `json:"method"`
	Params              param.Field[map[string]string]                                                 `json:"params"`
	RequestContentType  param.Field[AuthProviderUpdateRequestOauth2UserInfoRequestRequestContentType]  `json:"request_content_type"`
	ResponseContentType param.Field[AuthProviderUpdateRequestOauth2UserInfoRequestResponseContentType] `json:"response_content_type"`
	ResponseMap         param.Field[map[string]string]                                                 `json:"response_map"`
	Triggers            param.Field[AuthProviderUpdateRequestOauth2UserInfoRequestTriggersParam]       `json:"triggers"`
}

func (r AuthProviderUpdateRequestOauth2UserInfoRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthProviderUpdateRequestOauth2UserInfoRequestRequestContentType string

const (
	AuthProviderUpdateRequestOauth2UserInfoRequestRequestContentTypeApplicationXWwwFormUrlencoded AuthProviderUpdateRequestOauth2UserInfoRequestRequestContentType = "application/x-www-form-urlencoded"
	AuthProviderUpdateRequestOauth2UserInfoRequestRequestContentTypeApplicationJson               AuthProviderUpdateRequestOauth2UserInfoRequestRequestContentType = "application/json"
)

func (r AuthProviderUpdateRequestOauth2UserInfoRequestRequestContentType) IsKnown() bool {
	switch r {
	case AuthProviderUpdateRequestOauth2UserInfoRequestRequestContentTypeApplicationXWwwFormUrlencoded, AuthProviderUpdateRequestOauth2UserInfoRequestRequestContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderUpdateRequestOauth2UserInfoRequestResponseContentType string

const (
	AuthProviderUpdateRequestOauth2UserInfoRequestResponseContentTypeApplicationXWwwFormUrlencoded AuthProviderUpdateRequestOauth2UserInfoRequestResponseContentType = "application/x-www-form-urlencoded"
	AuthProviderUpdateRequestOauth2UserInfoRequestResponseContentTypeApplicationJson               AuthProviderUpdateRequestOauth2UserInfoRequestResponseContentType = "application/json"
)

func (r AuthProviderUpdateRequestOauth2UserInfoRequestResponseContentType) IsKnown() bool {
	switch r {
	case AuthProviderUpdateRequestOauth2UserInfoRequestResponseContentTypeApplicationXWwwFormUrlencoded, AuthProviderUpdateRequestOauth2UserInfoRequestResponseContentTypeApplicationJson:
		return true
	}
	return false
}

type AuthProviderUpdateRequestOauth2UserInfoRequestTriggersParam struct {
	OnTokenGrant   param.Field[bool] `json:"on_token_grant"`
	OnTokenRefresh param.Field[bool] `json:"on_token_refresh"`
}

func (r AuthProviderUpdateRequestOauth2UserInfoRequestTriggersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdminAuthProviderListResponse struct {
	Items      []AuthProviderResponse            `json:"items"`
	Limit      int64                             `json:"limit"`
	Offset     int64                             `json:"offset"`
	PageCount  int64                             `json:"page_count"`
	TotalCount int64                             `json:"total_count"`
	JSON       adminAuthProviderListResponseJSON `json:"-"`
}

// adminAuthProviderListResponseJSON contains the JSON metadata for the struct
// [AdminAuthProviderListResponse]
type adminAuthProviderListResponseJSON struct {
	Items       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	PageCount   apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdminAuthProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminAuthProviderListResponseJSON) RawJSON() string {
	return r.raw
}

type AdminAuthProviderNewParams struct {
	AuthProviderCreateRequest AuthProviderCreateRequestParam `json:"auth_provider_create_request,required"`
}

func (r AdminAuthProviderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AuthProviderCreateRequest)
}

type AdminAuthProviderPatchParams struct {
	AuthProviderUpdateRequest AuthProviderUpdateRequestParam `json:"auth_provider_update_request,required"`
}

func (r AdminAuthProviderPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AuthProviderUpdateRequest)
}
