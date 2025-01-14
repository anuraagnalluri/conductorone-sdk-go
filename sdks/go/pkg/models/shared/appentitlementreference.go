// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppEntitlementReference - This object references an app entitlement's ID and AppID.
type AppEntitlementReference struct {
	// The ID of the Entitlement.
	AppEntitlementID *string `json:"appEntitlementId,omitempty"`
	// The ID of the App this entitlement belongs to.
	AppID *string `json:"appId,omitempty"`
}
