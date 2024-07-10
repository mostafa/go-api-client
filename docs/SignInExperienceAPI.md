# \SignInExperienceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSignInExp**](SignInExperienceAPI.md#GetSignInExp) | **Get** /api/sign-in-exp | Get default sign-in experience settings
[**UpdateSignInExp**](SignInExperienceAPI.md#UpdateSignInExp) | **Patch** /api/sign-in-exp | Update default sign-in experience settings



## GetSignInExp

> GetSignInExp200Response GetSignInExp(ctx).Execute()

Get default sign-in experience settings



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

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	resp, r, err := apiClient.SignInExperienceAPI.GetSignInExp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignInExperienceAPI.GetSignInExp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignInExp`: GetSignInExp200Response
	fmt.Fprintf(os.Stdout, "Response from `SignInExperienceAPI.GetSignInExp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignInExpRequest struct via the builder pattern


### Return type

[**GetSignInExp200Response**](GetSignInExp200Response.md)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSignInExp

> UpdateSignInExp200Response UpdateSignInExp(ctx).UpdateSignInExpRequest(updateSignInExpRequest).RemoveUnusedDemoSocialConnector(removeUnusedDemoSocialConnector).Execute()

Update default sign-in experience settings



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
	updateSignInExpRequest := *logto.NewUpdateSignInExpRequest() // UpdateSignInExpRequest | 
	removeUnusedDemoSocialConnector := "removeUnusedDemoSocialConnector_example" // string | Whether to remove unused demo social connectors. (These demo social connectors are only used during cloud user onboarding) (optional)

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	resp, r, err := apiClient.SignInExperienceAPI.UpdateSignInExp(context.Background()).UpdateSignInExpRequest(updateSignInExpRequest).RemoveUnusedDemoSocialConnector(removeUnusedDemoSocialConnector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignInExperienceAPI.UpdateSignInExp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSignInExp`: UpdateSignInExp200Response
	fmt.Fprintf(os.Stdout, "Response from `SignInExperienceAPI.UpdateSignInExp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSignInExpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSignInExpRequest** | [**UpdateSignInExpRequest**](UpdateSignInExpRequest.md) |  | 
 **removeUnusedDemoSocialConnector** | **string** | Whether to remove unused demo social connectors. (These demo social connectors are only used during cloud user onboarding) | 

### Return type

[**UpdateSignInExp200Response**](UpdateSignInExp200Response.md)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

