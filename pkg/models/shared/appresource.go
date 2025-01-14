// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// AppResource - The app resource message is a single resource that can have entitlements.
type AppResource struct {
	// The app that this resource belongs to.
	AppID *string `json:"appId,omitempty"`
	// The resource type that this resource is.
	AppResourceTypeID *string    `json:"appResourceTypeId,omitempty"`
	CreatedAt         *time.Time `json:"createdAt,omitempty"`
	// A custom description that can be set for a resource.
	CustomDescription *string    `json:"customDescription,omitempty"`
	DeletedAt         *time.Time `json:"deletedAt,omitempty"`
	// The description set for the resource.
	Description *string `json:"description,omitempty"`
	// The display name for this resource.
	DisplayName *string `json:"displayName,omitempty"`
	// The number of grants to this resource.
	GrantCount *string `json:"grantCount,omitempty"`
	// The id of the resource.
	ID        *string    `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

func (o *AppResource) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppResource) GetAppResourceTypeID() *string {
	if o == nil {
		return nil
	}
	return o.AppResourceTypeID
}

func (o *AppResource) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AppResource) GetCustomDescription() *string {
	if o == nil {
		return nil
	}
	return o.CustomDescription
}

func (o *AppResource) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *AppResource) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AppResource) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppResource) GetGrantCount() *string {
	if o == nil {
		return nil
	}
	return o.GrantCount
}

func (o *AppResource) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AppResource) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
