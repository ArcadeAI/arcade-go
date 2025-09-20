// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadego

import (
	"context"
	"net/http"
	"slices"

	"github.com/ArcadeAI/arcade-go/internal/apijson"
	"github.com/ArcadeAI/arcade-go/internal/requestconfig"
	"github.com/ArcadeAI/arcade-go/option"
)

// ChatCompletionService contains methods and other services that help with
// interacting with the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatCompletionService] method instead.
type ChatCompletionService struct {
	Options []option.RequestOption
}

// NewChatCompletionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatCompletionService(opts ...option.RequestOption) (r *ChatCompletionService) {
	r = &ChatCompletionService{}
	r.Options = opts
	return
}

// Interact with language models via OpenAI's chat completions API
func (r *ChatCompletionService) New(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (res *ChatResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ChatCompletionNewParams struct {
	ChatRequest ChatRequestParam `json:"chat_request,required"`
}

func (r ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ChatRequest)
}
