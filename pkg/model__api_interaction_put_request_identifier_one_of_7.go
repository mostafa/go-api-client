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

// checks if the ApiInteractionPutRequestIdentifierOneOf7 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionPutRequestIdentifierOneOf7{}

// ApiInteractionPutRequestIdentifierOneOf7 struct for ApiInteractionPutRequestIdentifierOneOf7
type ApiInteractionPutRequestIdentifierOneOf7 struct {
	ConnectorId string `json:"connectorId"`
	Phone string `json:"phone"`
}

type _ApiInteractionPutRequestIdentifierOneOf7 ApiInteractionPutRequestIdentifierOneOf7

// NewApiInteractionPutRequestIdentifierOneOf7 instantiates a new ApiInteractionPutRequestIdentifierOneOf7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionPutRequestIdentifierOneOf7(connectorId string, phone string) *ApiInteractionPutRequestIdentifierOneOf7 {
	this := ApiInteractionPutRequestIdentifierOneOf7{}
	this.ConnectorId = connectorId
	this.Phone = phone
	return &this
}

// NewApiInteractionPutRequestIdentifierOneOf7WithDefaults instantiates a new ApiInteractionPutRequestIdentifierOneOf7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionPutRequestIdentifierOneOf7WithDefaults() *ApiInteractionPutRequestIdentifierOneOf7 {
	this := ApiInteractionPutRequestIdentifierOneOf7{}
	return &this
}

// GetConnectorId returns the ConnectorId field value
func (o *ApiInteractionPutRequestIdentifierOneOf7) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionPutRequestIdentifierOneOf7) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *ApiInteractionPutRequestIdentifierOneOf7) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetPhone returns the Phone field value
func (o *ApiInteractionPutRequestIdentifierOneOf7) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionPutRequestIdentifierOneOf7) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *ApiInteractionPutRequestIdentifierOneOf7) SetPhone(v string) {
	o.Phone = v
}

func (o ApiInteractionPutRequestIdentifierOneOf7) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionPutRequestIdentifierOneOf7) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectorId"] = o.ConnectorId
	toSerialize["phone"] = o.Phone
	return toSerialize, nil
}

func (o *ApiInteractionPutRequestIdentifierOneOf7) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectorId",
		"phone",
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

	varApiInteractionPutRequestIdentifierOneOf7 := _ApiInteractionPutRequestIdentifierOneOf7{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiInteractionPutRequestIdentifierOneOf7)

	if err != nil {
		return err
	}

	*o = ApiInteractionPutRequestIdentifierOneOf7(varApiInteractionPutRequestIdentifierOneOf7)

	return err
}

type NullableApiInteractionPutRequestIdentifierOneOf7 struct {
	value *ApiInteractionPutRequestIdentifierOneOf7
	isSet bool
}

func (v NullableApiInteractionPutRequestIdentifierOneOf7) Get() *ApiInteractionPutRequestIdentifierOneOf7 {
	return v.value
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf7) Set(val *ApiInteractionPutRequestIdentifierOneOf7) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionPutRequestIdentifierOneOf7) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionPutRequestIdentifierOneOf7(val *ApiInteractionPutRequestIdentifierOneOf7) *NullableApiInteractionPutRequestIdentifierOneOf7 {
	return &NullableApiInteractionPutRequestIdentifierOneOf7{value: val, isSet: true}
}

func (v NullableApiInteractionPutRequestIdentifierOneOf7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionPutRequestIdentifierOneOf7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

