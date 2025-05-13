// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadeengine

import (
	"github.com/ArcadeAI/arcade-go/internal/apijson"
	"github.com/ArcadeAI/arcade-go/option"
	"github.com/ArcadeAI/arcade-go/packages/param"
	"github.com/ArcadeAI/arcade-go/packages/respjson"
	"github.com/ArcadeAI/arcade-go/shared"
)

// ChatService contains methods and other services that help with interacting with
// the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	Options     []option.RequestOption
	Completions ChatCompletionService
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r ChatService) {
	r = ChatService{}
	r.Options = opts
	r.Completions = NewChatCompletionService(opts...)
	return
}

type ChatMessage struct {
	// The content of the message.
	Content string `json:"content,required"`
	// The role of the author of this message. One of system, user, tool, or assistant.
	Role string `json:"role,required"`
	// tool Name
	Name string `json:"name"`
	// tool_call_id
	ToolCallID string `json:"tool_call_id"`
	// tool calls if any
	ToolCalls []ChatMessageToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Name        respjson.Field
		ToolCallID  respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChatMessage to a ChatMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChatMessageParam.Overrides()
func (r ChatMessage) ToParam() ChatMessageParam {
	return param.Override[ChatMessageParam](r.RawJSON())
}

type ChatMessageToolCall struct {
	ID       string                      `json:"id"`
	Function ChatMessageToolCallFunction `json:"function"`
	// Any of "function".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Function    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageToolCallFunction struct {
	Arguments string `json:"arguments"`
	Name      string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatMessageToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *ChatMessageToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, Role are required.
type ChatMessageParam struct {
	// The content of the message.
	Content string `json:"content,required"`
	// The role of the author of this message. One of system, user, tool, or assistant.
	Role string `json:"role,required"`
	// tool Name
	Name param.Opt[string] `json:"name,omitzero"`
	// tool_call_id
	ToolCallID param.Opt[string] `json:"tool_call_id,omitzero"`
	// tool calls if any
	ToolCalls []ChatMessageToolCallParam `json:"tool_calls,omitzero"`
	paramObj
}

func (r ChatMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatMessageToolCallParam struct {
	ID       param.Opt[string]                `json:"id,omitzero"`
	Function ChatMessageToolCallFunctionParam `json:"function,omitzero"`
	// Any of "function".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatMessageToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatMessageToolCallParam](
		"type", "function",
	)
}

type ChatMessageToolCallFunctionParam struct {
	Arguments param.Opt[string] `json:"arguments,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ChatMessageToolCallFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatMessageToolCallFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatMessageToolCallFunctionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatRequestParam struct {
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// LogProbs indicates whether to return log probabilities of the output tokens or
	// not. If true, returns the log probabilities of each output token returned in the
	// content of message. This option is currently not available on the
	// gpt-4-vision-preview model.
	Logprobs  param.Opt[bool]   `json:"logprobs,omitzero"`
	MaxTokens param.Opt[int64]  `json:"max_tokens,omitzero"`
	Model     param.Opt[string] `json:"model,omitzero"`
	N         param.Opt[int64]  `json:"n,omitzero"`
	// Disable the default behavior of parallel tool calls by setting it: false.
	ParallelToolCalls param.Opt[bool]    `json:"parallel_tool_calls,omitzero"`
	PresencePenalty   param.Opt[float64] `json:"presence_penalty,omitzero"`
	Seed              param.Opt[int64]   `json:"seed,omitzero"`
	Stream            param.Opt[bool]    `json:"stream,omitzero"`
	Temperature       param.Opt[float64] `json:"temperature,omitzero"`
	// TopLogProbs is an integer between 0 and 5 specifying the number of most likely
	// tokens to return at each token position, each with an associated log
	// probability. logprobs must be set to true if this parameter is used.
	TopLogprobs param.Opt[int64]   `json:"top_logprobs,omitzero"`
	TopP        param.Opt[float64] `json:"top_p,omitzero"`
	User        param.Opt[string]  `json:"user,omitzero"`
	// LogitBias is must be a token id string (specified by their token ID in the
	// tokenizer), not a word string. incorrect: `"logit_bias":{"You": 6}`, correct:
	// `"logit_bias":{"1639": 6}` refs:
	// https://platform.openai.com/docs/api-reference/chat/create#chat/create-logit_bias
	LogitBias      map[string]int64               `json:"logit_bias,omitzero"`
	Messages       []ChatMessageParam             `json:"messages,omitzero"`
	ResponseFormat ChatRequestResponseFormatParam `json:"response_format,omitzero"`
	Stop           []string                       `json:"stop,omitzero"`
	// Options for streaming response. Only set this when you set stream: true.
	StreamOptions ChatRequestStreamOptionsParam `json:"stream_options,omitzero"`
	// This can be either a string or an ToolChoice object.
	ToolChoice any `json:"tool_choice,omitzero"`
	Tools      any `json:"tools,omitzero"`
	paramObj
}

func (r ChatRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatRequestResponseFormatParam struct {
	// Any of "json_object", "text".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ChatRequestResponseFormatParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatRequestResponseFormatParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatRequestResponseFormatParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatRequestResponseFormatParam](
		"type", "json_object", "text",
	)
}

// Options for streaming response. Only set this when you set stream: true.
type ChatRequestStreamOptionsParam struct {
	// If set, an additional chunk will be streamed before the data: [DONE] message.
	// The usage field on this chunk shows the token usage statistics for the entire
	// request, and the choices field will always be an empty array. All other chunks
	// will also include a usage field, but with a null value.
	IncludeUsage param.Opt[bool] `json:"include_usage,omitzero"`
	paramObj
}

func (r ChatRequestStreamOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatRequestStreamOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatRequestStreamOptionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatResponse struct {
	ID                string   `json:"id"`
	Choices           []Choice `json:"choices"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Object            string   `json:"object"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Usage             Usage    `json:"usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Choices           respjson.Field
		Created           respjson.Field
		Model             respjson.Field
		Object            respjson.Field
		SystemFingerprint respjson.Field
		Usage             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Choice struct {
	FinishReason       string                         `json:"finish_reason"`
	Index              int64                          `json:"index"`
	Logprobs           any                            `json:"logprobs"`
	Message            ChatMessage                    `json:"message"`
	ToolAuthorizations []shared.AuthorizationResponse `json:"tool_authorizations"`
	ToolMessages       []ChatMessage                  `json:"tool_messages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason       respjson.Field
		Index              respjson.Field
		Logprobs           respjson.Field
		Message            respjson.Field
		ToolAuthorizations respjson.Field
		ToolMessages       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Choice) RawJSON() string { return r.JSON.raw }
func (r *Choice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Usage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionTokens respjson.Field
		PromptTokens     respjson.Field
		TotalTokens      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Usage) RawJSON() string { return r.JSON.raw }
func (r *Usage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
