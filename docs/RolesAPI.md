# \RolesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](RolesAPI.md#CreateRole) | **Post** /api/roles | Create a role
[**CreateRoleApplication**](RolesAPI.md#CreateRoleApplication) | **Post** /api/roles/{id}/applications | Assign role to applications
[**CreateRoleScope**](RolesAPI.md#CreateRoleScope) | **Post** /api/roles/{id}/scopes | Link scopes to role
[**CreateRoleUser**](RolesAPI.md#CreateRoleUser) | **Post** /api/roles/{id}/users | Assign role to users
[**DeleteRole**](RolesAPI.md#DeleteRole) | **Delete** /api/roles/{id} | Delete role
[**DeleteRoleApplication**](RolesAPI.md#DeleteRoleApplication) | **Delete** /api/roles/{id}/applications/{applicationId} | Remove role from application
[**DeleteRoleScope**](RolesAPI.md#DeleteRoleScope) | **Delete** /api/roles/{id}/scopes/{scopeId} | Unlink scope from role
[**DeleteRoleUser**](RolesAPI.md#DeleteRoleUser) | **Delete** /api/roles/{id}/users/{userId} | Remove role from user
[**GetRole**](RolesAPI.md#GetRole) | **Get** /api/roles/{id} | Get role
[**ListRoleApplications**](RolesAPI.md#ListRoleApplications) | **Get** /api/roles/{id}/applications | Get role applications
[**ListRoleScopes**](RolesAPI.md#ListRoleScopes) | **Get** /api/roles/{id}/scopes | Get role scopes
[**ListRoleUsers**](RolesAPI.md#ListRoleUsers) | **Get** /api/roles/{id}/users | Get role users
[**ListRoles**](RolesAPI.md#ListRoles) | **Get** /api/roles | Get roles
[**UpdateRole**](RolesAPI.md#UpdateRole) | **Patch** /api/roles/{id} | Update role



## CreateRole

> ListApplicationRoles200ResponseInner CreateRole(ctx).CreateRoleRequest(createRoleRequest).Execute()

Create a role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	createRoleRequest := *logto.NewCreateRoleRequest("Name_example", "Description_example") // CreateRoleRequest | 

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.CreateRole(context.Background()).CreateRoleRequest(createRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.CreateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRole`: ListApplicationRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRoleRequest** | [**CreateRoleRequest**](CreateRoleRequest.md) |  | 

### Return type

[**ListApplicationRoles200ResponseInner**](ListApplicationRoles200ResponseInner.md)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoleApplication

> CreateRoleApplication(ctx, id).CreateRoleApplicationRequest(createRoleApplicationRequest).Execute()

Assign role to applications



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.
	createRoleApplicationRequest := *logto.NewCreateRoleApplicationRequest([]string{"ApplicationIds_example"}) // CreateRoleApplicationRequest | 

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.CreateRoleApplication(context.Background(), id).CreateRoleApplicationRequest(createRoleApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.CreateRoleApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRoleApplicationRequest** | [**CreateRoleApplicationRequest**](CreateRoleApplicationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoleScope

> CreateRoleScope(ctx, id).CreateRoleScopeRequest(createRoleScopeRequest).Execute()

Link scopes to role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.
	createRoleScopeRequest := *logto.NewCreateRoleScopeRequest([]string{"ScopeIds_example"}) // CreateRoleScopeRequest | 

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.CreateRoleScope(context.Background(), id).CreateRoleScopeRequest(createRoleScopeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.CreateRoleScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRoleScopeRequest** | [**CreateRoleScopeRequest**](CreateRoleScopeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoleUser

> CreateRoleUser(ctx, id).CreateRoleUserRequest(createRoleUserRequest).Execute()

Assign role to users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.
	createRoleUserRequest := *logto.NewCreateRoleUserRequest([]string{"UserIds_example"}) // CreateRoleUserRequest | 

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.CreateRoleUser(context.Background(), id).CreateRoleUserRequest(createRoleUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.CreateRoleUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRoleUserRequest** | [**CreateRoleUserRequest**](CreateRoleUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, id).Execute()

Delete role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.DeleteRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleApplication

> DeleteRoleApplication(ctx, id, applicationId).Execute()

Remove role from application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.
	applicationId := "applicationId_example" // string | The unique identifier of the application.

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.DeleteRoleApplication(context.Background(), id, applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.DeleteRoleApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 
**applicationId** | **string** | The unique identifier of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleScope

> DeleteRoleScope(ctx, id, scopeId).Execute()

Unlink scope from role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.
	scopeId := "scopeId_example" // string | The unique identifier of the scope.

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.DeleteRoleScope(context.Background(), id, scopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.DeleteRoleScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 
**scopeId** | **string** | The unique identifier of the scope. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleUser

> DeleteRoleUser(ctx, id, userId).Execute()

Remove role from user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.
	userId := "userId_example" // string | The unique identifier of the user.

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.DeleteRoleUser(context.Background(), id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.DeleteRoleUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> ListApplicationRoles200ResponseInner GetRole(ctx, id).Execute()

Get role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: ListApplicationRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListApplicationRoles200ResponseInner**](ListApplicationRoles200ResponseInner.md)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleApplications

> []ListApplications200ResponseInner ListRoleApplications(ctx, id).Page(page).PageSize(pageSize).Execute()

Get role applications



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ListRoleApplications(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ListRoleApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoleApplications`: []ListApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ListRoleApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleScopes

> []ListRoleScopes200ResponseInner ListRoleScopes(ctx, id).Page(page).PageSize(pageSize).Execute()

Get role scopes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ListRoleScopes(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ListRoleScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoleScopes`: []ListRoleScopes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ListRoleScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListRoleScopes200ResponseInner**](ListRoleScopes200ResponseInner.md)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleUsers

> []UpdateUser200Response ListRoleUsers(ctx, id).Page(page).PageSize(pageSize).Execute()

Get role users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ListRoleUsers(context.Background(), id).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ListRoleUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoleUsers`: []UpdateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ListRoleUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]UpdateUser200Response**](UpdateUser200Response.md)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> []ListRoles200ResponseInner ListRoles(ctx).ExcludeUserId(excludeUserId).ExcludeApplicationId(excludeApplicationId).Type_(type_).Page(page).PageSize(pageSize).Execute()

Get roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	excludeUserId := "excludeUserId_example" // string | Exclude roles assigned to a user. (optional)
	excludeApplicationId := "excludeApplicationId_example" // string | Exclude roles assigned to an application. (optional)
	type_ := "type__example" // string | Filter by role type. (optional)
	page := int32(56) // int32 | Page number (starts from 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Entries per page. (optional) (default to 20)

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ListRoles(context.Background()).ExcludeUserId(excludeUserId).ExcludeApplicationId(excludeApplicationId).Type_(type_).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoles`: []ListRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **excludeUserId** | **string** | Exclude roles assigned to a user. | 
 **excludeApplicationId** | **string** | Exclude roles assigned to an application. | 
 **type_** | **string** | Filter by role type. | 
 **page** | **int32** | Page number (starts from 1). | [default to 1]
 **pageSize** | **int32** | Entries per page. | [default to 20]

### Return type

[**[]ListRoles200ResponseInner**](ListRoles200ResponseInner.md)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> ListApplicationRoles200ResponseInner UpdateRole(ctx, id).UpdateRoleRequest(updateRoleRequest).Execute()

Update role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	logto "github.com/mostafa/go-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the role.
	updateRoleRequest := *logto.NewUpdateRoleRequest() // UpdateRoleRequest | 

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateRole(context.Background(), id).UpdateRoleRequest(updateRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRole`: ListApplicationRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleRequest** | [**UpdateRoleRequest**](UpdateRoleRequest.md) |  | 

### Return type

[**ListApplicationRoles200ResponseInner**](ListApplicationRoles200ResponseInner.md)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

