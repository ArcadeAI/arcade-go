// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadeengine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/arcade-engine-go/internal/apijson"
	"github.com/stainless-sdks/arcade-engine-go/internal/apiquery"
	"github.com/stainless-sdks/arcade-engine-go/internal/requestconfig"
	"github.com/stainless-sdks/arcade-engine-go/option"
	"github.com/stainless-sdks/arcade-engine-go/packages/pagination"
	"github.com/stainless-sdks/arcade-engine-go/packages/param"
	"github.com/stainless-sdks/arcade-engine-go/packages/respjson"
)

// WorkerService contains methods and other services that help with interacting
// with the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkerService] method instead.
type WorkerService struct {
	Options []option.RequestOption
}

// NewWorkerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWorkerService(opts ...option.RequestOption) (r WorkerService) {
	r = WorkerService{}
	r.Options = opts
	return
}

// Create a worker
func (r *WorkerService) New(ctx context.Context, body WorkerNewParams, opts ...option.RequestOption) (res *WorkerResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/workers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a worker
func (r *WorkerService) Update(ctx context.Context, id string, body WorkerUpdateParams, opts ...option.RequestOption) (res *WorkerResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/workers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all workers with their definitions
func (r *WorkerService) List(ctx context.Context, query WorkerListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WorkerResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/workers"
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

// List all workers with their definitions
func (r *WorkerService) ListAutoPaging(ctx context.Context, query WorkerListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WorkerResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a worker
func (r *WorkerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/workers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get a worker by ID
func (r *WorkerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WorkerResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/workers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the health of a worker
func (r *WorkerService) Health(ctx context.Context, id string, opts ...option.RequestOption) (res *WorkerHealthResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/workers/%s/health", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a page of tools
func (r *WorkerService) Tools(ctx context.Context, id string, query WorkerToolsParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ToolDefinition], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/workers/%s/tools", id)
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

// Returns a page of tools
func (r *WorkerService) ToolsAutoPaging(ctx context.Context, id string, query WorkerToolsParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ToolDefinition] {
	return pagination.NewOffsetPageAutoPager(r.Tools(ctx, id, query, opts...))
}

// The property ID is required.
type CreateWorkerRequestParam struct {
	ID      string                       `json:"id,required"`
	Enabled param.Opt[bool]              `json:"enabled,omitzero"`
	Type    param.Opt[string]            `json:"type,omitzero"`
	HTTP    CreateWorkerRequestHTTPParam `json:"http,omitzero"`
	Mcp     CreateWorkerRequestMcpParam  `json:"mcp,omitzero"`
	paramObj
}

func (r CreateWorkerRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateWorkerRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateWorkerRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Retry, Secret, Timeout, Uri are required.
type CreateWorkerRequestHTTPParam struct {
	Retry   int64  `json:"retry,required"`
	Secret  string `json:"secret,required"`
	Timeout int64  `json:"timeout,required"`
	Uri     string `json:"uri,required"`
	paramObj
}

func (r CreateWorkerRequestHTTPParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateWorkerRequestHTTPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateWorkerRequestHTTPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Retry, Timeout, Uri are required.
type CreateWorkerRequestMcpParam struct {
	Retry   int64  `json:"retry,required"`
	Timeout int64  `json:"timeout,required"`
	Uri     string `json:"uri,required"`
	paramObj
}

func (r CreateWorkerRequestMcpParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateWorkerRequestMcpParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateWorkerRequestMcpParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateWorkerRequestParam struct {
	Enabled param.Opt[bool]              `json:"enabled,omitzero"`
	HTTP    UpdateWorkerRequestHTTPParam `json:"http,omitzero"`
	Mcp     UpdateWorkerRequestMcpParam  `json:"mcp,omitzero"`
	paramObj
}

func (r UpdateWorkerRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateWorkerRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateWorkerRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateWorkerRequestHTTPParam struct {
	Retry   param.Opt[int64]  `json:"retry,omitzero"`
	Secret  param.Opt[string] `json:"secret,omitzero"`
	Timeout param.Opt[int64]  `json:"timeout,omitzero"`
	Uri     param.Opt[string] `json:"uri,omitzero"`
	paramObj
}

func (r UpdateWorkerRequestHTTPParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateWorkerRequestHTTPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateWorkerRequestHTTPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateWorkerRequestMcpParam struct {
	Retry   param.Opt[int64]  `json:"retry,omitzero"`
	Timeout param.Opt[int64]  `json:"timeout,omitzero"`
	Uri     param.Opt[string] `json:"uri,omitzero"`
	paramObj
}

func (r UpdateWorkerRequestMcpParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateWorkerRequestMcpParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateWorkerRequestMcpParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerHealthResponse struct {
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
	Healthy bool   `json:"healthy"`
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Enabled     respjson.Field
		Healthy     respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerHealthResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkerHealthResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerResponse struct {
	ID      string                `json:"id"`
	Binding WorkerResponseBinding `json:"binding"`
	Enabled bool                  `json:"enabled"`
	HTTP    WorkerResponseHTTP    `json:"http"`
	Mcp     WorkerResponseMcp     `json:"mcp"`
	Oxp     WorkerResponseOxp     `json:"oxp"`
	// Any of "http", "mcp", "unknown".
	Type WorkerResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Binding     respjson.Field
		Enabled     respjson.Field
		HTTP        respjson.Field
		Mcp         respjson.Field
		Oxp         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerResponseBinding struct {
	ID string `json:"id"`
	// Any of "static", "tenant", "project", "account".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerResponseBinding) RawJSON() string { return r.JSON.raw }
func (r *WorkerResponseBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerResponseHTTP struct {
	Retry   int64                    `json:"retry"`
	Secret  WorkerResponseHTTPSecret `json:"secret"`
	Timeout int64                    `json:"timeout"`
	Uri     string                   `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Retry       respjson.Field
		Secret      respjson.Field
		Timeout     respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerResponseHTTP) RawJSON() string { return r.JSON.raw }
func (r *WorkerResponseHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerResponseHTTPSecret struct {
	// Any of "static", "tenant", "project", "account".
	Binding  string `json:"binding"`
	Editable bool   `json:"editable"`
	Exists   bool   `json:"exists"`
	Hint     string `json:"hint"`
	Value    string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Binding     respjson.Field
		Editable    respjson.Field
		Exists      respjson.Field
		Hint        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerResponseHTTPSecret) RawJSON() string { return r.JSON.raw }
func (r *WorkerResponseHTTPSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerResponseMcp struct {
	Retry   int64  `json:"retry"`
	Timeout int64  `json:"timeout"`
	Uri     string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Retry       respjson.Field
		Timeout     respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerResponseMcp) RawJSON() string { return r.JSON.raw }
func (r *WorkerResponseMcp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerResponseOxp struct {
	Retry   int64                   `json:"retry"`
	Secret  WorkerResponseOxpSecret `json:"secret"`
	Timeout int64                   `json:"timeout"`
	Uri     string                  `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Retry       respjson.Field
		Secret      respjson.Field
		Timeout     respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerResponseOxp) RawJSON() string { return r.JSON.raw }
func (r *WorkerResponseOxp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerResponseOxpSecret struct {
	// Any of "static", "tenant", "project", "account".
	Binding  string `json:"binding"`
	Editable bool   `json:"editable"`
	Exists   bool   `json:"exists"`
	Hint     string `json:"hint"`
	Value    string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Binding     respjson.Field
		Editable    respjson.Field
		Exists      respjson.Field
		Hint        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerResponseOxpSecret) RawJSON() string { return r.JSON.raw }
func (r *WorkerResponseOxpSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerResponseType string

const (
	WorkerResponseTypeHTTP    WorkerResponseType = "http"
	WorkerResponseTypeMcp     WorkerResponseType = "mcp"
	WorkerResponseTypeUnknown WorkerResponseType = "unknown"
)

type WorkerNewParams struct {
	CreateWorkerRequest CreateWorkerRequestParam
	paramObj
}

func (r WorkerNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.CreateWorkerRequest)
}
func (r *WorkerNewParams) UnmarshalJSON(data []byte) error {
	return r.CreateWorkerRequest.UnmarshalJSON(data)
}

type WorkerUpdateParams struct {
	UpdateWorkerRequest UpdateWorkerRequestParam
	paramObj
}

func (r WorkerUpdateParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.UpdateWorkerRequest)
}
func (r *WorkerUpdateParams) UnmarshalJSON(data []byte) error {
	return r.UpdateWorkerRequest.UnmarshalJSON(data)
}

type WorkerListParams struct {
	// Number of items to return (default: 25, max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset from the start of the list (default: 0)
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkerListParams]'s query parameters as `url.Values`.
func (r WorkerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkerToolsParams struct {
	// Number of items to return (default: 25, max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset from the start of the list (default: 0)
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkerToolsParams]'s query parameters as `url.Values`.
func (r WorkerToolsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
