/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetUser200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUser200Response{}

// GetUser200Response struct for GetUser200Response
type GetUser200Response struct {
	Id string `json:"id"`
	Username NullableString `json:"username"`
	PrimaryEmail NullableString `json:"primaryEmail"`
	PrimaryPhone NullableString `json:"primaryPhone"`
	Name NullableString `json:"name"`
	Avatar NullableString `json:"avatar"`
	// arbitrary
	CustomData map[string]interface{} `json:"customData"`
	Identities map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue `json:"identities"`
	LastSignInAt NullableFloat32 `json:"lastSignInAt"`
	CreatedAt float32 `json:"createdAt"`
	UpdatedAt float32 `json:"updatedAt"`
	Profile GetJwtCustomizer200ResponseOneOfContextSampleUserProfile `json:"profile"`
	ApplicationId NullableString `json:"applicationId"`
	IsSuspended bool `json:"isSuspended"`
	HasPassword *bool `json:"hasPassword,omitempty"`
	// List of SSO identities associated with the user. Only available when the `includeSsoIdentities` query parameter is provided with a truthy value.
	SsoIdentities []GetUser200ResponseSsoIdentitiesInner `json:"ssoIdentities,omitempty"`
}

type _GetUser200Response GetUser200Response

// NewGetUser200Response instantiates a new GetUser200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUser200Response(id string, username NullableString, primaryEmail NullableString, primaryPhone NullableString, name NullableString, avatar NullableString, customData map[string]interface{}, identities map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue, lastSignInAt NullableFloat32, createdAt float32, updatedAt float32, profile GetJwtCustomizer200ResponseOneOfContextSampleUserProfile, applicationId NullableString, isSuspended bool) *GetUser200Response {
	this := GetUser200Response{}
	this.Id = id
	this.Username = username
	this.PrimaryEmail = primaryEmail
	this.PrimaryPhone = primaryPhone
	this.Name = name
	this.Avatar = avatar
	this.CustomData = customData
	this.Identities = identities
	this.LastSignInAt = lastSignInAt
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Profile = profile
	this.ApplicationId = applicationId
	this.IsSuspended = isSuspended
	return &this
}

// NewGetUser200ResponseWithDefaults instantiates a new GetUser200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUser200ResponseWithDefaults() *GetUser200Response {
	this := GetUser200Response{}
	return &this
}

// GetId returns the Id field value
func (o *GetUser200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetUser200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetUser200Response) SetId(v string) {
	o.Id = v
}

// GetUsername returns the Username field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetUser200Response) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}

	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUser200Response) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// SetUsername sets field value
func (o *GetUser200Response) SetUsername(v string) {
	o.Username.Set(&v)
}

// GetPrimaryEmail returns the PrimaryEmail field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetUser200Response) GetPrimaryEmail() string {
	if o == nil || o.PrimaryEmail.Get() == nil {
		var ret string
		return ret
	}

	return *o.PrimaryEmail.Get()
}

// GetPrimaryEmailOk returns a tuple with the PrimaryEmail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUser200Response) GetPrimaryEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryEmail.Get(), o.PrimaryEmail.IsSet()
}

// SetPrimaryEmail sets field value
func (o *GetUser200Response) SetPrimaryEmail(v string) {
	o.PrimaryEmail.Set(&v)
}

// GetPrimaryPhone returns the PrimaryPhone field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetUser200Response) GetPrimaryPhone() string {
	if o == nil || o.PrimaryPhone.Get() == nil {
		var ret string
		return ret
	}

	return *o.PrimaryPhone.Get()
}

// GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUser200Response) GetPrimaryPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryPhone.Get(), o.PrimaryPhone.IsSet()
}

// SetPrimaryPhone sets field value
func (o *GetUser200Response) SetPrimaryPhone(v string) {
	o.PrimaryPhone.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetUser200Response) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUser200Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *GetUser200Response) SetName(v string) {
	o.Name.Set(&v)
}

// GetAvatar returns the Avatar field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetUser200Response) GetAvatar() string {
	if o == nil || o.Avatar.Get() == nil {
		var ret string
		return ret
	}

	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUser200Response) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// SetAvatar sets field value
func (o *GetUser200Response) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// GetCustomData returns the CustomData field value
func (o *GetUser200Response) GetCustomData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value
// and a boolean to check if the value has been set.
func (o *GetUser200Response) GetCustomDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.CustomData, true
}

// SetCustomData sets field value
func (o *GetUser200Response) SetCustomData(v map[string]interface{}) {
	o.CustomData = v
}

// GetIdentities returns the Identities field value
func (o *GetUser200Response) GetIdentities() map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue {
	if o == nil {
		var ret map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue
		return ret
	}

	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value
// and a boolean to check if the value has been set.
func (o *GetUser200Response) GetIdentitiesOk() (*map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identities, true
}

// SetIdentities sets field value
func (o *GetUser200Response) SetIdentities(v map[string]GetJwtCustomizer200ResponseOneOfContextSampleUserIdentitiesValue) {
	o.Identities = v
}

// GetLastSignInAt returns the LastSignInAt field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *GetUser200Response) GetLastSignInAt() float32 {
	if o == nil || o.LastSignInAt.Get() == nil {
		var ret float32
		return ret
	}

	return *o.LastSignInAt.Get()
}

