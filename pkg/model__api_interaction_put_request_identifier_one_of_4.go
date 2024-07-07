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

// checks if the ApiInteractionPutRequestIdentifierOneOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionPutRequestIdentifierOneOf4{}

// ApiInteractionPutRequestIdentifierOneOf4 struct for ApiInteractionPutRequestIdentifierOneOf4
type ApiInteractionPutRequestIdentifierOneOf4 struct {
	Phone string `json:"phone"`
	VerificationCode string `json:"verificationCode"`
}

type _ApiInteractionPutRequestIdentifierOneOf4 ApiInteractionPutRequestIdentifierOneOf4

// NewApiInteractionPutRequestIdentifierOneOf4 instantiates a new ApiInteractionPutRequestIdentifierOneOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionPutRequestIdentifierOneOf4(phone string, verificationCode string) *ApiInteractionPutRequestIdentifierOneOf4 {
	this := ApiInteractionPutRequestIdentifierOneOf4{}
	this.Phone = phone
	this.VerificationCode = verificationCode
	return &this
}

// NewApiInteractionPutRequestIdentifierOneOf4WithDefaults instantiates a new ApiInteractionPutRequestIdentifierOneOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionPutRequestIdentifierOneOf4WithDefaults() *ApiInteractionPutRequestIdentifierOneOf4 {
	this := ApiInteractionPutRequestIdentifierOneOf4{}
	return &this
}

// GetPhone returns the Phone field value
func (o *ApiInteractionPutRequestIdentifierOneOf4) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionPutRequestIdentifierOneOf4) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *ApiInteractionPutRequestIdentifierOneOf4) SetPhone(v string) {
	o.Phone = v
}

// GetVerificationCode returns the VerificationCode field value
func (o *ApiInteractionPutRequestIdentifierOneOf4) GetVerificationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationCode
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionPutRequestIdentifierOneOf4) GetVerificationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationCode, true
}

// SetVerificationCode sets field value
func (o *ApiInteractionPutRequestIdentifierOneOf4) SetVerificationCode(v string) {
	o.VerificationCode = v
}

func (o ApiInteractionPutRequestIdentifierOneOf4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionPutRequestIdentifierOneOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phone"] = o.Phone
	toSerialize["verificationCode"] = o.VerificationCode
	return toSerialize, nil
}

func (o *ApiInteractionPutRequestIdentifierOneOf4) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phone",
		"verificationCode",
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

	varApiInteractionPutRequestIdentifierOneOf4 := _ApiInteractionPutRequestIdentifierOneOf4{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiInteractionPutRequestIdentifierOneOf4)

	if err != nil {
		return err
	}

	*o = ApiInteractionPutRequestIdentifierOneOf4(varApiInteractionPutRequestIdentifierOneOf4)

	return err
}

type NullableApiInteractionPutRequestIdentifierOneOf4 struct {
	value *ApiInteractionPutRequestIdentifierOneOf4
	isSet bool
}

func (v NullableApiInteractionPutRequestIdentifierOneOf4) Get() *ApiInteractionPutRequestIdentifierOneOf4 {
	return v.value
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf4) Set(val *ApiInteractionPutRequestIdentifierOneOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionPutRequestIdentifierOneOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionPutRequestIdentifierOneOf4(val *ApiInteractionPutRequestIdentifierOneOf4) *NullableApiInteractionPutRequestIdentifierOneOf4 {
	return &NullableApiInteractionPutRequestIdentifierOneOf4{value: val, isSet: true}
}

func (v NullableApiInteractionPutRequestIdentifierOneOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

