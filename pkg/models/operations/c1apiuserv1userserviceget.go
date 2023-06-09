// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone-api/pkg/models/shared"
	"net/http"
)

type C1APIUserV1UserServiceGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIUserV1UserServiceGetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	C1APIUserV1UserServiceGetResponse *shared.C1APIUserV1UserServiceGetResponse
}
