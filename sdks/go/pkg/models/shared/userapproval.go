// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UserApproval - The user approval object describes the approval configuration of a policy step that needs to be approved by a specific list of users.
type UserApproval struct {
	// Configuration to allow self approval of if the user is specified and also the target of the ticket.
	AllowSelfApproval *bool `json:"allowSelfApproval,omitempty"`
	// Array of users configured for approval.
	UserIds []string `json:"userIds,omitempty"`
}
