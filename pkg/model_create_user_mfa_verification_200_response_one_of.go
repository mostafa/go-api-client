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

// checks if the CreateUserMfaVerification200ResponseOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserMfaVerification200ResponseOneOf{}

// CreateUserMfaVerification200ResponseOneOf struct for CreateUserMfaVerification200ResponseOneOf
type CreateUserMfaVerification200ResponseOneOf struct {
	Type string `json:"type"`
	Secret string `json:"secret"`
	SecretQrCode string `json:"secretQrCode"`
}

type _CreateUserMfaVerification200ResponseOneOf CreateUserMfaVerification200ResponseOneOf

// NewCreateUserMfaVerification200ResponseOneOf instantiates a new CreateUserMfaVerification200ResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserMfaVerification200ResponseOneOf(type_ string, secret string, secretQrCode string) *CreateUserMfaVerification200ResponseOneOf {
	this := CreateUserMfaVerification200ResponseOneOf{}
	this.Type = type_
	this.Secret = secret
	this.SecretQrCode = secretQrCode
	return &this
}

// NewCreateUserMfaVerification200ResponseOneOfWithDefaults instantiates a new CreateUserMfaVerification200ResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserMfaVerification200ResponseOneOfWithDefaults() *CreateUserMfaVerification200ResponseOneOf {
	this := CreateUserMfaVerification200ResponseOneOf{}
	return &this
}

// GetType returns the Type field value
func (o *CreateUserMfaVerification200ResponseOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateUserMfaVerification200ResponseOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateUserMfaVerification200ResponseOneOf) SetType(v string) {
	o.Type = v
}

// GetSecret returns the Secret field value
func (o *CreateUserMfaVerification200ResponseOneOf) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CreateUserMfaVerification200ResponseOneOf) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CreateUserMfaVerification200ResponseOneOf) SetSecret(v string) {
	o.Secret = v
}

// GetSecretQrCode returns the SecretQrCode field value
func (o *CreateUserMfaVerification200ResponseOneOf) GetSecretQrCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretQrCode
}

// GetSecretQrCodeOk returns a tuple with the SecretQrCode field value
// and a boolean to check if the value has been set.
func (o *CreateUserMfaVerification200ResponseOneOf) GetSecretQrCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretQrCode, true
}

// SetSecretQrCode sets field value
func (o *CreateUserMfaVerification200ResponseOneOf) SetSecretQrCode(v string) {
	o.SecretQrCode = v
}

func (o CreateUserMfaVerification200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserMfaVerification200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["secret"] = o.Secret
	toSerialize["secretQrCode"] = o.SecretQrCode
	return toSerialize, nil
}

func (o *CreateUserMfaVerification200ResponseOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"secret",
		"secretQrCode",
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

	varCreateUserMfaVerification200ResponseOneOf := _CreateUserMfaVerification200ResponseOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUserMfaVerification200ResponseOneOf)

	if err != nil {
		return err
	}

	*o = CreateUserMfaVerification200ResponseOneOf(varCreateUserMfaVerification200ResponseOneOf)

	return err
}

type NullableCreateUserMfaVerification200ResponseOneOf struct {
	value *CreateUserMfaVerification200ResponseOneOf
	isSet bool
}

func (v NullableCreateUserMfaVerification200ResponseOneOf) Get() *CreateUserMfaVerification200ResponseOneOf {
	return v.value
}

func (v *NullableCreateUserMfaVerification200ResponseOneOf) Set(val *CreateUserMfaVerification200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserMfaVerification200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserMfaVerification200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserMfaVerification200ResponseOneOf(val *CreateUserMfaVerification200ResponseOneOf) *NullableCreateUserMfaVerification200ResponseOneOf {
	return &NullableCreateUserMfaVerification200ResponseOneOf{value: val, isSet: true}
}

func (v NullableCreateUserMfaVerification200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserMfaVerification200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

