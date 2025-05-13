// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadeengine_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/arcade-engine-go"
	"github.com/stainless-sdks/arcade-engine-go/internal/testutil"
	"github.com/stainless-sdks/arcade-engine-go/option"
)

func TestChatCompletionNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := arcadeengine.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chat.Completions.New(context.TODO(), arcadeengine.ChatCompletionNewParams{
		ChatRequest: arcadeengine.ChatRequestParam{
			FrequencyPenalty: arcadeengine.Float(0),
			LogitBias: map[string]int64{
				"foo": 0,
			},
			Logprobs:  arcadeengine.Bool(true),
			MaxTokens: arcadeengine.Int(0),
			Messages: []arcadeengine.ChatMessageParam{{
				Content:    "content",
				Role:       "role",
				Name:       arcadeengine.String("name"),
				ToolCallID: arcadeengine.String("tool_call_id"),
				ToolCalls: []arcadeengine.ChatMessageToolCallParam{{
					ID: arcadeengine.String("id"),
					Function: arcadeengine.ChatMessageToolCallFunctionParam{
						Arguments: arcadeengine.String("arguments"),
						Name:      arcadeengine.String("name"),
					},
					Type: "function",
				}},
			}},
			Model:             arcadeengine.String("model"),
			N:                 arcadeengine.Int(0),
			ParallelToolCalls: arcadeengine.Bool(true),
			PresencePenalty:   arcadeengine.Float(0),
			ResponseFormat: arcadeengine.ChatRequestResponseFormatParam{
				Type: "json_object",
			},
			Seed:   arcadeengine.Int(0),
			Stop:   []string{"string"},
			Stream: arcadeengine.Bool(true),
			StreamOptions: arcadeengine.ChatRequestStreamOptionsParam{
				IncludeUsage: arcadeengine.Bool(true),
			},
			Temperature: arcadeengine.Float(0),
			ToolChoice:  map[string]interface{}{},
			Tools:       map[string]interface{}{},
			TopLogprobs: arcadeengine.Int(0),
			TopP:        arcadeengine.Float(0),
			User:        arcadeengine.String("user"),
		},
	})
	if err != nil {
		var apierr *arcadeengine.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
