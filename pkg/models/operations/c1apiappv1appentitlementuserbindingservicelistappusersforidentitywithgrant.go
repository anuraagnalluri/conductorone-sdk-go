// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone-api/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantRequest struct {
	AppEntitlementID string `pathParam:"style=simple,explode=false,name=app_entitlement_id"`
	AppID            string `pathParam:"style=simple,explode=false,name=app_id"`
	IdentityUserID   string `pathParam:"style=simple,explode=false,name=identity_user_id"`
}

type C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	C1APIAppV1ListAppUsersForIdentityWithGrantResponse *shared.C1APIAppV1ListAppUsersForIdentityWithGrantResponse
}
