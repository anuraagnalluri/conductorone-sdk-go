# Directory

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List

## Create

Create a directory.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Directory.Create(ctx, shared.DirectoryServiceCreateRequest{
        AppID: conductoroneapi.String("accusantium"),
        ExpandMask: &shared.DirectoryExpandMask{
            Paths: []string{
                "sapiente",
                "dicta",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectoryServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.DirectoryServiceCreateRequest](../../models/shared/directoryservicecreaterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.C1APIDirectoryV1DirectoryServiceCreateResponse](../../models/operations/c1apidirectoryv1directoryservicecreateresponse.md), error**


## Delete

Delete a directory by app_id.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Directory.Delete(ctx, operations.C1APIDirectoryV1DirectoryServiceDeleteRequest{
        DirectoryServiceDeleteRequest: &shared.DirectoryServiceDeleteRequest{},
        AppID: "ullam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectoryServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APIDirectoryV1DirectoryServiceDeleteRequest](../../models/operations/c1apidirectoryv1directoryservicedeleterequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.C1APIDirectoryV1DirectoryServiceDeleteResponse](../../models/operations/c1apidirectoryv1directoryservicedeleteresponse.md), error**


## Get

Get a directory by app_id.

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
    res, err := s.Directory.Get(ctx, operations.C1APIDirectoryV1DirectoryServiceGetRequest{
        AppID: "reprehenderit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectoryServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APIDirectoryV1DirectoryServiceGetRequest](../../models/operations/c1apidirectoryv1directoryservicegetrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.C1APIDirectoryV1DirectoryServiceGetResponse](../../models/operations/c1apidirectoryv1directoryservicegetresponse.md), error**


## List

List directories.

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
    res, err := s.Directory.List(ctx, operations.C1APIDirectoryV1DirectoryServiceListRequest{
        PageSize: conductoroneapi.Float64(3567.07),
        PageToken: conductoroneapi.String("nisi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectoryServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIDirectoryV1DirectoryServiceListRequest](../../models/operations/c1apidirectoryv1directoryservicelistrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.C1APIDirectoryV1DirectoryServiceListResponse](../../models/operations/c1apidirectoryv1directoryservicelistresponse.md), error**

