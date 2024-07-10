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

// checks if the ListHookRecentLogs200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHookRecentLogs200ResponseInner{}

// ListHookRecentLogs200ResponseInner struct for ListHookRecentLogs200ResponseInner
type ListHookRecentLogs200ResponseInner struct {
	Id string `json:"id"`
	Key string `json:"key"`
	Payload ListLogs200ResponseInnerPayload `json:"payload"`
	CreatedAt float32 `json:"createdAt"`
}

type _ListHookRecentLogs200ResponseInner ListHookRecentLogs200ResponseInner

// NewListHookRecentLogs200ResponseInner instantiates a new ListHookRecentLogs200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHookRecentLogs200ResponseInner(id string, key string, payload ListLogs200ResponseInnerPayload, createdAt float32) *ListHookRecentLogs200ResponseInner {
	this := ListHookRecentLogs200ResponseInner{}
	this.Id = id
	this.Key = key
	this.Payload = payload
	this.CreatedAt = createdAt
	return &this
}

// NewListHookRecentLogs200ResponseInnerWithDefaults instantiates a new ListHookRecentLogs200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHookRecentLogs200ResponseInnerWithDefaults() *ListHookRecentLogs200ResponseInner {
	this := ListHookRecentLogs200ResponseInner{}
	return &this
}

// GetId returns the Id field value
func (o *ListHookRecentLogs200ResponseInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListHookRecentLogs200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListHookRecentLogs200ResponseInner) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *ListHookRecentLogs200ResponseInner) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ListHookRecentLogs200ResponseInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ListHookRecentLogs200ResponseInner) SetKey(v string) {
	o.Key = v
}

// GetPayload returns the Payload field value
func (o *ListHookRecentLogs200ResponseInner) GetPayload() ListLogs200ResponseInnerPayload {
	if o == nil {
		var ret ListLogs200ResponseInnerPayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *ListHookRecentLogs200ResponseInner) GetPayloadOk() (*ListLogs200ResponseInnerPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *ListHookRecentLogs200ResponseInner) SetPayload(v ListLogs200ResponseInnerPayload) {
	o.Payload = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ListHookRecentLogs200ResponseInner) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ListHookRecentLogs200ResponseInner) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ListHookRecentLogs200ResponseInner) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

func (o ListHookRecentLogs200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHookRecentLogs200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["payload"] = o.Payload
	toSerialize["createdAt"] = o.CreatedAt
	return toSerialize, nil
}

func (o *ListHookRecentLogs200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"key",
		"payload",
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

	varListHookRecentLogs200ResponseInner := _ListHookRecentLogs200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListHookRecentLogs200ResponseInner)

	if err != nil {
		return err
	}

	*o = ListHookRecentLogs200ResponseInner(varListHookRecentLogs200ResponseInner)

	return err
}

type NullableListHookRecentLogs200ResponseInner struct {
	value *ListHookRecentLogs200ResponseInner
	isSet bool
}

func (v NullableListHookRecentLogs200ResponseInner) Get() *ListHookRecentLogs200ResponseInner {
	return v.value
}

func (v *NullableListHookRecentLogs200ResponseInner) Set(val *ListHookRecentLogs200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListHookRecentLogs200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListHookRecentLogs200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHookRecentLogs200ResponseInner(val *ListHookRecentLogs200ResponseInner) *NullableListHookRecentLogs200ResponseInner {
	return &NullableListHookRecentLogs200ResponseInner{value: val, isSet: true}
}

func (v NullableListHookRecentLogs200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHookRecentLogs200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
