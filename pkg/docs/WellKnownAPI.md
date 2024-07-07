# \WellKnownAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSignInExperienceConfig**](WellKnownAPI.md#GetSignInExperienceConfig) | **Get** /api/.well-known/sign-in-exp | Get full sign-in experience
[**GetSignInExperiencePhrases**](WellKnownAPI.md#GetSignInExperiencePhrases) | **Get** /api/.well-known/phrases | Get localized phrases



## GetSignInExperienceConfig

> GetSignInExperienceConfig200Response GetSignInExperienceConfig(ctx).Execute()

Get full sign-in experience



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
	resp, r, err := apiClient.WellKnownAPI.GetSignInExperienceConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WellKnownAPI.GetSignInExperienceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignInExperienceConfig`: GetSignInExperienceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `WellKnownAPI.GetSignInExperienceConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignInExperienceConfigRequest struct via the builder pattern


### Return type

[**GetSignInExperienceConfig200Response**](GetSignInExperienceConfig200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignInExperiencePhrases

> map[string]GetSignInExperiencePhrases200ResponseValue GetSignInExperiencePhrases(ctx).Lng(lng).Execute()

Get localized phrases



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
	lng := "lng_example" // string | The language tag for localization. (optional)

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	resp, r, err := apiClient.WellKnownAPI.GetSignInExperiencePhrases(context.Background()).Lng(lng).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WellKnownAPI.GetSignInExperiencePhrases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignInExperiencePhrases`: map[string]GetSignInExperiencePhrases200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `WellKnownAPI.GetSignInExperiencePhrases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignInExperiencePhrasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lng** | **string** | The language tag for localization. | 

### Return type

[**map[string]GetSignInExperiencePhrases200ResponseValue**](GetSignInExperiencePhrases200ResponseValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

