// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// TaskTypeRevokeInput - The TaskTypeRevoke message indicates that a task is a revoke task and all related details.
type TaskTypeRevokeInput struct {
	// The TaskRevokeSource message indicates the source of the revoke task is one of expired, nonUsage, request, or review.
	//
	// This message contains a oneof named origin. Only a single field of the following list may be set at a time:
	//   - review
	//   - request
	//   - expired
	//   - nonUsage
	//
	Source *TaskRevokeSource `json:"source,omitempty"`
}

// TaskTypeRevokeOutcome - The outcome of the revoke.
type TaskTypeRevokeOutcome string

const (
	TaskTypeRevokeOutcomeRevokeOutcomeUnspecified TaskTypeRevokeOutcome = "REVOKE_OUTCOME_UNSPECIFIED"
	TaskTypeRevokeOutcomeRevokeOutcomeRevoked     TaskTypeRevokeOutcome = "REVOKE_OUTCOME_REVOKED"
	TaskTypeRevokeOutcomeRevokeOutcomeDenied      TaskTypeRevokeOutcome = "REVOKE_OUTCOME_DENIED"
	TaskTypeRevokeOutcomeRevokeOutcomeError       TaskTypeRevokeOutcome = "REVOKE_OUTCOME_ERROR"
	TaskTypeRevokeOutcomeRevokeOutcomeCancelled   TaskTypeRevokeOutcome = "REVOKE_OUTCOME_CANCELLED"
)

func (e TaskTypeRevokeOutcome) ToPointer() *TaskTypeRevokeOutcome {
	return &e
}

func (e *TaskTypeRevokeOutcome) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "REVOKE_OUTCOME_UNSPECIFIED":
		fallthrough
	case "REVOKE_OUTCOME_REVOKED":
		fallthrough
	case "REVOKE_OUTCOME_DENIED":
		fallthrough
	case "REVOKE_OUTCOME_ERROR":
		fallthrough
	case "REVOKE_OUTCOME_CANCELLED":
		*e = TaskTypeRevokeOutcome(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskTypeRevokeOutcome: %v", v)
	}
}

// TaskTypeRevoke - The TaskTypeRevoke message indicates that a task is a revoke task and all related details.
type TaskTypeRevoke struct {
	// The ID of the app entitlement.
	AppEntitlementID *string `json:"appEntitlementId,omitempty"`
	// The ID of the app.
	AppID *string `json:"appId,omitempty"`
	// The ID of the app user.
	AppUserID *string `json:"appUserId,omitempty"`
	// The ID of the user.
	IdentityUserID *string `json:"identityUserId,omitempty"`
	// The outcome of the revoke.
	Outcome     *TaskTypeRevokeOutcome `json:"outcome,omitempty"`
	OutcomeTime *time.Time             `json:"outcomeTime,omitempty"`
	// The TaskRevokeSource message indicates the source of the revoke task is one of expired, nonUsage, request, or review.
	//
	// This message contains a oneof named origin. Only a single field of the following list may be set at a time:
	//   - review
	//   - request
	//   - expired
	//   - nonUsage
	//
	Source *TaskRevokeSource `json:"source,omitempty"`
}
