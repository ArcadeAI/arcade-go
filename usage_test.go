// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadeengine_test

import (
	"context"
	"os"
	"testing"

	"github.com/ArcadeAI/arcade-go"
	"github.com/ArcadeAI/arcade-go/internal/testutil"
	"github.com/ArcadeAI/arcade-go/option"
)

func TestUsage(t *testing.T) {
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
	executeToolResponse, err := client.Tools.Execute(context.TODO(), arcadeengine.ToolExecuteParams{
		ExecuteToolRequest: arcadeengine.ExecuteToolRequestParam{
			ToolName: "tool_name",
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", executeToolResponse.ID)
}
