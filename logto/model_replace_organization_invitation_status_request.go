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

// checks if the ReplaceOrganizationInvitationStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceOrganizationInvitationStatusRequest{}

// ReplaceOrganizationInvitationStatusRequest struct for ReplaceOrganizationInvitationStatusRequest
type ReplaceOrganizationInvitationStatusRequest struct {
	// The ID of the user who accepted the organization invitation. Required if the status is \"Accepted\".
	AcceptedUserId NullableString `json:"acceptedUserId,omitempty"`
	// The status of the organization invitation.
	Status string `json:"status"`
}

type _ReplaceOrganizationInvitationStatusRequest ReplaceOrganizationInvitationStatusRequest

// NewReplaceOrganizationInvitationStatusRequest instantiates a new ReplaceOrganizationInvitationStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceOrganizationInvitationStatusRequest(status string) *ReplaceOrganizationInvitationStatusRequest {
	this := ReplaceOrganizationInvitationStatusRequest{}
	this.Status = status
	return &this
}

// NewReplaceOrganizationInvitationStatusRequestWithDefaults instantiates a new ReplaceOrganizationInvitationStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceOrganizationInvitationStatusRequestWithDefaults() *ReplaceOrganizationInvitationStatusRequest {
	this := ReplaceOrganizationInvitationStatusRequest{}
	return &this
}

// GetAcceptedUserId returns the AcceptedUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplaceOrganizationInvitationStatusRequest) GetAcceptedUserId() string {
	if o == nil || IsNil(o.AcceptedUserId.Get()) {
		var ret string
		return ret
	}
	return *o.AcceptedUserId.Get()
}

// GetAcceptedUserIdOk returns a tuple with the AcceptedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplaceOrganizationInvitationStatusRequest) GetAcceptedUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptedUserId.Get(), o.AcceptedUserId.IsSet()
}

// HasAcceptedUserId returns a boolean if a field has been set.
func (o *ReplaceOrganizationInvitationStatusRequest) HasAcceptedUserId() bool {
	if o != nil && o.AcceptedUserId.IsSet() {
		return true
	}

	return false
}

// SetAcceptedUserId gets a reference to the given NullableString and assigns it to the AcceptedUserId field.
func (o *ReplaceOrganizationInvitationStatusRequest) SetAcceptedUserId(v string) {
	o.AcceptedUserId.Set(&v)
}
// SetAcceptedUserIdNil sets the value for AcceptedUserId to be an explicit nil
func (o *ReplaceOrganizationInvitationStatusRequest) SetAcceptedUserIdNil() {
	o.AcceptedUserId.Set(nil)
}

// UnsetAcceptedUserId ensures that no value is present for AcceptedUserId, not even an explicit nil
func (o *ReplaceOrganizationInvitationStatusRequest) UnsetAcceptedUserId() {
	o.AcceptedUserId.Unset()
}

// GetStatus returns the Status field value
func (o *ReplaceOrganizationInvitationStatusRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReplaceOrganizationInvitationStatusRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReplaceOrganizationInvitationStatusRequest) SetStatus(v string) {
	o.Status = v
}

func (o ReplaceOrganizationInvitationStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceOrganizationInvitationStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptedUserId.IsSet() {
		toSerialize["acceptedUserId"] = o.AcceptedUserId.Get()
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ReplaceOrganizationInvitationStatusRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varReplaceOrganizationInvitationStatusRequest := _ReplaceOrganizationInvitationStatusRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplaceOrganizationInvitationStatusRequest)

	if err != nil {
		return err
	}

	*o = ReplaceOrganizationInvitationStatusRequest(varReplaceOrganizationInvitationStatusRequest)

	return err
}

type NullableReplaceOrganizationInvitationStatusRequest struct {
	value *ReplaceOrganizationInvitationStatusRequest
	isSet bool
}

func (v NullableReplaceOrganizationInvitationStatusRequest) Get() *ReplaceOrganizationInvitationStatusRequest {
	return v.value
}

func (v *NullableReplaceOrganizationInvitationStatusRequest) Set(val *ReplaceOrganizationInvitationStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceOrganizationInvitationStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceOrganizationInvitationStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceOrganizationInvitationStatusRequest(val *ReplaceOrganizationInvitationStatusRequest) *NullableReplaceOrganizationInvitationStatusRequest {
	return &NullableReplaceOrganizationInvitationStatusRequest{value: val, isSet: true}
}

func (v NullableReplaceOrganizationInvitationStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceOrganizationInvitationStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
