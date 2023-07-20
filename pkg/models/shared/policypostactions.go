// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PolicyPostActions - The PolicyPostActions message.
//
// This message contains a oneof named action. Only a single field of the following list may be set at a time:
//   - certifyRemediateImmediately
type PolicyPostActions struct {
	//  ONLY valid when used in a CERTIFY Ticket Type:
	//  Causes any deprovision or change in a grant to be applied when Certify Ticket is closed.
	//
	// This field is part of the `action` oneof.
	// See the documentation for `c1.api.policy.v1.PolicyPostActions` for more details.
	CertifyRemediateImmediately *bool `json:"certifyRemediateImmediately,omitempty"`
}

func (o *PolicyPostActions) GetCertifyRemediateImmediately() *bool {
	if o == nil {
		return nil
	}
	return o.CertifyRemediateImmediately
}
