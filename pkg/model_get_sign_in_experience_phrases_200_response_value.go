/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"encoding/json"
	"fmt"
)

// GetSignInExperiencePhrases200ResponseValue - struct for GetSignInExperiencePhrases200ResponseValue
type GetSignInExperiencePhrases200ResponseValue struct {
	MapmapOfStringAny *map[string]interface{}
	String *string
}

// map[string]interface{}AsGetSignInExperiencePhrases200ResponseValue is a convenience function that returns map[string]interface{} wrapped in GetSignInExperiencePhrases200ResponseValue
func MapmapOfStringAnyAsGetSignInExperiencePhrases200ResponseValue(v *map[string]interface{}) GetSignInExperiencePhrases200ResponseValue {
	return GetSignInExperiencePhrases200ResponseValue{
		MapmapOfStringAny: v,
	}
}

// stringAsGetSignInExperiencePhrases200ResponseValue is a convenience function that returns string wrapped in GetSignInExperiencePhrases200ResponseValue
func StringAsGetSignInExperiencePhrases200ResponseValue(v *string) GetSignInExperiencePhrases200ResponseValue {
	return GetSignInExperiencePhrases200ResponseValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetSignInExperiencePhrases200ResponseValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			match++
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MapmapOfStringAny = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetSignInExperiencePhrases200ResponseValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetSignInExperiencePhrases200ResponseValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSignInExperiencePhrases200ResponseValue) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSignInExperiencePhrases200ResponseValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableGetSignInExperiencePhrases200ResponseValue struct {
	value *GetSignInExperiencePhrases200ResponseValue
	isSet bool
}

func (v NullableGetSignInExperiencePhrases200ResponseValue) Get() *GetSignInExperiencePhrases200ResponseValue {
	return v.value
}

func (v *NullableGetSignInExperiencePhrases200ResponseValue) Set(val *GetSignInExperiencePhrases200ResponseValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignInExperiencePhrases200ResponseValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignInExperiencePhrases200ResponseValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignInExperiencePhrases200ResponseValue(val *GetSignInExperiencePhrases200ResponseValue) *NullableGetSignInExperiencePhrases200ResponseValue {
	return &NullableGetSignInExperiencePhrases200ResponseValue{value: val, isSet: true}
}

func (v NullableGetSignInExperiencePhrases200ResponseValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignInExperiencePhrases200ResponseValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


