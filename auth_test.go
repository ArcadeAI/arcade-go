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

func TestAuthAuthorizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Auth.Authorize(context.TODO(), arcadego.AuthAuthorizeParams{
		AuthRequest: arcadego.AuthRequestParam{
			AuthRequirement: arcadego.F(arcadego.AuthRequestAuthRequirementParam{
				ID: arcadego.F("id"),
				Oauth2: arcadego.F(arcadego.AuthRequestAuthRequirementOauth2Param{
					Scopes: arcadego.F([]string{"string"}),
				}),
				ProviderID:   arcadego.F("provider_id"),
				ProviderType: arcadego.F("provider_type"),
			}),
			UserID:            arcadego.F("user_id"),
			ForceVerification: arcadego.F(true),
			NextUri:           arcadego.F("next_uri"),
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

func TestAuthStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Auth.Status(context.TODO(), arcadego.AuthStatusParams{
		ID:   arcadego.F("id"),
		Wait: arcadego.F(int64(0)),
	})
	if err != nil {
		var apierr *arcadego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
