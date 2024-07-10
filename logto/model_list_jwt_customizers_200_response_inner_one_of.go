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

// checks if the ListJwtCustomizers200ResponseInnerOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListJwtCustomizers200ResponseInnerOneOf{}

// ListJwtCustomizers200ResponseInnerOneOf struct for ListJwtCustomizers200ResponseInnerOneOf
type ListJwtCustomizers200ResponseInnerOneOf struct {
	Key string `json:"key"`
	Value GetJwtCustomizer200ResponseOneOf `json:"value"`
}

type _ListJwtCustomizers200ResponseInnerOneOf ListJwtCustomizers200ResponseInnerOneOf

// NewListJwtCustomizers200ResponseInnerOneOf instantiates a new ListJwtCustomizers200ResponseInnerOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListJwtCustomizers200ResponseInnerOneOf(key string, value GetJwtCustomizer200ResponseOneOf) *ListJwtCustomizers200ResponseInnerOneOf {
	this := ListJwtCustomizers200ResponseInnerOneOf{}
	this.Key = key
	this.Value = value
	return &this
}

// NewListJwtCustomizers200ResponseInnerOneOfWithDefaults instantiates a new ListJwtCustomizers200ResponseInnerOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListJwtCustomizers200ResponseInnerOneOfWithDefaults() *ListJwtCustomizers200ResponseInnerOneOf {
	this := ListJwtCustomizers200ResponseInnerOneOf{}
	return &this
}

// GetKey returns the Key field value
func (o *ListJwtCustomizers200ResponseInnerOneOf) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ListJwtCustomizers200ResponseInnerOneOf) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ListJwtCustomizers200ResponseInnerOneOf) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *ListJwtCustomizers200ResponseInnerOneOf) GetValue() GetJwtCustomizer200ResponseOneOf {
	if o == nil {
		var ret GetJwtCustomizer200ResponseOneOf
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ListJwtCustomizers200ResponseInnerOneOf) GetValueOk() (*GetJwtCustomizer200ResponseOneOf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ListJwtCustomizers200ResponseInnerOneOf) SetValue(v GetJwtCustomizer200ResponseOneOf) {
	o.Value = v
}

func (o ListJwtCustomizers200ResponseInnerOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListJwtCustomizers200ResponseInnerOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ListJwtCustomizers200ResponseInnerOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
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

	varListJwtCustomizers200ResponseInnerOneOf := _ListJwtCustomizers200ResponseInnerOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListJwtCustomizers200ResponseInnerOneOf)

	if err != nil {
		return err
	}

	*o = ListJwtCustomizers200ResponseInnerOneOf(varListJwtCustomizers200ResponseInnerOneOf)

	return err
}

type NullableListJwtCustomizers200ResponseInnerOneOf struct {
	value *ListJwtCustomizers200ResponseInnerOneOf
	isSet bool
}

func (v NullableListJwtCustomizers200ResponseInnerOneOf) Get() *ListJwtCustomizers200ResponseInnerOneOf {
	return v.value
}

func (v *NullableListJwtCustomizers200ResponseInnerOneOf) Set(val *ListJwtCustomizers200ResponseInnerOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListJwtCustomizers200ResponseInnerOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListJwtCustomizers200ResponseInnerOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListJwtCustomizers200ResponseInnerOneOf(val *ListJwtCustomizers200ResponseInnerOneOf) *NullableListJwtCustomizers200ResponseInnerOneOf {
	return &NullableListJwtCustomizers200ResponseInnerOneOf{value: val, isSet: true}
}

func (v NullableListJwtCustomizers200ResponseInnerOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListJwtCustomizers200ResponseInnerOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
