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

// checks if the UpdateSignInExpRequestColor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSignInExpRequestColor{}

// UpdateSignInExpRequestColor Specify the primary branding color for the sign-in page (both light/dark mode).
type UpdateSignInExpRequestColor struct {
	PrimaryColor string `json:"primaryColor"`
	IsDarkModeEnabled bool `json:"isDarkModeEnabled"`
	DarkPrimaryColor string `json:"darkPrimaryColor"`
}

type _UpdateSignInExpRequestColor UpdateSignInExpRequestColor

// NewUpdateSignInExpRequestColor instantiates a new UpdateSignInExpRequestColor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSignInExpRequestColor(primaryColor string, isDarkModeEnabled bool, darkPrimaryColor string) *UpdateSignInExpRequestColor {
	this := UpdateSignInExpRequestColor{}
	this.PrimaryColor = primaryColor
	this.IsDarkModeEnabled = isDarkModeEnabled
	this.DarkPrimaryColor = darkPrimaryColor
	return &this
}

// NewUpdateSignInExpRequestColorWithDefaults instantiates a new UpdateSignInExpRequestColor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSignInExpRequestColorWithDefaults() *UpdateSignInExpRequestColor {
	this := UpdateSignInExpRequestColor{}
	return &this
}

// GetPrimaryColor returns the PrimaryColor field value
func (o *UpdateSignInExpRequestColor) GetPrimaryColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryColor
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequestColor) GetPrimaryColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryColor, true
}

// SetPrimaryColor sets field value
func (o *UpdateSignInExpRequestColor) SetPrimaryColor(v string) {
	o.PrimaryColor = v
}

// GetIsDarkModeEnabled returns the IsDarkModeEnabled field value
func (o *UpdateSignInExpRequestColor) GetIsDarkModeEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDarkModeEnabled
}

// GetIsDarkModeEnabledOk returns a tuple with the IsDarkModeEnabled field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequestColor) GetIsDarkModeEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDarkModeEnabled, true
}

// SetIsDarkModeEnabled sets field value
func (o *UpdateSignInExpRequestColor) SetIsDarkModeEnabled(v bool) {
	o.IsDarkModeEnabled = v
}

// GetDarkPrimaryColor returns the DarkPrimaryColor field value
func (o *UpdateSignInExpRequestColor) GetDarkPrimaryColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DarkPrimaryColor
}

// GetDarkPrimaryColorOk returns a tuple with the DarkPrimaryColor field value
// and a boolean to check if the value has been set.
func (o *UpdateSignInExpRequestColor) GetDarkPrimaryColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DarkPrimaryColor, true
}

// SetDarkPrimaryColor sets field value
func (o *UpdateSignInExpRequestColor) SetDarkPrimaryColor(v string) {
	o.DarkPrimaryColor = v
}

func (o UpdateSignInExpRequestColor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSignInExpRequestColor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["primaryColor"] = o.PrimaryColor
	toSerialize["isDarkModeEnabled"] = o.IsDarkModeEnabled
	toSerialize["darkPrimaryColor"] = o.DarkPrimaryColor
	return toSerialize, nil
}

func (o *UpdateSignInExpRequestColor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"primaryColor",
		"isDarkModeEnabled",
		"darkPrimaryColor",
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

	varUpdateSignInExpRequestColor := _UpdateSignInExpRequestColor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateSignInExpRequestColor)

	if err != nil {
		return err
	}

	*o = UpdateSignInExpRequestColor(varUpdateSignInExpRequestColor)

	return err
}

type NullableUpdateSignInExpRequestColor struct {
	value *UpdateSignInExpRequestColor
	isSet bool
}

func (v NullableUpdateSignInExpRequestColor) Get() *UpdateSignInExpRequestColor {
	return v.value
}

func (v *NullableUpdateSignInExpRequestColor) Set(val *UpdateSignInExpRequestColor) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSignInExpRequestColor) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSignInExpRequestColor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSignInExpRequestColor(val *UpdateSignInExpRequestColor) *NullableUpdateSignInExpRequestColor {
	return &NullableUpdateSignInExpRequestColor{value: val, isSet: true}
}

func (v NullableUpdateSignInExpRequestColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSignInExpRequestColor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
