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

// checks if the GetSignInExp200ResponsePasswordPolicyLength type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSignInExp200ResponsePasswordPolicyLength{}

// GetSignInExp200ResponsePasswordPolicyLength struct for GetSignInExp200ResponsePasswordPolicyLength
type GetSignInExp200ResponsePasswordPolicyLength struct {
	Min float32 `json:"min"`
	Max float32 `json:"max"`
}

type _GetSignInExp200ResponsePasswordPolicyLength GetSignInExp200ResponsePasswordPolicyLength

// NewGetSignInExp200ResponsePasswordPolicyLength instantiates a new GetSignInExp200ResponsePasswordPolicyLength object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignInExp200ResponsePasswordPolicyLength(min float32, max float32) *GetSignInExp200ResponsePasswordPolicyLength {
	this := GetSignInExp200ResponsePasswordPolicyLength{}
	this.Min = min
	this.Max = max
	return &this
}

// NewGetSignInExp200ResponsePasswordPolicyLengthWithDefaults instantiates a new GetSignInExp200ResponsePasswordPolicyLength object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignInExp200ResponsePasswordPolicyLengthWithDefaults() *GetSignInExp200ResponsePasswordPolicyLength {
	this := GetSignInExp200ResponsePasswordPolicyLength{}
	var min float32
	this.Min = min
	var max float32
	this.Max = max
	return &this
}

// GetMin returns the Min field value
func (o *GetSignInExp200ResponsePasswordPolicyLength) GetMin() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponsePasswordPolicyLength) GetMinOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value
func (o *GetSignInExp200ResponsePasswordPolicyLength) SetMin(v float32) {
	o.Min = v
}

// GetMax returns the Max field value
func (o *GetSignInExp200ResponsePasswordPolicyLength) GetMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *GetSignInExp200ResponsePasswordPolicyLength) GetMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value
func (o *GetSignInExp200ResponsePasswordPolicyLength) SetMax(v float32) {
	o.Max = v
}

func (o GetSignInExp200ResponsePasswordPolicyLength) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSignInExp200ResponsePasswordPolicyLength) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["min"] = o.Min
	toSerialize["max"] = o.Max
	return toSerialize, nil
}

func (o *GetSignInExp200ResponsePasswordPolicyLength) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"min",
		"max",
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

	varGetSignInExp200ResponsePasswordPolicyLength := _GetSignInExp200ResponsePasswordPolicyLength{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSignInExp200ResponsePasswordPolicyLength)

	if err != nil {
		return err
	}

	*o = GetSignInExp200ResponsePasswordPolicyLength(varGetSignInExp200ResponsePasswordPolicyLength)

	return err
}

type NullableGetSignInExp200ResponsePasswordPolicyLength struct {
	value *GetSignInExp200ResponsePasswordPolicyLength
	isSet bool
}

func (v NullableGetSignInExp200ResponsePasswordPolicyLength) Get() *GetSignInExp200ResponsePasswordPolicyLength {
	return v.value
}

func (v *NullableGetSignInExp200ResponsePasswordPolicyLength) Set(val *GetSignInExp200ResponsePasswordPolicyLength) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignInExp200ResponsePasswordPolicyLength) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignInExp200ResponsePasswordPolicyLength) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignInExp200ResponsePasswordPolicyLength(val *GetSignInExp200ResponsePasswordPolicyLength) *NullableGetSignInExp200ResponsePasswordPolicyLength {
	return &NullableGetSignInExp200ResponsePasswordPolicyLength{value: val, isSet: true}
}

func (v NullableGetSignInExp200ResponsePasswordPolicyLength) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignInExp200ResponsePasswordPolicyLength) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

