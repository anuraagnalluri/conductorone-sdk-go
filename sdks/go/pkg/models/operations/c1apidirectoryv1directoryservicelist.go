// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIDirectoryV1DirectoryServiceListRequest struct {
	PageSize  *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
}

type C1APIDirectoryV1DirectoryServiceListResponse struct {
	ContentType string
	// The DirectoryServiceListResponse message contains a list of results and a nextPageToken if applicable.
	DirectoryServiceListResponse *shared.DirectoryServiceListResponse
	StatusCode                   int
	RawResponse                  *http.Response
}