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

// checks if the ReplaceOrganizationUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceOrganizationUsersRequest{}

// ReplaceOrganizationUsersRequest struct for ReplaceOrganizationUsersRequest
type ReplaceOrganizationUsersRequest struct {
	// An array of user IDs to replace existing users.
	UserIds []string `json:"userIds"`
}

type _ReplaceOrganizationUsersRequest ReplaceOrganizationUsersRequest

// NewReplaceOrganizationUsersRequest instantiates a new ReplaceOrganizationUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceOrganizationUsersRequest(userIds []string) *ReplaceOrganizationUsersRequest {
	this := ReplaceOrganizationUsersRequest{}
	this.UserIds = userIds
	return &this
}

// NewReplaceOrganizationUsersRequestWithDefaults instantiates a new ReplaceOrganizationUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceOrganizationUsersRequestWithDefaults() *ReplaceOrganizationUsersRequest {
	this := ReplaceOrganizationUsersRequest{}
	return &this
}

// GetUserIds returns the UserIds field value
func (o *ReplaceOrganizationUsersRequest) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value
// and a boolean to check if the value has been set.
func (o *ReplaceOrganizationUsersRequest) GetUserIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserIds, true
}

// SetUserIds sets field value
func (o *ReplaceOrganizationUsersRequest) SetUserIds(v []string) {
	o.UserIds = v
}

func (o ReplaceOrganizationUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceOrganizationUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userIds"] = o.UserIds
	return toSerialize, nil
}

func (o *ReplaceOrganizationUsersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userIds",
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

	varReplaceOrganizationUsersRequest := _ReplaceOrganizationUsersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplaceOrganizationUsersRequest)

	if err != nil {
		return err
	}

	*o = ReplaceOrganizationUsersRequest(varReplaceOrganizationUsersRequest)

	return err
}

type NullableReplaceOrganizationUsersRequest struct {
	value *ReplaceOrganizationUsersRequest
	isSet bool
}

func (v NullableReplaceOrganizationUsersRequest) Get() *ReplaceOrganizationUsersRequest {
	return v.value
}

func (v *NullableReplaceOrganizationUsersRequest) Set(val *ReplaceOrganizationUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceOrganizationUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceOrganizationUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceOrganizationUsersRequest(val *ReplaceOrganizationUsersRequest) *NullableReplaceOrganizationUsersRequest {
	return &NullableReplaceOrganizationUsersRequest{value: val, isSet: true}
}

func (v NullableReplaceOrganizationUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceOrganizationUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

