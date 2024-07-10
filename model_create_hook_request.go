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

// checks if the CreateHookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateHookRequest{}

// CreateHookRequest struct for CreateHookRequest
type CreateHookRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	// The name of the hook.
	Name *string `json:"name,omitempty"`
	// Use `events` instead.
	// Deprecated
	Event *string `json:"event,omitempty"`
	// An array of hook events.
	Events []string `json:"events,omitempty"`
	Config CreateHookRequestConfig `json:"config"`
	Enabled *bool `json:"enabled,omitempty"`
	CreatedAt *float32 `json:"createdAt,omitempty"`
}

type _CreateHookRequest CreateHookRequest

// NewCreateHookRequest instantiates a new CreateHookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateHookRequest(config CreateHookRequestConfig) *CreateHookRequest {
	this := CreateHookRequest{}
	this.Config = config
	return &this
}

// NewCreateHookRequestWithDefaults instantiates a new CreateHookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateHookRequestWithDefaults() *CreateHookRequest {
	this := CreateHookRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CreateHookRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHookRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CreateHookRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CreateHookRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateHookRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHookRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateHookRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateHookRequest) SetName(v string) {
	o.Name = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
// Deprecated
func (o *CreateHookRequest) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateHookRequest) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *CreateHookRequest) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
// Deprecated
func (o *CreateHookRequest) SetEvent(v string) {
	o.Event = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CreateHookRequest) GetEvents() []string {
	if o == nil || IsNil(o.Events) {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHookRequest) GetEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *CreateHookRequest) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *CreateHookRequest) SetEvents(v []string) {
	o.Events = v
}

// GetConfig returns the Config field value
func (o *CreateHookRequest) GetConfig() CreateHookRequestConfig {
	if o == nil {
		var ret CreateHookRequestConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *CreateHookRequest) GetConfigOk() (*CreateHookRequestConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *CreateHookRequest) SetConfig(v CreateHookRequestConfig) {
	o.Config = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateHookRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHookRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateHookRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateHookRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateHookRequest) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateHookRequest) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateHookRequest) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *CreateHookRequest) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

func (o CreateHookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateHookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	toSerialize["config"] = o.Config
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

func (o *CreateHookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"config",
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

	varCreateHookRequest := _CreateHookRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateHookRequest)

	if err != nil {
		return err
	}

	*o = CreateHookRequest(varCreateHookRequest)

	return err
}

type NullableCreateHookRequest struct {
	value *CreateHookRequest
	isSet bool
}

func (v NullableCreateHookRequest) Get() *CreateHookRequest {
	return v.value
}

func (v *NullableCreateHookRequest) Set(val *CreateHookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateHookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateHookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateHookRequest(val *CreateHookRequest) *NullableCreateHookRequest {
	return &NullableCreateHookRequest{value: val, isSet: true}
}

func (v NullableCreateHookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateHookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