// GetLastSignInAtOk returns a tuple with the LastSignInAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUser200Response) GetLastSignInAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSignInAt.Get(), o.LastSignInAt.IsSet()
}

// SetLastSignInAt sets field value
func (o *GetUser200Response) SetLastSignInAt(v float32) {
	o.LastSignInAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetUser200Response) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetUser200Response) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetUser200Response) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetUser200Response) GetUpdatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetUser200Response) GetUpdatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetUser200Response) SetUpdatedAt(v float32) {
	o.UpdatedAt = v
}

// GetProfile returns the Profile field value
func (o *GetUser200Response) GetProfile() GetJwtCustomizer200ResponseOneOfContextSampleUserProfile {
	if o == nil {
		var ret GetJwtCustomizer200ResponseOneOfContextSampleUserProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *GetUser200Response) GetProfileOk() (*GetJwtCustomizer200ResponseOneOfContextSampleUserProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *GetUser200Response) SetProfile(v GetJwtCustomizer200ResponseOneOfContextSampleUserProfile) {
	o.Profile = v
}

// GetApplicationId returns the ApplicationId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetUser200Response) GetApplicationId() string {
	if o == nil || o.ApplicationId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ApplicationId.Get()
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUser200Response) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationId.Get(), o.ApplicationId.IsSet()
}

// SetApplicationId sets field value
func (o *GetUser200Response) SetApplicationId(v string) {
	o.ApplicationId.Set(&v)
}

// GetIsSuspended returns the IsSuspended field value
func (o *GetUser200Response) GetIsSuspended() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSuspended
}

// GetIsSuspendedOk returns a tuple with the IsSuspended field value
// and a boolean to check if the value has been set.
func (o *GetUser200Response) GetIsSuspendedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSuspended, true
}

// SetIsSuspended sets field value
func (o *GetUser200Response) SetIsSuspended(v bool) {
	o.IsSuspended = v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *GetUser200Response) GetHasPassword() bool {
	if o == nil || IsNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUser200Response) GetHasPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPassword) {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *GetUser200Response) HasHasPassword() bool {
	if o != nil && !IsNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *GetUser200Response) SetHasPassword(v bool) {
	o.HasPassword = &v
}

// GetSsoIdentities returns the SsoIdentities field value if set, zero value otherwise.
func (o *GetUser200Response) GetSsoIdentities() []GetUser200ResponseSsoIdentitiesInner {
	if o == nil || IsNil(o.SsoIdentities) {
		var ret []GetUser200ResponseSsoIdentitiesInner
		return ret
	}
	return o.SsoIdentities
}

// GetSsoIdentitiesOk returns a tuple with the SsoIdentities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUser200Response) GetSsoIdentitiesOk() ([]GetUser200ResponseSsoIdentitiesInner, bool) {
	if o == nil || IsNil(o.SsoIdentities) {
		return nil, false
	}
	return o.SsoIdentities, true
}

// HasSsoIdentities returns a boolean if a field has been set.
func (o *GetUser200Response) HasSsoIdentities() bool {
	if o != nil && !IsNil(o.SsoIdentities) {
		return true
	}

	return false
}

// SetSsoIdentities gets a reference to the given []GetUser200ResponseSsoIdentitiesInner and assigns it to the SsoIdentities field.
func (o *GetUser200Response) SetSsoIdentities(v []GetUser200ResponseSsoIdentitiesInner) {
	o.SsoIdentities = v
}

func (o GetUser200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUser200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["username"] = o.Username.Get()
	toSerialize["primaryEmail"] = o.PrimaryEmail.Get()
	toSerialize["primaryPhone"] = o.PrimaryPhone.Get()
	toSerialize["name"] = o.Name.Get()
	toSerialize["avatar"] = o.Avatar.Get()
	toSerialize["customData"] = o.CustomData
	toSerialize["identities"] = o.Identities
	toSerialize["lastSignInAt"] = o.LastSignInAt.Get()
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["profile"] = o.Profile
	toSerialize["applicationId"] = o.ApplicationId.Get()
	toSerialize["isSuspended"] = o.IsSuspended
	if !IsNil(o.HasPassword) {
		toSerialize["hasPassword"] = o.HasPassword
	}
	if !IsNil(o.SsoIdentities) {
		toSerialize["ssoIdentities"] = o.SsoIdentities
	}
	return toSerialize, nil
}

func (o *GetUser200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"username",
		"primaryEmail",
		"primaryPhone",
		"name",
		"avatar",
		"customData",
		"identities",
		"lastSignInAt",
		"createdAt",
		"updatedAt",
		"profile",
		"applicationId",
		"isSuspended",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetUser200Response := _GetUser200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetUser200Response)

	if err != nil {
		return err
	}

	*o = GetUser200Response(varGetUser200Response)

	return err
}

type NullableGetUser200Response struct {
	value *GetUser200Response
	isSet bool
}

func (v NullableGetUser200Response) Get() *GetUser200Response {
	return v.value
}

func (v *NullableGetUser200Response) Set(val *GetUser200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUser200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUser200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUser200Response(val *GetUser200Response) *NullableGetUser200Response {
	return &NullableGetUser200Response{value: val, isSet: true}
}

func (v NullableGetUser200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUser200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

