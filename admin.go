// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadego

import (
	"github.com/ArcadeAI/arcade-go/option"
)

// AdminService contains methods and other services that help with interacting with
// the Arcade API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminService] method instead.
type AdminService struct {
	Options         []option.RequestOption
	UserConnections *AdminUserConnectionService
	AuthProviders   *AdminAuthProviderService
	Secrets         *AdminSecretService
}

// NewAdminService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAdminService(opts ...option.RequestOption) (r *AdminService) {
	r = &AdminService{}
	r.Options = opts
	r.UserConnections = NewAdminUserConnectionService(opts...)
	r.AuthProviders = NewAdminAuthProviderService(opts...)
	r.Secrets = NewAdminSecretService(opts...)
	return
}
