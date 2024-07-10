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

// checks if the GetActiveUserCounts200ResponseDauCurveInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetActiveUserCounts200ResponseDauCurveInner{}

// GetActiveUserCounts200ResponseDauCurveInner struct for GetActiveUserCounts200ResponseDauCurveInner
type GetActiveUserCounts200ResponseDauCurveInner struct {
	Date string `json:"date"`
	Count float32 `json:"count"`
}

type _GetActiveUserCounts200ResponseDauCurveInner GetActiveUserCounts200ResponseDauCurveInner

// NewGetActiveUserCounts200ResponseDauCurveInner instantiates a new GetActiveUserCounts200ResponseDauCurveInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetActiveUserCounts200ResponseDauCurveInner(date string, count float32) *GetActiveUserCounts200ResponseDauCurveInner {
	this := GetActiveUserCounts200ResponseDauCurveInner{}
	this.Date = date
	this.Count = count
	return &this
}

// NewGetActiveUserCounts200ResponseDauCurveInnerWithDefaults instantiates a new GetActiveUserCounts200ResponseDauCurveInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetActiveUserCounts200ResponseDauCurveInnerWithDefaults() *GetActiveUserCounts200ResponseDauCurveInner {
	this := GetActiveUserCounts200ResponseDauCurveInner{}
	return &this
}

// GetDate returns the Date field value
func (o *GetActiveUserCounts200ResponseDauCurveInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *GetActiveUserCounts200ResponseDauCurveInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *GetActiveUserCounts200ResponseDauCurveInner) SetDate(v string) {
	o.Date = v
}

// GetCount returns the Count field value
func (o *GetActiveUserCounts200ResponseDauCurveInner) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetActiveUserCounts200ResponseDauCurveInner) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetActiveUserCounts200ResponseDauCurveInner) SetCount(v float32) {
	o.Count = v
}

func (o GetActiveUserCounts200ResponseDauCurveInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetActiveUserCounts200ResponseDauCurveInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

func (o *GetActiveUserCounts200ResponseDauCurveInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"count",
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

	varGetActiveUserCounts200ResponseDauCurveInner := _GetActiveUserCounts200ResponseDauCurveInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetActiveUserCounts200ResponseDauCurveInner)

	if err != nil {
		return err
	}

	*o = GetActiveUserCounts200ResponseDauCurveInner(varGetActiveUserCounts200ResponseDauCurveInner)

	return err
}

type NullableGetActiveUserCounts200ResponseDauCurveInner struct {
	value *GetActiveUserCounts200ResponseDauCurveInner
	isSet bool
}

func (v NullableGetActiveUserCounts200ResponseDauCurveInner) Get() *GetActiveUserCounts200ResponseDauCurveInner {
	return v.value
}

func (v *NullableGetActiveUserCounts200ResponseDauCurveInner) Set(val *GetActiveUserCounts200ResponseDauCurveInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActiveUserCounts200ResponseDauCurveInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActiveUserCounts200ResponseDauCurveInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActiveUserCounts200ResponseDauCurveInner(val *GetActiveUserCounts200ResponseDauCurveInner) *NullableGetActiveUserCounts200ResponseDauCurveInner {
	return &NullableGetActiveUserCounts200ResponseDauCurveInner{value: val, isSet: true}
}

func (v NullableGetActiveUserCounts200ResponseDauCurveInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActiveUserCounts200ResponseDauCurveInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
