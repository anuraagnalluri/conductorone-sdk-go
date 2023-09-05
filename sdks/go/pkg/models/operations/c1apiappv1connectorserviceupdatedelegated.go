// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1ConnectorServiceUpdateDelegatedRequest struct {
	ConnectorServiceUpdateDelegatedRequestInput *shared.ConnectorServiceUpdateDelegatedRequestInput `request:"mediaType=application/json"`
	ConnectorAppID                              string                                              `pathParam:"style=simple,explode=false,name=connector_app_id"`
	ConnectorID                                 string                                              `pathParam:"style=simple,explode=false,name=connector_id"`
}

type C1APIAppV1ConnectorServiceUpdateDelegatedResponse struct {
	// ConnectorServiceUpdateResponse is the response returned by the update method.
	ConnectorServiceUpdateResponse *shared.ConnectorServiceUpdateResponse
	ContentType                    string
	StatusCode                     int
	RawResponse                    *http.Response
}