// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppReportActionServiceGenerateReportRequest struct {
	AppActionsServiceGenerateReportRequest *shared.AppActionsServiceGenerateReportRequest `request:"mediaType=application/json"`
	AppID                                  string                                         `pathParam:"style=simple,explode=false,name=app_id"`
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportRequest) GetAppActionsServiceGenerateReportRequest() *shared.AppActionsServiceGenerateReportRequest {
	if o == nil {
		return nil
	}
	return o.AppActionsServiceGenerateReportRequest
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

type C1APIAppV1AppReportActionServiceGenerateReportResponse struct {
	// Empty response body. Status code indicates success.
	AppActionsServiceGenerateReportResponse *shared.AppActionsServiceGenerateReportResponse
	ContentType                             string
	StatusCode                              int
	RawResponse                             *http.Response
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportResponse) GetAppActionsServiceGenerateReportResponse() *shared.AppActionsServiceGenerateReportResponse {
	if o == nil {
		return nil
	}
	return o.AppActionsServiceGenerateReportResponse
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
