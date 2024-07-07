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

// checks if the ApiInteractionPutRequestIdentifierOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionPutRequestIdentifierOneOf{}

// ApiInteractionPutRequestIdentifierOneOf struct for ApiInteractionPutRequestIdentifierOneOf
type ApiInteractionPutRequestIdentifierOneOf struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type _ApiInteractionPutRequestIdentifierOneOf ApiInteractionPutRequestIdentifierOneOf

// NewApiInteractionPutRequestIdentifierOneOf instantiates a new ApiInteractionPutRequestIdentifierOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionPutRequestIdentifierOneOf(username string, password string) *ApiInteractionPutRequestIdentifierOneOf {
	this := ApiInteractionPutRequestIdentifierOneOf{}
	this.Username = username
	this.Password = password
	return &this
}

// NewApiInteractionPutRequestIdentifierOneOfWithDefaults instantiates a new ApiInteractionPutRequestIdentifierOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionPutRequestIdentifierOneOfWithDefaults() *ApiInteractionPutRequestIdentifierOneOf {
	this := ApiInteractionPutRequestIdentifierOneOf{}
	return &this
}

// GetUsername returns the Username field value
func (o *ApiInteractionPutRequestIdentifierOneOf) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionPutRequestIdentifierOneOf) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ApiInteractionPutRequestIdentifierOneOf) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *ApiInteractionPutRequestIdentifierOneOf) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionPutRequestIdentifierOneOf) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ApiInteractionPutRequestIdentifierOneOf) SetPassword(v string) {
	o.Password = v
}

func (o ApiInteractionPutRequestIdentifierOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionPutRequestIdentifierOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *ApiInteractionPutRequestIdentifierOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
		"password",
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

	varApiInteractionPutRequestIdentifierOneOf := _ApiInteractionPutRequestIdentifierOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiInteractionPutRequestIdentifierOneOf)

	if err != nil {
		return err
	}

	*o = ApiInteractionPutRequestIdentifierOneOf(varApiInteractionPutRequestIdentifierOneOf)

	return err
}

type NullableApiInteractionPutRequestIdentifierOneOf struct {
	value *ApiInteractionPutRequestIdentifierOneOf
	isSet bool
}

func (v NullableApiInteractionPutRequestIdentifierOneOf) Get() *ApiInteractionPutRequestIdentifierOneOf {
	return v.value
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf) Set(val *ApiInteractionPutRequestIdentifierOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionPutRequestIdentifierOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionPutRequestIdentifierOneOf(val *ApiInteractionPutRequestIdentifierOneOf) *NullableApiInteractionPutRequestIdentifierOneOf {
	return &NullableApiInteractionPutRequestIdentifierOneOf{value: val, isSet: true}
}

func (v NullableApiInteractionPutRequestIdentifierOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

