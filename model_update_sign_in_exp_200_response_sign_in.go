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

// checks if the UpdateSignInExp200ResponseSignIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSignInExp200ResponseSignIn{}

// UpdateSignInExp200ResponseSignIn struct for UpdateSignInExp200ResponseSignIn
type UpdateSignInExp200ResponseSignIn struct {
	Methods []GetSignInExp200ResponseSignInMethodsInner `json:"methods"`
}

type _UpdateSignInExp200ResponseSignIn UpdateSignInExp200ResponseSignIn

// NewUpdateSignInExp200ResponseSignIn instantiates a new UpdateSignInExp200ResponseSignIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSignInExp200ResponseSignIn(methods []GetSignInExp200ResponseSignInMethodsInner) *UpdateSignInExp200ResponseSignIn {
	this := UpdateSignInExp200ResponseSignIn{}
	this.Methods = methods
	return &this
}

// NewUpdateSignInExp200ResponseSignInWithDefaults instantiates a new UpdateSignInExp200ResponseSignIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSignInExp200ResponseSignInWithDefaults() *UpdateSignInExp200ResponseSignIn {
	this := UpdateSignInExp200ResponseSignIn{}
	return &this
}

// GetMethods returns the Methods field value
func (o *UpdateSignInExp200ResponseSignIn) GetMethods() []GetSignInExp200ResponseSignInMethodsInner {
	if o == nil {
		var ret []GetSignInExp200ResponseSignInMethodsInner
		return ret
	}

	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExp200ResponseSignIn) GetMethodsOk() ([]GetSignInExp200ResponseSignInMethodsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Methods, true
}

// SetMethods sets field value
func (o *UpdateSignInExp200ResponseSignIn) SetMethods(v []GetSignInExp200ResponseSignInMethodsInner) {
	o.Methods = v
}

func (o UpdateSignInExp200ResponseSignIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSignInExp200ResponseSignIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["methods"] = o.Methods
	return toSerialize, nil
}

func (o *UpdateSignInExp200ResponseSignIn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"methods",
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

	varUpdateSignInExp200ResponseSignIn := _UpdateSignInExp200ResponseSignIn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateSignInExp200ResponseSignIn)

	if err != nil {
		return err
	}

	*o = UpdateSignInExp200ResponseSignIn(varUpdateSignInExp200ResponseSignIn)

	return err
}

type NullableUpdateSignInExp200ResponseSignIn struct {
	value *UpdateSignInExp200ResponseSignIn
	isSet bool
}

func (v NullableUpdateSignInExp200ResponseSignIn) Get() *UpdateSignInExp200ResponseSignIn {
	return v.value
}

func (v *NullableUpdateSignInExp200ResponseSignIn) Set(val *UpdateSignInExp200ResponseSignIn) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSignInExp200ResponseSignIn) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSignInExp200ResponseSignIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSignInExp200ResponseSignIn(val *UpdateSignInExp200ResponseSignIn) *NullableUpdateSignInExp200ResponseSignIn {
	return &NullableUpdateSignInExp200ResponseSignIn{value: val, isSet: true}
}

func (v NullableUpdateSignInExp200ResponseSignIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSignInExp200ResponseSignIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
