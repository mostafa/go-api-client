# \VerificationCodesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerificationCode**](VerificationCodesAPI.md#CreateVerificationCode) | **Post** /api/verification-codes | Request and send a verification code
[**VerifyVerificationCode**](VerificationCodesAPI.md#VerifyVerificationCode) | **Post** /api/verification-codes/verify | Verify a verification code



## CreateVerificationCode

> CreateVerificationCode(ctx).ApiInteractionVerificationVerificationCodePostRequest(apiInteractionVerificationVerificationCodePostRequest).Execute()

Request and send a verification code



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
	apiInteractionVerificationVerificationCodePostRequest := logto._api_interaction_verification_verification_code_post_request{ApiInteractionVerificationVerificationCodePostRequestOneOf: logto.NewApiInteractionVerificationVerificationCodePostRequestOneOf("Email_example")} // ApiInteractionVerificationVerificationCodePostRequest | 

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	r, err := apiClient.VerificationCodesAPI.CreateVerificationCode(context.Background()).ApiInteractionVerificationVerificationCodePostRequest(apiInteractionVerificationVerificationCodePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationCodesAPI.CreateVerificationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerificationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInteractionVerificationVerificationCodePostRequest** | [**ApiInteractionVerificationVerificationCodePostRequest**](ApiInteractionVerificationVerificationCodePostRequest.md) |  | 

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


## VerifyVerificationCode

> VerifyVerificationCode(ctx).VerifyVerificationCodeRequest(verifyVerificationCodeRequest).Execute()

Verify a verification code



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
	verifyVerificationCodeRequest := logto.VerifyVerificationCode_request{ApiInteractionPutRequestIdentifierOneOf3: logto.NewApiInteractionPutRequestIdentifierOneOf3("Email_example", "VerificationCode_example")} // VerifyVerificationCodeRequest | 

	configuration := logto.NewConfiguration()
	apiClient := logto.NewAPIClient(configuration)
	r, err := apiClient.VerificationCodesAPI.VerifyVerificationCode(context.Background()).VerifyVerificationCodeRequest(verifyVerificationCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationCodesAPI.VerifyVerificationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyVerificationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyVerificationCodeRequest** | [**VerifyVerificationCodeRequest**](VerifyVerificationCodeRequest.md) |  | 

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

