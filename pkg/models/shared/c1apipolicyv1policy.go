// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// C1APIPolicyV1PolicyPolicyType - The policyType field.
type C1APIPolicyV1PolicyPolicyType string

const (
	C1APIPolicyV1PolicyPolicyTypePolicyTypeUnspecified   C1APIPolicyV1PolicyPolicyType = "POLICY_TYPE_UNSPECIFIED"
	C1APIPolicyV1PolicyPolicyTypePolicyTypeGrant         C1APIPolicyV1PolicyPolicyType = "POLICY_TYPE_GRANT"
	C1APIPolicyV1PolicyPolicyTypePolicyTypeRevoke        C1APIPolicyV1PolicyPolicyType = "POLICY_TYPE_REVOKE"
	C1APIPolicyV1PolicyPolicyTypePolicyTypeCertify       C1APIPolicyV1PolicyPolicyType = "POLICY_TYPE_CERTIFY"
	C1APIPolicyV1PolicyPolicyTypePolicyTypeAccessRequest C1APIPolicyV1PolicyPolicyType = "POLICY_TYPE_ACCESS_REQUEST"
	C1APIPolicyV1PolicyPolicyTypePolicyTypeProvision     C1APIPolicyV1PolicyPolicyType = "POLICY_TYPE_PROVISION"
)

func (e C1APIPolicyV1PolicyPolicyType) ToPointer() *C1APIPolicyV1PolicyPolicyType {
	return &e
}

func (e *C1APIPolicyV1PolicyPolicyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POLICY_TYPE_UNSPECIFIED":
		fallthrough
	case "POLICY_TYPE_GRANT":
		fallthrough
	case "POLICY_TYPE_REVOKE":
		fallthrough
	case "POLICY_TYPE_CERTIFY":
		fallthrough
	case "POLICY_TYPE_ACCESS_REQUEST":
		fallthrough
	case "POLICY_TYPE_PROVISION":
		*e = C1APIPolicyV1PolicyPolicyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for C1APIPolicyV1PolicyPolicyType: %v", v)
	}
}

// C1APIPolicyV1Policy - The Policy message.
type C1APIPolicyV1Policy struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	//  Properties
	//
	ID *string `json:"id,omitempty"`
	// The policySteps field.
	PolicySteps map[string]C1APIPolicyV1PolicySteps `json:"policySteps,omitempty"`
	// The policyType field.
	PolicyType *C1APIPolicyV1PolicyPolicyType `json:"policyType,omitempty"`
	// The postActions field.
	PostActions []C1APIPolicyV1PolicyPostActions `json:"postActions,omitempty"`
	// The reassignTasksToDelegates field.
	ReassignTasksToDelegates *bool `json:"reassignTasksToDelegates,omitempty"`
	// The systemBuiltin field.
	SystemBuiltin *bool      `json:"systemBuiltin,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
}