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

// checks if the CreateConnectorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateConnectorRequest{}

// CreateConnectorRequest struct for CreateConnectorRequest
type CreateConnectorRequest struct {
	// The connector config object that will be passed to the connector. The config object should be compatible with the connector factory.
	Config map[string]interface{} `json:"config,omitempty"`
	// The connector factory ID for creating the connector.
	ConnectorId string `json:"connectorId"`
	Metadata *CreateConnectorRequestMetadata `json:"metadata,omitempty"`
	// Whether to sync user profile from the identity provider to Logto at each sign-in. If `false`, the user profile will only be synced when the user is created.
	SyncProfile *bool `json:"syncProfile,omitempty"`
	// The unique ID for the connector. If not provided, a random ID will be generated.
	Id *string `json:"id,omitempty"`
}

type _CreateConnectorRequest CreateConnectorRequest

// NewCreateConnectorRequest instantiates a new CreateConnectorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConnectorRequest(connectorId string) *CreateConnectorRequest {
	this := CreateConnectorRequest{}
	this.ConnectorId = connectorId
	return &this
}

// NewCreateConnectorRequestWithDefaults instantiates a new CreateConnectorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConnectorRequestWithDefaults() *CreateConnectorRequest {
	this := CreateConnectorRequest{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CreateConnectorRequest) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequest) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CreateConnectorRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *CreateConnectorRequest) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetConnectorId returns the ConnectorId field value
func (o *CreateConnectorRequest) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequest) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *CreateConnectorRequest) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateConnectorRequest) GetMetadata() CreateConnectorRequestMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret CreateConnectorRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequest) GetMetadataOk() (*CreateConnectorRequestMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateConnectorRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given CreateConnectorRequestMetadata and assigns it to the Metadata field.
func (o *CreateConnectorRequest) SetMetadata(v CreateConnectorRequestMetadata) {
	o.Metadata = &v
}

// GetSyncProfile returns the SyncProfile field value if set, zero value otherwise.
func (o *CreateConnectorRequest) GetSyncProfile() bool {
	if o == nil || IsNil(o.SyncProfile) {
		var ret bool
		return ret
	}
	return *o.SyncProfile
}

// GetSyncProfileOk returns a tuple with the SyncProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequest) GetSyncProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncProfile) {
		return nil, false
	}
	return o.SyncProfile, true
}

// HasSyncProfile returns a boolean if a field has been set.
func (o *CreateConnectorRequest) HasSyncProfile() bool {
	if o != nil && !IsNil(o.SyncProfile) {
		return true
	}

	return false
}

// SetSyncProfile gets a reference to the given bool and assigns it to the SyncProfile field.
func (o *CreateConnectorRequest) SetSyncProfile(v bool) {
	o.SyncProfile = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateConnectorRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectorRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateConnectorRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateConnectorRequest) SetId(v string) {
	o.Id = &v
}

func (o CreateConnectorRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateConnectorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	toSerialize["connectorId"] = o.ConnectorId
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.SyncProfile) {
		toSerialize["syncProfile"] = o.SyncProfile
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

func (o *CreateConnectorRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectorId",
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

	varCreateConnectorRequest := _CreateConnectorRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateConnectorRequest)

	if err != nil {
		return err
	}

	*o = CreateConnectorRequest(varCreateConnectorRequest)

	return err
}

type NullableCreateConnectorRequest struct {
	value *CreateConnectorRequest
	isSet bool
}

func (v NullableCreateConnectorRequest) Get() *CreateConnectorRequest {
	return v.value
}

func (v *NullableCreateConnectorRequest) Set(val *CreateConnectorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConnectorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConnectorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConnectorRequest(val *CreateConnectorRequest) *NullableCreateConnectorRequest {
	return &NullableCreateConnectorRequest{value: val, isSet: true}
}

func (v NullableCreateConnectorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConnectorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

