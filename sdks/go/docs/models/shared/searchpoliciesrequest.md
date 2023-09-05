# SearchPoliciesRequest

Search Policies by a few properties.


## Fields

| Field                                                                                                                                                       | Type                                                                                                                                                        | Required                                                                                                                                                    | Description                                                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `DisplayName`                                                                                                                                               | **string*                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                          | Search for policies with a case insensitive match on the display name.                                                                                      |
| `PageSize`                                                                                                                                                  | **float64*                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                          | The pageSize where 0 <= pageSize <= 100. Values < 10 will be set to 10. A value of 0 returns the default page size (currently 25)                           |
| `PageToken`                                                                                                                                                 | **string*                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                          | The pageToken field.                                                                                                                                        |
| `PolicyTypes`                                                                                                                                               | [][SearchPoliciesRequestPolicyTypes](../../models/shared/searchpoliciesrequestpolicytypes.md)                                                               | :heavy_minus_sign:                                                                                                                                          | The policy type to search on. This can be POLICY_TYPE_GRANT, POLICY_TYPE_REVOKE, POLICY_TYPE_CERTIFY, POLICY_TYPE_ACCESS_REQUEST, or POLICY_TYPE_PROVISION. |
| `Query`                                                                                                                                                     | **string*                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                          | Query the policies with a fuzzy search on display name and description.                                                                                     |
| `Refs`                                                                                                                                                      | [][PolicyRef](../../models/shared/policyref.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                          | The refs field.                                                                                                                                             |