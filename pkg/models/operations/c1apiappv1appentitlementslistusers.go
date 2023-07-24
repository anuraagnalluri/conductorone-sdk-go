// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsListUsersRequest struct {
	AppEntitlementID string `pathParam:"style=simple,explode=false,name=app_entitlement_id"`
	AppID            string `pathParam:"style=simple,explode=false,name=app_id"`
}

func (o *C1APIAppV1AppEntitlementsListUsersRequest) GetAppEntitlementID() string {
	if o == nil {
		return ""
	}
	return o.AppEntitlementID
}

func (o *C1APIAppV1AppEntitlementsListUsersRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

type C1APIAppV1AppEntitlementsListUsersResponse struct {
	ContentType string
	// Successful response
	ListAppEntitlementUsersResponse *shared.ListAppEntitlementUsersResponse
	StatusCode                      int
	RawResponse                     *http.Response
}

func (o *C1APIAppV1AppEntitlementsListUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementsListUsersResponse) GetListAppEntitlementUsersResponse() *shared.ListAppEntitlementUsersResponse {
	if o == nil {
		return nil
	}
	return o.ListAppEntitlementUsersResponse
}

func (o *C1APIAppV1AppEntitlementsListUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementsListUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}