// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// IntrospectResponse - IntrospectResponse contains information about the current user who is authenticated.
type IntrospectResponse struct {
	// The list of feature flags enabled for the tenant the logged in user belongs to.
	Features []string `json:"features,omitempty"`
	// The list of permissions that the current logged in user has.
	Permissions []string `json:"permissions,omitempty"`
	// The principleID of the current logged in user.
	PrincipleID *string `json:"principleId,omitempty"`
	// The list of roles that the current logged in user has.
	Roles []string `json:"roles,omitempty"`
	// The userID of the current logged in user.
	UserID *string `json:"userId,omitempty"`
}