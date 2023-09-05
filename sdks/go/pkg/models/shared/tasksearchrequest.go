// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// TaskSearchRequestCurrentStep - Search tasks that have this type of step as the current step.
type TaskSearchRequestCurrentStep string

const (
	TaskSearchRequestCurrentStepTaskSearchCurrentStepUnspecified TaskSearchRequestCurrentStep = "TASK_SEARCH_CURRENT_STEP_UNSPECIFIED"
	TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval    TaskSearchRequestCurrentStep = "TASK_SEARCH_CURRENT_STEP_APPROVAL"
	TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision   TaskSearchRequestCurrentStep = "TASK_SEARCH_CURRENT_STEP_PROVISION"
)

func (e TaskSearchRequestCurrentStep) ToPointer() *TaskSearchRequestCurrentStep {
	return &e
}

func (e *TaskSearchRequestCurrentStep) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_SEARCH_CURRENT_STEP_UNSPECIFIED":
		fallthrough
	case "TASK_SEARCH_CURRENT_STEP_APPROVAL":
		fallthrough
	case "TASK_SEARCH_CURRENT_STEP_PROVISION":
		*e = TaskSearchRequestCurrentStep(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskSearchRequestCurrentStep: %v", v)
	}
}

// TaskSearchRequestEmergencyStatus - Search tasks that are or are not emergency access.
type TaskSearchRequestEmergencyStatus string

const (
	TaskSearchRequestEmergencyStatusUnspecified  TaskSearchRequestEmergencyStatus = "UNSPECIFIED"
	TaskSearchRequestEmergencyStatusAll          TaskSearchRequestEmergencyStatus = "ALL"
	TaskSearchRequestEmergencyStatusNonEmergency TaskSearchRequestEmergencyStatus = "NON_EMERGENCY"
	TaskSearchRequestEmergencyStatusEmergency    TaskSearchRequestEmergencyStatus = "EMERGENCY"
)

func (e TaskSearchRequestEmergencyStatus) ToPointer() *TaskSearchRequestEmergencyStatus {
	return &e
}

func (e *TaskSearchRequestEmergencyStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNSPECIFIED":
		fallthrough
	case "ALL":
		fallthrough
	case "NON_EMERGENCY":
		fallthrough
	case "EMERGENCY":
		*e = TaskSearchRequestEmergencyStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskSearchRequestEmergencyStatus: %v", v)
	}
}

// TaskSearchRequestSortBy - Sort tasks in a specific order.
type TaskSearchRequestSortBy string

const (
	TaskSearchRequestSortByTaskSearchSortByUnspecified     TaskSearchRequestSortBy = "TASK_SEARCH_SORT_BY_UNSPECIFIED"
	TaskSearchRequestSortByTaskSearchSortByAccount         TaskSearchRequestSortBy = "TASK_SEARCH_SORT_BY_ACCOUNT"
	TaskSearchRequestSortByTaskSearchSortByResource        TaskSearchRequestSortBy = "TASK_SEARCH_SORT_BY_RESOURCE"
	TaskSearchRequestSortByTaskSearchSortByAccountOwner    TaskSearchRequestSortBy = "TASK_SEARCH_SORT_BY_ACCOUNT_OWNER"
	TaskSearchRequestSortByTaskSearchSortByReverseTicketID TaskSearchRequestSortBy = "TASK_SEARCH_SORT_BY_REVERSE_TICKET_ID"
)

func (e TaskSearchRequestSortBy) ToPointer() *TaskSearchRequestSortBy {
	return &e
}

func (e *TaskSearchRequestSortBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_SEARCH_SORT_BY_UNSPECIFIED":
		fallthrough
	case "TASK_SEARCH_SORT_BY_ACCOUNT":
		fallthrough
	case "TASK_SEARCH_SORT_BY_RESOURCE":
		fallthrough
	case "TASK_SEARCH_SORT_BY_ACCOUNT_OWNER":
		fallthrough
	case "TASK_SEARCH_SORT_BY_REVERSE_TICKET_ID":
		*e = TaskSearchRequestSortBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskSearchRequestSortBy: %v", v)
	}
}

type TaskSearchRequestTaskStates string

const (
	TaskSearchRequestTaskStatesTaskStateUnspecified TaskSearchRequestTaskStates = "TASK_STATE_UNSPECIFIED"
	TaskSearchRequestTaskStatesTaskStateOpen        TaskSearchRequestTaskStates = "TASK_STATE_OPEN"
	TaskSearchRequestTaskStatesTaskStateClosed      TaskSearchRequestTaskStates = "TASK_STATE_CLOSED"
)

