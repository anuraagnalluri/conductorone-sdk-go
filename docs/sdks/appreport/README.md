# AppReport

### Available Operations

* [List](#list) - List

## List

Get a list of reports for the given app.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppReport.List(ctx, operations.C1APIAppV1AppReportServiceListRequest{
        AppID: "quasi",
        PageSize: conductoroneapi.Float64(9719.45),
        PageToken: conductoroneapi.String("voluptatibus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppReportServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.C1APIAppV1AppReportServiceListRequest](../../models/operations/c1apiappv1appreportservicelistrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.C1APIAppV1AppReportServiceListResponse](../../models/operations/c1apiappv1appreportservicelistresponse.md), error**

