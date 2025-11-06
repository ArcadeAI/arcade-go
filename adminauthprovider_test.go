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

func TestAdminAuthProviderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Admin.AuthProviders.New(context.TODO(), arcadego.AdminAuthProviderNewParams{
		AuthProviderCreateRequest: arcadego.AuthProviderCreateRequestParam{
			ID:          arcadego.F("id"),
			Description: arcadego.F("description"),
			ExternalID:  arcadego.F("external_id"),
			Oauth2: arcadego.F(arcadego.AuthProviderCreateRequestOauth2Param{
				ClientID: arcadego.F("client_id"),
				AuthorizeRequest: arcadego.F(arcadego.AuthProviderCreateRequestOauth2AuthorizeRequestParam{
					Endpoint:              arcadego.F("endpoint"),
					AuthHeaderValueFormat: arcadego.F("auth_header_value_format"),
					AuthMethod:            arcadego.F("auth_method"),
					Method:                arcadego.F("method"),
					Params: arcadego.F(map[string]string{
						"foo": "string",
					}),
					RequestContentType:  arcadego.F(arcadego.AuthProviderCreateRequestOauth2AuthorizeRequestRequestContentTypeApplicationXWwwFormUrlencoded),
					ResponseContentType: arcadego.F(arcadego.AuthProviderCreateRequestOauth2AuthorizeRequestResponseContentTypeApplicationXWwwFormUrlencoded),
					ResponseMap: arcadego.F(map[string]string{
						"foo": "string",
					}),
				}),
				ClientSecret: arcadego.F("client_secret"),
				Pkce: arcadego.F(arcadego.AuthProviderCreateRequestOauth2PkceParam{
					CodeChallengeMethod: arcadego.F("code_challenge_method"),
					Enabled:             arcadego.F(true),
				}),
				RefreshRequest: arcadego.F(arcadego.AuthProviderCreateRequestOauth2RefreshRequestParam{
					Endpoint:              arcadego.F("endpoint"),
					AuthHeaderValueFormat: arcadego.F("auth_header_value_format"),
					AuthMethod:            arcadego.F("auth_method"),
					Method:                arcadego.F("method"),
					Params: arcadego.F(map[string]string{
						"foo": "string",
					}),
					RequestContentType:  arcadego.F(arcadego.AuthProviderCreateRequestOauth2RefreshRequestRequestContentTypeApplicationXWwwFormUrlencoded),
					ResponseContentType: arcadego.F(arcadego.AuthProviderCreateRequestOauth2RefreshRequestResponseContentTypeApplicationXWwwFormUrlencoded),
					ResponseMap: arcadego.F(map[string]string{
						"foo": "string",
					}),
				}),
				ScopeDelimiter: arcadego.F(arcadego.AuthProviderCreateRequestOauth2ScopeDelimiterUnknown0),
				TokenIntrospectionRequest: arcadego.F(arcadego.AuthProviderCreateRequestOauth2TokenIntrospectionRequestParam{
					Endpoint: arcadego.F("endpoint"),
					Triggers: arcadego.F(arcadego.AuthProviderCreateRequestOauth2TokenIntrospectionRequestTriggersParam{
						OnTokenGrant:   arcadego.F(true),
						OnTokenRefresh: arcadego.F(true),
					}),
					AuthHeaderValueFormat: arcadego.F("auth_header_value_format"),
					AuthMethod:            arcadego.F("auth_method"),
					Method:                arcadego.F("method"),
					Params: arcadego.F(map[string]string{
						"foo": "string",
					}),
					RequestContentType:  arcadego.F(arcadego.AuthProviderCreateRequestOauth2TokenIntrospectionRequestRequestContentTypeApplicationXWwwFormUrlencoded),
					ResponseContentType: arcadego.F(arcadego.AuthProviderCreateRequestOauth2TokenIntrospectionRequestResponseContentTypeApplicationXWwwFormUrlencoded),
					ResponseMap: arcadego.F(map[string]string{
						"foo": "string",
					}),
				}),
				TokenRequest: arcadego.F(arcadego.AuthProviderCreateRequestOauth2TokenRequestParam{
					Endpoint:              arcadego.F("endpoint"),
					AuthHeaderValueFormat: arcadego.F("auth_header_value_format"),
					AuthMethod:            arcadego.F("auth_method"),
					Method:                arcadego.F("method"),
					Params: arcadego.F(map[string]string{
						"foo": "string",
					}),
					RequestContentType:  arcadego.F(arcadego.AuthProviderCreateRequestOauth2TokenRequestRequestContentTypeApplicationXWwwFormUrlencoded),
					ResponseContentType: arcadego.F(arcadego.AuthProviderCreateRequestOauth2TokenRequestResponseContentTypeApplicationXWwwFormUrlencoded),
					ResponseMap: arcadego.F(map[string]string{
						"foo": "string",
					}),
				}),
				UserInfoRequest: arcadego.F(arcadego.AuthProviderCreateRequestOauth2UserInfoRequestParam{
					Endpoint: arcadego.F("endpoint"),
					Triggers: arcadego.F(arcadego.AuthProviderCreateRequestOauth2UserInfoRequestTriggersParam{
						OnTokenGrant:   arcadego.F(true),
						OnTokenRefresh: arcadego.F(true),
					}),
					AuthHeaderValueFormat: arcadego.F("auth_header_value_format"),
					AuthMethod:            arcadego.F("auth_method"),
					Method:                arcadego.F("method"),
					Params: arcadego.F(map[string]string{
						"foo": "string",
					}),
					RequestContentType:  arcadego.F(arcadego.AuthProviderCreateRequestOauth2UserInfoRequestRequestContentTypeApplicationXWwwFormUrlencoded),
					ResponseContentType: arcadego.F(arcadego.AuthProviderCreateRequestOauth2UserInfoRequestResponseContentTypeApplicationXWwwFormUrlencoded),
					ResponseMap: arcadego.F(map[string]string{
						"foo": "string",
					}),
				}),
			}),
			ProviderID: arcadego.F("provider_id"),
			Status:     arcadego.F("status"),
			Type:       arcadego.F("type"),
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

func TestAdminAuthProviderList(t *testing.T) {
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
	_, err := client.Admin.AuthProviders.List(context.TODO())
	if err != nil {
		var apierr *arcadego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdminAuthProviderDelete(t *testing.T) {
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
	_, err := client.Admin.AuthProviders.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *arcadego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdminAuthProviderGet(t *testing.T) {
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
	_, err := client.Admin.AuthProviders.Get(context.TODO(), "id")
	if err != nil {
		var apierr *arcadego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdminAuthProviderPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Admin.AuthProviders.Patch(
		context.TODO(),
		"id",
		arcadego.AdminAuthProviderPatchParams{
			AuthProviderUpdateRequest: arcadego.AuthProviderUpdateRequestParam{
				ID:          arcadego.F("id"),
				Description: arcadego.F("description"),
				Oauth2: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2Param{
					AuthorizeRequest: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2AuthorizeRequestParam{
						AuthHeaderValueFormat: arcadego.F("auth_header_value_format"),
						AuthMethod:            arcadego.F("auth_method"),
						Endpoint:              arcadego.F("endpoint"),
						Method:                arcadego.F("method"),
						Params: arcadego.F(map[string]string{
							"foo": "string",
						}),
						RequestContentType:  arcadego.F(arcadego.AuthProviderUpdateRequestOauth2AuthorizeRequestRequestContentTypeApplicationXWwwFormUrlencoded),
						ResponseContentType: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2AuthorizeRequestResponseContentTypeApplicationXWwwFormUrlencoded),
						ResponseMap: arcadego.F(map[string]string{
							"foo": "string",
						}),
					}),
					ClientID:     arcadego.F("client_id"),
					ClientSecret: arcadego.F("client_secret"),
					Pkce: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2PkceParam{
						CodeChallengeMethod: arcadego.F("code_challenge_method"),
						Enabled:             arcadego.F(true),
					}),
					RefreshRequest: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2RefreshRequestParam{
						AuthHeaderValueFormat: arcadego.F("auth_header_value_format"),
						AuthMethod:            arcadego.F("auth_method"),
						Endpoint:              arcadego.F("endpoint"),
						Method:                arcadego.F("method"),
						Params: arcadego.F(map[string]string{
							"foo": "string",
						}),
						RequestContentType:  arcadego.F(arcadego.AuthProviderUpdateRequestOauth2RefreshRequestRequestContentTypeApplicationXWwwFormUrlencoded),
						ResponseContentType: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2RefreshRequestResponseContentTypeApplicationXWwwFormUrlencoded),
						ResponseMap: arcadego.F(map[string]string{
							"foo": "string",
						}),
					}),
					ScopeDelimiter: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2ScopeDelimiterUnknown1),
					TokenRequest: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2TokenRequestParam{
						AuthHeaderValueFormat: arcadego.F("auth_header_value_format"),
						AuthMethod:            arcadego.F("auth_method"),
						Endpoint:              arcadego.F("endpoint"),
						Method:                arcadego.F("method"),
						Params: arcadego.F(map[string]string{
							"foo": "string",
						}),
						RequestContentType:  arcadego.F(arcadego.AuthProviderUpdateRequestOauth2TokenRequestRequestContentTypeApplicationXWwwFormUrlencoded),
						ResponseContentType: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2TokenRequestResponseContentTypeApplicationXWwwFormUrlencoded),
						ResponseMap: arcadego.F(map[string]string{
							"foo": "string",
						}),
					}),
					UserInfoRequest: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2UserInfoRequestParam{
						AuthHeaderValueFormat: arcadego.F("auth_header_value_format"),
						AuthMethod:            arcadego.F("auth_method"),
						Endpoint:              arcadego.F("endpoint"),
						Method:                arcadego.F("method"),
						Params: arcadego.F(map[string]string{
							"foo": "string",
						}),
						RequestContentType:  arcadego.F(arcadego.AuthProviderUpdateRequestOauth2UserInfoRequestRequestContentTypeApplicationXWwwFormUrlencoded),
						ResponseContentType: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2UserInfoRequestResponseContentTypeApplicationXWwwFormUrlencoded),
						ResponseMap: arcadego.F(map[string]string{
							"foo": "string",
						}),
						Triggers: arcadego.F(arcadego.AuthProviderUpdateRequestOauth2UserInfoRequestTriggersParam{
							OnTokenGrant:   arcadego.F(true),
							OnTokenRefresh: arcadego.F(true),
						}),
					}),
				}),
				ProviderID: arcadego.F("provider_id"),
				Status:     arcadego.F("status"),
				Type:       arcadego.F("type"),
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
