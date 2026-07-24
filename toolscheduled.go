// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadego

import (
	"github.com/ArcadeAI/arcade-go/option"
)

// ToolScheduledService contains methods and other services that help with
// interacting with the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolScheduledService] method instead.
type ToolScheduledService struct {
	Options []option.RequestOption
}

// NewToolScheduledService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolScheduledService(opts ...option.RequestOption) (r *ToolScheduledService) {
	r = &ToolScheduledService{}
	r.Options = opts
	return
}
