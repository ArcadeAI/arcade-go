// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadego_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/ArcadeAI/arcade-go"
	"github.com/ArcadeAI/arcade-go/internal/testutil"
	"github.com/ArcadeAI/arcade-go/option"
)

func TestChatCompletionNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := arcadego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chat.Completions.New(context.TODO(), arcadego.ChatCompletionNewParams{
		ChatRequest: arcadego.ChatRequestParam{
			FrequencyPenalty: arcadego.F(0.000000),
			LogitBias: arcadego.F(map[string]int64{
				"foo": int64(0),
			}),
			Logprobs:  arcadego.F(true),
			MaxTokens: arcadego.F(int64(0)),
			Messages: arcadego.F([]arcadego.ChatMessageParam{{
				Content:    arcadego.F("content"),
				Role:       arcadego.F("role"),
				Name:       arcadego.F("name"),
				ToolCallID: arcadego.F("tool_call_id"),
				ToolCalls: arcadego.F([]arcadego.ChatMessageToolCallParam{{
					ID: arcadego.F("id"),
					Function: arcadego.F(arcadego.ChatMessageToolCallsFunctionParam{
						Arguments: arcadego.F("arguments"),
						Name:      arcadego.F("name"),
					}),
					Type: arcadego.F(arcadego.ChatMessageToolCallsTypeFunction),
				}}),
			}}),
			Model:             arcadego.F("model"),
			N:                 arcadego.F(int64(0)),
			ParallelToolCalls: arcadego.F(true),
			PresencePenalty:   arcadego.F(0.000000),
			ResponseFormat: arcadego.F(arcadego.ChatRequestResponseFormatParam{
				Type: arcadego.F(arcadego.ChatRequestResponseFormatTypeJsonObject),
			}),
			Seed:   arcadego.F(int64(0)),
			Stop:   arcadego.F([]string{"string"}),
			Stream: arcadego.F(true),
			StreamOptions: arcadego.F(arcadego.ChatRequestStreamOptionsParam{
				IncludeUsage: arcadego.F(true),
			}),
			Temperature: arcadego.F(0.000000),
			ToolChoice:  arcadego.F[any](map[string]interface{}{}),
			Tools:       arcadego.F[any](map[string]interface{}{}),
			TopLogprobs: arcadego.F(int64(0)),
			TopP:        arcadego.F(0.000000),
			User:        arcadego.F("user"),
		},
	})
	if err != nil {
		var apierr *arcadego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
