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
func NewWorkerService(opts ...option.RequestOption) (r *WorkerService) {
	r = &WorkerService{}
	r.Options = opts
	return
}

// Create a worker
func (r *WorkerService) New(ctx context.Context, body WorkerNewParams, opts ...option.RequestOption) (res *WorkerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/workers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a worker
func (r *WorkerService) Update(ctx context.Context, id string, body WorkerUpdateParams, opts ...option.RequestOption) (res *WorkerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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

type CreateWorkerRequestParam struct {
	ID      param.Field[string]                       `json:"id,required"`
	Enabled param.Field[bool]                         `json:"enabled"`
	HTTP    param.Field[CreateWorkerRequestHTTPParam] `json:"http"`
	Mcp     param.Field[CreateWorkerRequestMcpParam]  `json:"mcp"`
	Type    param.Field[string]                       `json:"type"`
}

func (r CreateWorkerRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreateWorkerRequestHTTPParam struct {
	Retry   param.Field[int64]  `json:"retry,required"`
	Secret  param.Field[string] `json:"secret,required"`
	Timeout param.Field[int64]  `json:"timeout,required"`
	Uri     param.Field[string] `json:"uri,required"`
}

func (r CreateWorkerRequestHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreateWorkerRequestMcpParam struct {
	Retry   param.Field[int64]                             `json:"retry,required"`
	Timeout param.Field[int64]                             `json:"timeout,required"`
	Uri     param.Field[string]                            `json:"uri,required"`
	Headers param.Field[map[string]string]                 `json:"headers"`
	Oauth2  param.Field[CreateWorkerRequestMcpOauth2Param] `json:"oauth2"`
	Secrets param.Field[map[string]string]                 `json:"secrets"`
}

func (r CreateWorkerRequestMcpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreateWorkerRequestMcpOauth2Param struct {
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	ClientID         param.Field[string] `json:"client_id"`
	ClientSecret     param.Field[string] `json:"client_secret"`
	ExternalID       param.Field[string] `json:"external_id"`
}

func (r CreateWorkerRequestMcpOauth2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateWorkerRequestParam struct {
	Enabled param.Field[bool]                         `json:"enabled"`
	HTTP    param.Field[UpdateWorkerRequestHTTPParam] `json:"http"`
	Mcp     param.Field[UpdateWorkerRequestMcpParam]  `json:"mcp"`
}

func (r UpdateWorkerRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateWorkerRequestHTTPParam struct {
	Retry   param.Field[int64]  `json:"retry"`
	Secret  param.Field[string] `json:"secret"`
	Timeout param.Field[int64]  `json:"timeout"`
	Uri     param.Field[string] `json:"uri"`
}

func (r UpdateWorkerRequestHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateWorkerRequestMcpParam struct {
	Headers param.Field[map[string]string]                 `json:"headers"`
	Oauth2  param.Field[UpdateWorkerRequestMcpOauth2Param] `json:"oauth2"`
	Retry   param.Field[int64]                             `json:"retry"`
	Secrets param.Field[map[string]string]                 `json:"secrets"`
	Timeout param.Field[int64]                             `json:"timeout"`
	Uri     param.Field[string]                            `json:"uri"`
}

func (r UpdateWorkerRequestMcpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateWorkerRequestMcpOauth2Param struct {
	AuthorizationURL param.Field[string] `json:"authorization_url"`
	ClientID         param.Field[string] `json:"client_id"`
	ClientSecret     param.Field[string] `json:"client_secret"`
}

func (r UpdateWorkerRequestMcpOauth2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerHealthResponse struct {
	ID      string                   `json:"id"`
	Enabled bool                     `json:"enabled"`
	Healthy bool                     `json:"healthy"`
	Message string                   `json:"message"`
	JSON    workerHealthResponseJSON `json:"-"`
}

// workerHealthResponseJSON contains the JSON metadata for the struct
// [WorkerHealthResponse]
type workerHealthResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Healthy     apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerHealthResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerHealthResponseJSON) RawJSON() string {
	return r.raw
}

type WorkerResponse struct {
	ID           string                     `json:"id"`
	Binding      WorkerResponseBinding      `json:"binding"`
	Enabled      bool                       `json:"enabled"`
	HTTP         WorkerResponseHTTP         `json:"http"`
	Managed      bool                       `json:"managed"`
	Mcp          WorkerResponseMcp          `json:"mcp"`
	Requirements WorkerResponseRequirements `json:"requirements"`
	Type         WorkerResponseType         `json:"type"`
	JSON         workerResponseJSON         `json:"-"`
}

// workerResponseJSON contains the JSON metadata for the struct [WorkerResponse]
type workerResponseJSON struct {
	ID           apijson.Field
	Binding      apijson.Field
	Enabled      apijson.Field
	HTTP         apijson.Field
	Managed      apijson.Field
	Mcp          apijson.Field
	Requirements apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WorkerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerResponseJSON) RawJSON() string {
	return r.raw
}

type WorkerResponseBinding struct {
	ID   string                    `json:"id"`
	Type WorkerResponseBindingType `json:"type"`
	JSON workerResponseBindingJSON `json:"-"`
}

// workerResponseBindingJSON contains the JSON metadata for the struct
// [WorkerResponseBinding]
type workerResponseBindingJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerResponseBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerResponseBindingJSON) RawJSON() string {
	return r.raw
}

type WorkerResponseBindingType string

const (
	WorkerResponseBindingTypeStatic  WorkerResponseBindingType = "static"
	WorkerResponseBindingTypeTenant  WorkerResponseBindingType = "tenant"
	WorkerResponseBindingTypeProject WorkerResponseBindingType = "project"
	WorkerResponseBindingTypeAccount WorkerResponseBindingType = "account"
)

func (r WorkerResponseBindingType) IsKnown() bool {
	switch r {
	case WorkerResponseBindingTypeStatic, WorkerResponseBindingTypeTenant, WorkerResponseBindingTypeProject, WorkerResponseBindingTypeAccount:
		return true
	}
	return false
}

type WorkerResponseHTTP struct {
	Retry   int64                    `json:"retry"`
	Secret  WorkerResponseHTTPSecret `json:"secret"`
	Timeout int64                    `json:"timeout"`
	Uri     string                   `json:"uri"`
	JSON    workerResponseHTTPJSON   `json:"-"`
}

// workerResponseHTTPJSON contains the JSON metadata for the struct
// [WorkerResponseHTTP]
type workerResponseHTTPJSON struct {
	Retry       apijson.Field
	Secret      apijson.Field
	Timeout     apijson.Field
	Uri         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerResponseHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerResponseHTTPJSON) RawJSON() string {
	return r.raw
}

type WorkerResponseHTTPSecret struct {
	Binding  WorkerResponseHTTPSecretBinding `json:"binding"`
	Editable bool                            `json:"editable"`
	Exists   bool                            `json:"exists"`
	Hint     string                          `json:"hint"`
	Value    string                          `json:"value"`
	JSON     workerResponseHTTPSecretJSON    `json:"-"`
}

// workerResponseHTTPSecretJSON contains the JSON metadata for the struct
// [WorkerResponseHTTPSecret]
type workerResponseHTTPSecretJSON struct {
	Binding     apijson.Field
	Editable    apijson.Field
	Exists      apijson.Field
	Hint        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerResponseHTTPSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerResponseHTTPSecretJSON) RawJSON() string {
	return r.raw
}

type WorkerResponseHTTPSecretBinding string

const (
	WorkerResponseHTTPSecretBindingStatic  WorkerResponseHTTPSecretBinding = "static"
	WorkerResponseHTTPSecretBindingTenant  WorkerResponseHTTPSecretBinding = "tenant"
	WorkerResponseHTTPSecretBindingProject WorkerResponseHTTPSecretBinding = "project"
	WorkerResponseHTTPSecretBindingAccount WorkerResponseHTTPSecretBinding = "account"
)

func (r WorkerResponseHTTPSecretBinding) IsKnown() bool {
	switch r {
	case WorkerResponseHTTPSecretBindingStatic, WorkerResponseHTTPSecretBindingTenant, WorkerResponseHTTPSecretBindingProject, WorkerResponseHTTPSecretBindingAccount:
		return true
	}
	return false
}

type WorkerResponseMcp struct {
	Headers map[string]string                  `json:"headers"`
	Oauth2  WorkerResponseMcpOauth2            `json:"oauth2"`
	Retry   int64                              `json:"retry"`
	Secrets map[string]WorkerResponseMcpSecret `json:"secrets"`
	Timeout int64                              `json:"timeout"`
	Uri     string                             `json:"uri"`
	JSON    workerResponseMcpJSON              `json:"-"`
}

// workerResponseMcpJSON contains the JSON metadata for the struct
// [WorkerResponseMcp]
type workerResponseMcpJSON struct {
	Headers     apijson.Field
	Oauth2      apijson.Field
	Retry       apijson.Field
	Secrets     apijson.Field
	Timeout     apijson.Field
	Uri         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerResponseMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerResponseMcpJSON) RawJSON() string {
	return r.raw
}

type WorkerResponseMcpOauth2 struct {
	AuthorizationURL string                              `json:"authorization_url"`
	ClientID         string                              `json:"client_id"`
	ClientSecret     WorkerResponseMcpOauth2ClientSecret `json:"client_secret"`
	RedirectUri      string                              `json:"redirect_uri"`
	JSON             workerResponseMcpOauth2JSON         `json:"-"`
}

// workerResponseMcpOauth2JSON contains the JSON metadata for the struct
// [WorkerResponseMcpOauth2]
type workerResponseMcpOauth2JSON struct {
	AuthorizationURL apijson.Field
	ClientID         apijson.Field
	ClientSecret     apijson.Field
	RedirectUri      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkerResponseMcpOauth2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerResponseMcpOauth2JSON) RawJSON() string {
	return r.raw
}

type WorkerResponseMcpOauth2ClientSecret struct {
	Binding  WorkerResponseMcpOauth2ClientSecretBinding `json:"binding"`
	Editable bool                                       `json:"editable"`
	Exists   bool                                       `json:"exists"`
	Hint     string                                     `json:"hint"`
	Value    string                                     `json:"value"`
	JSON     workerResponseMcpOauth2ClientSecretJSON    `json:"-"`
}

// workerResponseMcpOauth2ClientSecretJSON contains the JSON metadata for the
// struct [WorkerResponseMcpOauth2ClientSecret]
type workerResponseMcpOauth2ClientSecretJSON struct {
	Binding     apijson.Field
	Editable    apijson.Field
	Exists      apijson.Field
	Hint        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerResponseMcpOauth2ClientSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerResponseMcpOauth2ClientSecretJSON) RawJSON() string {
	return r.raw
}

type WorkerResponseMcpOauth2ClientSecretBinding string

const (
	WorkerResponseMcpOauth2ClientSecretBindingStatic  WorkerResponseMcpOauth2ClientSecretBinding = "static"
	WorkerResponseMcpOauth2ClientSecretBindingTenant  WorkerResponseMcpOauth2ClientSecretBinding = "tenant"
	WorkerResponseMcpOauth2ClientSecretBindingProject WorkerResponseMcpOauth2ClientSecretBinding = "project"
	WorkerResponseMcpOauth2ClientSecretBindingAccount WorkerResponseMcpOauth2ClientSecretBinding = "account"
)

func (r WorkerResponseMcpOauth2ClientSecretBinding) IsKnown() bool {
	switch r {
	case WorkerResponseMcpOauth2ClientSecretBindingStatic, WorkerResponseMcpOauth2ClientSecretBindingTenant, WorkerResponseMcpOauth2ClientSecretBindingProject, WorkerResponseMcpOauth2ClientSecretBindingAccount:
		return true
	}
	return false
}

type WorkerResponseMcpSecret struct {
	Binding  WorkerResponseMcpSecretsBinding `json:"binding"`
	Editable bool                            `json:"editable"`
	Exists   bool                            `json:"exists"`
	Hint     string                          `json:"hint"`
	Value    string                          `json:"value"`
	JSON     workerResponseMcpSecretJSON     `json:"-"`
}

// workerResponseMcpSecretJSON contains the JSON metadata for the struct
// [WorkerResponseMcpSecret]
type workerResponseMcpSecretJSON struct {
	Binding     apijson.Field
	Editable    apijson.Field
	Exists      apijson.Field
	Hint        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerResponseMcpSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerResponseMcpSecretJSON) RawJSON() string {
	return r.raw
}

type WorkerResponseMcpSecretsBinding string

const (
	WorkerResponseMcpSecretsBindingStatic  WorkerResponseMcpSecretsBinding = "static"
	WorkerResponseMcpSecretsBindingTenant  WorkerResponseMcpSecretsBinding = "tenant"
	WorkerResponseMcpSecretsBindingProject WorkerResponseMcpSecretsBinding = "project"
	WorkerResponseMcpSecretsBindingAccount WorkerResponseMcpSecretsBinding = "account"
)

func (r WorkerResponseMcpSecretsBinding) IsKnown() bool {
	switch r {
	case WorkerResponseMcpSecretsBindingStatic, WorkerResponseMcpSecretsBindingTenant, WorkerResponseMcpSecretsBindingProject, WorkerResponseMcpSecretsBindingAccount:
		return true
	}
	return false
}

type WorkerResponseRequirements struct {
	Authorization WorkerResponseRequirementsAuthorization `json:"authorization"`
	Met           bool                                    `json:"met"`
	JSON          workerResponseRequirementsJSON          `json:"-"`
}

// workerResponseRequirementsJSON contains the JSON metadata for the struct
// [WorkerResponseRequirements]
type workerResponseRequirementsJSON struct {
	Authorization apijson.Field
	Met           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerResponseRequirements) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerResponseRequirementsJSON) RawJSON() string {
	return r.raw
}

type WorkerResponseRequirementsAuthorization struct {
	Met    bool                                          `json:"met"`
	Oauth2 WorkerResponseRequirementsAuthorizationOauth2 `json:"oauth2"`
	JSON   workerResponseRequirementsAuthorizationJSON   `json:"-"`
}

// workerResponseRequirementsAuthorizationJSON contains the JSON metadata for the
// struct [WorkerResponseRequirementsAuthorization]
type workerResponseRequirementsAuthorizationJSON struct {
	Met         apijson.Field
	Oauth2      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerResponseRequirementsAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerResponseRequirementsAuthorizationJSON) RawJSON() string {
	return r.raw
}

type WorkerResponseRequirementsAuthorizationOauth2 struct {
	Met  bool                                              `json:"met"`
	JSON workerResponseRequirementsAuthorizationOauth2JSON `json:"-"`
}

// workerResponseRequirementsAuthorizationOauth2JSON contains the JSON metadata for
// the struct [WorkerResponseRequirementsAuthorizationOauth2]
type workerResponseRequirementsAuthorizationOauth2JSON struct {
	Met         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerResponseRequirementsAuthorizationOauth2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerResponseRequirementsAuthorizationOauth2JSON) RawJSON() string {
	return r.raw
}

type WorkerResponseType string

const (
	WorkerResponseTypeHTTP    WorkerResponseType = "http"
	WorkerResponseTypeMcp     WorkerResponseType = "mcp"
	WorkerResponseTypeUnknown WorkerResponseType = "unknown"
)

func (r WorkerResponseType) IsKnown() bool {
	switch r {
	case WorkerResponseTypeHTTP, WorkerResponseTypeMcp, WorkerResponseTypeUnknown:
		return true
	}
	return false
}

type WorkerNewParams struct {
	CreateWorkerRequest CreateWorkerRequestParam `json:"create_worker_request,required"`
}

func (r WorkerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateWorkerRequest)
}

type WorkerUpdateParams struct {
	UpdateWorkerRequest UpdateWorkerRequestParam `json:"update_worker_request,required"`
}

func (r WorkerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateWorkerRequest)
}

type WorkerListParams struct {
	// Number of items to return (default: 25, max: 100)
	Limit param.Field[int64] `query:"limit"`
	// Offset from the start of the list (default: 0)
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [WorkerListParams]'s query parameters as `url.Values`.
func (r WorkerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkerToolsParams struct {
	// Number of items to return (default: 25, max: 100)
	Limit param.Field[int64] `query:"limit"`
	// Offset from the start of the list (default: 0)
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [WorkerToolsParams]'s query parameters as `url.Values`.
func (r WorkerToolsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
