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

func TestWorkerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.New(context.TODO(), arcadego.WorkerNewParams{
		CreateWorkerRequest: arcadego.CreateWorkerRequestParam{
			ID:      arcadego.F("id"),
			Enabled: arcadego.F(true),
			HTTP: arcadego.F(arcadego.CreateWorkerRequestHTTPParam{
				Retry:   arcadego.F(int64(0)),
				Secret:  arcadego.F("secret"),
				Timeout: arcadego.F(int64(1)),
				Uri:     arcadego.F("uri"),
			}),
			Mcp: arcadego.F(arcadego.CreateWorkerRequestMcpParam{
				Retry:   arcadego.F(int64(0)),
				Timeout: arcadego.F(int64(1)),
				Uri:     arcadego.F("uri"),
				Headers: arcadego.F(map[string]string{
					"foo": "string",
				}),
				Oauth2: arcadego.F(arcadego.CreateWorkerRequestMcpOauth2Param{
					AuthorizationURL: arcadego.F("authorization_url"),
					ClientID:         arcadego.F("client_id"),
					ClientSecret:     arcadego.F("client_secret"),
					ExternalID:       arcadego.F("external_id"),
				}),
				Secrets: arcadego.F(map[string]string{
					"foo": "string",
				}),
			}),
			Type: arcadego.F("type"),
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

func TestWorkerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Update(
		context.TODO(),
		"id",
		arcadego.WorkerUpdateParams{
			UpdateWorkerRequest: arcadego.UpdateWorkerRequestParam{
				Enabled: arcadego.F(true),
				HTTP: arcadego.F(arcadego.UpdateWorkerRequestHTTPParam{
					Retry:   arcadego.F(int64(0)),
					Secret:  arcadego.F("secret"),
					Timeout: arcadego.F(int64(1)),
					Uri:     arcadego.F("uri"),
				}),
				Mcp: arcadego.F(arcadego.UpdateWorkerRequestMcpParam{
					Headers: arcadego.F(map[string]string{
						"foo": "string",
					}),
					Oauth2: arcadego.F(arcadego.UpdateWorkerRequestMcpOauth2Param{
						AuthorizationURL: arcadego.F("authorization_url"),
						ClientID:         arcadego.F("client_id"),
						ClientSecret:     arcadego.F("client_secret"),
					}),
					Retry: arcadego.F(int64(0)),
					Secrets: arcadego.F(map[string]string{
						"foo": "string",
					}),
					Timeout: arcadego.F(int64(1)),
					Uri:     arcadego.F("uri"),
				}),
			},
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

func TestWorkerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.List(context.TODO(), arcadego.WorkerListParams{
		Limit:  arcadego.F(int64(0)),
		Offset: arcadego.F(int64(0)),
	})
	if err != nil {
		var apierr *arcadego.Error
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
	client := arcadego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Workers.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *arcadego.Error
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
	client := arcadego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workers.Get(context.TODO(), "id")
	if err != nil {
		var apierr *arcadego.Error
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
	client := arcadego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workers.Health(context.TODO(), "id")
	if err != nil {
		var apierr *arcadego.Error
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
	client := arcadego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workers.Tools(
		context.TODO(),
		"id",
		arcadego.WorkerToolsParams{
			Limit:  arcadego.F(int64(0)),
			Offset: arcadego.F(int64(0)),
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
