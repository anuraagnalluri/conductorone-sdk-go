// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetAppUsageControlsResponse - The GetAppUsageControlsResponse message contains the retrieved AppUsageControls object.
type GetAppUsageControlsResponse struct {
	// The AppUsageControls object describes some peripheral configuration for an app.
	AppUsageControls *AppUsageControls `json:"appUsageControls,omitempty"`
	// HasUsageData is false if the access entitlement for this app has no usage data.
	HasUsageData *bool `json:"hasUsageData,omitempty"`
}