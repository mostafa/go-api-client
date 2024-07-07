/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"encoding/json"
)

// checks if the UpdateConnectorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateConnectorRequest{}

// UpdateConnectorRequest struct for UpdateConnectorRequest
type UpdateConnectorRequest struct {
	// The connector config object that will be passed to the connector. The config object should be compatible with the connector factory.
	Config map[string]interface{} `json:"config,omitempty"`
	Metadata *UpdateConnectorRequestMetadata `json:"metadata,omitempty"`
	// Whether to sync user profile from the identity provider to Logto at each sign-in. If `false`, the user profile will only be synced when the user is created.
	SyncProfile *bool `json:"syncProfile,omitempty"`
}

// NewUpdateConnectorRequest instantiates a new UpdateConnectorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConnectorRequest() *UpdateConnectorRequest {
	this := UpdateConnectorRequest{}
	return &this
}

// NewUpdateConnectorRequestWithDefaults instantiates a new UpdateConnectorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConnectorRequestWithDefaults() *UpdateConnectorRequest {
	this := UpdateConnectorRequest{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UpdateConnectorRequest) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorRequest) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *UpdateConnectorRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *UpdateConnectorRequest) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateConnectorRequest) GetMetadata() UpdateConnectorRequestMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret UpdateConnectorRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorRequest) GetMetadataOk() (*UpdateConnectorRequestMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateConnectorRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given UpdateConnectorRequestMetadata and assigns it to the Metadata field.
func (o *UpdateConnectorRequest) SetMetadata(v UpdateConnectorRequestMetadata) {
	o.Metadata = &v
}

// GetSyncProfile returns the SyncProfile field value if set, zero value otherwise.
func (o *UpdateConnectorRequest) GetSyncProfile() bool {
	if o == nil || IsNil(o.SyncProfile) {
		var ret bool
		return ret
	}
	return *o.SyncProfile
}

// GetSyncProfileOk returns a tuple with the SyncProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorRequest) GetSyncProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncProfile) {
		return nil, false
	}
	return o.SyncProfile, true
}

// HasSyncProfile returns a boolean if a field has been set.
func (o *UpdateConnectorRequest) HasSyncProfile() bool {
	if o != nil && !IsNil(o.SyncProfile) {
		return true
	}

	return false
}

// SetSyncProfile gets a reference to the given bool and assigns it to the SyncProfile field.
func (o *UpdateConnectorRequest) SetSyncProfile(v bool) {
	o.SyncProfile = &v
}

func (o UpdateConnectorRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateConnectorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.SyncProfile) {
		toSerialize["syncProfile"] = o.SyncProfile
	}
	return toSerialize, nil
}

type NullableUpdateConnectorRequest struct {
	value *UpdateConnectorRequest
	isSet bool
}

func (v NullableUpdateConnectorRequest) Get() *UpdateConnectorRequest {
	return v.value
}

func (v *NullableUpdateConnectorRequest) Set(val *UpdateConnectorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConnectorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConnectorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConnectorRequest(val *UpdateConnectorRequest) *NullableUpdateConnectorRequest {
	return &NullableUpdateConnectorRequest{value: val, isSet: true}
}

func (v NullableUpdateConnectorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConnectorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

