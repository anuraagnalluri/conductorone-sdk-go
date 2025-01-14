// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppOwnersListRequest struct {
	AppID     string   `pathParam:"style=simple,explode=false,name=app_id"`
	PageSize  *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppOwnersListRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppOwnersListRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppOwnersListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppOwnersListResponse struct {
	ContentType string
	// Successful response
	ListAppOwnersResponse *shared.ListAppOwnersResponse
	StatusCode            int
	RawResponse           *http.Response
}

func (o *C1APIAppV1AppOwnersListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppOwnersListResponse) GetListAppOwnersResponse() *shared.ListAppOwnersResponse {
	if o == nil {
		return nil
	}
	return o.ListAppOwnersResponse
}

func (o *C1APIAppV1AppOwnersListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppOwnersListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
