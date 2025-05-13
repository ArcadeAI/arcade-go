// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadeengine

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ArcadeAI/arcade-go/internal/apijson"
	"github.com/ArcadeAI/arcade-go/internal/apiquery"
	"github.com/ArcadeAI/arcade-go/internal/requestconfig"
	"github.com/ArcadeAI/arcade-go/option"
	"github.com/ArcadeAI/arcade-go/packages/pagination"
	"github.com/ArcadeAI/arcade-go/packages/param"
	"github.com/ArcadeAI/arcade-go/packages/respjson"
)

// ToolScheduledService contains methods and other services that help with
// interacting with the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolScheduledService] method instead.
type ToolScheduledService struct {
	Options []option.RequestOption
}

// NewToolScheduledService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolScheduledService(opts ...option.RequestOption) (r ToolScheduledService) {
	r = ToolScheduledService{}
	r.Options = opts
	return
}

// Returns a page of scheduled tool executions
func (r *ToolScheduledService) List(ctx context.Context, query ToolScheduledListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ToolExecution], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/scheduled_tools"
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

// Returns a page of scheduled tool executions
func (r *ToolScheduledService) ListAutoPaging(ctx context.Context, query ToolScheduledListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ToolExecution] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Returns the details for a specific scheduled tool execution
func (r *ToolScheduledService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ToolScheduledGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/scheduled_tools/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ToolScheduledGetResponse struct {
	ID              string                 `json:"id"`
	Attempts        []ToolExecutionAttempt `json:"attempts"`
	CreatedAt       string                 `json:"created_at"`
	ExecutionStatus string                 `json:"execution_status"`
	ExecutionType   string                 `json:"execution_type"`
	FinishedAt      string                 `json:"finished_at"`
	Input           map[string]any         `json:"input"`
	RunAt           string                 `json:"run_at"`
	StartedAt       string                 `json:"started_at"`
	ToolName        string                 `json:"tool_name"`
	ToolkitName     string                 `json:"toolkit_name"`
	ToolkitVersion  string                 `json:"toolkit_version"`
	UpdatedAt       string                 `json:"updated_at"`
	UserID          string                 `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Attempts        respjson.Field
		CreatedAt       respjson.Field
		ExecutionStatus respjson.Field
		ExecutionType   respjson.Field
		FinishedAt      respjson.Field
		Input           respjson.Field
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
func (r ToolScheduledGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ToolScheduledGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolScheduledListParams struct {
	// Number of items to return (default: 25, max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset from the start of the list (default: 0)
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolScheduledListParams]'s query parameters as
// `url.Values`.
func (r ToolScheduledListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
