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

func TestToolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.List(context.TODO(), arcadego.ToolListParams{
		IncludeFormat: arcadego.F([]arcadego.ToolListParamsIncludeFormat{arcadego.ToolListParamsIncludeFormatArcade}),
		Limit:         arcadego.F(int64(0)),
		Offset:        arcadego.F(int64(0)),
		Toolkit:       arcadego.F("toolkit"),
	})
	if err != nil {
		var apierr *arcadego.Error
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
	client := arcadego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Tools.Authorize(context.TODO(), arcadego.ToolAuthorizeParams{
		AuthorizeToolRequest: arcadego.AuthorizeToolRequestParam{
			ToolName:    arcadego.F("tool_name"),
			ToolVersion: arcadego.F("tool_version"),
			UserID:      arcadego.F("user_id"),
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

func TestToolExecuteWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.Execute(context.TODO(), arcadego.ToolExecuteParams{
		ExecuteToolRequest: arcadego.ExecuteToolRequestParam{
			ToolName: arcadego.F("tool_name"),
			Input: arcadego.F(map[string]interface{}{
				"foo": "bar",
			}),
			RunAt:       arcadego.F("run_at"),
			ToolVersion: arcadego.F("tool_version"),
			UserID:      arcadego.F("user_id"),
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

func TestToolGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.Get(
		context.TODO(),
		"name",
		arcadego.ToolGetParams{
			IncludeFormat: arcadego.F([]arcadego.ToolGetParamsIncludeFormat{arcadego.ToolGetParamsIncludeFormatArcade}),
		},
	)
	if err != nil {
		var apierr *arcadego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
