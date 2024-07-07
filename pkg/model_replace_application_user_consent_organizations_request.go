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

// checks if the ReplaceApplicationUserConsentOrganizationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceApplicationUserConsentOrganizationsRequest{}

// ReplaceApplicationUserConsentOrganizationsRequest struct for ReplaceApplicationUserConsentOrganizationsRequest
type ReplaceApplicationUserConsentOrganizationsRequest struct {
	// A list of organization ids to be granted. <br/> All the existing organizations' access will be revoked if not in the list. <br/> If the list is empty, all the organizations' access will be revoked.
	OrganizationIds []string `json:"organizationIds"`
}

type _ReplaceApplicationUserConsentOrganizationsRequest ReplaceApplicationUserConsentOrganizationsRequest

// NewReplaceApplicationUserConsentOrganizationsRequest instantiates a new ReplaceApplicationUserConsentOrganizationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceApplicationUserConsentOrganizationsRequest(organizationIds []string) *ReplaceApplicationUserConsentOrganizationsRequest {
	this := ReplaceApplicationUserConsentOrganizationsRequest{}
	this.OrganizationIds = organizationIds
	return &this
}

// NewReplaceApplicationUserConsentOrganizationsRequestWithDefaults instantiates a new ReplaceApplicationUserConsentOrganizationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceApplicationUserConsentOrganizationsRequestWithDefaults() *ReplaceApplicationUserConsentOrganizationsRequest {
	this := ReplaceApplicationUserConsentOrganizationsRequest{}
	return &this
}

// GetOrganizationIds returns the OrganizationIds field value
func (o *ReplaceApplicationUserConsentOrganizationsRequest) GetOrganizationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OrganizationIds
}

// GetOrganizationIdsOk returns a tuple with the OrganizationIds field value
// and a boolean to check if the value has been set.
func (o *ReplaceApplicationUserConsentOrganizationsRequest) GetOrganizationIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationIds, true
}

// SetOrganizationIds sets field value
func (o *ReplaceApplicationUserConsentOrganizationsRequest) SetOrganizationIds(v []string) {
	o.OrganizationIds = v
}

func (o ReplaceApplicationUserConsentOrganizationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceApplicationUserConsentOrganizationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["organizationIds"] = o.OrganizationIds
	return toSerialize, nil
}

func (o *ReplaceApplicationUserConsentOrganizationsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"organizationIds",
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

	varReplaceApplicationUserConsentOrganizationsRequest := _ReplaceApplicationUserConsentOrganizationsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplaceApplicationUserConsentOrganizationsRequest)

	if err != nil {
		return err
	}

	*o = ReplaceApplicationUserConsentOrganizationsRequest(varReplaceApplicationUserConsentOrganizationsRequest)

	return err
}

type NullableReplaceApplicationUserConsentOrganizationsRequest struct {
	value *ReplaceApplicationUserConsentOrganizationsRequest
	isSet bool
}

func (v NullableReplaceApplicationUserConsentOrganizationsRequest) Get() *ReplaceApplicationUserConsentOrganizationsRequest {
	return v.value
}

func (v *NullableReplaceApplicationUserConsentOrganizationsRequest) Set(val *ReplaceApplicationUserConsentOrganizationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceApplicationUserConsentOrganizationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceApplicationUserConsentOrganizationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceApplicationUserConsentOrganizationsRequest(val *ReplaceApplicationUserConsentOrganizationsRequest) *NullableReplaceApplicationUserConsentOrganizationsRequest {
	return &NullableReplaceApplicationUserConsentOrganizationsRequest{value: val, isSet: true}
}

func (v NullableReplaceApplicationUserConsentOrganizationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceApplicationUserConsentOrganizationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
