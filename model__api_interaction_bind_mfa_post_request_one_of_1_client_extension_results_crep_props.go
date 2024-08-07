/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"encoding/json"
)

// checks if the ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps{}

// ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps struct for ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps
type ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps struct {
	Rk *bool `json:"rk,omitempty"`
}

// NewApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps instantiates a new ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps() *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps {
	this := ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps{}
	return &this
}

// NewApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepPropsWithDefaults instantiates a new ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepPropsWithDefaults() *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps {
	this := ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps{}
	return &this
}

// GetRk returns the Rk field value if set, zero value otherwise.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) GetRk() bool {
	if o == nil || IsNil(o.Rk) {
		var ret bool
		return ret
	}
	return *o.Rk
}

// GetRkOk returns a tuple with the Rk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) GetRkOk() (*bool, bool) {
	if o == nil || IsNil(o.Rk) {
		return nil, false
	}
	return o.Rk, true
}

// HasRk returns a boolean if a field has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) HasRk() bool {
	if o != nil && !IsNil(o.Rk) {
		return true
	}

	return false
}

// SetRk gets a reference to the given bool and assigns it to the Rk field.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) SetRk(v bool) {
	o.Rk = &v
}

func (o ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rk) {
		toSerialize["rk"] = o.Rk
	}
	return toSerialize, nil
}

type NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps struct {
	value *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps
	isSet bool
}

func (v NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) Get() *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps {
	return v.value
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) Set(val *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps(val *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) *NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps {
	return &NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps{value: val, isSet: true}
}

func (v NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

