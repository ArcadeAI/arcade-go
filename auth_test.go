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

func TestAuthAuthorize(t *testing.T) {
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
	_, err := client.Auth.Authorize(context.TODO(), arcadeengine.AuthAuthorizeParams{
		AuthRequest: arcadeengine.AuthRequestParam{
			AuthRequirement: arcadeengine.AuthRequestAuthRequirementParam{
				ID: arcadeengine.String("id"),
				Oauth2: arcadeengine.AuthRequestAuthRequirementOauth2Param{
					Scopes: []string{"string"},
				},
				ProviderID:   arcadeengine.String("provider_id"),
				ProviderType: arcadeengine.String("provider_type"),
			},
			UserID: "user_id",
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

func TestAuthStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Auth.Status(context.TODO(), arcadeengine.AuthStatusParams{
		ID:   "id",
		Wait: arcadeengine.Int(0),
	})
	if err != nil {
		var apierr *arcadeengine.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
