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

func TestWorkerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.New(context.TODO(), arcadeengine.WorkerNewParams{
		CreateWorkerRequest: arcadeengine.CreateWorkerRequestParam{
			ID:      "id",
			Enabled: arcadeengine.Bool(true),
			HTTP: arcadeengine.CreateWorkerRequestHTTPParam{
				Retry:   0,
				Secret:  "secret",
				Timeout: 1,
				Uri:     "uri",
			},
			Mcp: arcadeengine.CreateWorkerRequestMcpParam{
				Retry:   0,
				Timeout: 1,
				Uri:     "uri",
			},
			Type: arcadeengine.String("type"),
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

func TestWorkerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Update(
		context.TODO(),
		"id",
		arcadeengine.WorkerUpdateParams{
			UpdateWorkerRequest: arcadeengine.UpdateWorkerRequestParam{
				Enabled: arcadeengine.Bool(true),
				HTTP: arcadeengine.UpdateWorkerRequestHTTPParam{
					Retry:   arcadeengine.Int(0),
					Secret:  arcadeengine.String("secret"),
					Timeout: arcadeengine.Int(1),
					Uri:     arcadeengine.String("uri"),
				},
				Mcp: arcadeengine.UpdateWorkerRequestMcpParam{
					Retry:   arcadeengine.Int(0),
					Timeout: arcadeengine.Int(1),
					Uri:     arcadeengine.String("uri"),
				},
			},
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

func TestWorkerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.List(context.TODO(), arcadeengine.WorkerListParams{
		Limit:  arcadeengine.Int(0),
		Offset: arcadeengine.Int(0),
	})
	if err != nil {
		var apierr *arcadeengine.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerDelete(t *testing.T) {
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
	err := client.Workers.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *arcadeengine.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerGet(t *testing.T) {
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
	_, err := client.Workers.Get(context.TODO(), "id")
	if err != nil {
		var apierr *arcadeengine.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerHealth(t *testing.T) {
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
	_, err := client.Workers.Health(context.TODO(), "id")
	if err != nil {
		var apierr *arcadeengine.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerToolsWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Tools(
		context.TODO(),
		"id",
		arcadeengine.WorkerToolsParams{
			Limit:  arcadeengine.Int(0),
			Offset: arcadeengine.Int(0),
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
