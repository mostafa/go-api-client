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

// checks if the GetJwtCustomizer200ResponseOneOfContextSample type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetJwtCustomizer200ResponseOneOfContextSample{}

// GetJwtCustomizer200ResponseOneOfContextSample struct for GetJwtCustomizer200ResponseOneOfContextSample
type GetJwtCustomizer200ResponseOneOfContextSample struct {
	User GetJwtCustomizer200ResponseOneOfContextSampleUser `json:"user"`
}

type _GetJwtCustomizer200ResponseOneOfContextSample GetJwtCustomizer200ResponseOneOfContextSample

// NewGetJwtCustomizer200ResponseOneOfContextSample instantiates a new GetJwtCustomizer200ResponseOneOfContextSample object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetJwtCustomizer200ResponseOneOfContextSample(user GetJwtCustomizer200ResponseOneOfContextSampleUser) *GetJwtCustomizer200ResponseOneOfContextSample {
	this := GetJwtCustomizer200ResponseOneOfContextSample{}
	this.User = user
	return &this
}

// NewGetJwtCustomizer200ResponseOneOfContextSampleWithDefaults instantiates a new GetJwtCustomizer200ResponseOneOfContextSample object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetJwtCustomizer200ResponseOneOfContextSampleWithDefaults() *GetJwtCustomizer200ResponseOneOfContextSample {
	this := GetJwtCustomizer200ResponseOneOfContextSample{}
	return &this
}

// GetUser returns the User field value
func (o *GetJwtCustomizer200ResponseOneOfContextSample) GetUser() GetJwtCustomizer200ResponseOneOfContextSampleUser {
	if o == nil {
		var ret GetJwtCustomizer200ResponseOneOfContextSampleUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOfContextSample) GetUserOk() (*GetJwtCustomizer200ResponseOneOfContextSampleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GetJwtCustomizer200ResponseOneOfContextSample) SetUser(v GetJwtCustomizer200ResponseOneOfContextSampleUser) {
	o.User = v
}

func (o GetJwtCustomizer200ResponseOneOfContextSample) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetJwtCustomizer200ResponseOneOfContextSample) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *GetJwtCustomizer200ResponseOneOfContextSample) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
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

	varGetJwtCustomizer200ResponseOneOfContextSample := _GetJwtCustomizer200ResponseOneOfContextSample{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetJwtCustomizer200ResponseOneOfContextSample)

	if err != nil {
		return err
	}

	*o = GetJwtCustomizer200ResponseOneOfContextSample(varGetJwtCustomizer200ResponseOneOfContextSample)

	return err
}

type NullableGetJwtCustomizer200ResponseOneOfContextSample struct {
	value *GetJwtCustomizer200ResponseOneOfContextSample
	isSet bool
}

func (v NullableGetJwtCustomizer200ResponseOneOfContextSample) Get() *GetJwtCustomizer200ResponseOneOfContextSample {
	return v.value
}

func (v *NullableGetJwtCustomizer200ResponseOneOfContextSample) Set(val *GetJwtCustomizer200ResponseOneOfContextSample) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJwtCustomizer200ResponseOneOfContextSample) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJwtCustomizer200ResponseOneOfContextSample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJwtCustomizer200ResponseOneOfContextSample(val *GetJwtCustomizer200ResponseOneOfContextSample) *NullableGetJwtCustomizer200ResponseOneOfContextSample {
	return &NullableGetJwtCustomizer200ResponseOneOfContextSample{value: val, isSet: true}
}

func (v NullableGetJwtCustomizer200ResponseOneOfContextSample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJwtCustomizer200ResponseOneOfContextSample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

