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


// SignInExperienceAPIService SignInExperienceAPI service
type SignInExperienceAPIService service

type ApiGetSignInExpRequest struct {
	ctx context.Context
	ApiService *SignInExperienceAPIService
}

func (r ApiGetSignInExpRequest) Execute() (*GetSignInExp200Response, *http.Response, error) {
	return r.ApiService.GetSignInExpExecute(r)
}

/*
GetSignInExp Get default sign-in experience settings

Get the default sign-in experience settings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSignInExpRequest
*/
func (a *SignInExperienceAPIService) GetSignInExp(ctx context.Context) ApiGetSignInExpRequest {
	return ApiGetSignInExpRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSignInExp200Response
func (a *SignInExperienceAPIService) GetSignInExpExecute(r ApiGetSignInExpRequest) (*GetSignInExp200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSignInExp200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SignInExperienceAPIService.GetSignInExp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/sign-in-exp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateSignInExpRequest struct {
	ctx context.Context
	ApiService *SignInExperienceAPIService
	updateSignInExpRequest *UpdateSignInExpRequest
	removeUnusedDemoSocialConnector *string
}

func (r ApiUpdateSignInExpRequest) UpdateSignInExpRequest(updateSignInExpRequest UpdateSignInExpRequest) ApiUpdateSignInExpRequest {
	r.updateSignInExpRequest = &updateSignInExpRequest
	return r
}

// Whether to remove unused demo social connectors. (These demo social connectors are only used during cloud user onboarding)
func (r ApiUpdateSignInExpRequest) RemoveUnusedDemoSocialConnector(removeUnusedDemoSocialConnector string) ApiUpdateSignInExpRequest {
	r.removeUnusedDemoSocialConnector = &removeUnusedDemoSocialConnector
	return r
}

func (r ApiUpdateSignInExpRequest) Execute() (*UpdateSignInExp200Response, *http.Response, error) {
	return r.ApiService.UpdateSignInExpExecute(r)
}

/*
UpdateSignInExp Update default sign-in experience settings

Update the default sign-in experience settings with the provided data.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateSignInExpRequest
*/
func (a *SignInExperienceAPIService) UpdateSignInExp(ctx context.Context) ApiUpdateSignInExpRequest {
	return ApiUpdateSignInExpRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UpdateSignInExp200Response
func (a *SignInExperienceAPIService) UpdateSignInExpExecute(r ApiUpdateSignInExpRequest) (*UpdateSignInExp200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateSignInExp200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SignInExperienceAPIService.UpdateSignInExp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/sign-in-exp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateSignInExpRequest == nil {
		return localVarReturnValue, nil, reportError("updateSignInExpRequest is required and must be specified")
	}

	if r.removeUnusedDemoSocialConnector != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "removeUnusedDemoSocialConnector", r.removeUnusedDemoSocialConnector, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateSignInExpRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
