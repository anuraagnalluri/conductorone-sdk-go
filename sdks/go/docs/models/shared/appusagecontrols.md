# AppUsageControls

The AppUsageControls object describes some peripheral configuration for an app.


## Fields

| Field                                                                                                                      | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `AppID`                                                                                                                    | **string*                                                                                                                  | :heavy_minus_sign:                                                                                                         | The app that this object belongs to.                                                                                       |
| `Notify`                                                                                                                   | **bool*                                                                                                                    | :heavy_minus_sign:                                                                                                         | Whether or not to notify some if they have access to the app, but has not used it within a configurable amount of time.    |
| `NotifyAfterDays`                                                                                                          | **float64*                                                                                                                 | :heavy_minus_sign:                                                                                                         | The duration in days after which we notify users of nonusage.                                                              |
| `Revoke`                                                                                                                   | **bool*                                                                                                                    | :heavy_minus_sign:                                                                                                         | Whether or not to revoke a grant if they have access to the app, but has not used it within a configurable amount of time. |
| `RevokeAfterDays`                                                                                                          | **float64*                                                                                                                 | :heavy_minus_sign:                                                                                                         | The duration in days after which we revoke users that have not used that grant.                                            |