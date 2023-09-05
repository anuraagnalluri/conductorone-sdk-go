// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributesDeleteAttributeValueRequest struct {
	DeleteAttributeValueRequest *shared.DeleteAttributeValueRequest `request:"mediaType=application/json"`
	ID                          string                              `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAttributeV1AttributesDeleteAttributeValueResponse struct {
	ContentType string
	// DeleteAttributeValueResponse is the empty response for deleting an attribute value.
	DeleteAttributeValueResponse *shared.DeleteAttributeValueResponse
	StatusCode                   int
	RawResponse                  *http.Response
}
