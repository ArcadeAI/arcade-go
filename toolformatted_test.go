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

func TestToolFormattedListWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.Formatted.List(context.TODO(), arcadeengine.ToolFormattedListParams{
		Format:  arcadeengine.String("format"),
		Limit:   arcadeengine.Int(0),
		Offset:  arcadeengine.Int(0),
		Toolkit: arcadeengine.String("toolkit"),
	})
	if err != nil {
		var apierr *arcadeengine.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolFormattedGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.Formatted.Get(
		context.TODO(),
		"name",
		arcadeengine.ToolFormattedGetParams{
			Format: arcadeengine.String("format"),
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
