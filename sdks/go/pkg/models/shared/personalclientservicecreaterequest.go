// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PersonalClientServiceCreateRequest - The PersonalClientServiceCreateRequest message contains the fields for creating a new personal client.
type PersonalClientServiceCreateRequest struct {
	// A list of CIDRs to restrict this credential to.
	AllowSourceCidr []string `json:"allowSourceCidr,omitempty"`
	// The display name for the new personal client.
	DisplayName *string `json:"displayName,omitempty"`
	Expires     *string `json:"expires,omitempty"`
	// The list of roles to restrict the credential to.
	ScopedRoles []string `json:"scopedRoles,omitempty"`
}
