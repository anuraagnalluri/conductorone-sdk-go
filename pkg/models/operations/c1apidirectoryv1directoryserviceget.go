// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIDirectoryV1DirectoryServiceGetRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIDirectoryV1DirectoryServiceGetResponse struct {
	ContentType string
	// Successful response
	DirectoryServiceGetResponse *shared.DirectoryServiceGetResponse
	StatusCode                  int
	RawResponse                 *http.Response
}