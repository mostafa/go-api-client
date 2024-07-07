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

// checks if the ListApplicationOrganizations200ResponseInnerOrganizationRolesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListApplicationOrganizations200ResponseInnerOrganizationRolesInner{}

// ListApplicationOrganizations200ResponseInnerOrganizationRolesInner struct for ListApplicationOrganizations200ResponseInnerOrganizationRolesInner
type ListApplicationOrganizations200ResponseInnerOrganizationRolesInner struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type _ListApplicationOrganizations200ResponseInnerOrganizationRolesInner ListApplicationOrganizations200ResponseInnerOrganizationRolesInner

// NewListApplicationOrganizations200ResponseInnerOrganizationRolesInner instantiates a new ListApplicationOrganizations200ResponseInnerOrganizationRolesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApplicationOrganizations200ResponseInnerOrganizationRolesInner(id string, name string) *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner {
	this := ListApplicationOrganizations200ResponseInnerOrganizationRolesInner{}
	this.Id = id
	this.Name = name
	return &this
}

// NewListApplicationOrganizations200ResponseInnerOrganizationRolesInnerWithDefaults instantiates a new ListApplicationOrganizations200ResponseInnerOrganizationRolesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApplicationOrganizations200ResponseInnerOrganizationRolesInnerWithDefaults() *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner {
	this := ListApplicationOrganizations200ResponseInnerOrganizationRolesInner{}
	return &this
}

// GetId returns the Id field value
func (o *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) SetName(v string) {
	o.Name = v
}

func (o ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varListApplicationOrganizations200ResponseInnerOrganizationRolesInner := _ListApplicationOrganizations200ResponseInnerOrganizationRolesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListApplicationOrganizations200ResponseInnerOrganizationRolesInner)

	if err != nil {
		return err
	}

	*o = ListApplicationOrganizations200ResponseInnerOrganizationRolesInner(varListApplicationOrganizations200ResponseInnerOrganizationRolesInner)

	return err
}

type NullableListApplicationOrganizations200ResponseInnerOrganizationRolesInner struct {
	value *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner
	isSet bool
}

func (v NullableListApplicationOrganizations200ResponseInnerOrganizationRolesInner) Get() *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner {
	return v.value
}

func (v *NullableListApplicationOrganizations200ResponseInnerOrganizationRolesInner) Set(val *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplicationOrganizations200ResponseInnerOrganizationRolesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplicationOrganizations200ResponseInnerOrganizationRolesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplicationOrganizations200ResponseInnerOrganizationRolesInner(val *ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) *NullableListApplicationOrganizations200ResponseInnerOrganizationRolesInner {
	return &NullableListApplicationOrganizations200ResponseInnerOrganizationRolesInner{value: val, isSet: true}
}

func (v NullableListApplicationOrganizations200ResponseInnerOrganizationRolesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplicationOrganizations200ResponseInnerOrganizationRolesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

