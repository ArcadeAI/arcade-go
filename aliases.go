// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadeengine

import (
	"github.com/ArcadeAI/arcade-go/internal/apierror"
	"github.com/ArcadeAI/arcade-go/packages/param"
	"github.com/ArcadeAI/arcade-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type AuthorizationContext = shared.AuthorizationContext

// This is an alias to an internal type.
type AuthorizationResponse = shared.AuthorizationResponse

// This is an alias to an internal type.
type AuthorizationResponseStatus = shared.AuthorizationResponseStatus

// Equals "pending"
const AuthorizationResponseStatusPending = shared.AuthorizationResponseStatusPending

// Equals "completed"
const AuthorizationResponseStatusCompleted = shared.AuthorizationResponseStatusCompleted

// Equals "failed"
const AuthorizationResponseStatusFailed = shared.AuthorizationResponseStatusFailed
