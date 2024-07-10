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

// checks if the GetJwtCustomizer200ResponseOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetJwtCustomizer200ResponseOneOf{}

// GetJwtCustomizer200ResponseOneOf struct for GetJwtCustomizer200ResponseOneOf
type GetJwtCustomizer200ResponseOneOf struct {
	Script string `json:"script"`
	EnvironmentVariables *map[string]string `json:"environmentVariables,omitempty"`
	ContextSample *GetJwtCustomizer200ResponseOneOfContextSample `json:"contextSample,omitempty"`
	TokenSample *GetJwtCustomizer200ResponseOneOfTokenSample `json:"tokenSample,omitempty"`
}

type _GetJwtCustomizer200ResponseOneOf GetJwtCustomizer200ResponseOneOf

// NewGetJwtCustomizer200ResponseOneOf instantiates a new GetJwtCustomizer200ResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetJwtCustomizer200ResponseOneOf(script string) *GetJwtCustomizer200ResponseOneOf {
	this := GetJwtCustomizer200ResponseOneOf{}
	this.Script = script
	return &this
}

// NewGetJwtCustomizer200ResponseOneOfWithDefaults instantiates a new GetJwtCustomizer200ResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetJwtCustomizer200ResponseOneOfWithDefaults() *GetJwtCustomizer200ResponseOneOf {
	this := GetJwtCustomizer200ResponseOneOf{}
	return &this
}

// GetScript returns the Script field value
func (o *GetJwtCustomizer200ResponseOneOf) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOf) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *GetJwtCustomizer200ResponseOneOf) SetScript(v string) {
	o.Script = v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOf) GetEnvironmentVariables() map[string]string {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret map[string]string
		return ret
	}
	return *o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOf) GetEnvironmentVariablesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOf) HasEnvironmentVariables() bool {
	if o != nil && !IsNil(o.EnvironmentVariables) {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given map[string]string and assigns it to the EnvironmentVariables field.
func (o *GetJwtCustomizer200ResponseOneOf) SetEnvironmentVariables(v map[string]string) {
	o.EnvironmentVariables = &v
}

// GetContextSample returns the ContextSample field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOf) GetContextSample() GetJwtCustomizer200ResponseOneOfContextSample {
	if o == nil || IsNil(o.ContextSample) {
		var ret GetJwtCustomizer200ResponseOneOfContextSample
		return ret
	}
	return *o.ContextSample
}

// GetContextSampleOk returns a tuple with the ContextSample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOf) GetContextSampleOk() (*GetJwtCustomizer200ResponseOneOfContextSample, bool) {
	if o == nil || IsNil(o.ContextSample) {
		return nil, false
	}
	return o.ContextSample, true
}

// HasContextSample returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOf) HasContextSample() bool {
	if o != nil && !IsNil(o.ContextSample) {
		return true
	}

	return false
}

// SetContextSample gets a reference to the given GetJwtCustomizer200ResponseOneOfContextSample and assigns it to the ContextSample field.
func (o *GetJwtCustomizer200ResponseOneOf) SetContextSample(v GetJwtCustomizer200ResponseOneOfContextSample) {
	o.ContextSample = &v
}

// GetTokenSample returns the TokenSample field value if set, zero value otherwise.
func (o *GetJwtCustomizer200ResponseOneOf) GetTokenSample() GetJwtCustomizer200ResponseOneOfTokenSample {
	if o == nil || IsNil(o.TokenSample) {
		var ret GetJwtCustomizer200ResponseOneOfTokenSample
		return ret
	}
	return *o.TokenSample
}

// GetTokenSampleOk returns a tuple with the TokenSample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJwtCustomizer200ResponseOneOf) GetTokenSampleOk() (*GetJwtCustomizer200ResponseOneOfTokenSample, bool) {
	if o == nil || IsNil(o.TokenSample) {
		return nil, false
	}
	return o.TokenSample, true
}

// HasTokenSample returns a boolean if a field has been set.
func (o *GetJwtCustomizer200ResponseOneOf) HasTokenSample() bool {
	if o != nil && !IsNil(o.TokenSample) {
		return true
	}

	return false
}

// SetTokenSample gets a reference to the given GetJwtCustomizer200ResponseOneOfTokenSample and assigns it to the TokenSample field.
func (o *GetJwtCustomizer200ResponseOneOf) SetTokenSample(v GetJwtCustomizer200ResponseOneOfTokenSample) {
	o.TokenSample = &v
}

func (o GetJwtCustomizer200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetJwtCustomizer200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["script"] = o.Script
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	if !IsNil(o.ContextSample) {
		toSerialize["contextSample"] = o.ContextSample
	}
	if !IsNil(o.TokenSample) {
		toSerialize["tokenSample"] = o.TokenSample
	}
	return toSerialize, nil
}

func (o *GetJwtCustomizer200ResponseOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"script",
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

	varGetJwtCustomizer200ResponseOneOf := _GetJwtCustomizer200ResponseOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetJwtCustomizer200ResponseOneOf)

	if err != nil {
		return err
	}

	*o = GetJwtCustomizer200ResponseOneOf(varGetJwtCustomizer200ResponseOneOf)

	return err
}

type NullableGetJwtCustomizer200ResponseOneOf struct {
	value *GetJwtCustomizer200ResponseOneOf
	isSet bool
}

func (v NullableGetJwtCustomizer200ResponseOneOf) Get() *GetJwtCustomizer200ResponseOneOf {
	return v.value
}

func (v *NullableGetJwtCustomizer200ResponseOneOf) Set(val *GetJwtCustomizer200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJwtCustomizer200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJwtCustomizer200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJwtCustomizer200ResponseOneOf(val *GetJwtCustomizer200ResponseOneOf) *NullableGetJwtCustomizer200ResponseOneOf {
	return &NullableGetJwtCustomizer200ResponseOneOf{value: val, isSet: true}
}

func (v NullableGetJwtCustomizer200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJwtCustomizer200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
