// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectorView - The ConnectorView object provides a connector response object, as well as JSONPATHs to related objects provided by expanders.
type ConnectorView struct {
	// JSONPATH expression indicating the location of the App object in the expanded array.
	AppPath *string `json:"appPath,omitempty"`
	// A Connector is used to sync objects into Apps
	Connector *Connector `json:"connector,omitempty"`
	// JSONPATH expression indicating the location of the User object in the expanded array. This is the user that is a direct target of the ticket without a specific relationship to a potentially non-existent app user.
	UsersPath *string `json:"usersPath,omitempty"`
}
