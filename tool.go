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
	"github.com/ArcadeAI/arcade-go/shared"
)

// ToolService contains methods and other services that help with interacting with
// the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolService] method instead.
type ToolService struct {
	Options   []option.RequestOption
	Scheduled *ToolScheduledService
	Formatted *ToolFormattedService
}

// NewToolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolService(opts ...option.RequestOption) (r *ToolService) {
	r = &ToolService{}
	r.Options = opts
	r.Scheduled = NewToolScheduledService(opts...)
	r.Formatted = NewToolFormattedService(opts...)
	return
}

// Returns a page of tools from the engine configuration, optionally filtered by
// toolkit
func (r *ToolService) List(ctx context.Context, query ToolListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ToolDefinition], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	path := "v1/tools/authorize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Executes a tool by name and arguments
func (r *ToolService) Execute(ctx context.Context, body ToolExecuteParams, opts ...option.RequestOption) (res *ExecuteToolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tools/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the arcade tool specification for a specific tool
func (r *ToolService) Get(ctx context.Context, name string, query ToolGetParams, opts ...option.RequestOption) (res *ToolDefinition, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("v1/tools/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AuthorizeToolRequestParam struct {
	ToolName param.Field[string] `json:"tool_name,required"`
	// Optional: if provided, the user will be redirected to this URI after
	// authorization
	NextUri param.Field[string] `json:"next_uri"`
	// Optional: if not provided, any version is used
	ToolVersion param.Field[string] `json:"tool_version"`
	// Required only when calling with an API key
	UserID param.Field[string] `json:"user_id"`
}

func (r AuthorizeToolRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExecuteToolRequestParam struct {
	ToolName param.Field[string] `json:"tool_name,required"`
	// Whether to include the error stacktrace in the response. If not provided, the
	// error stacktrace is not included.
	IncludeErrorStacktrace param.Field[bool] `json:"include_error_stacktrace"`
	// JSON input to the tool, if any
	Input param.Field[map[string]interface{}] `json:"input"`
	// The time at which the tool should be run (optional). If not provided, the tool
	// is run immediately. Format ISO 8601: YYYY-MM-DDTHH:MM:SS
	RunAt param.Field[string] `json:"run_at"`
	// The tool version to use (optional). If not provided, any version is used
	ToolVersion param.Field[string] `json:"tool_version"`
	UserID      param.Field[string] `json:"user_id"`
}

func (r ExecuteToolRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Success bool                    `json:"success"`
	JSON    executeToolResponseJSON `json:"-"`
}

// executeToolResponseJSON contains the JSON metadata for the struct
// [ExecuteToolResponse]
type executeToolResponseJSON struct {
	ID            apijson.Field
	Duration      apijson.Field
	ExecutionID   apijson.Field
	ExecutionType apijson.Field
	FinishedAt    apijson.Field
	Output        apijson.Field
	RunAt         apijson.Field
	Status        apijson.Field
	Success       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ExecuteToolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeToolResponseJSON) RawJSON() string {
	return r.raw
}

type ExecuteToolResponseOutput struct {
	Authorization shared.AuthorizationResponse   `json:"authorization"`
	Error         ExecuteToolResponseOutputError `json:"error"`
	Logs          []ExecuteToolResponseOutputLog `json:"logs"`
	Value         interface{}                    `json:"value"`
	JSON          executeToolResponseOutputJSON  `json:"-"`
}

// executeToolResponseOutputJSON contains the JSON metadata for the struct
// [ExecuteToolResponseOutput]
type executeToolResponseOutputJSON struct {
	Authorization apijson.Field
	Error         apijson.Field
	Logs          apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ExecuteToolResponseOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeToolResponseOutputJSON) RawJSON() string {
	return r.raw
}

type ExecuteToolResponseOutputError struct {
	CanRetry                bool                               `json:"can_retry,required"`
	Kind                    ExecuteToolResponseOutputErrorKind `json:"kind,required"`
	Message                 string                             `json:"message,required"`
	AdditionalPromptContent string                             `json:"additional_prompt_content"`
	DeveloperMessage        string                             `json:"developer_message"`
	Extra                   map[string]interface{}             `json:"extra"`
	RetryAfterMs            int64                              `json:"retry_after_ms"`
	Stacktrace              string                             `json:"stacktrace"`
	StatusCode              int64                              `json:"status_code"`
	JSON                    executeToolResponseOutputErrorJSON `json:"-"`
}

// executeToolResponseOutputErrorJSON contains the JSON metadata for the struct
// [ExecuteToolResponseOutputError]
type executeToolResponseOutputErrorJSON struct {
	CanRetry                apijson.Field
	Kind                    apijson.Field
	Message                 apijson.Field
	AdditionalPromptContent apijson.Field
	DeveloperMessage        apijson.Field
	Extra                   apijson.Field
	RetryAfterMs            apijson.Field
	Stacktrace              apijson.Field
	StatusCode              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ExecuteToolResponseOutputError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeToolResponseOutputErrorJSON) RawJSON() string {
	return r.raw
}

type ExecuteToolResponseOutputErrorKind string

const (
	ExecuteToolResponseOutputErrorKindToolkitLoadFailed              ExecuteToolResponseOutputErrorKind = "TOOLKIT_LOAD_FAILED"
	ExecuteToolResponseOutputErrorKindToolDefinitionBadDefinition    ExecuteToolResponseOutputErrorKind = "TOOL_DEFINITION_BAD_DEFINITION"
	ExecuteToolResponseOutputErrorKindToolDefinitionBadInputSchema   ExecuteToolResponseOutputErrorKind = "TOOL_DEFINITION_BAD_INPUT_SCHEMA"
	ExecuteToolResponseOutputErrorKindToolDefinitionBadOutputSchema  ExecuteToolResponseOutputErrorKind = "TOOL_DEFINITION_BAD_OUTPUT_SCHEMA"
	ExecuteToolResponseOutputErrorKindToolRequirementsNotMet         ExecuteToolResponseOutputErrorKind = "TOOL_REQUIREMENTS_NOT_MET"
	ExecuteToolResponseOutputErrorKindToolRuntimeBadInputValue       ExecuteToolResponseOutputErrorKind = "TOOL_RUNTIME_BAD_INPUT_VALUE"
	ExecuteToolResponseOutputErrorKindToolRuntimeBadOutputValue      ExecuteToolResponseOutputErrorKind = "TOOL_RUNTIME_BAD_OUTPUT_VALUE"
	ExecuteToolResponseOutputErrorKindToolRuntimeRetry               ExecuteToolResponseOutputErrorKind = "TOOL_RUNTIME_RETRY"
	ExecuteToolResponseOutputErrorKindToolRuntimeContextRequired     ExecuteToolResponseOutputErrorKind = "TOOL_RUNTIME_CONTEXT_REQUIRED"
	ExecuteToolResponseOutputErrorKindToolRuntimeFatal               ExecuteToolResponseOutputErrorKind = "TOOL_RUNTIME_FATAL"
	ExecuteToolResponseOutputErrorKindUpstreamRuntimeBadRequest      ExecuteToolResponseOutputErrorKind = "UPSTREAM_RUNTIME_BAD_REQUEST"
	ExecuteToolResponseOutputErrorKindUpstreamRuntimeAuthError       ExecuteToolResponseOutputErrorKind = "UPSTREAM_RUNTIME_AUTH_ERROR"
	ExecuteToolResponseOutputErrorKindUpstreamRuntimeNotFound        ExecuteToolResponseOutputErrorKind = "UPSTREAM_RUNTIME_NOT_FOUND"
	ExecuteToolResponseOutputErrorKindUpstreamRuntimeValidationError ExecuteToolResponseOutputErrorKind = "UPSTREAM_RUNTIME_VALIDATION_ERROR"
	ExecuteToolResponseOutputErrorKindUpstreamRuntimeRateLimit       ExecuteToolResponseOutputErrorKind = "UPSTREAM_RUNTIME_RATE_LIMIT"
	ExecuteToolResponseOutputErrorKindUpstreamRuntimeServerError     ExecuteToolResponseOutputErrorKind = "UPSTREAM_RUNTIME_SERVER_ERROR"
	ExecuteToolResponseOutputErrorKindUpstreamRuntimeUnmapped        ExecuteToolResponseOutputErrorKind = "UPSTREAM_RUNTIME_UNMAPPED"
	ExecuteToolResponseOutputErrorKindUnknown                        ExecuteToolResponseOutputErrorKind = "UNKNOWN"
)

func (r ExecuteToolResponseOutputErrorKind) IsKnown() bool {
	switch r {
	case ExecuteToolResponseOutputErrorKindToolkitLoadFailed, ExecuteToolResponseOutputErrorKindToolDefinitionBadDefinition, ExecuteToolResponseOutputErrorKindToolDefinitionBadInputSchema, ExecuteToolResponseOutputErrorKindToolDefinitionBadOutputSchema, ExecuteToolResponseOutputErrorKindToolRequirementsNotMet, ExecuteToolResponseOutputErrorKindToolRuntimeBadInputValue, ExecuteToolResponseOutputErrorKindToolRuntimeBadOutputValue, ExecuteToolResponseOutputErrorKindToolRuntimeRetry, ExecuteToolResponseOutputErrorKindToolRuntimeContextRequired, ExecuteToolResponseOutputErrorKindToolRuntimeFatal, ExecuteToolResponseOutputErrorKindUpstreamRuntimeBadRequest, ExecuteToolResponseOutputErrorKindUpstreamRuntimeAuthError, ExecuteToolResponseOutputErrorKindUpstreamRuntimeNotFound, ExecuteToolResponseOutputErrorKindUpstreamRuntimeValidationError, ExecuteToolResponseOutputErrorKindUpstreamRuntimeRateLimit, ExecuteToolResponseOutputErrorKindUpstreamRuntimeServerError, ExecuteToolResponseOutputErrorKindUpstreamRuntimeUnmapped, ExecuteToolResponseOutputErrorKindUnknown:
		return true
	}
	return false
}

type ExecuteToolResponseOutputLog struct {
	Level   string                           `json:"level,required"`
	Message string                           `json:"message,required"`
	Subtype string                           `json:"subtype"`
	JSON    executeToolResponseOutputLogJSON `json:"-"`
}

// executeToolResponseOutputLogJSON contains the JSON metadata for the struct
// [ExecuteToolResponseOutputLog]
type executeToolResponseOutputLogJSON struct {
	Level       apijson.Field
	Message     apijson.Field
	Subtype     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecuteToolResponseOutputLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeToolResponseOutputLogJSON) RawJSON() string {
	return r.raw
}

type ToolDefinition struct {
	FullyQualifiedName string                     `json:"fully_qualified_name,required"`
	Input              ToolDefinitionInput        `json:"input,required"`
	Name               string                     `json:"name,required"`
	QualifiedName      string                     `json:"qualified_name,required"`
	Toolkit            ToolDefinitionToolkit      `json:"toolkit,required"`
	Description        string                     `json:"description"`
	FormattedSchema    map[string]interface{}     `json:"formatted_schema"`
	Output             ToolDefinitionOutput       `json:"output"`
	Requirements       ToolDefinitionRequirements `json:"requirements"`
	JSON               toolDefinitionJSON         `json:"-"`
}

// toolDefinitionJSON contains the JSON metadata for the struct [ToolDefinition]
type toolDefinitionJSON struct {
	FullyQualifiedName apijson.Field
	Input              apijson.Field
	Name               apijson.Field
	QualifiedName      apijson.Field
	Toolkit            apijson.Field
	Description        apijson.Field
	FormattedSchema    apijson.Field
	Output             apijson.Field
	Requirements       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ToolDefinition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolDefinitionJSON) RawJSON() string {
	return r.raw
}

type ToolDefinitionInput struct {
	Parameters []ToolDefinitionInputParameter `json:"parameters"`
	JSON       toolDefinitionInputJSON        `json:"-"`
}

// toolDefinitionInputJSON contains the JSON metadata for the struct
// [ToolDefinitionInput]
type toolDefinitionInputJSON struct {
	Parameters  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolDefinitionInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolDefinitionInputJSON) RawJSON() string {
	return r.raw
}

type ToolDefinitionInputParameter struct {
	Name        string                           `json:"name,required"`
	ValueSchema ValueSchema                      `json:"value_schema,required"`
	Description string                           `json:"description"`
	Inferrable  bool                             `json:"inferrable"`
	Required    bool                             `json:"required"`
	JSON        toolDefinitionInputParameterJSON `json:"-"`
}

// toolDefinitionInputParameterJSON contains the JSON metadata for the struct
// [ToolDefinitionInputParameter]
type toolDefinitionInputParameterJSON struct {
	Name        apijson.Field
	ValueSchema apijson.Field
	Description apijson.Field
	Inferrable  apijson.Field
	Required    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolDefinitionInputParameter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolDefinitionInputParameterJSON) RawJSON() string {
	return r.raw
}

type ToolDefinitionToolkit struct {
	Name        string                    `json:"name,required"`
	Description string                    `json:"description"`
	Version     string                    `json:"version"`
	JSON        toolDefinitionToolkitJSON `json:"-"`
}

// toolDefinitionToolkitJSON contains the JSON metadata for the struct
// [ToolDefinitionToolkit]
type toolDefinitionToolkitJSON struct {
	Name        apijson.Field
	Description apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolDefinitionToolkit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolDefinitionToolkitJSON) RawJSON() string {
	return r.raw
}

type ToolDefinitionOutput struct {
	AvailableModes []string                 `json:"available_modes"`
	Description    string                   `json:"description"`
	ValueSchema    ValueSchema              `json:"value_schema"`
	JSON           toolDefinitionOutputJSON `json:"-"`
}

// toolDefinitionOutputJSON contains the JSON metadata for the struct
// [ToolDefinitionOutput]
type toolDefinitionOutputJSON struct {
	AvailableModes apijson.Field
	Description    apijson.Field
	ValueSchema    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ToolDefinitionOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolDefinitionOutputJSON) RawJSON() string {
	return r.raw
}

type ToolDefinitionRequirements struct {
	Authorization ToolDefinitionRequirementsAuthorization `json:"authorization"`
	Met           bool                                    `json:"met"`
	Secrets       []ToolDefinitionRequirementsSecret      `json:"secrets"`
	JSON          toolDefinitionRequirementsJSON          `json:"-"`
}

// toolDefinitionRequirementsJSON contains the JSON metadata for the struct
// [ToolDefinitionRequirements]
type toolDefinitionRequirementsJSON struct {
	Authorization apijson.Field
	Met           apijson.Field
	Secrets       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ToolDefinitionRequirements) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolDefinitionRequirementsJSON) RawJSON() string {
	return r.raw
}

type ToolDefinitionRequirementsAuthorization struct {
	ID           string                                             `json:"id"`
	Oauth2       ToolDefinitionRequirementsAuthorizationOauth2      `json:"oauth2"`
	ProviderID   string                                             `json:"provider_id"`
	ProviderType string                                             `json:"provider_type"`
	Status       ToolDefinitionRequirementsAuthorizationStatus      `json:"status"`
	StatusReason string                                             `json:"status_reason"`
	TokenStatus  ToolDefinitionRequirementsAuthorizationTokenStatus `json:"token_status"`
	JSON         toolDefinitionRequirementsAuthorizationJSON        `json:"-"`
}

// toolDefinitionRequirementsAuthorizationJSON contains the JSON metadata for the
// struct [ToolDefinitionRequirementsAuthorization]
type toolDefinitionRequirementsAuthorizationJSON struct {
	ID           apijson.Field
	Oauth2       apijson.Field
	ProviderID   apijson.Field
	ProviderType apijson.Field
	Status       apijson.Field
	StatusReason apijson.Field
	TokenStatus  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ToolDefinitionRequirementsAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolDefinitionRequirementsAuthorizationJSON) RawJSON() string {
	return r.raw
}

type ToolDefinitionRequirementsAuthorizationOauth2 struct {
	Scopes []string                                          `json:"scopes"`
	JSON   toolDefinitionRequirementsAuthorizationOauth2JSON `json:"-"`
}

// toolDefinitionRequirementsAuthorizationOauth2JSON contains the JSON metadata for
// the struct [ToolDefinitionRequirementsAuthorizationOauth2]
type toolDefinitionRequirementsAuthorizationOauth2JSON struct {
	Scopes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolDefinitionRequirementsAuthorizationOauth2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolDefinitionRequirementsAuthorizationOauth2JSON) RawJSON() string {
	return r.raw
}

type ToolDefinitionRequirementsAuthorizationStatus string

const (
	ToolDefinitionRequirementsAuthorizationStatusActive   ToolDefinitionRequirementsAuthorizationStatus = "active"
	ToolDefinitionRequirementsAuthorizationStatusInactive ToolDefinitionRequirementsAuthorizationStatus = "inactive"
)

func (r ToolDefinitionRequirementsAuthorizationStatus) IsKnown() bool {
	switch r {
	case ToolDefinitionRequirementsAuthorizationStatusActive, ToolDefinitionRequirementsAuthorizationStatusInactive:
		return true
	}
	return false
}

type ToolDefinitionRequirementsAuthorizationTokenStatus string

const (
	ToolDefinitionRequirementsAuthorizationTokenStatusNotStarted ToolDefinitionRequirementsAuthorizationTokenStatus = "not_started"
	ToolDefinitionRequirementsAuthorizationTokenStatusPending    ToolDefinitionRequirementsAuthorizationTokenStatus = "pending"
	ToolDefinitionRequirementsAuthorizationTokenStatusCompleted  ToolDefinitionRequirementsAuthorizationTokenStatus = "completed"
	ToolDefinitionRequirementsAuthorizationTokenStatusFailed     ToolDefinitionRequirementsAuthorizationTokenStatus = "failed"
)

func (r ToolDefinitionRequirementsAuthorizationTokenStatus) IsKnown() bool {
	switch r {
	case ToolDefinitionRequirementsAuthorizationTokenStatusNotStarted, ToolDefinitionRequirementsAuthorizationTokenStatusPending, ToolDefinitionRequirementsAuthorizationTokenStatusCompleted, ToolDefinitionRequirementsAuthorizationTokenStatusFailed:
		return true
	}
	return false
}

type ToolDefinitionRequirementsSecret struct {
	Key          string                               `json:"key,required"`
	Met          bool                                 `json:"met"`
	StatusReason string                               `json:"status_reason"`
	JSON         toolDefinitionRequirementsSecretJSON `json:"-"`
}

// toolDefinitionRequirementsSecretJSON contains the JSON metadata for the struct
// [ToolDefinitionRequirementsSecret]
type toolDefinitionRequirementsSecretJSON struct {
	Key          apijson.Field
	Met          apijson.Field
	StatusReason apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ToolDefinitionRequirementsSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolDefinitionRequirementsSecretJSON) RawJSON() string {
	return r.raw
}

type ToolExecution struct {
	ID              string            `json:"id"`
	CreatedAt       string            `json:"created_at"`
	ExecutionStatus string            `json:"execution_status"`
	ExecutionType   string            `json:"execution_type"`
	FinishedAt      string            `json:"finished_at"`
	RunAt           string            `json:"run_at"`
	StartedAt       string            `json:"started_at"`
	ToolName        string            `json:"tool_name"`
	ToolkitName     string            `json:"toolkit_name"`
	ToolkitVersion  string            `json:"toolkit_version"`
	UpdatedAt       string            `json:"updated_at"`
	UserID          string            `json:"user_id"`
	JSON            toolExecutionJSON `json:"-"`
}

// toolExecutionJSON contains the JSON metadata for the struct [ToolExecution]
type toolExecutionJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	ExecutionStatus apijson.Field
	ExecutionType   apijson.Field
	FinishedAt      apijson.Field
	RunAt           apijson.Field
	StartedAt       apijson.Field
	ToolName        apijson.Field
	ToolkitName     apijson.Field
	ToolkitVersion  apijson.Field
	UpdatedAt       apijson.Field
	UserID          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ToolExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolExecutionJSON) RawJSON() string {
	return r.raw
}

type ToolExecutionAttempt struct {
	ID                 string                     `json:"id"`
	FinishedAt         string                     `json:"finished_at"`
	Output             ToolExecutionAttemptOutput `json:"output"`
	StartedAt          string                     `json:"started_at"`
	Success            bool                       `json:"success"`
	SystemErrorMessage string                     `json:"system_error_message"`
	JSON               toolExecutionAttemptJSON   `json:"-"`
}

// toolExecutionAttemptJSON contains the JSON metadata for the struct
// [ToolExecutionAttempt]
type toolExecutionAttemptJSON struct {
	ID                 apijson.Field
	FinishedAt         apijson.Field
	Output             apijson.Field
	StartedAt          apijson.Field
	Success            apijson.Field
	SystemErrorMessage apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ToolExecutionAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolExecutionAttemptJSON) RawJSON() string {
	return r.raw
}

type ToolExecutionAttemptOutput struct {
	Authorization shared.AuthorizationResponse    `json:"authorization"`
	Error         ToolExecutionAttemptOutputError `json:"error"`
	Logs          []ToolExecutionAttemptOutputLog `json:"logs"`
	Value         interface{}                     `json:"value"`
	JSON          toolExecutionAttemptOutputJSON  `json:"-"`
}

// toolExecutionAttemptOutputJSON contains the JSON metadata for the struct
// [ToolExecutionAttemptOutput]
type toolExecutionAttemptOutputJSON struct {
	Authorization apijson.Field
	Error         apijson.Field
	Logs          apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ToolExecutionAttemptOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolExecutionAttemptOutputJSON) RawJSON() string {
	return r.raw
}

type ToolExecutionAttemptOutputError struct {
	CanRetry                bool                                `json:"can_retry,required"`
	Kind                    ToolExecutionAttemptOutputErrorKind `json:"kind,required"`
	Message                 string                              `json:"message,required"`
	AdditionalPromptContent string                              `json:"additional_prompt_content"`
	DeveloperMessage        string                              `json:"developer_message"`
	Extra                   map[string]interface{}              `json:"extra"`
	RetryAfterMs            int64                               `json:"retry_after_ms"`
	Stacktrace              string                              `json:"stacktrace"`
	StatusCode              int64                               `json:"status_code"`
	JSON                    toolExecutionAttemptOutputErrorJSON `json:"-"`
}

// toolExecutionAttemptOutputErrorJSON contains the JSON metadata for the struct
// [ToolExecutionAttemptOutputError]
type toolExecutionAttemptOutputErrorJSON struct {
	CanRetry                apijson.Field
	Kind                    apijson.Field
	Message                 apijson.Field
	AdditionalPromptContent apijson.Field
	DeveloperMessage        apijson.Field
	Extra                   apijson.Field
	RetryAfterMs            apijson.Field
	Stacktrace              apijson.Field
	StatusCode              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ToolExecutionAttemptOutputError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolExecutionAttemptOutputErrorJSON) RawJSON() string {
	return r.raw
}

type ToolExecutionAttemptOutputErrorKind string

const (
	ToolExecutionAttemptOutputErrorKindToolkitLoadFailed              ToolExecutionAttemptOutputErrorKind = "TOOLKIT_LOAD_FAILED"
	ToolExecutionAttemptOutputErrorKindToolDefinitionBadDefinition    ToolExecutionAttemptOutputErrorKind = "TOOL_DEFINITION_BAD_DEFINITION"
	ToolExecutionAttemptOutputErrorKindToolDefinitionBadInputSchema   ToolExecutionAttemptOutputErrorKind = "TOOL_DEFINITION_BAD_INPUT_SCHEMA"
	ToolExecutionAttemptOutputErrorKindToolDefinitionBadOutputSchema  ToolExecutionAttemptOutputErrorKind = "TOOL_DEFINITION_BAD_OUTPUT_SCHEMA"
	ToolExecutionAttemptOutputErrorKindToolRequirementsNotMet         ToolExecutionAttemptOutputErrorKind = "TOOL_REQUIREMENTS_NOT_MET"
	ToolExecutionAttemptOutputErrorKindToolRuntimeBadInputValue       ToolExecutionAttemptOutputErrorKind = "TOOL_RUNTIME_BAD_INPUT_VALUE"
	ToolExecutionAttemptOutputErrorKindToolRuntimeBadOutputValue      ToolExecutionAttemptOutputErrorKind = "TOOL_RUNTIME_BAD_OUTPUT_VALUE"
	ToolExecutionAttemptOutputErrorKindToolRuntimeRetry               ToolExecutionAttemptOutputErrorKind = "TOOL_RUNTIME_RETRY"
	ToolExecutionAttemptOutputErrorKindToolRuntimeContextRequired     ToolExecutionAttemptOutputErrorKind = "TOOL_RUNTIME_CONTEXT_REQUIRED"
	ToolExecutionAttemptOutputErrorKindToolRuntimeFatal               ToolExecutionAttemptOutputErrorKind = "TOOL_RUNTIME_FATAL"
	ToolExecutionAttemptOutputErrorKindUpstreamRuntimeBadRequest      ToolExecutionAttemptOutputErrorKind = "UPSTREAM_RUNTIME_BAD_REQUEST"
	ToolExecutionAttemptOutputErrorKindUpstreamRuntimeAuthError       ToolExecutionAttemptOutputErrorKind = "UPSTREAM_RUNTIME_AUTH_ERROR"
	ToolExecutionAttemptOutputErrorKindUpstreamRuntimeNotFound        ToolExecutionAttemptOutputErrorKind = "UPSTREAM_RUNTIME_NOT_FOUND"
	ToolExecutionAttemptOutputErrorKindUpstreamRuntimeValidationError ToolExecutionAttemptOutputErrorKind = "UPSTREAM_RUNTIME_VALIDATION_ERROR"
	ToolExecutionAttemptOutputErrorKindUpstreamRuntimeRateLimit       ToolExecutionAttemptOutputErrorKind = "UPSTREAM_RUNTIME_RATE_LIMIT"
	ToolExecutionAttemptOutputErrorKindUpstreamRuntimeServerError     ToolExecutionAttemptOutputErrorKind = "UPSTREAM_RUNTIME_SERVER_ERROR"
	ToolExecutionAttemptOutputErrorKindUpstreamRuntimeUnmapped        ToolExecutionAttemptOutputErrorKind = "UPSTREAM_RUNTIME_UNMAPPED"
	ToolExecutionAttemptOutputErrorKindUnknown                        ToolExecutionAttemptOutputErrorKind = "UNKNOWN"
)

func (r ToolExecutionAttemptOutputErrorKind) IsKnown() bool {
	switch r {
	case ToolExecutionAttemptOutputErrorKindToolkitLoadFailed, ToolExecutionAttemptOutputErrorKindToolDefinitionBadDefinition, ToolExecutionAttemptOutputErrorKindToolDefinitionBadInputSchema, ToolExecutionAttemptOutputErrorKindToolDefinitionBadOutputSchema, ToolExecutionAttemptOutputErrorKindToolRequirementsNotMet, ToolExecutionAttemptOutputErrorKindToolRuntimeBadInputValue, ToolExecutionAttemptOutputErrorKindToolRuntimeBadOutputValue, ToolExecutionAttemptOutputErrorKindToolRuntimeRetry, ToolExecutionAttemptOutputErrorKindToolRuntimeContextRequired, ToolExecutionAttemptOutputErrorKindToolRuntimeFatal, ToolExecutionAttemptOutputErrorKindUpstreamRuntimeBadRequest, ToolExecutionAttemptOutputErrorKindUpstreamRuntimeAuthError, ToolExecutionAttemptOutputErrorKindUpstreamRuntimeNotFound, ToolExecutionAttemptOutputErrorKindUpstreamRuntimeValidationError, ToolExecutionAttemptOutputErrorKindUpstreamRuntimeRateLimit, ToolExecutionAttemptOutputErrorKindUpstreamRuntimeServerError, ToolExecutionAttemptOutputErrorKindUpstreamRuntimeUnmapped, ToolExecutionAttemptOutputErrorKindUnknown:
		return true
	}
	return false
}

type ToolExecutionAttemptOutputLog struct {
	Level   string                            `json:"level,required"`
	Message string                            `json:"message,required"`
	Subtype string                            `json:"subtype"`
	JSON    toolExecutionAttemptOutputLogJSON `json:"-"`
}

// toolExecutionAttemptOutputLogJSON contains the JSON metadata for the struct
// [ToolExecutionAttemptOutputLog]
type toolExecutionAttemptOutputLogJSON struct {
	Level       apijson.Field
	Message     apijson.Field
	Subtype     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolExecutionAttemptOutputLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolExecutionAttemptOutputLogJSON) RawJSON() string {
	return r.raw
}

type ValueSchema struct {
	ValType      string          `json:"val_type,required"`
	Enum         []string        `json:"enum"`
	InnerValType string          `json:"inner_val_type"`
	JSON         valueSchemaJSON `json:"-"`
}

// valueSchemaJSON contains the JSON metadata for the struct [ValueSchema]
type valueSchemaJSON struct {
	ValType      apijson.Field
	Enum         apijson.Field
	InnerValType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ValueSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r valueSchemaJSON) RawJSON() string {
	return r.raw
}

type ToolListParams struct {
	// Comma separated tool formats that will be included in the response.
	IncludeFormat param.Field[[]ToolListParamsIncludeFormat] `query:"include_format"`
	// Number of items to return (default: 25, max: 100)
	Limit param.Field[int64] `query:"limit"`
	// Offset from the start of the list (default: 0)
	Offset param.Field[int64] `query:"offset"`
	// Toolkit name
	Toolkit param.Field[string] `query:"toolkit"`
	// User ID
	UserID param.Field[string] `query:"user_id"`
}

// URLQuery serializes [ToolListParams]'s query parameters as `url.Values`.
func (r ToolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolListParamsIncludeFormat string

const (
	ToolListParamsIncludeFormatArcade    ToolListParamsIncludeFormat = "arcade"
	ToolListParamsIncludeFormatOpenAI    ToolListParamsIncludeFormat = "openai"
	ToolListParamsIncludeFormatAnthropic ToolListParamsIncludeFormat = "anthropic"
)

func (r ToolListParamsIncludeFormat) IsKnown() bool {
	switch r {
	case ToolListParamsIncludeFormatArcade, ToolListParamsIncludeFormatOpenAI, ToolListParamsIncludeFormatAnthropic:
		return true
	}
	return false
}

type ToolAuthorizeParams struct {
	AuthorizeToolRequest AuthorizeToolRequestParam `json:"authorize_tool_request,required"`
}

func (r ToolAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AuthorizeToolRequest)
}

type ToolExecuteParams struct {
	ExecuteToolRequest ExecuteToolRequestParam `json:"execute_tool_request,required"`
}

func (r ToolExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ExecuteToolRequest)
}

type ToolGetParams struct {
	// Comma separated tool formats that will be included in the response.
	IncludeFormat param.Field[[]ToolGetParamsIncludeFormat] `query:"include_format"`
	// User ID
	UserID param.Field[string] `query:"user_id"`
}

// URLQuery serializes [ToolGetParams]'s query parameters as `url.Values`.
func (r ToolGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolGetParamsIncludeFormat string

const (
	ToolGetParamsIncludeFormatArcade    ToolGetParamsIncludeFormat = "arcade"
	ToolGetParamsIncludeFormatOpenAI    ToolGetParamsIncludeFormat = "openai"
	ToolGetParamsIncludeFormatAnthropic ToolGetParamsIncludeFormat = "anthropic"
)

func (r ToolGetParamsIncludeFormat) IsKnown() bool {
	switch r {
	case ToolGetParamsIncludeFormatArcade, ToolGetParamsIncludeFormatOpenAI, ToolGetParamsIncludeFormatAnthropic:
		return true
	}
	return false
}
