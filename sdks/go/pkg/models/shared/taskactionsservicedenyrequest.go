// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskActionsServiceDenyRequest - The TaskActionsServiceDenyRequest object lets you deny a task.
type TaskActionsServiceDenyRequest struct {
	// The comment attached to the request.
	Comment *string `json:"comment,omitempty"`
	// The task expand mask is an array of strings that specifes the related objects the requester wishes to have returned when making a request where the expand mask is part of the input. Use '*' to view all possible responses.
	ExpandMask *TaskExpandMask `json:"expandMask,omitempty"`
	// The ID of the currently policy step. This is the step you want to deny.
	PolicyStepID *string `json:"policyStepId,omitempty"`
}