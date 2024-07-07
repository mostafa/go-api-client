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

// checks if the CreateOrganizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationRequest{}

// CreateOrganizationRequest struct for CreateOrganizationRequest
type CreateOrganizationRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	// The name of the organization.
	Name string `json:"name"`
	// The description of the organization.
	Description NullableString `json:"description,omitempty"`
	// arbitrary
	CustomData map[string]interface{} `json:"customData,omitempty"`
	IsMfaRequired *bool `json:"isMfaRequired,omitempty"`
	CreatedAt *float32 `json:"createdAt,omitempty"`
}

type _CreateOrganizationRequest CreateOrganizationRequest

// NewCreateOrganizationRequest instantiates a new CreateOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationRequest(name string) *CreateOrganizationRequest {
	this := CreateOrganizationRequest{}
	this.Name = name
	return &this
}

// NewCreateOrganizationRequestWithDefaults instantiates a new CreateOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationRequestWithDefaults() *CreateOrganizationRequest {
	this := CreateOrganizationRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CreateOrganizationRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CreateOrganizationRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CreateOrganizationRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetName returns the Name field value
func (o *CreateOrganizationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrganizationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrganizationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrganizationRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateOrganizationRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateOrganizationRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateOrganizationRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *CreateOrganizationRequest) GetCustomData() map[string]interface{} {
	if o == nil || IsNil(o.CustomData) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationRequest) GetCustomDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomData) {
		return map[string]interface{}{}, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *CreateOrganizationRequest) HasCustomData() bool {
	if o != nil && !IsNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]interface{} and assigns it to the CustomData field.
func (o *CreateOrganizationRequest) SetCustomData(v map[string]interface{}) {
	o.CustomData = v
}

// GetIsMfaRequired returns the IsMfaRequired field value if set, zero value otherwise.
func (o *CreateOrganizationRequest) GetIsMfaRequired() bool {
	if o == nil || IsNil(o.IsMfaRequired) {
		var ret bool
		return ret
	}
	return *o.IsMfaRequired
}

// GetIsMfaRequiredOk returns a tuple with the IsMfaRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationRequest) GetIsMfaRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMfaRequired) {
		return nil, false
	}
	return o.IsMfaRequired, true
}

// HasIsMfaRequired returns a boolean if a field has been set.
func (o *CreateOrganizationRequest) HasIsMfaRequired() bool {
	if o != nil && !IsNil(o.IsMfaRequired) {
		return true
	}

	return false
}

// SetIsMfaRequired gets a reference to the given bool and assigns it to the IsMfaRequired field.
func (o *CreateOrganizationRequest) SetIsMfaRequired(v bool) {
	o.IsMfaRequired = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateOrganizationRequest) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationRequest) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateOrganizationRequest) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *CreateOrganizationRequest) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

func (o CreateOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.CustomData) {
		toSerialize["customData"] = o.CustomData
	}
	if !IsNil(o.IsMfaRequired) {
		toSerialize["isMfaRequired"] = o.IsMfaRequired
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

func (o *CreateOrganizationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateOrganizationRequest := _CreateOrganizationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrganizationRequest)

	if err != nil {
		return err
	}

	*o = CreateOrganizationRequest(varCreateOrganizationRequest)

	return err
}

type NullableCreateOrganizationRequest struct {
	value *CreateOrganizationRequest
	isSet bool
}

func (v NullableCreateOrganizationRequest) Get() *CreateOrganizationRequest {
	return v.value
}

func (v *NullableCreateOrganizationRequest) Set(val *CreateOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationRequest(val *CreateOrganizationRequest) *NullableCreateOrganizationRequest {
	return &NullableCreateOrganizationRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
