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

// checks if the ReplaceOrganizationJitSsoConnectorsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceOrganizationJitSsoConnectorsRequest{}

// ReplaceOrganizationJitSsoConnectorsRequest struct for ReplaceOrganizationJitSsoConnectorsRequest
type ReplaceOrganizationJitSsoConnectorsRequest struct {
	// An array of SSO connector IDs to replace existing SSO connectors.
	SsoConnectorIds []string `json:"ssoConnectorIds"`
}

type _ReplaceOrganizationJitSsoConnectorsRequest ReplaceOrganizationJitSsoConnectorsRequest

// NewReplaceOrganizationJitSsoConnectorsRequest instantiates a new ReplaceOrganizationJitSsoConnectorsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceOrganizationJitSsoConnectorsRequest(ssoConnectorIds []string) *ReplaceOrganizationJitSsoConnectorsRequest {
	this := ReplaceOrganizationJitSsoConnectorsRequest{}
	this.SsoConnectorIds = ssoConnectorIds
	return &this
}

// NewReplaceOrganizationJitSsoConnectorsRequestWithDefaults instantiates a new ReplaceOrganizationJitSsoConnectorsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceOrganizationJitSsoConnectorsRequestWithDefaults() *ReplaceOrganizationJitSsoConnectorsRequest {
	this := ReplaceOrganizationJitSsoConnectorsRequest{}
	return &this
}

// GetSsoConnectorIds returns the SsoConnectorIds field value
func (o *ReplaceOrganizationJitSsoConnectorsRequest) GetSsoConnectorIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SsoConnectorIds
}

// GetSsoConnectorIdsOk returns a tuple with the SsoConnectorIds field value
// and a boolean to check if the value has been set.
func (o *ReplaceOrganizationJitSsoConnectorsRequest) GetSsoConnectorIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoConnectorIds, true
}

// SetSsoConnectorIds sets field value
func (o *ReplaceOrganizationJitSsoConnectorsRequest) SetSsoConnectorIds(v []string) {
	o.SsoConnectorIds = v
}

func (o ReplaceOrganizationJitSsoConnectorsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceOrganizationJitSsoConnectorsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ssoConnectorIds"] = o.SsoConnectorIds
	return toSerialize, nil
}

func (o *ReplaceOrganizationJitSsoConnectorsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ssoConnectorIds",
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

	varReplaceOrganizationJitSsoConnectorsRequest := _ReplaceOrganizationJitSsoConnectorsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplaceOrganizationJitSsoConnectorsRequest)

	if err != nil {
		return err
	}

	*o = ReplaceOrganizationJitSsoConnectorsRequest(varReplaceOrganizationJitSsoConnectorsRequest)

	return err
}

type NullableReplaceOrganizationJitSsoConnectorsRequest struct {
	value *ReplaceOrganizationJitSsoConnectorsRequest
	isSet bool
}

func (v NullableReplaceOrganizationJitSsoConnectorsRequest) Get() *ReplaceOrganizationJitSsoConnectorsRequest {
	return v.value
}

func (v *NullableReplaceOrganizationJitSsoConnectorsRequest) Set(val *ReplaceOrganizationJitSsoConnectorsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceOrganizationJitSsoConnectorsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceOrganizationJitSsoConnectorsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceOrganizationJitSsoConnectorsRequest(val *ReplaceOrganizationJitSsoConnectorsRequest) *NullableReplaceOrganizationJitSsoConnectorsRequest {
	return &NullableReplaceOrganizationJitSsoConnectorsRequest{value: val, isSet: true}
}

func (v NullableReplaceOrganizationJitSsoConnectorsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceOrganizationJitSsoConnectorsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
