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

// checks if the UpdateResourceIsDefaultRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateResourceIsDefaultRequest{}

// UpdateResourceIsDefaultRequest struct for UpdateResourceIsDefaultRequest
type UpdateResourceIsDefaultRequest struct {
	// The updated value of the `isDefault` property.
	IsDefault bool `json:"isDefault"`
}

type _UpdateResourceIsDefaultRequest UpdateResourceIsDefaultRequest

// NewUpdateResourceIsDefaultRequest instantiates a new UpdateResourceIsDefaultRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateResourceIsDefaultRequest(isDefault bool) *UpdateResourceIsDefaultRequest {
	this := UpdateResourceIsDefaultRequest{}
	this.IsDefault = isDefault
	return &this
}

// NewUpdateResourceIsDefaultRequestWithDefaults instantiates a new UpdateResourceIsDefaultRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateResourceIsDefaultRequestWithDefaults() *UpdateResourceIsDefaultRequest {
	this := UpdateResourceIsDefaultRequest{}
	return &this
}

// GetIsDefault returns the IsDefault field value
func (o *UpdateResourceIsDefaultRequest) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *UpdateResourceIsDefaultRequest) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *UpdateResourceIsDefaultRequest) SetIsDefault(v bool) {
	o.IsDefault = v
}

func (o UpdateResourceIsDefaultRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateResourceIsDefaultRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isDefault"] = o.IsDefault
	return toSerialize, nil
}

func (o *UpdateResourceIsDefaultRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isDefault",
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

	varUpdateResourceIsDefaultRequest := _UpdateResourceIsDefaultRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateResourceIsDefaultRequest)

	if err != nil {
		return err
	}

	*o = UpdateResourceIsDefaultRequest(varUpdateResourceIsDefaultRequest)

	return err
}

type NullableUpdateResourceIsDefaultRequest struct {
	value *UpdateResourceIsDefaultRequest
	isSet bool
}

func (v NullableUpdateResourceIsDefaultRequest) Get() *UpdateResourceIsDefaultRequest {
	return v.value
}

func (v *NullableUpdateResourceIsDefaultRequest) Set(val *UpdateResourceIsDefaultRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateResourceIsDefaultRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateResourceIsDefaultRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateResourceIsDefaultRequest(val *UpdateResourceIsDefaultRequest) *NullableUpdateResourceIsDefaultRequest {
	return &NullableUpdateResourceIsDefaultRequest{value: val, isSet: true}
}

func (v NullableUpdateResourceIsDefaultRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateResourceIsDefaultRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
