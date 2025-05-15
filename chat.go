// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadego

import (
	"github.com/ArcadeAI/arcade-go/internal/apijson"
	"github.com/ArcadeAI/arcade-go/internal/param"
	"github.com/ArcadeAI/arcade-go/option"
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
	Completions *ChatCompletionService
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r *ChatService) {
	r = &ChatService{}
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
	JSON      chatMessageJSON       `json:"-"`
}

// chatMessageJSON contains the JSON metadata for the struct [ChatMessage]
type chatMessageJSON struct {
	Content     apijson.Field
	Role        apijson.Field
	Name        apijson.Field
	ToolCallID  apijson.Field
	ToolCalls   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChatMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatMessageJSON) RawJSON() string {
	return r.raw
}

type ChatMessageToolCall struct {
	ID       string                       `json:"id"`
	Function ChatMessageToolCallsFunction `json:"function"`
	Type     ChatMessageToolCallsType     `json:"type"`
	JSON     chatMessageToolCallJSON      `json:"-"`
}

// chatMessageToolCallJSON contains the JSON metadata for the struct
// [ChatMessageToolCall]
type chatMessageToolCallJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChatMessageToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatMessageToolCallJSON) RawJSON() string {
	return r.raw
}

type ChatMessageToolCallsFunction struct {
	Arguments string                           `json:"arguments"`
	Name      string                           `json:"name"`
	JSON      chatMessageToolCallsFunctionJSON `json:"-"`
}

// chatMessageToolCallsFunctionJSON contains the JSON metadata for the struct
// [ChatMessageToolCallsFunction]
type chatMessageToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChatMessageToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatMessageToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type ChatMessageToolCallsType string

const (
	ChatMessageToolCallsTypeFunction ChatMessageToolCallsType = "function"
)

func (r ChatMessageToolCallsType) IsKnown() bool {
	switch r {
	case ChatMessageToolCallsTypeFunction:
		return true
	}
	return false
}

type ChatMessageParam struct {
	// The content of the message.
	Content param.Field[string] `json:"content,required"`
	// The role of the author of this message. One of system, user, tool, or assistant.
	Role param.Field[string] `json:"role,required"`
	// tool Name
	Name param.Field[string] `json:"name"`
	// tool_call_id
	ToolCallID param.Field[string] `json:"tool_call_id"`
	// tool calls if any
	ToolCalls param.Field[[]ChatMessageToolCallParam] `json:"tool_calls"`
}

