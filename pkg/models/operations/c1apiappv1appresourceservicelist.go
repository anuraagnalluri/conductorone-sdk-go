// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceServiceListRequest struct {
	AppID             string   `pathParam:"style=simple,explode=false,name=app_id"`
	AppResourceTypeID string   `pathParam:"style=simple,explode=false,name=app_resource_type_id"`
	PageSize          *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken         *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppResourceServiceListRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppResourceServiceListRequest) GetAppResourceTypeID() string {
	if o == nil {
		return ""
	}
	return o.AppResourceTypeID
}

func (o *C1APIAppV1AppResourceServiceListRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppResourceServiceListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppResourceServiceListResponse struct {
	// The AppResourceServiceListResponse message contains a list of results and a nextPageToken if applicable.
	AppResourceServiceListResponse *shared.AppResourceServiceListResponse
	ContentType                    string
	StatusCode                     int
	RawResponse                    *http.Response
}

func (o *C1APIAppV1AppResourceServiceListResponse) GetAppResourceServiceListResponse() *shared.AppResourceServiceListResponse {
	if o == nil {
		return nil
	}
	return o.AppResourceServiceListResponse
}

func (o *C1APIAppV1AppResourceServiceListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppResourceServiceListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppResourceServiceListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
