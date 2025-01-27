// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsListRequest struct {
	AppID     string   `pathParam:"style=simple,explode=false,name=app_id"`
	PageSize  *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppEntitlementsListRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppEntitlementsListRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppEntitlementsListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppEntitlementsListResponse struct {
	ContentType string
	// The ListAppEntitlementsResponse message contains a list of results and a nextPageToken if applicable.
	ListAppEntitlementsResponse *shared.ListAppEntitlementsResponse
	StatusCode                  int
	RawResponse                 *http.Response
}

func (o *C1APIAppV1AppEntitlementsListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementsListResponse) GetListAppEntitlementsResponse() *shared.ListAppEntitlementsResponse {
	if o == nil {
		return nil
	}
	return o.ListAppEntitlementsResponse
}

func (o *C1APIAppV1AppEntitlementsListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementsListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
