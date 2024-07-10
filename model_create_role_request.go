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

// checks if the CreateRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleRequest{}

// CreateRoleRequest struct for CreateRoleRequest
type CreateRoleRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	// The name of the role. It should be unique within the tenant.
	Name string `json:"name"`
	Description string `json:"description"`
	// The type of the role. It cannot be changed after creation.
	Type *string `json:"type,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	// The initial API resource scopes assigned to the role.
	ScopeIds []string `json:"scopeIds,omitempty"`
}

type _CreateRoleRequest CreateRoleRequest

// NewCreateRoleRequest instantiates a new CreateRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleRequest(name string, description string) *CreateRoleRequest {
	this := CreateRoleRequest{}
	this.Name = name
	this.Description = description
	return &this
}

// NewCreateRoleRequestWithDefaults instantiates a new CreateRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleRequestWithDefaults() *CreateRoleRequest {
	this := CreateRoleRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CreateRoleRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CreateRoleRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CreateRoleRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetName returns the Name field value
func (o *CreateRoleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRoleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRoleRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *CreateRoleRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateRoleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateRoleRequest) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateRoleRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateRoleRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateRoleRequest) SetType(v string) {
	o.Type = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *CreateRoleRequest) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleRequest) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *CreateRoleRequest) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *CreateRoleRequest) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetScopeIds returns the ScopeIds field value if set, zero value otherwise.
func (o *CreateRoleRequest) GetScopeIds() []string {
	if o == nil || IsNil(o.ScopeIds) {
		var ret []string
		return ret
	}
	return o.ScopeIds
}

// GetScopeIdsOk returns a tuple with the ScopeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleRequest) GetScopeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ScopeIds) {
		return nil, false
	}
	return o.ScopeIds, true
}

// HasScopeIds returns a boolean if a field has been set.
func (o *CreateRoleRequest) HasScopeIds() bool {
	if o != nil && !IsNil(o.ScopeIds) {
		return true
	}

	return false
}

// SetScopeIds gets a reference to the given []string and assigns it to the ScopeIds field.
func (o *CreateRoleRequest) SetScopeIds(v []string) {
	o.ScopeIds = v
}

func (o CreateRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.ScopeIds) {
		toSerialize["scopeIds"] = o.ScopeIds
	}
	return toSerialize, nil
}

func (o *CreateRoleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
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

	varCreateRoleRequest := _CreateRoleRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRoleRequest)

	if err != nil {
		return err
	}

	*o = CreateRoleRequest(varCreateRoleRequest)

	return err
}

type NullableCreateRoleRequest struct {
	value *CreateRoleRequest
	isSet bool
}

func (v NullableCreateRoleRequest) Get() *CreateRoleRequest {
	return v.value
}

func (v *NullableCreateRoleRequest) Set(val *CreateRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleRequest(val *CreateRoleRequest) *NullableCreateRoleRequest {
	return &NullableCreateRoleRequest{value: val, isSet: true}
}

func (v NullableCreateRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
