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

// checks if the UpdateSignInExp200ResponseSignUp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSignInExp200ResponseSignUp{}

// UpdateSignInExp200ResponseSignUp struct for UpdateSignInExp200ResponseSignUp
type UpdateSignInExp200ResponseSignUp struct {
	Identifiers []string `json:"identifiers"`
	Password bool `json:"password"`
	Verify bool `json:"verify"`
}

type _UpdateSignInExp200ResponseSignUp UpdateSignInExp200ResponseSignUp

// NewUpdateSignInExp200ResponseSignUp instantiates a new UpdateSignInExp200ResponseSignUp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSignInExp200ResponseSignUp(identifiers []string, password bool, verify bool) *UpdateSignInExp200ResponseSignUp {
	this := UpdateSignInExp200ResponseSignUp{}
	this.Identifiers = identifiers
	this.Password = password
	this.Verify = verify
	return &this
}

// NewUpdateSignInExp200ResponseSignUpWithDefaults instantiates a new UpdateSignInExp200ResponseSignUp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSignInExp200ResponseSignUpWithDefaults() *UpdateSignInExp200ResponseSignUp {
	this := UpdateSignInExp200ResponseSignUp{}
	return &this
}

// GetIdentifiers returns the Identifiers field value
func (o *UpdateSignInExp200ResponseSignUp) GetIdentifiers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExp200ResponseSignUp) GetIdentifiersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifiers, true
}

// SetIdentifiers sets field value
func (o *UpdateSignInExp200ResponseSignUp) SetIdentifiers(v []string) {
	o.Identifiers = v
}

// GetPassword returns the Password field value
func (o *UpdateSignInExp200ResponseSignUp) GetPassword() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExp200ResponseSignUp) GetPasswordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UpdateSignInExp200ResponseSignUp) SetPassword(v bool) {
	o.Password = v
}

// GetVerify returns the Verify field value
func (o *UpdateSignInExp200ResponseSignUp) GetVerify() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExp200ResponseSignUp) GetVerifyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verify, true
}

// SetVerify sets field value
func (o *UpdateSignInExp200ResponseSignUp) SetVerify(v bool) {
	o.Verify = v
}

func (o UpdateSignInExp200ResponseSignUp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSignInExp200ResponseSignUp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifiers"] = o.Identifiers
	toSerialize["password"] = o.Password
	toSerialize["verify"] = o.Verify
	return toSerialize, nil
}

func (o *UpdateSignInExp200ResponseSignUp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifiers",
		"password",
		"verify",
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

	varUpdateSignInExp200ResponseSignUp := _UpdateSignInExp200ResponseSignUp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateSignInExp200ResponseSignUp)

	if err != nil {
		return err
	}

	*o = UpdateSignInExp200ResponseSignUp(varUpdateSignInExp200ResponseSignUp)

	return err
}

type NullableUpdateSignInExp200ResponseSignUp struct {
	value *UpdateSignInExp200ResponseSignUp
	isSet bool
}

func (v NullableUpdateSignInExp200ResponseSignUp) Get() *UpdateSignInExp200ResponseSignUp {
	return v.value
}

func (v *NullableUpdateSignInExp200ResponseSignUp) Set(val *UpdateSignInExp200ResponseSignUp) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSignInExp200ResponseSignUp) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSignInExp200ResponseSignUp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSignInExp200ResponseSignUp(val *UpdateSignInExp200ResponseSignUp) *NullableUpdateSignInExp200ResponseSignUp {
	return &NullableUpdateSignInExp200ResponseSignUp{value: val, isSet: true}
}

func (v NullableUpdateSignInExp200ResponseSignUp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSignInExp200ResponseSignUp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

