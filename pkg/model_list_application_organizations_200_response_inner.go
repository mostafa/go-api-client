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

// checks if the ListApplicationOrganizations200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListApplicationOrganizations200ResponseInner{}

// ListApplicationOrganizations200ResponseInner struct for ListApplicationOrganizations200ResponseInner
type ListApplicationOrganizations200ResponseInner struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
	// arbitrary
	CustomData map[string]interface{} `json:"customData"`
	IsMfaRequired bool `json:"isMfaRequired"`
	CreatedAt float32 `json:"createdAt"`
	OrganizationRoles []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner `json:"organizationRoles"`
}

type _ListApplicationOrganizations200ResponseInner ListApplicationOrganizations200ResponseInner

// NewListApplicationOrganizations200ResponseInner instantiates a new ListApplicationOrganizations200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApplicationOrganizations200ResponseInner(tenantId string, id string, name string, description NullableString, customData map[string]interface{}, isMfaRequired bool, createdAt float32, organizationRoles []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) *ListApplicationOrganizations200ResponseInner {
	this := ListApplicationOrganizations200ResponseInner{}
	this.TenantId = tenantId
	this.Id = id
	this.Name = name
	this.Description = description
	this.CustomData = customData
	this.IsMfaRequired = isMfaRequired
	this.CreatedAt = createdAt
	this.OrganizationRoles = organizationRoles
	return &this
}

// NewListApplicationOrganizations200ResponseInnerWithDefaults instantiates a new ListApplicationOrganizations200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApplicationOrganizations200ResponseInnerWithDefaults() *ListApplicationOrganizations200ResponseInner {
	this := ListApplicationOrganizations200ResponseInner{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListApplicationOrganizations200ResponseInner) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListApplicationOrganizations200ResponseInner) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListApplicationOrganizations200ResponseInner) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *ListApplicationOrganizations200ResponseInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListApplicationOrganizations200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListApplicationOrganizations200ResponseInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListApplicationOrganizations200ResponseInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListApplicationOrganizations200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListApplicationOrganizations200ResponseInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListApplicationOrganizations200ResponseInner) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListApplicationOrganizations200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *ListApplicationOrganizations200ResponseInner) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetCustomData returns the CustomData field value
func (o *ListApplicationOrganizations200ResponseInner) GetCustomData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value
// and a boolean to check if the value has been set.
func (o *ListApplicationOrganizations200ResponseInner) GetCustomDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.CustomData, true
}

// SetCustomData sets field value
func (o *ListApplicationOrganizations200ResponseInner) SetCustomData(v map[string]interface{}) {
	o.CustomData = v
}

// GetIsMfaRequired returns the IsMfaRequired field value
func (o *ListApplicationOrganizations200ResponseInner) GetIsMfaRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMfaRequired
}

// GetIsMfaRequiredOk returns a tuple with the IsMfaRequired field value
// and a boolean to check if the value has been set.
func (o *ListApplicationOrganizations200ResponseInner) GetIsMfaRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMfaRequired, true
}

// SetIsMfaRequired sets field value
func (o *ListApplicationOrganizations200ResponseInner) SetIsMfaRequired(v bool) {
	o.IsMfaRequired = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ListApplicationOrganizations200ResponseInner) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ListApplicationOrganizations200ResponseInner) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ListApplicationOrganizations200ResponseInner) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetOrganizationRoles returns the OrganizationRoles field value
func (o *ListApplicationOrganizations200ResponseInner) GetOrganizationRoles() []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner {
	if o == nil {
		var ret []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner
		return ret
	}

	return o.OrganizationRoles
}

// GetOrganizationRolesOk returns a tuple with the OrganizationRoles field value
// and a boolean to check if the value has been set.
func (o *ListApplicationOrganizations200ResponseInner) GetOrganizationRolesOk() ([]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationRoles, true
}

// SetOrganizationRoles sets field value
func (o *ListApplicationOrganizations200ResponseInner) SetOrganizationRoles(v []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) {
	o.OrganizationRoles = v
}

func (o ListApplicationOrganizations200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListApplicationOrganizations200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["customData"] = o.CustomData
	toSerialize["isMfaRequired"] = o.IsMfaRequired
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["organizationRoles"] = o.OrganizationRoles
	return toSerialize, nil
}

func (o *ListApplicationOrganizations200ResponseInner) UnmarshalJSON(data []byte) (err error) {
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
		"organizationRoles",
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

	varListApplicationOrganizations200ResponseInner := _ListApplicationOrganizations200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListApplicationOrganizations200ResponseInner)

	if err != nil {
		return err
	}

	*o = ListApplicationOrganizations200ResponseInner(varListApplicationOrganizations200ResponseInner)

	return err
}

type NullableListApplicationOrganizations200ResponseInner struct {
	value *ListApplicationOrganizations200ResponseInner
	isSet bool
}

func (v NullableListApplicationOrganizations200ResponseInner) Get() *ListApplicationOrganizations200ResponseInner {
	return v.value
}

func (v *NullableListApplicationOrganizations200ResponseInner) Set(val *ListApplicationOrganizations200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplicationOrganizations200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplicationOrganizations200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplicationOrganizations200ResponseInner(val *ListApplicationOrganizations200ResponseInner) *NullableListApplicationOrganizations200ResponseInner {
	return &NullableListApplicationOrganizations200ResponseInner{value: val, isSet: true}
}

func (v NullableListApplicationOrganizations200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplicationOrganizations200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

