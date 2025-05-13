// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadeengine

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/arcade-engine-go/internal/apijson"
	"github.com/stainless-sdks/arcade-engine-go/internal/requestconfig"
	"github.com/stainless-sdks/arcade-engine-go/option"
	"github.com/stainless-sdks/arcade-engine-go/packages/respjson"
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
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
	r = HealthService{}
	r.Options = opts
	return
}

// Check if Arcade Engine is healthy
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (res *HealthSchema, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HealthSchema struct {
	Healthy bool `json:"healthy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Healthy     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthSchema) RawJSON() string { return r.JSON.raw }
func (r *HealthSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
