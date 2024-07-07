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

// ApiInteractionMfaPutRequest - struct for ApiInteractionMfaPutRequest
type ApiInteractionMfaPutRequest struct {
	ApiInteractionBindMfaPostRequestOneOf *ApiInteractionBindMfaPostRequestOneOf
	ApiInteractionMfaPutRequestOneOf *ApiInteractionMfaPutRequestOneOf
	ApiInteractionMfaPutRequestOneOf1 *ApiInteractionMfaPutRequestOneOf1
}

// ApiInteractionBindMfaPostRequestOneOfAsApiInteractionMfaPutRequest is a convenience function that returns ApiInteractionBindMfaPostRequestOneOf wrapped in ApiInteractionMfaPutRequest
func ApiInteractionBindMfaPostRequestOneOfAsApiInteractionMfaPutRequest(v *ApiInteractionBindMfaPostRequestOneOf) ApiInteractionMfaPutRequest {
	return ApiInteractionMfaPutRequest{
		ApiInteractionBindMfaPostRequestOneOf: v,
	}
}

// ApiInteractionMfaPutRequestOneOfAsApiInteractionMfaPutRequest is a convenience function that returns ApiInteractionMfaPutRequestOneOf wrapped in ApiInteractionMfaPutRequest
func ApiInteractionMfaPutRequestOneOfAsApiInteractionMfaPutRequest(v *ApiInteractionMfaPutRequestOneOf) ApiInteractionMfaPutRequest {
	return ApiInteractionMfaPutRequest{
		ApiInteractionMfaPutRequestOneOf: v,
	}
}

// ApiInteractionMfaPutRequestOneOf1AsApiInteractionMfaPutRequest is a convenience function that returns ApiInteractionMfaPutRequestOneOf1 wrapped in ApiInteractionMfaPutRequest
func ApiInteractionMfaPutRequestOneOf1AsApiInteractionMfaPutRequest(v *ApiInteractionMfaPutRequestOneOf1) ApiInteractionMfaPutRequest {
	return ApiInteractionMfaPutRequest{
		ApiInteractionMfaPutRequestOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApiInteractionMfaPutRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApiInteractionBindMfaPostRequestOneOf
	err = newStrictDecoder(data).Decode(&dst.ApiInteractionBindMfaPostRequestOneOf)
	if err == nil {
		jsonApiInteractionBindMfaPostRequestOneOf, _ := json.Marshal(dst.ApiInteractionBindMfaPostRequestOneOf)
		if string(jsonApiInteractionBindMfaPostRequestOneOf) == "{}" { // empty struct
			dst.ApiInteractionBindMfaPostRequestOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ApiInteractionBindMfaPostRequestOneOf = nil
	}

	// try to unmarshal data into ApiInteractionMfaPutRequestOneOf
	err = newStrictDecoder(data).Decode(&dst.ApiInteractionMfaPutRequestOneOf)
	if err == nil {
		jsonApiInteractionMfaPutRequestOneOf, _ := json.Marshal(dst.ApiInteractionMfaPutRequestOneOf)
		if string(jsonApiInteractionMfaPutRequestOneOf) == "{}" { // empty struct
			dst.ApiInteractionMfaPutRequestOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ApiInteractionMfaPutRequestOneOf = nil
	}

	// try to unmarshal data into ApiInteractionMfaPutRequestOneOf1
	err = newStrictDecoder(data).Decode(&dst.ApiInteractionMfaPutRequestOneOf1)
	if err == nil {
		jsonApiInteractionMfaPutRequestOneOf1, _ := json.Marshal(dst.ApiInteractionMfaPutRequestOneOf1)
		if string(jsonApiInteractionMfaPutRequestOneOf1) == "{}" { // empty struct
			dst.ApiInteractionMfaPutRequestOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.ApiInteractionMfaPutRequestOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApiInteractionBindMfaPostRequestOneOf = nil
		dst.ApiInteractionMfaPutRequestOneOf = nil
		dst.ApiInteractionMfaPutRequestOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApiInteractionMfaPutRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApiInteractionMfaPutRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApiInteractionMfaPutRequest) MarshalJSON() ([]byte, error) {
	if src.ApiInteractionBindMfaPostRequestOneOf != nil {
		return json.Marshal(&src.ApiInteractionBindMfaPostRequestOneOf)
	}

	if src.ApiInteractionMfaPutRequestOneOf != nil {
		return json.Marshal(&src.ApiInteractionMfaPutRequestOneOf)
	}

	if src.ApiInteractionMfaPutRequestOneOf1 != nil {
		return json.Marshal(&src.ApiInteractionMfaPutRequestOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApiInteractionMfaPutRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApiInteractionBindMfaPostRequestOneOf != nil {
		return obj.ApiInteractionBindMfaPostRequestOneOf
	}

	if obj.ApiInteractionMfaPutRequestOneOf != nil {
		return obj.ApiInteractionMfaPutRequestOneOf
	}

	if obj.ApiInteractionMfaPutRequestOneOf1 != nil {
		return obj.ApiInteractionMfaPutRequestOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableApiInteractionMfaPutRequest struct {
	value *ApiInteractionMfaPutRequest
	isSet bool
}

func (v NullableApiInteractionMfaPutRequest) Get() *ApiInteractionMfaPutRequest {
	return v.value
}

func (v *NullableApiInteractionMfaPutRequest) Set(val *ApiInteractionMfaPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionMfaPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionMfaPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionMfaPutRequest(val *ApiInteractionMfaPutRequest) *NullableApiInteractionMfaPutRequest {
	return &NullableApiInteractionMfaPutRequest{value: val, isSet: true}
}

func (v NullableApiInteractionMfaPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionMfaPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


