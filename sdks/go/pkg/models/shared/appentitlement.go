// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type AppEntitlementDurationUnset struct {
}

// AppEntitlement - The app entitlement represents one permission in a downstream App (SAAS) that can be granted. For example, GitHub Read vs GitHub Write.
//
// This message contains a oneof named max_grant_duration. Only a single field of the following list may be set at a time:
//   - durationUnset
//   - durationGrant
type AppEntitlement struct {
	// The alias of the app entitlement used by Cone. Also exact-match queryable.
	Alias *string `json:"alias,omitempty"`
	// The ID of the app that is associated with the app entitlement.
	AppID *string `json:"appId,omitempty"`
	// The ID of the app resource that is associated with the app entitlement
	AppResourceID *string `json:"appResourceId,omitempty"`
	// The ID of the app resource type that is associated with the app entitlement
	AppResourceTypeID *string `json:"appResourceTypeId,omitempty"`
	// The ID of the policy that will be used for certify tickets related to the app entitlement.
	CertifyPolicyID *string `json:"certifyPolicyId,omitempty"`
	// The IDs of different compliance frameworks associated with this app entitlement ex (SOX, HIPAA, PCI, etc.)
	ComplianceFrameworkValueIds []string   `json:"complianceFrameworkValueIds,omitempty"`
	CreatedAt                   *time.Time `json:"createdAt,omitempty"`
	DeletedAt                   *time.Time `json:"deletedAt,omitempty"`
	// The description of the app entitlement.
	Description *string `json:"description,omitempty"`
	// The display name of the app entitlement.
	DisplayName   *string                      `json:"displayName,omitempty"`
	DurationGrant *string                      `json:"durationGrant,omitempty"`
	DurationUnset *AppEntitlementDurationUnset `json:"durationUnset,omitempty"`
	// This enables tasks to be created in an emergency and use a selected emergency access policy.
	EmergencyGrantEnabled *bool `json:"emergencyGrantEnabled,omitempty"`
	// The ID of the policy that will be used for emergency access grant tasks.
	EmergencyGrantPolicyID *string `json:"emergencyGrantPolicyId,omitempty"`
	// The amount of grants open for this entitlement
	GrantCount *string `json:"grantCount,omitempty"`
	// The ID of the policy that will be used for grant tickets related to the app entitlement.
	GrantPolicyID *string `json:"grantPolicyId,omitempty"`
	// The unique ID for the App Entitlement.
	ID *string `json:"id,omitempty"`
	// ProvisionPolicy is a oneOf that indicates how a provision step should be processed.
	//
	// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
	//   - connector
	//   - manual
	//   - delegated
	//
	ProvisionerPolicy *ProvisionPolicy `json:"provisionerPolicy,omitempty"`
	// The ID of the policy that will be used for revoke tickets related to the app entitlement
	RevokePolicyID *string `json:"revokePolicyId,omitempty"`
	// The riskLevelValueId field.
	RiskLevelValueID *string `json:"riskLevelValueId,omitempty"`
	// The slug is displayed as an oval next to the name in the frontend of C1, it tells you what permission the entitlement grants. See https://www.conductorone.com/docs/product/manage-access/entitlements/
	Slug *string `json:"slug,omitempty"`
	// This field indicates if this is a system builtin entitlement.
	SystemBuiltin *bool      `json:"systemBuiltin,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
}

// AppEntitlementInput - The app entitlement represents one permission in a downstream App (SAAS) that can be granted. For example, GitHub Read vs GitHub Write.
//
// This message contains a oneof named max_grant_duration. Only a single field of the following list may be set at a time:
//   - durationUnset
//   - durationGrant
type AppEntitlementInput struct {
	// The ID of the app that is associated with the app entitlement.
	AppID *string `json:"appId,omitempty"`
	// The ID of the app resource that is associated with the app entitlement
	AppResourceID *string `json:"appResourceId,omitempty"`
	// The ID of the app resource type that is associated with the app entitlement
	AppResourceTypeID *string `json:"appResourceTypeId,omitempty"`
	// The ID of the policy that will be used for certify tickets related to the app entitlement.
	CertifyPolicyID *string `json:"certifyPolicyId,omitempty"`
	// The IDs of different compliance frameworks associated with this app entitlement ex (SOX, HIPAA, PCI, etc.)
	ComplianceFrameworkValueIds []string `json:"complianceFrameworkValueIds,omitempty"`
	// The description of the app entitlement.
	Description *string `json:"description,omitempty"`
	// The display name of the app entitlement.
	DisplayName   *string                      `json:"displayName,omitempty"`
	DurationGrant *string                      `json:"durationGrant,omitempty"`
	DurationUnset *AppEntitlementDurationUnset `json:"durationUnset,omitempty"`
	// This enables tasks to be created in an emergency and use a selected emergency access policy.
	EmergencyGrantEnabled *bool `json:"emergencyGrantEnabled,omitempty"`
	// The ID of the policy that will be used for emergency access grant tasks.
	EmergencyGrantPolicyID *string `json:"emergencyGrantPolicyId,omitempty"`
	// The ID of the policy that will be used for grant tickets related to the app entitlement.
	GrantPolicyID *string `json:"grantPolicyId,omitempty"`
	// ProvisionPolicy is a oneOf that indicates how a provision step should be processed.
	//
	// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
	//   - connector
	//   - manual
	//   - delegated
	//
	ProvisionerPolicy *ProvisionPolicy `json:"provisionerPolicy,omitempty"`
	// The ID of the policy that will be used for revoke tickets related to the app entitlement
	RevokePolicyID *string `json:"revokePolicyId,omitempty"`
	// The riskLevelValueId field.
	RiskLevelValueID *string `json:"riskLevelValueId,omitempty"`
	// The slug is displayed as an oval next to the name in the frontend of C1, it tells you what permission the entitlement grants. See https://www.conductorone.com/docs/product/manage-access/entitlements/
	Slug *string `json:"slug,omitempty"`
}
