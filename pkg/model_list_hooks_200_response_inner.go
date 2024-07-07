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

// checks if the ListHooks200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHooks200ResponseInner{}

// ListHooks200ResponseInner struct for ListHooks200ResponseInner
type ListHooks200ResponseInner struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	Name string `json:"name"`
	Event NullableString `json:"event"`
	Events []string `json:"events"`
	Config ListHooks200ResponseInnerConfig `json:"config"`
	SigningKey string `json:"signingKey"`
	Enabled bool `json:"enabled"`
	CreatedAt float32 `json:"createdAt"`
	ExecutionStats *ListHooks200ResponseInnerExecutionStats `json:"executionStats,omitempty"`
}

type _ListHooks200ResponseInner ListHooks200ResponseInner

// NewListHooks200ResponseInner instantiates a new ListHooks200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHooks200ResponseInner(tenantId string, id string, name string, event NullableString, events []string, config ListHooks200ResponseInnerConfig, signingKey string, enabled bool, createdAt float32) *ListHooks200ResponseInner {
	this := ListHooks200ResponseInner{}
	this.TenantId = tenantId
	this.Id = id
	this.Name = name
	this.Event = event
	this.Events = events
	this.Config = config
	this.SigningKey = signingKey
	this.Enabled = enabled
	this.CreatedAt = createdAt
	return &this
}

// NewListHooks200ResponseInnerWithDefaults instantiates a new ListHooks200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHooks200ResponseInnerWithDefaults() *ListHooks200ResponseInner {
	this := ListHooks200ResponseInner{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListHooks200ResponseInner) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListHooks200ResponseInner) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListHooks200ResponseInner) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *ListHooks200ResponseInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListHooks200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListHooks200ResponseInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ListHooks200ResponseInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListHooks200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListHooks200ResponseInner) SetName(v string) {
	o.Name = v
}

// GetEvent returns the Event field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListHooks200ResponseInner) GetEvent() string {
	if o == nil || o.Event.Get() == nil {
		var ret string
		return ret
	}

	return *o.Event.Get()
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListHooks200ResponseInner) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Event.Get(), o.Event.IsSet()
}

// SetEvent sets field value
func (o *ListHooks200ResponseInner) SetEvent(v string) {
	o.Event.Set(&v)
}

// GetEvents returns the Events field value
func (o *ListHooks200ResponseInner) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *ListHooks200ResponseInner) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *ListHooks200ResponseInner) SetEvents(v []string) {
	o.Events = v
}

// GetConfig returns the Config field value
func (o *ListHooks200ResponseInner) GetConfig() ListHooks200ResponseInnerConfig {
	if o == nil {
		var ret ListHooks200ResponseInnerConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ListHooks200ResponseInner) GetConfigOk() (*ListHooks200ResponseInnerConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ListHooks200ResponseInner) SetConfig(v ListHooks200ResponseInnerConfig) {
	o.Config = v
}

// GetSigningKey returns the SigningKey field value
func (o *ListHooks200ResponseInner) GetSigningKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SigningKey
}

// GetSigningKeyOk returns a tuple with the SigningKey field value
// and a boolean to check if the value has been set.
func (o *ListHooks200ResponseInner) GetSigningKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningKey, true
}

// SetSigningKey sets field value
func (o *ListHooks200ResponseInner) SetSigningKey(v string) {
	o.SigningKey = v
}

// GetEnabled returns the Enabled field value
func (o *ListHooks200ResponseInner) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ListHooks200ResponseInner) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ListHooks200ResponseInner) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ListHooks200ResponseInner) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ListHooks200ResponseInner) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ListHooks200ResponseInner) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetExecutionStats returns the ExecutionStats field value if set, zero value otherwise.
func (o *ListHooks200ResponseInner) GetExecutionStats() ListHooks200ResponseInnerExecutionStats {
	if o == nil || IsNil(o.ExecutionStats) {
		var ret ListHooks200ResponseInnerExecutionStats
		return ret
	}
	return *o.ExecutionStats
}

// GetExecutionStatsOk returns a tuple with the ExecutionStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHooks200ResponseInner) GetExecutionStatsOk() (*ListHooks200ResponseInnerExecutionStats, bool) {
	if o == nil || IsNil(o.ExecutionStats) {
		return nil, false
	}
	return o.ExecutionStats, true
}

// HasExecutionStats returns a boolean if a field has been set.
func (o *ListHooks200ResponseInner) HasExecutionStats() bool {
	if o != nil && !IsNil(o.ExecutionStats) {
		return true
	}

	return false
}

// SetExecutionStats gets a reference to the given ListHooks200ResponseInnerExecutionStats and assigns it to the ExecutionStats field.
func (o *ListHooks200ResponseInner) SetExecutionStats(v ListHooks200ResponseInnerExecutionStats) {
	o.ExecutionStats = &v
}

func (o ListHooks200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHooks200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["event"] = o.Event.Get()
	toSerialize["events"] = o.Events
	toSerialize["config"] = o.Config
	toSerialize["signingKey"] = o.SigningKey
	toSerialize["enabled"] = o.Enabled
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.ExecutionStats) {
		toSerialize["executionStats"] = o.ExecutionStats
	}
	return toSerialize, nil
}

func (o *ListHooks200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"name",
		"event",
		"events",
		"config",
		"signingKey",
		"enabled",
		"createdAt",
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

	varListHooks200ResponseInner := _ListHooks200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListHooks200ResponseInner)

	if err != nil {
		return err
	}

	*o = ListHooks200ResponseInner(varListHooks200ResponseInner)

	return err
}

type NullableListHooks200ResponseInner struct {
	value *ListHooks200ResponseInner
	isSet bool
}

func (v NullableListHooks200ResponseInner) Get() *ListHooks200ResponseInner {
	return v.value
}

func (v *NullableListHooks200ResponseInner) Set(val *ListHooks200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListHooks200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListHooks200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHooks200ResponseInner(val *ListHooks200ResponseInner) *NullableListHooks200ResponseInner {
	return &NullableListHooks200ResponseInner{value: val, isSet: true}
}

func (v NullableListHooks200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHooks200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

