# \SSOConnectorProvidersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSsoConnectorProviders**](SSOConnectorProvidersAPI.md#ListSsoConnectorProviders) | **Get** /api/sso-connector-providers | List all the supported SSO connector provider details



## ListSsoConnectorProviders

> []ListSsoConnectorProviders200ResponseInner ListSsoConnectorProviders(ctx).Execute()

List all the supported SSO connector provider details



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
	resp, r, err := apiClient.SSOConnectorProvidersAPI.ListSsoConnectorProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOConnectorProvidersAPI.ListSsoConnectorProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSsoConnectorProviders`: []ListSsoConnectorProviders200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SSOConnectorProvidersAPI.ListSsoConnectorProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSsoConnectorProvidersRequest struct via the builder pattern


### Return type

[**[]ListSsoConnectorProviders200ResponseInner**](ListSsoConnectorProviders200ResponseInner.md)

### Authorization

[ManagementApi](../README.md#ManagementApi)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

