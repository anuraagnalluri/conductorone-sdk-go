// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectorExpandMask - The ConnectorExpandMask is used to expand related objects on a connector.
type ConnectorExpandMask struct {
	// Paths that you want expanded in the response. Possible values are "app_id" and "*".
	Paths []string `json:"paths,omitempty"`
}
