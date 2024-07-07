# UpdateApplicationRequestOidcClientMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUris** | Pointer to [**[]ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner**](ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner.md) |  | [optional] 
**PostLogoutRedirectUris** | Pointer to **[]string** |  | [optional] 
**BackchannelLogoutUri** | Pointer to **string** |  | [optional] 
**BackchannelLogoutSessionRequired** | Pointer to **bool** |  | [optional] 
**LogoUri** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateApplicationRequestOidcClientMetadata

`func NewUpdateApplicationRequestOidcClientMetadata() *UpdateApplicationRequestOidcClientMetadata`

NewUpdateApplicationRequestOidcClientMetadata instantiates a new UpdateApplicationRequestOidcClientMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplicationRequestOidcClientMetadataWithDefaults

`func NewUpdateApplicationRequestOidcClientMetadataWithDefaults() *UpdateApplicationRequestOidcClientMetadata`

NewUpdateApplicationRequestOidcClientMetadataWithDefaults instantiates a new UpdateApplicationRequestOidcClientMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUris

`func (o *UpdateApplicationRequestOidcClientMetadata) GetRedirectUris() []ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *UpdateApplicationRequestOidcClientMetadata) GetRedirectUrisOk() (*[]ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *UpdateApplicationRequestOidcClientMetadata) SetRedirectUris(v []ListApplications200ResponseInnerOidcClientMetadataRedirectUrisInner)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *UpdateApplicationRequestOidcClientMetadata) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetPostLogoutRedirectUris

`func (o *UpdateApplicationRequestOidcClientMetadata) GetPostLogoutRedirectUris() []string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *UpdateApplicationRequestOidcClientMetadata) GetPostLogoutRedirectUrisOk() (*[]string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *UpdateApplicationRequestOidcClientMetadata) SetPostLogoutRedirectUris(v []string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *UpdateApplicationRequestOidcClientMetadata) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### GetBackchannelLogoutUri

`func (o *UpdateApplicationRequestOidcClientMetadata) GetBackchannelLogoutUri() string`

GetBackchannelLogoutUri returns the BackchannelLogoutUri field if non-nil, zero value otherwise.

### GetBackchannelLogoutUriOk

`func (o *UpdateApplicationRequestOidcClientMetadata) GetBackchannelLogoutUriOk() (*string, bool)`

GetBackchannelLogoutUriOk returns a tuple with the BackchannelLogoutUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelLogoutUri

`func (o *UpdateApplicationRequestOidcClientMetadata) SetBackchannelLogoutUri(v string)`

SetBackchannelLogoutUri sets BackchannelLogoutUri field to given value.

### HasBackchannelLogoutUri

`func (o *UpdateApplicationRequestOidcClientMetadata) HasBackchannelLogoutUri() bool`

HasBackchannelLogoutUri returns a boolean if a field has been set.

### GetBackchannelLogoutSessionRequired

`func (o *UpdateApplicationRequestOidcClientMetadata) GetBackchannelLogoutSessionRequired() bool`

GetBackchannelLogoutSessionRequired returns the BackchannelLogoutSessionRequired field if non-nil, zero value otherwise.

### GetBackchannelLogoutSessionRequiredOk

`func (o *UpdateApplicationRequestOidcClientMetadata) GetBackchannelLogoutSessionRequiredOk() (*bool, bool)`

GetBackchannelLogoutSessionRequiredOk returns a tuple with the BackchannelLogoutSessionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelLogoutSessionRequired

`func (o *UpdateApplicationRequestOidcClientMetadata) SetBackchannelLogoutSessionRequired(v bool)`

SetBackchannelLogoutSessionRequired sets BackchannelLogoutSessionRequired field to given value.

### HasBackchannelLogoutSessionRequired

`func (o *UpdateApplicationRequestOidcClientMetadata) HasBackchannelLogoutSessionRequired() bool`

HasBackchannelLogoutSessionRequired returns a boolean if a field has been set.

### GetLogoUri

`func (o *UpdateApplicationRequestOidcClientMetadata) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *UpdateApplicationRequestOidcClientMetadata) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *UpdateApplicationRequestOidcClientMetadata) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *UpdateApplicationRequestOidcClientMetadata) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


