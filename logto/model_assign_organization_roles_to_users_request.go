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

// checks if the AssignOrganizationRolesToUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignOrganizationRolesToUsersRequest{}

// AssignOrganizationRolesToUsersRequest struct for AssignOrganizationRolesToUsersRequest
type AssignOrganizationRolesToUsersRequest struct {
	// An array of user IDs to assign roles.
	UserIds []string `json:"userIds"`
	// An array of organization role IDs to assign. User existed roles assignment will be ignored.
	OrganizationRoleIds []string `json:"organizationRoleIds"`
}

type _AssignOrganizationRolesToUsersRequest AssignOrganizationRolesToUsersRequest

// NewAssignOrganizationRolesToUsersRequest instantiates a new AssignOrganizationRolesToUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignOrganizationRolesToUsersRequest(userIds []string, organizationRoleIds []string) *AssignOrganizationRolesToUsersRequest {
	this := AssignOrganizationRolesToUsersRequest{}
	this.UserIds = userIds
	this.OrganizationRoleIds = organizationRoleIds
	return &this
}

// NewAssignOrganizationRolesToUsersRequestWithDefaults instantiates a new AssignOrganizationRolesToUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignOrganizationRolesToUsersRequestWithDefaults() *AssignOrganizationRolesToUsersRequest {
	this := AssignOrganizationRolesToUsersRequest{}
	return &this
}

// GetUserIds returns the UserIds field value
func (o *AssignOrganizationRolesToUsersRequest) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value
// and a boolean to check if the value has been set.
func (o *AssignOrganizationRolesToUsersRequest) GetUserIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserIds, true
}

// SetUserIds sets field value
func (o *AssignOrganizationRolesToUsersRequest) SetUserIds(v []string) {
	o.UserIds = v
}

// GetOrganizationRoleIds returns the OrganizationRoleIds field value
func (o *AssignOrganizationRolesToUsersRequest) GetOrganizationRoleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OrganizationRoleIds
}

// GetOrganizationRoleIdsOk returns a tuple with the OrganizationRoleIds field value
// and a boolean to check if the value has been set.
func (o *AssignOrganizationRolesToUsersRequest) GetOrganizationRoleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationRoleIds, true
}

// SetOrganizationRoleIds sets field value
func (o *AssignOrganizationRolesToUsersRequest) SetOrganizationRoleIds(v []string) {
	o.OrganizationRoleIds = v
}

func (o AssignOrganizationRolesToUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignOrganizationRolesToUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userIds"] = o.UserIds
	toSerialize["organizationRoleIds"] = o.OrganizationRoleIds
	return toSerialize, nil
}

func (o *AssignOrganizationRolesToUsersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userIds",
		"organizationRoleIds",
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

	varAssignOrganizationRolesToUsersRequest := _AssignOrganizationRolesToUsersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssignOrganizationRolesToUsersRequest)

	if err != nil {
		return err
	}

	*o = AssignOrganizationRolesToUsersRequest(varAssignOrganizationRolesToUsersRequest)

	return err
}

type NullableAssignOrganizationRolesToUsersRequest struct {
	value *AssignOrganizationRolesToUsersRequest
	isSet bool
}

func (v NullableAssignOrganizationRolesToUsersRequest) Get() *AssignOrganizationRolesToUsersRequest {
	return v.value
}

func (v *NullableAssignOrganizationRolesToUsersRequest) Set(val *AssignOrganizationRolesToUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignOrganizationRolesToUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignOrganizationRolesToUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignOrganizationRolesToUsersRequest(val *AssignOrganizationRolesToUsersRequest) *NullableAssignOrganizationRolesToUsersRequest {
	return &NullableAssignOrganizationRolesToUsersRequest{value: val, isSet: true}
}

func (v NullableAssignOrganizationRolesToUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignOrganizationRolesToUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
