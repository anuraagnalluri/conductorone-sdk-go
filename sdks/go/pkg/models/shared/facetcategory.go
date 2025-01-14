// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// FacetCategory - The FacetCategory indicates a grouping of facets by type. For example, facets "OnePassword" and "Okta" would group under an "Apps" category.
//
// This message contains a oneof named item. Only a single field of the following list may be set at a time:
//   - value
//   - range
type FacetCategory struct {
	// The display name of the category.
	DisplayName *string `json:"displayName,omitempty"`
	// An icon for the category.
	IconURL *string `json:"iconUrl,omitempty"`
	// The param that is being set when checking a facet in this category.
	Param *string `json:"param,omitempty"`
	// The FacetRangeItem message.
	Range *FacetRangeItem `json:"range,omitempty"`
	// The FacetValueItem message.
	Value *FacetValueItem `json:"value,omitempty"`
}
