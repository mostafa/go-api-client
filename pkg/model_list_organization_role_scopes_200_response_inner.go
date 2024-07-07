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

// checks if the ListOrganizationRoleScopes200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationRoleScopes200ResponseInner{}

// ListOrganizationRoleScopes200ResponseInner struct for ListOrganizationRoleScopes200ResponseInner
type ListOrganizationRoleScopes200ResponseInner struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
}

type _ListOrganizationRoleScopes200ResponseInner ListOrganizationRoleScopes200ResponseInner

// NewListOrganizationRoleScopes200ResponseInner instantiates a new ListOrganizationRoleScopes200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationRoleScopes200ResponseInner(tenantId string, id string, name string, description NullableString) *ListOrganizationRoleScopes200ResponseInner {
	this := ListOrganizationRoleScopes200ResponseInner{}
	this.TenantId = tenantId
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewListOrganizationRoleScopes200ResponseInnerWithDefaults instantiates a new ListOrganizationRoleScopes200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationRoleScopes200ResponseInnerWithDefaults() *ListOrganizationRoleScopes200ResponseInner {
	this := ListOrganizationRoleScopes200ResponseInner{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListOrganizationRoleScopes200ResponseInner) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationRoleScopes200ResponseInner) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListOrganizationRoleScopes200ResponseInner) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *ListOrganizationRoleScopes200ResponseInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationRoleScopes200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListOrganizationRoleScopes200ResponseInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListOrganizationRoleScopes200ResponseInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationRoleScopes200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListOrganizationRoleScopes200ResponseInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListOrganizationRoleScopes200ResponseInner) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListOrganizationRoleScopes200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *ListOrganizationRoleScopes200ResponseInner) SetDescription(v string) {
	o.Description.Set(&v)
}

func (o ListOrganizationRoleScopes200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationRoleScopes200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	return toSerialize, nil
}

func (o *ListOrganizationRoleScopes200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"name",
		"description",
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

	varListOrganizationRoleScopes200ResponseInner := _ListOrganizationRoleScopes200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListOrganizationRoleScopes200ResponseInner)

	if err != nil {
		return err
	}

	*o = ListOrganizationRoleScopes200ResponseInner(varListOrganizationRoleScopes200ResponseInner)

	return err
}

type NullableListOrganizationRoleScopes200ResponseInner struct {
	value *ListOrganizationRoleScopes200ResponseInner
	isSet bool
}

func (v NullableListOrganizationRoleScopes200ResponseInner) Get() *ListOrganizationRoleScopes200ResponseInner {
	return v.value
}

func (v *NullableListOrganizationRoleScopes200ResponseInner) Set(val *ListOrganizationRoleScopes200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationRoleScopes200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationRoleScopes200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationRoleScopes200ResponseInner(val *ListOrganizationRoleScopes200ResponseInner) *NullableListOrganizationRoleScopes200ResponseInner {
	return &NullableListOrganizationRoleScopes200ResponseInner{value: val, isSet: true}
}

func (v NullableListOrganizationRoleScopes200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationRoleScopes200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

