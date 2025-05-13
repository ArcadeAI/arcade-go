// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadeengine

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/arcade-engine-go/internal/apiquery"
	"github.com/stainless-sdks/arcade-engine-go/internal/requestconfig"
	"github.com/stainless-sdks/arcade-engine-go/option"
	"github.com/stainless-sdks/arcade-engine-go/packages/pagination"
	"github.com/stainless-sdks/arcade-engine-go/packages/param"
)

// ToolFormattedService contains methods and other services that help with
// interacting with the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolFormattedService] method instead.
type ToolFormattedService struct {
	Options []option.RequestOption
}

// NewToolFormattedService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolFormattedService(opts ...option.RequestOption) (r ToolFormattedService) {
	r = ToolFormattedService{}
	r.Options = opts
	return
}

// Returns a page of tools from the engine configuration, optionally filtered by
// toolkit, formatted for a specific provider
func (r *ToolFormattedService) List(ctx context.Context, query ToolFormattedListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ToolFormattedListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/formatted_tools"
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

// Returns a page of tools from the engine configuration, optionally filtered by
// toolkit, formatted for a specific provider
func (r *ToolFormattedService) ListAutoPaging(ctx context.Context, query ToolFormattedListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ToolFormattedListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Returns the formatted tool specification for a specific tool, given a provider
func (r *ToolFormattedService) Get(ctx context.Context, name string, query ToolFormattedGetParams, opts ...option.RequestOption) (res *ToolFormattedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("v1/formatted_tools/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ToolFormattedListResponse = any

type ToolFormattedGetResponse = any

type ToolFormattedListParams struct {
	// Provider format
	Format param.Opt[string] `query:"format,omitzero" json:"-"`
	// Number of items to return (default: 25, max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset from the start of the list (default: 0)
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Toolkit name
	Toolkit param.Opt[string] `query:"toolkit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolFormattedListParams]'s query parameters as
// `url.Values`.
func (r ToolFormattedListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolFormattedGetParams struct {
	// Provider format
	Format param.Opt[string] `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolFormattedGetParams]'s query parameters as `url.Values`.
func (r ToolFormattedGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
