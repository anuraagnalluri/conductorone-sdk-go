// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestCatalogManagementServiceAddAccessEntitlementsRequest - The RequestCatalogManagementServiceAddAccessEntitlementsRequest message is used to add access entitlements to a request
//
//	catalog to determine which users can view the request catalog.
type RequestCatalogManagementServiceAddAccessEntitlementsRequest struct {
	// List of entitlements to add to the request catalog as access entitlements.
	AccessEntitlements []AppEntitlementRef `json:"accessEntitlements,omitempty"`
}
