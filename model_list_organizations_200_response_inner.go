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

// checks if the ListOrganizations200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizations200ResponseInner{}

// ListOrganizations200ResponseInner struct for ListOrganizations200ResponseInner
type ListOrganizations200ResponseInner struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
	// arbitrary
	CustomData map[string]interface{} `json:"customData"`
	IsMfaRequired bool `json:"isMfaRequired"`
	CreatedAt float32 `json:"createdAt"`
	UsersCount *float32 `json:"usersCount,omitempty"`
	FeaturedUsers []ListRoles200ResponseInnerFeaturedUsersInner `json:"featuredUsers,omitempty"`
}

type _ListOrganizations200ResponseInner ListOrganizations200ResponseInner

// NewListOrganizations200ResponseInner instantiates a new ListOrganizations200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizations200ResponseInner(tenantId string, id string, name string, description NullableString, customData map[string]interface{}, isMfaRequired bool, createdAt float32) *ListOrganizations200ResponseInner {
	this := ListOrganizations200ResponseInner{}
	this.TenantId = tenantId
	this.Id = id
	this.Name = name
	this.Description = description
	this.CustomData = customData
	this.IsMfaRequired = isMfaRequired
	this.CreatedAt = createdAt
	return &this
}

// NewListOrganizations200ResponseInnerWithDefaults instantiates a new ListOrganizations200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizations200ResponseInnerWithDefaults() *ListOrganizations200ResponseInner {
	this := ListOrganizations200ResponseInner{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListOrganizations200ResponseInner) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListOrganizations200ResponseInner) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListOrganizations200ResponseInner) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *ListOrganizations200ResponseInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListOrganizations200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListOrganizations200ResponseInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListOrganizations200ResponseInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListOrganizations200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListOrganizations200ResponseInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListOrganizations200ResponseInner) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListOrganizations200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *ListOrganizations200ResponseInner) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetCustomData returns the CustomData field value
func (o *ListOrganizations200ResponseInner) GetCustomData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value
// and a boolean to check if the value has been set.
func (o *ListOrganizations200ResponseInner) GetCustomDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.CustomData, true
}

// SetCustomData sets field value
func (o *ListOrganizations200ResponseInner) SetCustomData(v map[string]interface{}) {
	o.CustomData = v
}

// GetIsMfaRequired returns the IsMfaRequired field value
func (o *ListOrganizations200ResponseInner) GetIsMfaRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMfaRequired
}

// GetIsMfaRequiredOk returns a tuple with the IsMfaRequired field value
// and a boolean to check if the value has been set.
func (o *ListOrganizations200ResponseInner) GetIsMfaRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMfaRequired, true
}

// SetIsMfaRequired sets field value
func (o *ListOrganizations200ResponseInner) SetIsMfaRequired(v bool) {
	o.IsMfaRequired = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ListOrganizations200ResponseInner) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ListOrganizations200ResponseInner) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ListOrganizations200ResponseInner) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetUsersCount returns the UsersCount field value if set, zero value otherwise.
func (o *ListOrganizations200ResponseInner) GetUsersCount() float32 {
	if o == nil || IsNil(o.UsersCount) {
		var ret float32
		return ret
	}
	return *o.UsersCount
}

// GetUsersCountOk returns a tuple with the UsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizations200ResponseInner) GetUsersCountOk() (*float32, bool) {
	if o == nil || IsNil(o.UsersCount) {
		return nil, false
	}
	return o.UsersCount, true
}

// HasUsersCount returns a boolean if a field has been set.
func (o *ListOrganizations200ResponseInner) HasUsersCount() bool {
	if o != nil && !IsNil(o.UsersCount) {
		return true
	}

	return false
}

// SetUsersCount gets a reference to the given float32 and assigns it to the UsersCount field.
func (o *ListOrganizations200ResponseInner) SetUsersCount(v float32) {
	o.UsersCount = &v
}

// GetFeaturedUsers returns the FeaturedUsers field value if set, zero value otherwise.
func (o *ListOrganizations200ResponseInner) GetFeaturedUsers() []ListRoles200ResponseInnerFeaturedUsersInner {
	if o == nil || IsNil(o.FeaturedUsers) {
		var ret []ListRoles200ResponseInnerFeaturedUsersInner
		return ret
	}
	return o.FeaturedUsers
}

// GetFeaturedUsersOk returns a tuple with the FeaturedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizations200ResponseInner) GetFeaturedUsersOk() ([]ListRoles200ResponseInnerFeaturedUsersInner, bool) {
	if o == nil || IsNil(o.FeaturedUsers) {
		return nil, false
	}
	return o.FeaturedUsers, true
}

// HasFeaturedUsers returns a boolean if a field has been set.
func (o *ListOrganizations200ResponseInner) HasFeaturedUsers() bool {
	if o != nil && !IsNil(o.FeaturedUsers) {
		return true
	}

	return false
}

// SetFeaturedUsers gets a reference to the given []ListRoles200ResponseInnerFeaturedUsersInner and assigns it to the FeaturedUsers field.
func (o *ListOrganizations200ResponseInner) SetFeaturedUsers(v []ListRoles200ResponseInnerFeaturedUsersInner) {
	o.FeaturedUsers = v
}

func (o ListOrganizations200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizations200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["customData"] = o.CustomData
	toSerialize["isMfaRequired"] = o.IsMfaRequired
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.UsersCount) {
		toSerialize["usersCount"] = o.UsersCount
	}
	if !IsNil(o.FeaturedUsers) {
		toSerialize["featuredUsers"] = o.FeaturedUsers
	}
	return toSerialize, nil
}

func (o *ListOrganizations200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"name",
		"description",
		"customData",
		"isMfaRequired",
		"createdAt",
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

	varListOrganizations200ResponseInner := _ListOrganizations200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListOrganizations200ResponseInner)

	if err != nil {
		return err
	}

	*o = ListOrganizations200ResponseInner(varListOrganizations200ResponseInner)

	return err
}

type NullableListOrganizations200ResponseInner struct {
	value *ListOrganizations200ResponseInner
	isSet bool
}

func (v NullableListOrganizations200ResponseInner) Get() *ListOrganizations200ResponseInner {
	return v.value
}

func (v *NullableListOrganizations200ResponseInner) Set(val *ListOrganizations200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizations200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizations200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizations200ResponseInner(val *ListOrganizations200ResponseInner) *NullableListOrganizations200ResponseInner {
	return &NullableListOrganizations200ResponseInner{value: val, isSet: true}
}

func (v NullableListOrganizations200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizations200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
