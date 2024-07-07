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

// checks if the GetOrganizationRole200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationRole200Response{}

// GetOrganizationRole200Response struct for GetOrganizationRole200Response
type GetOrganizationRole200Response struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
	Type string `json:"type"`
}

type _GetOrganizationRole200Response GetOrganizationRole200Response

// NewGetOrganizationRole200Response instantiates a new GetOrganizationRole200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationRole200Response(tenantId string, id string, name string, description NullableString, type_ string) *GetOrganizationRole200Response {
	this := GetOrganizationRole200Response{}
	this.TenantId = tenantId
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	return &this
}

// NewGetOrganizationRole200ResponseWithDefaults instantiates a new GetOrganizationRole200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationRole200ResponseWithDefaults() *GetOrganizationRole200Response {
	this := GetOrganizationRole200Response{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *GetOrganizationRole200Response) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationRole200Response) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *GetOrganizationRole200Response) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *GetOrganizationRole200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationRole200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetOrganizationRole200Response) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GetOrganizationRole200Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationRole200Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetOrganizationRole200Response) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetOrganizationRole200Response) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOrganizationRole200Response) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *GetOrganizationRole200Response) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetType returns the Type field value
func (o *GetOrganizationRole200Response) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationRole200Response) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetOrganizationRole200Response) SetType(v string) {
	o.Type = v
}

func (o GetOrganizationRole200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationRole200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *GetOrganizationRole200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"name",
		"description",
		"type",
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

	varGetOrganizationRole200Response := _GetOrganizationRole200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetOrganizationRole200Response)

	if err != nil {
		return err
	}

	*o = GetOrganizationRole200Response(varGetOrganizationRole200Response)

	return err
}

type NullableGetOrganizationRole200Response struct {
	value *GetOrganizationRole200Response
	isSet bool
}

func (v NullableGetOrganizationRole200Response) Get() *GetOrganizationRole200Response {
	return v.value
}

func (v *NullableGetOrganizationRole200Response) Set(val *GetOrganizationRole200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationRole200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationRole200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationRole200Response(val *GetOrganizationRole200Response) *NullableGetOrganizationRole200Response {
	return &NullableGetOrganizationRole200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationRole200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationRole200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

