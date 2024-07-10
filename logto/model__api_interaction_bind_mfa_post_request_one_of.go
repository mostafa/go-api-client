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

// checks if the ApiInteractionBindMfaPostRequestOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionBindMfaPostRequestOneOf{}

// ApiInteractionBindMfaPostRequestOneOf struct for ApiInteractionBindMfaPostRequestOneOf
type ApiInteractionBindMfaPostRequestOneOf struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type _ApiInteractionBindMfaPostRequestOneOf ApiInteractionBindMfaPostRequestOneOf

// NewApiInteractionBindMfaPostRequestOneOf instantiates a new ApiInteractionBindMfaPostRequestOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionBindMfaPostRequestOneOf(type_ string, code string) *ApiInteractionBindMfaPostRequestOneOf {
	this := ApiInteractionBindMfaPostRequestOneOf{}
	this.Type = type_
	this.Code = code
	return &this
}

// NewApiInteractionBindMfaPostRequestOneOfWithDefaults instantiates a new ApiInteractionBindMfaPostRequestOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionBindMfaPostRequestOneOfWithDefaults() *ApiInteractionBindMfaPostRequestOneOf {
	this := ApiInteractionBindMfaPostRequestOneOf{}
	return &this
}

// GetType returns the Type field value
func (o *ApiInteractionBindMfaPostRequestOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiInteractionBindMfaPostRequestOneOf) SetType(v string) {
	o.Type = v
}

// GetCode returns the Code field value
func (o *ApiInteractionBindMfaPostRequestOneOf) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ApiInteractionBindMfaPostRequestOneOf) SetCode(v string) {
	o.Code = v
}

func (o ApiInteractionBindMfaPostRequestOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionBindMfaPostRequestOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

func (o *ApiInteractionBindMfaPostRequestOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"code",
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

	varApiInteractionBindMfaPostRequestOneOf := _ApiInteractionBindMfaPostRequestOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiInteractionBindMfaPostRequestOneOf)

	if err != nil {
		return err
	}

	*o = ApiInteractionBindMfaPostRequestOneOf(varApiInteractionBindMfaPostRequestOneOf)

	return err
}

type NullableApiInteractionBindMfaPostRequestOneOf struct {
	value *ApiInteractionBindMfaPostRequestOneOf
	isSet bool
}

func (v NullableApiInteractionBindMfaPostRequestOneOf) Get() *ApiInteractionBindMfaPostRequestOneOf {
	return v.value
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf) Set(val *ApiInteractionBindMfaPostRequestOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionBindMfaPostRequestOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionBindMfaPostRequestOneOf(val *ApiInteractionBindMfaPostRequestOneOf) *NullableApiInteractionBindMfaPostRequestOneOf {
	return &NullableApiInteractionBindMfaPostRequestOneOf{value: val, isSet: true}
}

func (v NullableApiInteractionBindMfaPostRequestOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
