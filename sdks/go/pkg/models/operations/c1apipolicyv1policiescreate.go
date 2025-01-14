// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIPolicyV1PoliciesCreateResponse struct {
	ContentType string
	// The CreatePolicyResponse message contains the created policy object.
	CreatePolicyResponse *shared.CreatePolicyResponse
	StatusCode           int
	RawResponse          *http.Response
}
