// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadego

import (
	"context"
	"net/http"
	"slices"

	"github.com/ArcadeAI/arcade-go/internal/apijson"
	"github.com/ArcadeAI/arcade-go/internal/requestconfig"
	"github.com/ArcadeAI/arcade-go/option"
)

// HealthService contains methods and other services that help with interacting
// with the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	Options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r *HealthService) {
	r = &HealthService{}
	r.Options = opts
	return
}

// Check if Arcade Engine is healthy
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (res *HealthSchema, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HealthSchema struct {
	Healthy bool             `json:"healthy"`
	JSON    healthSchemaJSON `json:"-"`
}

// healthSchemaJSON contains the JSON metadata for the struct [HealthSchema]
type healthSchemaJSON struct {
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthSchemaJSON) RawJSON() string {
	return r.raw
}
