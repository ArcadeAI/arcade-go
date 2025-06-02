// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/ArcadeAI/arcade-go/internal/apijson"
)

type AuthorizationContext struct {
	Token    string                   `json:"token"`
	UserInfo map[string]interface{}   `json:"user_info"`
	JSON     authorizationContextJSON `json:"-"`
}

// authorizationContextJSON contains the JSON metadata for the struct
// [AuthorizationContext]
type authorizationContextJSON struct {
	Token       apijson.Field
	UserInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizationContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizationContextJSON) RawJSON() string {
	return r.raw
}

type AuthorizationResponse struct {
	ID         string                      `json:"id"`
	Context    AuthorizationContext        `json:"context"`
	ProviderID string                      `json:"provider_id"`
	Scopes     []string                    `json:"scopes"`
	Status     AuthorizationResponseStatus `json:"status"`
	URL        string                      `json:"url"`
	UserID     string                      `json:"user_id"`
	JSON       authorizationResponseJSON   `json:"-"`
}

// authorizationResponseJSON contains the JSON metadata for the struct
// [AuthorizationResponse]
type authorizationResponseJSON struct {
	ID          apijson.Field
	Context     apijson.Field
	ProviderID  apijson.Field
	Scopes      apijson.Field
	Status      apijson.Field
	URL         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizationResponseJSON) RawJSON() string {
	return r.raw
}

type AuthorizationResponseStatus string

const (
	AuthorizationResponseStatusNotStarted AuthorizationResponseStatus = "not_started"
	AuthorizationResponseStatusPending    AuthorizationResponseStatus = "pending"
	AuthorizationResponseStatusCompleted  AuthorizationResponseStatus = "completed"
	AuthorizationResponseStatusFailed     AuthorizationResponseStatus = "failed"
)

func (r AuthorizationResponseStatus) IsKnown() bool {
	switch r {
	case AuthorizationResponseStatusNotStarted, AuthorizationResponseStatusPending, AuthorizationResponseStatusCompleted, AuthorizationResponseStatusFailed:
		return true
	}
	return false
}
