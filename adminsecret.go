// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadego

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ArcadeAI/arcade-go/internal/apijson"
	"github.com/ArcadeAI/arcade-go/internal/requestconfig"
	"github.com/ArcadeAI/arcade-go/option"
)

// AdminSecretService contains methods and other services that help with
// interacting with the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminSecretService] method instead.
type AdminSecretService struct {
	Options []option.RequestOption
}

// NewAdminSecretService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminSecretService(opts ...option.RequestOption) (r *AdminSecretService) {
	r = &AdminSecretService{}
	r.Options = opts
	return
}

// List all secrets that are visible to the caller
func (r *AdminSecretService) List(ctx context.Context, opts ...option.RequestOption) (res *AdminSecretListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/admin/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a secret by its ID
func (r *AdminSecretService) Delete(ctx context.Context, secretID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return
	}
	path := fmt.Sprintf("v1/admin/secrets/%s", secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type SecretResponse struct {
	ID             string                `json:"id"`
	Binding        SecretResponseBinding `json:"binding"`
	CreatedAt      string                `json:"created_at"`
	Description    string                `json:"description"`
	Hint           string                `json:"hint"`
	Key            string                `json:"key"`
	LastAccessedAt string                `json:"last_accessed_at"`
	UpdatedAt      string                `json:"updated_at"`
	JSON           secretResponseJSON    `json:"-"`
}

// secretResponseJSON contains the JSON metadata for the struct [SecretResponse]
type secretResponseJSON struct {
	ID             apijson.Field
	Binding        apijson.Field
	CreatedAt      apijson.Field
	Description    apijson.Field
	Hint           apijson.Field
	Key            apijson.Field
	LastAccessedAt apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretResponseJSON) RawJSON() string {
	return r.raw
}

type SecretResponseBinding struct {
	ID   string                    `json:"id"`
	Type SecretResponseBindingType `json:"type"`
	JSON secretResponseBindingJSON `json:"-"`
}

// secretResponseBindingJSON contains the JSON metadata for the struct
// [SecretResponseBinding]
type secretResponseBindingJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretResponseBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretResponseBindingJSON) RawJSON() string {
	return r.raw
}

type SecretResponseBindingType string

const (
	SecretResponseBindingTypeStatic  SecretResponseBindingType = "static"
	SecretResponseBindingTypeTenant  SecretResponseBindingType = "tenant"
	SecretResponseBindingTypeProject SecretResponseBindingType = "project"
	SecretResponseBindingTypeAccount SecretResponseBindingType = "account"
)

func (r SecretResponseBindingType) IsKnown() bool {
	switch r {
	case SecretResponseBindingTypeStatic, SecretResponseBindingTypeTenant, SecretResponseBindingTypeProject, SecretResponseBindingTypeAccount:
		return true
	}
	return false
}

type AdminSecretListResponse struct {
	Items      []SecretResponse            `json:"items"`
	Limit      int64                       `json:"limit"`
	Offset     int64                       `json:"offset"`
	PageCount  int64                       `json:"page_count"`
	TotalCount int64                       `json:"total_count"`
	JSON       adminSecretListResponseJSON `json:"-"`
}

// adminSecretListResponseJSON contains the JSON metadata for the struct
// [AdminSecretListResponse]
type adminSecretListResponseJSON struct {
	Items       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	PageCount   apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdminSecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminSecretListResponseJSON) RawJSON() string {
	return r.raw
}
