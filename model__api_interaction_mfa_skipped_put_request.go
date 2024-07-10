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

// checks if the ApiInteractionMfaSkippedPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionMfaSkippedPutRequest{}

// ApiInteractionMfaSkippedPutRequest struct for ApiInteractionMfaSkippedPutRequest
type ApiInteractionMfaSkippedPutRequest struct {
	MfaSkipped bool `json:"mfaSkipped"`
}

type _ApiInteractionMfaSkippedPutRequest ApiInteractionMfaSkippedPutRequest

// NewApiInteractionMfaSkippedPutRequest instantiates a new ApiInteractionMfaSkippedPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionMfaSkippedPutRequest(mfaSkipped bool) *ApiInteractionMfaSkippedPutRequest {
	this := ApiInteractionMfaSkippedPutRequest{}
	this.MfaSkipped = mfaSkipped
	return &this
}

// NewApiInteractionMfaSkippedPutRequestWithDefaults instantiates a new ApiInteractionMfaSkippedPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionMfaSkippedPutRequestWithDefaults() *ApiInteractionMfaSkippedPutRequest {
	this := ApiInteractionMfaSkippedPutRequest{}
	return &this
}

// GetMfaSkipped returns the MfaSkipped field value
func (o *ApiInteractionMfaSkippedPutRequest) GetMfaSkipped() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MfaSkipped
}

// GetMfaSkippedOk returns a tuple with the MfaSkipped field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionMfaSkippedPutRequest) GetMfaSkippedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MfaSkipped, true
}

// SetMfaSkipped sets field value
func (o *ApiInteractionMfaSkippedPutRequest) SetMfaSkipped(v bool) {
	o.MfaSkipped = v
}

func (o ApiInteractionMfaSkippedPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionMfaSkippedPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mfaSkipped"] = o.MfaSkipped
	return toSerialize, nil
}

func (o *ApiInteractionMfaSkippedPutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mfaSkipped",
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

	varApiInteractionMfaSkippedPutRequest := _ApiInteractionMfaSkippedPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiInteractionMfaSkippedPutRequest)

	if err != nil {
		return err
	}

	*o = ApiInteractionMfaSkippedPutRequest(varApiInteractionMfaSkippedPutRequest)

	return err
}

type NullableApiInteractionMfaSkippedPutRequest struct {
	value *ApiInteractionMfaSkippedPutRequest
	isSet bool
}

func (v NullableApiInteractionMfaSkippedPutRequest) Get() *ApiInteractionMfaSkippedPutRequest {
	return v.value
}

func (v *NullableApiInteractionMfaSkippedPutRequest) Set(val *ApiInteractionMfaSkippedPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionMfaSkippedPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionMfaSkippedPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionMfaSkippedPutRequest(val *ApiInteractionMfaSkippedPutRequest) *NullableApiInteractionMfaSkippedPutRequest {
	return &NullableApiInteractionMfaSkippedPutRequest{value: val, isSet: true}
}

func (v NullableApiInteractionMfaSkippedPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionMfaSkippedPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
