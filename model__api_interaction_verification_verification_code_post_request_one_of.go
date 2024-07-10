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

// checks if the ApiInteractionVerificationVerificationCodePostRequestOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionVerificationVerificationCodePostRequestOneOf{}

// ApiInteractionVerificationVerificationCodePostRequestOneOf struct for ApiInteractionVerificationVerificationCodePostRequestOneOf
type ApiInteractionVerificationVerificationCodePostRequestOneOf struct {
	Email string `json:"email"`
}

type _ApiInteractionVerificationVerificationCodePostRequestOneOf ApiInteractionVerificationVerificationCodePostRequestOneOf

// NewApiInteractionVerificationVerificationCodePostRequestOneOf instantiates a new ApiInteractionVerificationVerificationCodePostRequestOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionVerificationVerificationCodePostRequestOneOf(email string) *ApiInteractionVerificationVerificationCodePostRequestOneOf {
	this := ApiInteractionVerificationVerificationCodePostRequestOneOf{}
	this.Email = email
	return &this
}

// NewApiInteractionVerificationVerificationCodePostRequestOneOfWithDefaults instantiates a new ApiInteractionVerificationVerificationCodePostRequestOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionVerificationVerificationCodePostRequestOneOfWithDefaults() *ApiInteractionVerificationVerificationCodePostRequestOneOf {
	this := ApiInteractionVerificationVerificationCodePostRequestOneOf{}
	return &this
}

// GetEmail returns the Email field value
func (o *ApiInteractionVerificationVerificationCodePostRequestOneOf) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionVerificationVerificationCodePostRequestOneOf) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ApiInteractionVerificationVerificationCodePostRequestOneOf) SetEmail(v string) {
	o.Email = v
}

func (o ApiInteractionVerificationVerificationCodePostRequestOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionVerificationVerificationCodePostRequestOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *ApiInteractionVerificationVerificationCodePostRequestOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varApiInteractionVerificationVerificationCodePostRequestOneOf := _ApiInteractionVerificationVerificationCodePostRequestOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiInteractionVerificationVerificationCodePostRequestOneOf)

	if err != nil {
		return err
	}

	*o = ApiInteractionVerificationVerificationCodePostRequestOneOf(varApiInteractionVerificationVerificationCodePostRequestOneOf)

	return err
}

type NullableApiInteractionVerificationVerificationCodePostRequestOneOf struct {
	value *ApiInteractionVerificationVerificationCodePostRequestOneOf
	isSet bool
}

func (v NullableApiInteractionVerificationVerificationCodePostRequestOneOf) Get() *ApiInteractionVerificationVerificationCodePostRequestOneOf {
	return v.value
}

func (v *NullableApiInteractionVerificationVerificationCodePostRequestOneOf) Set(val *ApiInteractionVerificationVerificationCodePostRequestOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionVerificationVerificationCodePostRequestOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionVerificationVerificationCodePostRequestOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionVerificationVerificationCodePostRequestOneOf(val *ApiInteractionVerificationVerificationCodePostRequestOneOf) *NullableApiInteractionVerificationVerificationCodePostRequestOneOf {
	return &NullableApiInteractionVerificationVerificationCodePostRequestOneOf{value: val, isSet: true}
}

func (v NullableApiInteractionVerificationVerificationCodePostRequestOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionVerificationVerificationCodePostRequestOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
