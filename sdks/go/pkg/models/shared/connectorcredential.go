// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// ConnectorCredential - ConnectorCredential is used by a connector to authenticate with conductor one.
type ConnectorCredential struct {
	// The appId of the app the connector is attached to.
	AppID *string `json:"appId,omitempty"`
	// The client id of the ConnectorCredential.
	ClientID *string `json:"clientId,omitempty"`
	// The connectorId of the connector the credential is associated with.
	ConnectorID *string    `json:"connectorId,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
	// The display name of the ConnectorCredential.
	DisplayName *string    `json:"displayName,omitempty"`
	ExpiresTime *time.Time `json:"expiresTime,omitempty"`
	// The id of the ConnectorCredential.
	ID         *string    `json:"id,omitempty"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
}
