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

// checks if the ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults{}

// ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults struct for ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults
type ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults struct {
	Appid *bool `json:"appid,omitempty"`
	CrepProps *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps `json:"crepProps,omitempty"`
	HmacCreateSecret *bool `json:"hmacCreateSecret,omitempty"`
}

// NewApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults instantiates a new ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults() *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults {
	this := ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults{}
	return &this
}

// NewApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsWithDefaults instantiates a new ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsWithDefaults() *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults {
	this := ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults{}
	return &this
}

// GetAppid returns the Appid field value if set, zero value otherwise.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) GetAppid() bool {
	if o == nil || IsNil(o.Appid) {
		var ret bool
		return ret
	}
	return *o.Appid
}

// GetAppidOk returns a tuple with the Appid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) GetAppidOk() (*bool, bool) {
	if o == nil || IsNil(o.Appid) {
		return nil, false
	}
	return o.Appid, true
}

// HasAppid returns a boolean if a field has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) HasAppid() bool {
	if o != nil && !IsNil(o.Appid) {
		return true
	}

	return false
}

// SetAppid gets a reference to the given bool and assigns it to the Appid field.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) SetAppid(v bool) {
	o.Appid = &v
}

// GetCrepProps returns the CrepProps field value if set, zero value otherwise.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) GetCrepProps() ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps {
	if o == nil || IsNil(o.CrepProps) {
		var ret ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps
		return ret
	}
	return *o.CrepProps
}

// GetCrepPropsOk returns a tuple with the CrepProps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) GetCrepPropsOk() (*ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps, bool) {
	if o == nil || IsNil(o.CrepProps) {
		return nil, false
	}
	return o.CrepProps, true
}

// HasCrepProps returns a boolean if a field has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) HasCrepProps() bool {
	if o != nil && !IsNil(o.CrepProps) {
		return true
	}

	return false
}

// SetCrepProps gets a reference to the given ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps and assigns it to the CrepProps field.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) SetCrepProps(v ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResultsCrepProps) {
	o.CrepProps = &v
}

// GetHmacCreateSecret returns the HmacCreateSecret field value if set, zero value otherwise.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) GetHmacCreateSecret() bool {
	if o == nil || IsNil(o.HmacCreateSecret) {
		var ret bool
		return ret
	}
	return *o.HmacCreateSecret
}

// GetHmacCreateSecretOk returns a tuple with the HmacCreateSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) GetHmacCreateSecretOk() (*bool, bool) {
	if o == nil || IsNil(o.HmacCreateSecret) {
		return nil, false
	}
	return o.HmacCreateSecret, true
}

// HasHmacCreateSecret returns a boolean if a field has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) HasHmacCreateSecret() bool {
	if o != nil && !IsNil(o.HmacCreateSecret) {
		return true
	}

	return false
}

// SetHmacCreateSecret gets a reference to the given bool and assigns it to the HmacCreateSecret field.
func (o *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) SetHmacCreateSecret(v bool) {
	o.HmacCreateSecret = &v
}

func (o ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Appid) {
		toSerialize["appid"] = o.Appid
	}
	if !IsNil(o.CrepProps) {
		toSerialize["crepProps"] = o.CrepProps
	}
	if !IsNil(o.HmacCreateSecret) {
		toSerialize["hmacCreateSecret"] = o.HmacCreateSecret
	}
	return toSerialize, nil
}

type NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults struct {
	value *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults
	isSet bool
}

func (v NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) Get() *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults {
	return v.value
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) Set(val *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults(val *ApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) *NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults {
	return &NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults{value: val, isSet: true}
}

func (v NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf1ClientExtensionResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
