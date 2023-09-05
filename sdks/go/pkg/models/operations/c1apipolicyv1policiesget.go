// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIPolicyV1PoliciesGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIPolicyV1PoliciesGetResponse struct {
	ContentType string
	// The GetPolicyResponse message contains the policy object.
	GetPolicyResponse *shared.GetPolicyResponse
	StatusCode        int
	RawResponse       *http.Response
}
