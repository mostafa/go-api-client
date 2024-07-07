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

// checks if the CreateRoleApplicationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleApplicationRequest{}

// CreateRoleApplicationRequest struct for CreateRoleApplicationRequest
type CreateRoleApplicationRequest struct {
	// An array of application IDs to be assigned.
	ApplicationIds []string `json:"applicationIds"`
}

type _CreateRoleApplicationRequest CreateRoleApplicationRequest

// NewCreateRoleApplicationRequest instantiates a new CreateRoleApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleApplicationRequest(applicationIds []string) *CreateRoleApplicationRequest {
	this := CreateRoleApplicationRequest{}
	this.ApplicationIds = applicationIds
	return &this
}

// NewCreateRoleApplicationRequestWithDefaults instantiates a new CreateRoleApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleApplicationRequestWithDefaults() *CreateRoleApplicationRequest {
	this := CreateRoleApplicationRequest{}
	return &this
}

// GetApplicationIds returns the ApplicationIds field value
func (o *CreateRoleApplicationRequest) GetApplicationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ApplicationIds
}

// GetApplicationIdsOk returns a tuple with the ApplicationIds field value
// and a boolean to check if the value has been set.
func (o *CreateRoleApplicationRequest) GetApplicationIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationIds, true
}

// SetApplicationIds sets field value
func (o *CreateRoleApplicationRequest) SetApplicationIds(v []string) {
	o.ApplicationIds = v
}

func (o CreateRoleApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleApplicationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applicationIds"] = o.ApplicationIds
	return toSerialize, nil
}

func (o *CreateRoleApplicationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applicationIds",
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

	varCreateRoleApplicationRequest := _CreateRoleApplicationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRoleApplicationRequest)

	if err != nil {
		return err
	}

	*o = CreateRoleApplicationRequest(varCreateRoleApplicationRequest)

	return err
}

type NullableCreateRoleApplicationRequest struct {
	value *CreateRoleApplicationRequest
	isSet bool
}

func (v NullableCreateRoleApplicationRequest) Get() *CreateRoleApplicationRequest {
	return v.value
}

func (v *NullableCreateRoleApplicationRequest) Set(val *CreateRoleApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleApplicationRequest(val *CreateRoleApplicationRequest) *NullableCreateRoleApplicationRequest {
	return &NullableCreateRoleApplicationRequest{value: val, isSet: true}
}

func (v NullableCreateRoleApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

