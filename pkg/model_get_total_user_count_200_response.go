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

// checks if the GetTotalUserCount200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTotalUserCount200Response{}

// GetTotalUserCount200Response struct for GetTotalUserCount200Response
type GetTotalUserCount200Response struct {
	TotalUserCount float32 `json:"totalUserCount"`
}

type _GetTotalUserCount200Response GetTotalUserCount200Response

// NewGetTotalUserCount200Response instantiates a new GetTotalUserCount200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTotalUserCount200Response(totalUserCount float32) *GetTotalUserCount200Response {
	this := GetTotalUserCount200Response{}
	this.TotalUserCount = totalUserCount
	return &this
}

// NewGetTotalUserCount200ResponseWithDefaults instantiates a new GetTotalUserCount200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTotalUserCount200ResponseWithDefaults() *GetTotalUserCount200Response {
	this := GetTotalUserCount200Response{}
	return &this
}

// GetTotalUserCount returns the TotalUserCount field value
func (o *GetTotalUserCount200Response) GetTotalUserCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalUserCount
}

// GetTotalUserCountOk returns a tuple with the TotalUserCount field value
// and a boolean to check if the value has been set.
func (o *GetTotalUserCount200Response) GetTotalUserCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalUserCount, true
}

// SetTotalUserCount sets field value
func (o *GetTotalUserCount200Response) SetTotalUserCount(v float32) {
	o.TotalUserCount = v
}

func (o GetTotalUserCount200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTotalUserCount200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["totalUserCount"] = o.TotalUserCount
	return toSerialize, nil
}

func (o *GetTotalUserCount200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"totalUserCount",
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

	varGetTotalUserCount200Response := _GetTotalUserCount200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTotalUserCount200Response)

	if err != nil {
		return err
	}

	*o = GetTotalUserCount200Response(varGetTotalUserCount200Response)

	return err
}

type NullableGetTotalUserCount200Response struct {
	value *GetTotalUserCount200Response
	isSet bool
}

func (v NullableGetTotalUserCount200Response) Get() *GetTotalUserCount200Response {
	return v.value
}

func (v *NullableGetTotalUserCount200Response) Set(val *GetTotalUserCount200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTotalUserCount200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTotalUserCount200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTotalUserCount200Response(val *GetTotalUserCount200Response) *NullableGetTotalUserCount200Response {
	return &NullableGetTotalUserCount200Response{value: val, isSet: true}
}

func (v NullableGetTotalUserCount200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTotalUserCount200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

