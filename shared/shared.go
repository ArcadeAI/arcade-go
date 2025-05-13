// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/arcade-engine-go/internal/apijson"
	"github.com/stainless-sdks/arcade-engine-go/packages/param"
	"github.com/stainless-sdks/arcade-engine-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type AuthorizationContext struct {
	Token    string         `json:"token"`
	UserInfo map[string]any `json:"user_info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		UserInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthorizationContext) RawJSON() string { return r.JSON.raw }
func (r *AuthorizationContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthorizationResponse struct {
	ID         string               `json:"id"`
	Context    AuthorizationContext `json:"context"`
	ProviderID string               `json:"provider_id"`
	Scopes     []string             `json:"scopes"`
	// Any of "pending", "completed", "failed".
	Status AuthorizationResponseStatus `json:"status"`
	URL    string                      `json:"url"`
	UserID string                      `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Context     respjson.Field
		ProviderID  respjson.Field
		Scopes      respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthorizationResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthorizationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthorizationResponseStatus string

const (
	AuthorizationResponseStatusPending   AuthorizationResponseStatus = "pending"
	AuthorizationResponseStatusCompleted AuthorizationResponseStatus = "completed"
	AuthorizationResponseStatusFailed    AuthorizationResponseStatus = "failed"
)
