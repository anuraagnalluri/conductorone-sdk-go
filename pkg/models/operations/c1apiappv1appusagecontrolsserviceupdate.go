// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppUsageControlsServiceUpdateRequest struct {
	UpdateAppUsageControlsRequest *shared.UpdateAppUsageControlsRequest `request:"mediaType=application/json"`
	AppID                         string                                `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1AppUsageControlsServiceUpdateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	UpdateAppUsageControlsResponse *shared.UpdateAppUsageControlsResponse
}