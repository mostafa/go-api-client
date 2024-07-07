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

// checks if the CreateOrganizationRoleScopeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationRoleScopeRequest{}

// CreateOrganizationRoleScopeRequest struct for CreateOrganizationRoleScopeRequest
type CreateOrganizationRoleScopeRequest struct {
	// An array of organization scope IDs to be assigned. Existed scope IDs assignments will be ignored.
	OrganizationScopeIds []string `json:"organizationScopeIds"`
}

type _CreateOrganizationRoleScopeRequest CreateOrganizationRoleScopeRequest

// NewCreateOrganizationRoleScopeRequest instantiates a new CreateOrganizationRoleScopeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationRoleScopeRequest(organizationScopeIds []string) *CreateOrganizationRoleScopeRequest {
	this := CreateOrganizationRoleScopeRequest{}
	this.OrganizationScopeIds = organizationScopeIds
	return &this
}

// NewCreateOrganizationRoleScopeRequestWithDefaults instantiates a new CreateOrganizationRoleScopeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationRoleScopeRequestWithDefaults() *CreateOrganizationRoleScopeRequest {
	this := CreateOrganizationRoleScopeRequest{}
	return &this
}

// GetOrganizationScopeIds returns the OrganizationScopeIds field value
func (o *CreateOrganizationRoleScopeRequest) GetOrganizationScopeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OrganizationScopeIds
}

// GetOrganizationScopeIdsOk returns a tuple with the OrganizationScopeIds field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationRoleScopeRequest) GetOrganizationScopeIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationScopeIds, true
}

// SetOrganizationScopeIds sets field value
func (o *CreateOrganizationRoleScopeRequest) SetOrganizationScopeIds(v []string) {
	o.OrganizationScopeIds = v
}

func (o CreateOrganizationRoleScopeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationRoleScopeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["organizationScopeIds"] = o.OrganizationScopeIds
	return toSerialize, nil
}

func (o *CreateOrganizationRoleScopeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"organizationScopeIds",
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

	varCreateOrganizationRoleScopeRequest := _CreateOrganizationRoleScopeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrganizationRoleScopeRequest)

	if err != nil {
		return err
	}

	*o = CreateOrganizationRoleScopeRequest(varCreateOrganizationRoleScopeRequest)

	return err
}

type NullableCreateOrganizationRoleScopeRequest struct {
	value *CreateOrganizationRoleScopeRequest
	isSet bool
}

func (v NullableCreateOrganizationRoleScopeRequest) Get() *CreateOrganizationRoleScopeRequest {
	return v.value
}

func (v *NullableCreateOrganizationRoleScopeRequest) Set(val *CreateOrganizationRoleScopeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationRoleScopeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationRoleScopeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationRoleScopeRequest(val *CreateOrganizationRoleScopeRequest) *NullableCreateOrganizationRoleScopeRequest {
	return &NullableCreateOrganizationRoleScopeRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationRoleScopeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationRoleScopeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

