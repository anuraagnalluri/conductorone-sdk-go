# C1APITaskV1TaskActionsServiceDenyResponse


## Fields

| Field                                                                                                                     | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ContentType`                                                                                                             | *string*                                                                                                                  | :heavy_check_mark:                                                                                                        | N/A                                                                                                                       |
| `StatusCode`                                                                                                              | *int*                                                                                                                     | :heavy_check_mark:                                                                                                        | N/A                                                                                                                       |
| `RawResponse`                                                                                                             | [*http.Response](https://pkg.go.dev/net/http#Response)                                                                    | :heavy_minus_sign:                                                                                                        | N/A                                                                                                                       |
| `TaskActionsServiceDenyResponse`                                                                                          | [*shared.TaskActionsServiceDenyResponse](../../models/shared/taskactionsservicedenyresponse.md)                           | :heavy_minus_sign:                                                                                                        | The TaskActionsServiceDenyResponse returns a task view with paths indicating the location of expanded items in the array. |