func (r ChatMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatMessageToolCallParam struct {
	ID       param.Field[string]                            `json:"id"`
	Function param.Field[ChatMessageToolCallsFunctionParam] `json:"function"`
	Type     param.Field[ChatMessageToolCallsType]          `json:"type"`
}

func (r ChatMessageToolCallParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatMessageToolCallsFunctionParam struct {
	Arguments param.Field[string] `json:"arguments"`
	Name      param.Field[string] `json:"name"`
}

func (r ChatMessageToolCallsFunctionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatRequestParam struct {
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// LogitBias is must be a token id string (specified by their token ID in the
	// tokenizer), not a word string. incorrect: `"logit_bias":{"You": 6}`, correct:
	// `"logit_bias":{"1639": 6}` refs:
	// https://platform.openai.com/docs/api-reference/chat/create#chat/create-logit_bias
	LogitBias param.Field[map[string]int64] `json:"logit_bias"`
	// LogProbs indicates whether to return log probabilities of the output tokens or
	// not. If true, returns the log probabilities of each output token returned in the
	// content of message. This option is currently not available on the
	// gpt-4-vision-preview model.
	Logprobs  param.Field[bool]               `json:"logprobs"`
	MaxTokens param.Field[int64]              `json:"max_tokens"`
	Messages  param.Field[[]ChatMessageParam] `json:"messages"`
	Model     param.Field[string]             `json:"model"`
	N         param.Field[int64]              `json:"n"`
	// Disable the default behavior of parallel tool calls by setting it: false.
	ParallelToolCalls param.Field[bool]                           `json:"parallel_tool_calls"`
	PresencePenalty   param.Field[float64]                        `json:"presence_penalty"`
	ResponseFormat    param.Field[ChatRequestResponseFormatParam] `json:"response_format"`
	Seed              param.Field[int64]                          `json:"seed"`
	Stop              param.Field[[]string]                       `json:"stop"`
	Stream            param.Field[bool]                           `json:"stream"`
	// Options for streaming response. Only set this when you set stream: true.
	StreamOptions param.Field[ChatRequestStreamOptionsParam] `json:"stream_options"`
	Temperature   param.Field[float64]                       `json:"temperature"`
	// This can be either a string or an ToolChoice object.
	ToolChoice param.Field[interface{}] `json:"tool_choice"`
	Tools      param.Field[interface{}] `json:"tools"`
	// TopLogProbs is an integer between 0 and 5 specifying the number of most likely
	// tokens to return at each token position, each with an associated log
	// probability. logprobs must be set to true if this parameter is used.
	TopLogprobs param.Field[int64]   `json:"top_logprobs"`
	TopP        param.Field[float64] `json:"top_p"`
	User        param.Field[string]  `json:"user"`
}

func (r ChatRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatRequestResponseFormatParam struct {
	Type param.Field[ChatRequestResponseFormatType] `json:"type"`
}

func (r ChatRequestResponseFormatParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatRequestResponseFormatType string

const (
	ChatRequestResponseFormatTypeJsonObject ChatRequestResponseFormatType = "json_object"
	ChatRequestResponseFormatTypeText       ChatRequestResponseFormatType = "text"
)

func (r ChatRequestResponseFormatType) IsKnown() bool {
	switch r {
	case ChatRequestResponseFormatTypeJsonObject, ChatRequestResponseFormatTypeText:
		return true
	}
	return false
}

// Options for streaming response. Only set this when you set stream: true.
type ChatRequestStreamOptionsParam struct {
	// If set, an additional chunk will be streamed before the data: [DONE] message.
	// The usage field on this chunk shows the token usage statistics for the entire
	// request, and the choices field will always be an empty array. All other chunks
	// will also include a usage field, but with a null value.
	IncludeUsage param.Field[bool] `json:"include_usage"`
}

func (r ChatRequestStreamOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatResponse struct {
	ID                string           `json:"id"`
	Choices           []Choice         `json:"choices"`
	Created           int64            `json:"created"`
	Model             string           `json:"model"`
	Object            string           `json:"object"`
	SystemFingerprint string           `json:"system_fingerprint"`
	Usage             Usage            `json:"usage"`
	JSON              chatResponseJSON `json:"-"`
}

// chatResponseJSON contains the JSON metadata for the struct [ChatResponse]
type chatResponseJSON struct {
	ID                apijson.Field
	Choices           apijson.Field
	Created           apijson.Field
	Model             apijson.Field
	Object            apijson.Field
	SystemFingerprint apijson.Field
	Usage             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ChatResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatResponseJSON) RawJSON() string {
	return r.raw
}

type Choice struct {
	FinishReason       string                         `json:"finish_reason"`
	Index              int64                          `json:"index"`
	Logprobs           interface{}                    `json:"logprobs"`
	Message            ChatMessage                    `json:"message"`
	ToolAuthorizations []shared.AuthorizationResponse `json:"tool_authorizations"`
	ToolMessages       []ChatMessage                  `json:"tool_messages"`
	JSON               choiceJSON                     `json:"-"`
}

// choiceJSON contains the JSON metadata for the struct [Choice]
type choiceJSON struct {
	FinishReason       apijson.Field
	Index              apijson.Field
	Logprobs           apijson.Field
	Message            apijson.Field
	ToolAuthorizations apijson.Field
	ToolMessages       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Choice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r choiceJSON) RawJSON() string {
	return r.raw
}

type Usage struct {
	CompletionTokens int64     `json:"completion_tokens"`
	PromptTokens     int64     `json:"prompt_tokens"`
	TotalTokens      int64     `json:"total_tokens"`
	JSON             usageJSON `json:"-"`
}

// usageJSON contains the JSON metadata for the struct [Usage]
type usageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Usage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageJSON) RawJSON() string {
	return r.raw
}
