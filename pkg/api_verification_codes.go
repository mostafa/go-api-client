/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// VerificationCodesAPIService VerificationCodesAPI service
type VerificationCodesAPIService service

type ApiCreateVerificationCodeRequest struct {
	ctx context.Context
	ApiService *VerificationCodesAPIService
	apiInteractionVerificationVerificationCodePostRequest *ApiInteractionVerificationVerificationCodePostRequest
}

func (r ApiCreateVerificationCodeRequest) ApiInteractionVerificationVerificationCodePostRequest(apiInteractionVerificationVerificationCodePostRequest ApiInteractionVerificationVerificationCodePostRequest) ApiCreateVerificationCodeRequest {
	r.apiInteractionVerificationVerificationCodePostRequest = &apiInteractionVerificationVerificationCodePostRequest
	return r
}

func (r ApiCreateVerificationCodeRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateVerificationCodeExecute(r)
}

/*
CreateVerificationCode Request and send a verification code

Request a verification code for the provided identifier (email/phone).
if you're using email as the identifier, you need to setup your email connector first.
if you're using phone as the identifier, you need to setup your SMS connector first.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateVerificationCodeRequest
*/
func (a *VerificationCodesAPIService) CreateVerificationCode(ctx context.Context) ApiCreateVerificationCodeRequest {
	return ApiCreateVerificationCodeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *VerificationCodesAPIService) CreateVerificationCodeExecute(r ApiCreateVerificationCodeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VerificationCodesAPIService.CreateVerificationCode")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/verification-codes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiInteractionVerificationVerificationCodePostRequest == nil {
		return nil, reportError("apiInteractionVerificationVerificationCodePostRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiInteractionVerificationVerificationCodePostRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiVerifyVerificationCodeRequest struct {
	ctx context.Context
	ApiService *VerificationCodesAPIService
	verifyVerificationCodeRequest *VerifyVerificationCodeRequest
}

func (r ApiVerifyVerificationCodeRequest) VerifyVerificationCodeRequest(verifyVerificationCodeRequest VerifyVerificationCodeRequest) ApiVerifyVerificationCodeRequest {
	r.verifyVerificationCodeRequest = &verifyVerificationCodeRequest
	return r
}

func (r ApiVerifyVerificationCodeRequest) Execute() (*http.Response, error) {
	return r.ApiService.VerifyVerificationCodeExecute(r)
}

/*
VerifyVerificationCode Verify a verification code

Verify a verification code for a specified identifier.
if you're using email as the identifier, you need to setup your email connector first.
if you're using phone as the identifier, you need to setup your SMS connector first.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifyVerificationCodeRequest
*/
func (a *VerificationCodesAPIService) VerifyVerificationCode(ctx context.Context) ApiVerifyVerificationCodeRequest {
	return ApiVerifyVerificationCodeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *VerificationCodesAPIService) VerifyVerificationCodeExecute(r ApiVerifyVerificationCodeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VerificationCodesAPIService.VerifyVerificationCode")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/verification-codes/verify"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.verifyVerificationCodeRequest == nil {
		return nil, reportError("verifyVerificationCodeRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.verifyVerificationCodeRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
