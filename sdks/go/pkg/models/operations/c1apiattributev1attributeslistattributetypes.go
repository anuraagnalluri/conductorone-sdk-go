// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributesListAttributeTypesRequest struct {
	PageSize  *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
}

type C1APIAttributeV1AttributesListAttributeTypesResponse struct {
	ContentType string
	// ListAttributeTypesResponse is the response for listing attribute types.
	ListAttributeTypesResponse *shared.ListAttributeTypesResponse
	StatusCode                 int
	RawResponse                *http.Response
}