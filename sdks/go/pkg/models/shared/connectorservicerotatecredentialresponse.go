// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectorServiceRotateCredentialResponse - ConnectorServiceRotateCredentialResponse is the response returned by the rotate method.
type ConnectorServiceRotateCredentialResponse struct {
	// The new clientSecret returned after rotating the connector credential.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// ConnectorCredential is used by a connector to authenticate with conductor one.
	Credential *ConnectorCredential `json:"credential,omitempty"`
}