func (e TaskSearchRequestTaskStates) ToPointer() *TaskSearchRequestTaskStates {
	return &e
}

func (e *TaskSearchRequestTaskStates) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_STATE_UNSPECIFIED":
		fallthrough
	case "TASK_STATE_OPEN":
		fallthrough
	case "TASK_STATE_CLOSED":
		*e = TaskSearchRequestTaskStates(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskSearchRequestTaskStates: %v", v)
	}
}

// TaskSearchRequestInput - Search for tasks based on a plethora filters.
type TaskSearchRequestInput struct {
	// Search tasks that belong to any of the access reviews included in this list.
	AccessReviewIds []string `json:"accessReviewIds,omitempty"`
	// Search tasks that have any of these account owners.
	AccountOwnerIds []string `json:"accountOwnerIds,omitempty"`
	// Search tasks that have this actor ID.
	ActorID *string `json:"actorId,omitempty"`
	// Search tasks that have any of these app entitlement IDs.
	AppEntitlementIds []string `json:"appEntitlementIds,omitempty"`
	// Search tasks that have any of these app resource IDs.
	AppResourceIds []string `json:"appResourceIds,omitempty"`
	// Search tasks that have any of these app resource type IDs.
	AppResourceTypeIds []string `json:"appResourceTypeIds,omitempty"`
	// Search tasks that have any of these app users as subjects.
	AppUserSubjectIds []string `json:"appUserSubjectIds,omitempty"`
	// Search tasks that have any of these apps as targets.
	ApplicationIds []string `json:"applicationIds,omitempty"`
	// Search tasks by  List of UserIDs which are currently assigned these Tasks
	AssigneesInIds []string   `json:"assigneesInIds,omitempty"`
	CreatedAfter   *time.Time `json:"createdAfter,omitempty"`
	CreatedBefore  *time.Time `json:"createdBefore,omitempty"`
	// Search tasks that have this type of step as the current step.
	CurrentStep *TaskSearchRequestCurrentStep `json:"currentStep,omitempty"`
	// Search tasks that are or are not emergency access.
	EmergencyStatus *TaskSearchRequestEmergencyStatus `json:"emergencyStatus,omitempty"`
	// Search tasks that do not have any of these app entitlement IDs.
	ExcludeAppEntitlementIds []string `json:"excludeAppEntitlementIds,omitempty"`
	// Exclude Specific TaskIDs from this serach result.
	ExcludeIds []string `json:"excludeIds,omitempty"`
	// The task expand mask is an array of strings that specifes the related objects the requester wishes to have returned when making a request where the expand mask is part of the input. Use '*' to view all possible responses.
	ExpandMask *TaskExpandMask `json:"expandMask,omitempty"`
	// Whether or not to include deleted tasks.
	IncludeDeleted *bool `json:"includeDeleted,omitempty"`
	// Search tasks where the user would see this task in the My Work section
	MyWorkUserIds []string `json:"myWorkUserIds,omitempty"`
	// Search tasks that were created by any of the users in this array.
	OpenerIds []string `json:"openerIds,omitempty"`
	// The pageSize where 0 <= pageSize <= 100. Values < 10 will be set to 10. A value of 0 returns the default page size (currently 25)
	PageSize *float64 `json:"pageSize,omitempty"`
	// The pageToken field.
	PageToken *string `json:"pageToken,omitempty"`
	// Search tasks that were acted on by any of these users.
	PreviouslyActedOnIds []string `json:"previouslyActedOnIds,omitempty"`
	// Fuzzy search tasks by display name or description. Also can search by numeric ID.
	Query *string `json:"query,omitempty"`
	// Query tasks by display name, description, or numeric ID.
	Refs []TaskRef `json:"refs,omitempty"`
	// Sort tasks in a specific order.
	SortBy *TaskSearchRequestSortBy `json:"sortBy,omitempty"`
	// Search tasks where these users are the subject.
	SubjectIds []string `json:"subjectIds,omitempty"`
	// Search tasks with this task state.
	TaskStates []TaskSearchRequestTaskStates `json:"taskStates,omitempty"`
	// Search tasks with this task type. This is a oneOf, and needs an object, which can be empty, to sort.
	TaskTypes []TaskTypeInput `json:"taskTypes,omitempty"`
}
