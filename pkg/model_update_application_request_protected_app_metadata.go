/*
Logto API references

API references for Logto services.  Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.

API version: Cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logto

import (
	"encoding/json"
)

// checks if the UpdateApplicationRequestProtectedAppMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateApplicationRequestProtectedAppMetadata{}

// UpdateApplicationRequestProtectedAppMetadata struct for UpdateApplicationRequestProtectedAppMetadata
type UpdateApplicationRequestProtectedAppMetadata struct {
	Origin *string `json:"origin,omitempty"`
	SessionDuration *float32 `json:"sessionDuration,omitempty"`
	PageRules []ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner `json:"pageRules,omitempty"`
}

// NewUpdateApplicationRequestProtectedAppMetadata instantiates a new UpdateApplicationRequestProtectedAppMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApplicationRequestProtectedAppMetadata() *UpdateApplicationRequestProtectedAppMetadata {
	this := UpdateApplicationRequestProtectedAppMetadata{}
	return &this
}

// NewUpdateApplicationRequestProtectedAppMetadataWithDefaults instantiates a new UpdateApplicationRequestProtectedAppMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApplicationRequestProtectedAppMetadataWithDefaults() *UpdateApplicationRequestProtectedAppMetadata {
	this := UpdateApplicationRequestProtectedAppMetadata{}
	return &this
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *UpdateApplicationRequestProtectedAppMetadata) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplicationRequestProtectedAppMetadata) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *UpdateApplicationRequestProtectedAppMetadata) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *UpdateApplicationRequestProtectedAppMetadata) SetOrigin(v string) {
	o.Origin = &v
}

// GetSessionDuration returns the SessionDuration field value if set, zero value otherwise.
func (o *UpdateApplicationRequestProtectedAppMetadata) GetSessionDuration() float32 {
	if o == nil || IsNil(o.SessionDuration) {
		var ret float32
		return ret
	}
	return *o.SessionDuration
}

// GetSessionDurationOk returns a tuple with the SessionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplicationRequestProtectedAppMetadata) GetSessionDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.SessionDuration) {
		return nil, false
	}
	return o.SessionDuration, true
}

// HasSessionDuration returns a boolean if a field has been set.
func (o *UpdateApplicationRequestProtectedAppMetadata) HasSessionDuration() bool {
	if o != nil && !IsNil(o.SessionDuration) {
		return true
	}

	return false
}

// SetSessionDuration gets a reference to the given float32 and assigns it to the SessionDuration field.
func (o *UpdateApplicationRequestProtectedAppMetadata) SetSessionDuration(v float32) {
	o.SessionDuration = &v
}

// GetPageRules returns the PageRules field value if set, zero value otherwise.
func (o *UpdateApplicationRequestProtectedAppMetadata) GetPageRules() []ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner {
	if o == nil || IsNil(o.PageRules) {
		var ret []ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner
		return ret
	}
	return o.PageRules
}

// GetPageRulesOk returns a tuple with the PageRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplicationRequestProtectedAppMetadata) GetPageRulesOk() ([]ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner, bool) {
	if o == nil || IsNil(o.PageRules) {
		return nil, false
	}
	return o.PageRules, true
}

// HasPageRules returns a boolean if a field has been set.
func (o *UpdateApplicationRequestProtectedAppMetadata) HasPageRules() bool {
	if o != nil && !IsNil(o.PageRules) {
		return true
	}

	return false
}

// SetPageRules gets a reference to the given []ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner and assigns it to the PageRules field.
func (o *UpdateApplicationRequestProtectedAppMetadata) SetPageRules(v []ListApplications200ResponseInnerProtectedAppMetadataPageRulesInner) {
	o.PageRules = v
}

func (o UpdateApplicationRequestProtectedAppMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateApplicationRequestProtectedAppMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.SessionDuration) {
		toSerialize["sessionDuration"] = o.SessionDuration
	}
	if !IsNil(o.PageRules) {
		toSerialize["pageRules"] = o.PageRules
	}
	return toSerialize, nil
}

type NullableUpdateApplicationRequestProtectedAppMetadata struct {
	value *UpdateApplicationRequestProtectedAppMetadata
	isSet bool
}

func (v NullableUpdateApplicationRequestProtectedAppMetadata) Get() *UpdateApplicationRequestProtectedAppMetadata {
	return v.value
}

func (v *NullableUpdateApplicationRequestProtectedAppMetadata) Set(val *UpdateApplicationRequestProtectedAppMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApplicationRequestProtectedAppMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApplicationRequestProtectedAppMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApplicationRequestProtectedAppMetadata(val *UpdateApplicationRequestProtectedAppMetadata) *NullableUpdateApplicationRequestProtectedAppMetadata {
	return &NullableUpdateApplicationRequestProtectedAppMetadata{value: val, isSet: true}
}

func (v NullableUpdateApplicationRequestProtectedAppMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApplicationRequestProtectedAppMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

