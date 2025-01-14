// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1ConnectorServiceCreateDelegatedRequest struct {
	ConnectorServiceCreateDelegatedRequest *shared.ConnectorServiceCreateDelegatedRequest `request:"mediaType=application/json"`
	AppID                                  string                                         `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1ConnectorServiceCreateDelegatedResponse struct {
	// The ConnectorServiceCreateResponse is the response returned from creating a connector.
	ConnectorServiceCreateResponse *shared.ConnectorServiceCreateResponse
	ContentType                    string
	StatusCode                     int
	RawResponse                    *http.Response
}
