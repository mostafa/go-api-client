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

// checks if the ListApplicationUserConsentScopes200ResponseOrganizationScopesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListApplicationUserConsentScopes200ResponseOrganizationScopesInner{}

// ListApplicationUserConsentScopes200ResponseOrganizationScopesInner struct for ListApplicationUserConsentScopes200ResponseOrganizationScopesInner
type ListApplicationUserConsentScopes200ResponseOrganizationScopesInner struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
}

type _ListApplicationUserConsentScopes200ResponseOrganizationScopesInner ListApplicationUserConsentScopes200ResponseOrganizationScopesInner

// NewListApplicationUserConsentScopes200ResponseOrganizationScopesInner instantiates a new ListApplicationUserConsentScopes200ResponseOrganizationScopesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApplicationUserConsentScopes200ResponseOrganizationScopesInner(id string, name string, description NullableString) *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner {
	this := ListApplicationUserConsentScopes200ResponseOrganizationScopesInner{}
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewListApplicationUserConsentScopes200ResponseOrganizationScopesInnerWithDefaults instantiates a new ListApplicationUserConsentScopes200ResponseOrganizationScopesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApplicationUserConsentScopes200ResponseOrganizationScopesInnerWithDefaults() *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner {
	this := ListApplicationUserConsentScopes200ResponseOrganizationScopesInner{}
	return &this
}

// GetId returns the Id field value
func (o *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) SetDescription(v string) {
	o.Description.Set(&v)
}

func (o ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	return toSerialize, nil
}

func (o *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varListApplicationUserConsentScopes200ResponseOrganizationScopesInner := _ListApplicationUserConsentScopes200ResponseOrganizationScopesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListApplicationUserConsentScopes200ResponseOrganizationScopesInner)

	if err != nil {
		return err
	}

	*o = ListApplicationUserConsentScopes200ResponseOrganizationScopesInner(varListApplicationUserConsentScopes200ResponseOrganizationScopesInner)

	return err
}

type NullableListApplicationUserConsentScopes200ResponseOrganizationScopesInner struct {
	value *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner
	isSet bool
}

func (v NullableListApplicationUserConsentScopes200ResponseOrganizationScopesInner) Get() *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner {
	return v.value
}

func (v *NullableListApplicationUserConsentScopes200ResponseOrganizationScopesInner) Set(val *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplicationUserConsentScopes200ResponseOrganizationScopesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplicationUserConsentScopes200ResponseOrganizationScopesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplicationUserConsentScopes200ResponseOrganizationScopesInner(val *ListApplicationUserConsentScopes200ResponseOrganizationScopesInner) *NullableListApplicationUserConsentScopes200ResponseOrganizationScopesInner {
	return &NullableListApplicationUserConsentScopes200ResponseOrganizationScopesInner{value: val, isSet: true}
}

func (v NullableListApplicationUserConsentScopes200ResponseOrganizationScopesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplicationUserConsentScopes200ResponseOrganizationScopesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

