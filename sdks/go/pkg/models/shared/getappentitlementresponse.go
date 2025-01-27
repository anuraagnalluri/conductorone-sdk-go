// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetAppEntitlementResponse - The get app entitlement response returns an entitlement view containing paths in the expanded array for the objects expanded as indicated by the expand mask in the request.
type GetAppEntitlementResponse struct {
	// The app entitlement view contains the serialized app entitlement and paths to objects referenced by the app entitlement.
	AppEntitlementView *AppEntitlementView `json:"appEntitlementView,omitempty"`
	// List of serialized related objects.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
}
