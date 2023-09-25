# ExpressionApproval

The ExpressionApproval message.


## Fields

| Field                                                                                                                            | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `AllowSelfApproval`                                                                                                              | **bool*                                                                                                                          | :heavy_minus_sign:                                                                                                               | Configuration to allow self approval of if the user is specified and also the target of the ticket.                              |
| `AssignedUserIds`                                                                                                                | []*string*                                                                                                                       | :heavy_minus_sign:                                                                                                               | The assignedUserIds field.                                                                                                       |
| `Expressions`                                                                                                                    | []*string*                                                                                                                       | :heavy_minus_sign:                                                                                                               | Array of dynamic expressions to determine the approvers.  The first expression to return a non-empty list of users will be used. |
| `Fallback`                                                                                                                       | **bool*                                                                                                                          | :heavy_minus_sign:                                                                                                               | Configuration to allow a fallback if the expression does not return a valid list of users.                                       |
| `FallbackUserIds`                                                                                                                | []*string*                                                                                                                       | :heavy_minus_sign:                                                                                                               | Configuration to specific which users to fallback to if and the expression does not return a valid list of users.                |