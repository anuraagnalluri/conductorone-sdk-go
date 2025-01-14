// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestCatalogExpandMask - The RequestCatalogExpandMask includes the paths in the catalog view to expand in the return value of this call.
type RequestCatalogExpandMask struct {
	// An array of paths to be expanded in the response. May be any combination of "*", "created_by_user_id", "app_ids", and "access_entitlements".
	Paths []string `json:"paths,omitempty"`
}
