// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest struct {
	RequestCatalogManagementServiceAddAccessEntitlementsRequest *shared.RequestCatalogManagementServiceAddAccessEntitlementsRequest `request:"mediaType=application/json"`
	CatalogID                                                   string                                                              `pathParam:"style=simple,explode=false,name=catalog_id"`
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsResponse struct {
	ContentType string
	// Empty response with a status code indicating success.
	RequestCatalogManagementServiceAddAccessEntitlementsResponse *shared.RequestCatalogManagementServiceAddAccessEntitlementsResponse
	StatusCode                                                   int
	RawResponse                                                  *http.Response
}