# AppEntitlementUserBinding

### Available Operations

* [ListAppUsersForIdentityWithGrant](#listappusersforidentitywithgrant) - List App Users For Identity With Grant

## ListAppUsersForIdentityWithGrant

Returns a list of app users for the identity in the app. If that app user also has a grant to the entitlement from the request, data about the grant is also returned. It will always return ALL app users for this identity, but only SOME may have grant data.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
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
    res, err := s.AppEntitlementUserBinding.ListAppUsersForIdentityWithGrant(ctx, operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantRequest{
        AppEntitlementID: "occaecati",
        AppID: "fugit",
        IdentityUserID: "deleniti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppUsersForIdentityWithGrantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                    | Type                                                                                                                                                                                                         | Required                                                                                                                                                                                                     | Description                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                          |
| `request`                                                                                                                                                                                                    | [operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantRequest](../../models/operations/c1apiappv1appentitlementuserbindingservicelistappusersforidentitywithgrantrequest.md) | :heavy_check_mark:                                                                                                                                                                                           | The request object to use for the request.                                                                                                                                                                   |


### Response

**[*operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantResponse](../../models/operations/c1apiappv1appentitlementuserbindingservicelistappusersforidentitywithgrantresponse.md), error**

