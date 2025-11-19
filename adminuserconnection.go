// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadego

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/ArcadeAI/arcade-go/internal/apijson"
	"github.com/ArcadeAI/arcade-go/internal/apiquery"
	"github.com/ArcadeAI/arcade-go/internal/param"
	"github.com/ArcadeAI/arcade-go/internal/requestconfig"
	"github.com/ArcadeAI/arcade-go/option"
	"github.com/ArcadeAI/arcade-go/packages/pagination"
)

// AdminUserConnectionService contains methods and other services that help with
// interacting with the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminUserConnectionService] method instead.
type AdminUserConnectionService struct {
	Options []option.RequestOption
}

// NewAdminUserConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminUserConnectionService(opts ...option.RequestOption) (r *AdminUserConnectionService) {
	r = &AdminUserConnectionService{}
	r.Options = opts
	return
}

// List all auth connections
func (r *AdminUserConnectionService) List(ctx context.Context, query AdminUserConnectionListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[UserConnectionResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/admin/user_connections"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all auth connections
func (r *AdminUserConnectionService) ListAutoPaging(ctx context.Context, query AdminUserConnectionListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[UserConnectionResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a user/auth provider connection
func (r *AdminUserConnectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/admin/user_connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type UserConnectionResponse struct {
	ID                  string                     `json:"id"`
	ConnectionID        string                     `json:"connection_id"`
	ConnectionStatus    string                     `json:"connection_status"`
	ProviderDescription string                     `json:"provider_description"`
	ProviderID          string                     `json:"provider_id"`
	ProviderType        string                     `json:"provider_type"`
	ProviderUserInfo    interface{}                `json:"provider_user_info"`
	Scopes              []string                   `json:"scopes"`
	UserID              string                     `json:"user_id"`
	JSON                userConnectionResponseJSON `json:"-"`
}

// userConnectionResponseJSON contains the JSON metadata for the struct
// [UserConnectionResponse]
type userConnectionResponseJSON struct {
	ID                  apijson.Field
	ConnectionID        apijson.Field
	ConnectionStatus    apijson.Field
	ProviderDescription apijson.Field
	ProviderID          apijson.Field
	ProviderType        apijson.Field
	ProviderUserInfo    apijson.Field
	Scopes              apijson.Field
	UserID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *UserConnectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userConnectionResponseJSON) RawJSON() string {
	return r.raw
}

type AdminUserConnectionListParams struct {
	// Page size
	Limit param.Field[int64] `query:"limit"`
	// Page offset
	Offset   param.Field[int64]                                 `query:"offset"`
	Provider param.Field[AdminUserConnectionListParamsProvider] `query:"provider"`
	User     param.Field[AdminUserConnectionListParamsUser]     `query:"user"`
}

// URLQuery serializes [AdminUserConnectionListParams]'s query parameters as
// `url.Values`.
func (r AdminUserConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminUserConnectionListParamsProvider struct {
	// Provider ID
	ID param.Field[string] `query:"id"`
}

// URLQuery serializes [AdminUserConnectionListParamsProvider]'s query parameters
// as `url.Values`.
func (r AdminUserConnectionListParamsProvider) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminUserConnectionListParamsUser struct {
	// User ID
	ID param.Field[string] `query:"id"`
}

// URLQuery serializes [AdminUserConnectionListParamsUser]'s query parameters as
// `url.Values`.
func (r AdminUserConnectionListParamsUser) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
