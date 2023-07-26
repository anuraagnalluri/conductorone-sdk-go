// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdatePolicyResponse -  The UpdatePolicyResponse message contains the updated policy object.
type UpdatePolicyResponse struct {
	//  A policy describes the behavior of the ConductorOne system when processing a task. You can describe the type, approvers, fallback behavior, and escalation processes.
	//
	Policy *Policy `json:"policy,omitempty"`
}

func (o *UpdatePolicyResponse) GetPolicy() *Policy {
	if o == nil {
		return nil
	}
	return o.Policy
}
