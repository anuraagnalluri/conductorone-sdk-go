// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskServiceCreateGrantRequest -  Create a grant task.
type TaskServiceCreateGrantRequest struct {
	//  The task expand mask is an array of strings that specifes the related objects the requester wishes to have returned when making a request where the expand mask is part of the input. Use '*' to view all possible responses.
	//
	TaskExpandMask *TaskExpandMask `json:"expandMask,omitempty"`
	//  The ID of the app entitlement to grant access to.
	//
	AppEntitlementID string `json:"appEntitlementId"`
	//  The ID of the app that is associated with the entitlement.
	//
	AppID string `json:"appId"`
	//  The ID of the app user to grant access for. This field and identityUserId cannot both be set for a given request.
	//
	AppUserID *string `json:"appUserId,omitempty"`
	//  The description of the request.
	//
	Description *string `json:"description,omitempty"`
	//  Boolean stating whether or not the task is marked as emergency access.
	//
	EmergencyAccess *bool   `json:"emergencyAccess,omitempty"`
	GrantDuration   *string `json:"grantDuration,omitempty"`
	//  The ID of the user associated with the app user we are granting access for. This field cannot be set if appUserID is also set.
	//
	IdentityUserID *string `json:"identityUserId,omitempty"`
}

func (o *TaskServiceCreateGrantRequest) GetTaskExpandMask() *TaskExpandMask {
	if o == nil {
		return nil
	}
	return o.TaskExpandMask
}

func (o *TaskServiceCreateGrantRequest) GetAppEntitlementID() string {
	if o == nil {
		return ""
	}
	return o.AppEntitlementID
}

func (o *TaskServiceCreateGrantRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *TaskServiceCreateGrantRequest) GetAppUserID() *string {
	if o == nil {
		return nil
	}
	return o.AppUserID
}

func (o *TaskServiceCreateGrantRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TaskServiceCreateGrantRequest) GetEmergencyAccess() *bool {
	if o == nil {
		return nil
	}
	return o.EmergencyAccess
}

func (o *TaskServiceCreateGrantRequest) GetGrantDuration() *string {
	if o == nil {
		return nil
	}
	return o.GrantDuration
}

func (o *TaskServiceCreateGrantRequest) GetIdentityUserID() *string {
	if o == nil {
		return nil
	}
	return o.IdentityUserID
}
