// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadeengine_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/ArcadeAI/arcade-go"
	"github.com/ArcadeAI/arcade-go/internal/testutil"
	"github.com/ArcadeAI/arcade-go/option"
)

func TestToolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.List(context.TODO(), arcadeengine.ToolListParams{
		IncludeFormat: []string{"arcade"},
		Limit:         arcadeengine.Int(0),
		Offset:        arcadeengine.Int(0),
		Toolkit:       arcadeengine.String("toolkit"),
	})
	if err != nil {
		var apierr *arcadeengine.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolAuthorizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.Authorize(context.TODO(), arcadeengine.ToolAuthorizeParams{
		AuthorizeToolRequest: arcadeengine.AuthorizeToolRequestParam{
			ToolName:    "tool_name",
			ToolVersion: arcadeengine.String("tool_version"),
			UserID:      arcadeengine.String("user_id"),
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

func TestToolExecuteWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.Execute(context.TODO(), arcadeengine.ToolExecuteParams{
		ExecuteToolRequest: arcadeengine.ExecuteToolRequestParam{
			ToolName: "tool_name",
			Input: map[string]any{
				"foo": "bar",
			},
			RunAt:       arcadeengine.String("run_at"),
			ToolVersion: arcadeengine.String("tool_version"),
			UserID:      arcadeengine.String("user_id"),
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

func TestToolGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.Get(
		context.TODO(),
		"name",
		arcadeengine.ToolGetParams{
			IncludeFormat: []string{"arcade"},
		},
	)
	if err != nil {
		var apierr *arcadeengine.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
