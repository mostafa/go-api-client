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

// checks if the ListLogs200ResponseInnerPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListLogs200ResponseInnerPayload{}

// ListLogs200ResponseInnerPayload struct for ListLogs200ResponseInnerPayload
type ListLogs200ResponseInnerPayload struct {
	Key string `json:"key"`
	Result string `json:"result"`
	Error *ListLogs200ResponseInnerPayloadError `json:"error,omitempty"`
	Ip *string `json:"ip,omitempty"`
	UserAgent *string `json:"userAgent,omitempty"`
	UserId *string `json:"userId,omitempty"`
	ApplicationId *string `json:"applicationId,omitempty"`
	SessionId *string `json:"sessionId,omitempty"`
}

type _ListLogs200ResponseInnerPayload ListLogs200ResponseInnerPayload

// NewListLogs200ResponseInnerPayload instantiates a new ListLogs200ResponseInnerPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLogs200ResponseInnerPayload(key string, result string) *ListLogs200ResponseInnerPayload {
	this := ListLogs200ResponseInnerPayload{}
	this.Key = key
	this.Result = result
	return &this
}

// NewListLogs200ResponseInnerPayloadWithDefaults instantiates a new ListLogs200ResponseInnerPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLogs200ResponseInnerPayloadWithDefaults() *ListLogs200ResponseInnerPayload {
	this := ListLogs200ResponseInnerPayload{}
	return &this
}

// GetKey returns the Key field value
func (o *ListLogs200ResponseInnerPayload) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ListLogs200ResponseInnerPayload) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ListLogs200ResponseInnerPayload) SetKey(v string) {
	o.Key = v
}

// GetResult returns the Result field value
func (o *ListLogs200ResponseInnerPayload) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ListLogs200ResponseInnerPayload) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ListLogs200ResponseInnerPayload) SetResult(v string) {
	o.Result = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ListLogs200ResponseInnerPayload) GetError() ListLogs200ResponseInnerPayloadError {
	if o == nil || IsNil(o.Error) {
		var ret ListLogs200ResponseInnerPayloadError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogs200ResponseInnerPayload) GetErrorOk() (*ListLogs200ResponseInnerPayloadError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ListLogs200ResponseInnerPayload) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ListLogs200ResponseInnerPayloadError and assigns it to the Error field.
func (o *ListLogs200ResponseInnerPayload) SetError(v ListLogs200ResponseInnerPayloadError) {
	o.Error = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *ListLogs200ResponseInnerPayload) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogs200ResponseInnerPayload) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ListLogs200ResponseInnerPayload) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *ListLogs200ResponseInnerPayload) SetIp(v string) {
	o.Ip = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *ListLogs200ResponseInnerPayload) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogs200ResponseInnerPayload) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *ListLogs200ResponseInnerPayload) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *ListLogs200ResponseInnerPayload) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ListLogs200ResponseInnerPayload) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogs200ResponseInnerPayload) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ListLogs200ResponseInnerPayload) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ListLogs200ResponseInnerPayload) SetUserId(v string) {
	o.UserId = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ListLogs200ResponseInnerPayload) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogs200ResponseInnerPayload) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ListLogs200ResponseInnerPayload) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ListLogs200ResponseInnerPayload) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *ListLogs200ResponseInnerPayload) GetSessionId() string {
	if o == nil || IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogs200ResponseInnerPayload) GetSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *ListLogs200ResponseInnerPayload) HasSessionId() bool {
	if o != nil && !IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *ListLogs200ResponseInnerPayload) SetSessionId(v string) {
	o.SessionId = &v
}

func (o ListLogs200ResponseInnerPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListLogs200ResponseInnerPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["result"] = o.Result
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.SessionId) {
		toSerialize["sessionId"] = o.SessionId
	}
	return toSerialize, nil
}

func (o *ListLogs200ResponseInnerPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"result",
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

	varListLogs200ResponseInnerPayload := _ListLogs200ResponseInnerPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListLogs200ResponseInnerPayload)

	if err != nil {
		return err
	}

	*o = ListLogs200ResponseInnerPayload(varListLogs200ResponseInnerPayload)

	return err
}

type NullableListLogs200ResponseInnerPayload struct {
	value *ListLogs200ResponseInnerPayload
	isSet bool
}

func (v NullableListLogs200ResponseInnerPayload) Get() *ListLogs200ResponseInnerPayload {
	return v.value
}

func (v *NullableListLogs200ResponseInnerPayload) Set(val *ListLogs200ResponseInnerPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableListLogs200ResponseInnerPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableListLogs200ResponseInnerPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLogs200ResponseInnerPayload(val *ListLogs200ResponseInnerPayload) *NullableListLogs200ResponseInnerPayload {
	return &NullableListLogs200ResponseInnerPayload{value: val, isSet: true}
}

func (v NullableListLogs200ResponseInnerPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLogs200ResponseInnerPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

