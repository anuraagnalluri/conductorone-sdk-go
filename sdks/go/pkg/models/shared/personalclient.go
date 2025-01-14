// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// PersonalClient - The PersonalClient message contains information about a presonal client credential.
type PersonalClient struct {
	// If set, only allows the CIDRs in the array to use the credential.
	AllowSourceCidr []string `json:"allowSourceCidr,omitempty"`
	// The clientID of the credential.
	ClientID  *string    `json:"clientId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The display name of the personal client credential.
	DisplayName *string    `json:"displayName,omitempty"`
	ExpiresTime *time.Time `json:"expiresTime,omitempty"`
	// The unique ID of the personal client credential.
	ID         *string    `json:"id,omitempty"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	// scoped_roles provides a list of IAM Roles
	//  that this OAuth2 Client's API permissions
	//  are reduced to. The permissions granted to OAuth2 Client
	//  are AND'ed against the owning User's own permissions.
	ScopedRoles []string   `json:"scopedRoles,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	// The ID of the user that this credential is created for.
	UserID *string `json:"userId,omitempty"`
}
