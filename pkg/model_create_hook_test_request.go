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

// checks if the CreateHookTestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateHookTestRequest{}

// CreateHookTestRequest struct for CreateHookTestRequest
type CreateHookTestRequest struct {
	// An array of hook events for testing.
	Events []string `json:"events"`
	Config CreateHookTestRequestConfig `json:"config"`
	// Use `events` instead.
	// Deprecated
	Event interface{} `json:"event,omitempty"`
}

type _CreateHookTestRequest CreateHookTestRequest

// NewCreateHookTestRequest instantiates a new CreateHookTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateHookTestRequest(events []string, config CreateHookTestRequestConfig) *CreateHookTestRequest {
	this := CreateHookTestRequest{}
	this.Events = events
	this.Config = config
	return &this
}

// NewCreateHookTestRequestWithDefaults instantiates a new CreateHookTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateHookTestRequestWithDefaults() *CreateHookTestRequest {
	this := CreateHookTestRequest{}
	return &this
}

// GetEvents returns the Events field value
func (o *CreateHookTestRequest) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *CreateHookTestRequest) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *CreateHookTestRequest) SetEvents(v []string) {
	o.Events = v
}

// GetConfig returns the Config field value
func (o *CreateHookTestRequest) GetConfig() CreateHookTestRequestConfig {
	if o == nil {
		var ret CreateHookTestRequestConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *CreateHookTestRequest) GetConfigOk() (*CreateHookTestRequestConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *CreateHookTestRequest) SetConfig(v CreateHookTestRequestConfig) {
	o.Config = v
}

// GetEvent returns the Event field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *CreateHookTestRequest) GetEvent() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *CreateHookTestRequest) GetEventOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return &o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *CreateHookTestRequest) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given interface{} and assigns it to the Event field.
// Deprecated
func (o *CreateHookTestRequest) SetEvent(v interface{}) {
	o.Event = v
}

func (o CreateHookTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateHookTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	toSerialize["config"] = o.Config
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	return toSerialize, nil
}

func (o *CreateHookTestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
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

	varCreateHookTestRequest := _CreateHookTestRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateHookTestRequest)

	if err != nil {
		return err
	}

	*o = CreateHookTestRequest(varCreateHookTestRequest)

	return err
}

type NullableCreateHookTestRequest struct {
	value *CreateHookTestRequest
	isSet bool
}

func (v NullableCreateHookTestRequest) Get() *CreateHookTestRequest {
	return v.value
}

func (v *NullableCreateHookTestRequest) Set(val *CreateHookTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateHookTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateHookTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateHookTestRequest(val *CreateHookTestRequest) *NullableCreateHookTestRequest {
	return &NullableCreateHookTestRequest{value: val, isSet: true}
}

func (v NullableCreateHookTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateHookTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
