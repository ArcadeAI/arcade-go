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
	"github.com/stainless-sdks/arcade-engine-go/shared"
)

// ToolService contains methods and other services that help with interacting with
// the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolService] method instead.
type ToolService struct {
	Options   []option.RequestOption
	Scheduled ToolScheduledService
	Formatted ToolFormattedService
}

// NewToolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolService(opts ...option.RequestOption) (r ToolService) {
	r = ToolService{}
	r.Options = opts
	r.Scheduled = NewToolScheduledService(opts...)
	r.Formatted = NewToolFormattedService(opts...)
	return
}

// Returns a page of tools from the engine configuration, optionally filtered by
// toolkit
func (r *ToolService) List(ctx context.Context, query ToolListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ToolDefinition], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/tools"
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
// toolkit
func (r *ToolService) ListAutoPaging(ctx context.Context, query ToolListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ToolDefinition] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Authorizes a user for a specific tool by name
func (r *ToolService) Authorize(ctx context.Context, body ToolAuthorizeParams, opts ...option.RequestOption) (res *shared.AuthorizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tools/authorize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Executes a tool by name and arguments
func (r *ToolService) Execute(ctx context.Context, body ToolExecuteParams, opts ...option.RequestOption) (res *ExecuteToolResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tools/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the arcade tool specification for a specific tool
func (r *ToolService) Get(ctx context.Context, name string, query ToolGetParams, opts ...option.RequestOption) (res *ToolDefinition, err error) {
	opts = append(r.Options[:], opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("v1/tools/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The property ToolName is required.
type AuthorizeToolRequestParam struct {
	ToolName string `json:"tool_name,required"`
	// Optional: if not provided, any version is used
	ToolVersion param.Opt[string] `json:"tool_version,omitzero"`
	// Required only when calling with an API key
	UserID param.Opt[string] `json:"user_id,omitzero"`
	paramObj
}

func (r AuthorizeToolRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AuthorizeToolRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthorizeToolRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ToolName is required.
type ExecuteToolRequestParam struct {
	ToolName string `json:"tool_name,required"`
	// The time at which the tool should be run (optional). If not provided, the tool
	// is run immediately. Format ISO 8601: YYYY-MM-DDTHH:MM:SS
	RunAt param.Opt[string] `json:"run_at,omitzero"`
	// The tool version to use (optional). If not provided, any version is used
	ToolVersion param.Opt[string] `json:"tool_version,omitzero"`
	UserID      param.Opt[string] `json:"user_id,omitzero"`
	// JSON input to the tool, if any
	Input map[string]any `json:"input,omitzero"`
	paramObj
}

func (r ExecuteToolRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExecuteToolRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExecuteToolRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecuteToolResponse struct {
	ID            string                    `json:"id"`
	Duration      float64                   `json:"duration"`
	ExecutionID   string                    `json:"execution_id"`
	ExecutionType string                    `json:"execution_type"`
	FinishedAt    string                    `json:"finished_at"`
	Output        ExecuteToolResponseOutput `json:"output"`
	RunAt         string                    `json:"run_at"`
	Status        string                    `json:"status"`
	// Whether the request was successful. For immediately-executed requests, this will
	// be true if the tool call succeeded. For scheduled requests, this will be true if
	// the request was scheduled successfully.
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Duration      respjson.Field
		ExecutionID   respjson.Field
		ExecutionType respjson.Field
		FinishedAt    respjson.Field
		Output        respjson.Field
		RunAt         respjson.Field
		Status        respjson.Field
		Success       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteToolResponse) RawJSON() string { return r.JSON.raw }
func (r *ExecuteToolResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecuteToolResponseOutput struct {
	Authorization shared.AuthorizationResponse   `json:"authorization"`
	Error         ExecuteToolResponseOutputError `json:"error"`
	Logs          []ExecuteToolResponseOutputLog `json:"logs"`
	Value         any                            `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Authorization respjson.Field
		Error         respjson.Field
		Logs          respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteToolResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *ExecuteToolResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecuteToolResponseOutputError struct {
	Message                 string `json:"message,required"`
	AdditionalPromptContent string `json:"additional_prompt_content"`
	CanRetry                bool   `json:"can_retry"`
	DeveloperMessage        string `json:"developer_message"`
	RetryAfterMs            int64  `json:"retry_after_ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message                 respjson.Field
		AdditionalPromptContent respjson.Field
		CanRetry                respjson.Field
		DeveloperMessage        respjson.Field
		RetryAfterMs            respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteToolResponseOutputError) RawJSON() string { return r.JSON.raw }
func (r *ExecuteToolResponseOutputError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecuteToolResponseOutputLog struct {
	Level   string `json:"level,required"`
	Message string `json:"message,required"`
	Subtype string `json:"subtype"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		Message     respjson.Field
		Subtype     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteToolResponseOutputLog) RawJSON() string { return r.JSON.raw }
func (r *ExecuteToolResponseOutputLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolDefinition struct {
	FullyQualifiedName string                     `json:"fully_qualified_name,required"`
	Input              ToolDefinitionInput        `json:"input,required"`
	Name               string                     `json:"name,required"`
	QualifiedName      string                     `json:"qualified_name,required"`
	Toolkit            ToolDefinitionToolkit      `json:"toolkit,required"`
	Description        string                     `json:"description"`
	FormattedSchema    map[string]any             `json:"formatted_schema"`
	Output             ToolDefinitionOutput       `json:"output"`
	Requirements       ToolDefinitionRequirements `json:"requirements"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullyQualifiedName respjson.Field
		Input              respjson.Field
		Name               respjson.Field
		QualifiedName      respjson.Field
		Toolkit            respjson.Field
		Description        respjson.Field
		FormattedSchema    respjson.Field
		Output             respjson.Field
		Requirements       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDefinition) RawJSON() string { return r.JSON.raw }
func (r *ToolDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolDefinitionInput struct {
	Parameters []ToolDefinitionInputParameter `json:"parameters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDefinitionInput) RawJSON() string { return r.JSON.raw }
func (r *ToolDefinitionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolDefinitionInputParameter struct {
	Name        string      `json:"name,required"`
	ValueSchema ValueSchema `json:"value_schema,required"`
	Description string      `json:"description"`
	Inferrable  bool        `json:"inferrable"`
	Required    bool        `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ValueSchema respjson.Field
		Description respjson.Field
		Inferrable  respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDefinitionInputParameter) RawJSON() string { return r.JSON.raw }
func (r *ToolDefinitionInputParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolDefinitionToolkit struct {
	Name        string `json:"name,required"`
	Description string `json:"description"`
	Version     string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDefinitionToolkit) RawJSON() string { return r.JSON.raw }
func (r *ToolDefinitionToolkit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolDefinitionOutput struct {
	AvailableModes []string    `json:"available_modes"`
	Description    string      `json:"description"`
	ValueSchema    ValueSchema `json:"value_schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableModes respjson.Field
		Description    respjson.Field
		ValueSchema    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDefinitionOutput) RawJSON() string { return r.JSON.raw }
func (r *ToolDefinitionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolDefinitionRequirements struct {
	Authorization ToolDefinitionRequirementsAuthorization `json:"authorization"`
	Secrets       []ToolDefinitionRequirementsSecret      `json:"secrets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Authorization respjson.Field
		Secrets       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDefinitionRequirements) RawJSON() string { return r.JSON.raw }
func (r *ToolDefinitionRequirements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolDefinitionRequirementsAuthorization struct {
	ID           string                                        `json:"id"`
	Oauth2       ToolDefinitionRequirementsAuthorizationOauth2 `json:"oauth2"`
	ProviderID   string                                        `json:"provider_id"`
	ProviderType string                                        `json:"provider_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Oauth2       respjson.Field
		ProviderID   respjson.Field
		ProviderType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDefinitionRequirementsAuthorization) RawJSON() string { return r.JSON.raw }
func (r *ToolDefinitionRequirementsAuthorization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolDefinitionRequirementsAuthorizationOauth2 struct {
	Scopes []string `json:"scopes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDefinitionRequirementsAuthorizationOauth2) RawJSON() string { return r.JSON.raw }
func (r *ToolDefinitionRequirementsAuthorizationOauth2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolDefinitionRequirementsSecret struct {
	Key string `json:"key,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDefinitionRequirementsSecret) RawJSON() string { return r.JSON.raw }
func (r *ToolDefinitionRequirementsSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolExecution struct {
	ID              string `json:"id"`
	CreatedAt       string `json:"created_at"`
	ExecutionStatus string `json:"execution_status"`
	ExecutionType   string `json:"execution_type"`
	FinishedAt      string `json:"finished_at"`
	RunAt           string `json:"run_at"`
	StartedAt       string `json:"started_at"`
	ToolName        string `json:"tool_name"`
	ToolkitName     string `json:"toolkit_name"`
	ToolkitVersion  string `json:"toolkit_version"`
	UpdatedAt       string `json:"updated_at"`
	UserID          string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		ExecutionStatus respjson.Field
		ExecutionType   respjson.Field
		FinishedAt      respjson.Field
		RunAt           respjson.Field
		StartedAt       respjson.Field
		ToolName        respjson.Field
		ToolkitName     respjson.Field
		ToolkitVersion  respjson.Field
		UpdatedAt       respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolExecution) RawJSON() string { return r.JSON.raw }
func (r *ToolExecution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolExecutionAttempt struct {
	ID                 string                     `json:"id"`
	FinishedAt         string                     `json:"finished_at"`
	Output             ToolExecutionAttemptOutput `json:"output"`
	StartedAt          string                     `json:"started_at"`
	Success            bool                       `json:"success"`
	SystemErrorMessage string                     `json:"system_error_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		FinishedAt         respjson.Field
		Output             respjson.Field
		StartedAt          respjson.Field
		Success            respjson.Field
		SystemErrorMessage respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolExecutionAttempt) RawJSON() string { return r.JSON.raw }
func (r *ToolExecutionAttempt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolExecutionAttemptOutput struct {
	Authorization shared.AuthorizationResponse    `json:"authorization"`
	Error         ToolExecutionAttemptOutputError `json:"error"`
	Logs          []ToolExecutionAttemptOutputLog `json:"logs"`
	Value         any                             `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Authorization respjson.Field
		Error         respjson.Field
		Logs          respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolExecutionAttemptOutput) RawJSON() string { return r.JSON.raw }
func (r *ToolExecutionAttemptOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolExecutionAttemptOutputError struct {
	Message                 string `json:"message,required"`
	AdditionalPromptContent string `json:"additional_prompt_content"`
	CanRetry                bool   `json:"can_retry"`
	DeveloperMessage        string `json:"developer_message"`
	RetryAfterMs            int64  `json:"retry_after_ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message                 respjson.Field
		AdditionalPromptContent respjson.Field
		CanRetry                respjson.Field
		DeveloperMessage        respjson.Field
		RetryAfterMs            respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolExecutionAttemptOutputError) RawJSON() string { return r.JSON.raw }
func (r *ToolExecutionAttemptOutputError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolExecutionAttemptOutputLog struct {
	Level   string `json:"level,required"`
	Message string `json:"message,required"`
	Subtype string `json:"subtype"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		Message     respjson.Field
		Subtype     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolExecutionAttemptOutputLog) RawJSON() string { return r.JSON.raw }
func (r *ToolExecutionAttemptOutputLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ValueSchema struct {
	ValType      string   `json:"val_type,required"`
	Enum         []string `json:"enum"`
	InnerValType string   `json:"inner_val_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ValType      respjson.Field
		Enum         respjson.Field
		InnerValType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ValueSchema) RawJSON() string { return r.JSON.raw }
func (r *ValueSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolListParams struct {
	// Number of items to return (default: 25, max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset from the start of the list (default: 0)
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Toolkit name
	Toolkit param.Opt[string] `query:"toolkit,omitzero" json:"-"`
	// Comma separated tool formats that will be included in the response.
	//
	// Any of "arcade", "openai", "anthropic".
	IncludeFormat []string `query:"include_format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolListParams]'s query parameters as `url.Values`.
func (r ToolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolAuthorizeParams struct {
	AuthorizeToolRequest AuthorizeToolRequestParam
	paramObj
}

func (r ToolAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.AuthorizeToolRequest)
}
func (r *ToolAuthorizeParams) UnmarshalJSON(data []byte) error {
	return r.AuthorizeToolRequest.UnmarshalJSON(data)
}

type ToolExecuteParams struct {
	ExecuteToolRequest ExecuteToolRequestParam
	paramObj
}

func (r ToolExecuteParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.ExecuteToolRequest)
}
func (r *ToolExecuteParams) UnmarshalJSON(data []byte) error {
	return r.ExecuteToolRequest.UnmarshalJSON(data)
}

type ToolGetParams struct {
	// Comma separated tool formats that will be included in the response.
	//
	// Any of "arcade", "openai", "anthropic".
	IncludeFormat []string `query:"include_format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolGetParams]'s query parameters as `url.Values`.
func (r ToolGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
