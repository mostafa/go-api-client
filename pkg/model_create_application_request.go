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

// checks if the CreateApplicationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApplicationRequest{}

// CreateApplicationRequest struct for CreateApplicationRequest
type CreateApplicationRequest struct {
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Type string `json:"type"`
	OidcClientMetadata *ListApplications200ResponseInnerOidcClientMetadata `json:"oidcClientMetadata,omitempty"`
	CustomClientMetadata *ListApplications200ResponseInnerCustomClientMetadata `json:"customClientMetadata,omitempty"`
	IsThirdParty *bool `json:"isThirdParty,omitempty"`
	ProtectedAppMetadata *CreateApplicationRequestProtectedAppMetadata `json:"protectedAppMetadata,omitempty"`
}

type _CreateApplicationRequest CreateApplicationRequest

// NewCreateApplicationRequest instantiates a new CreateApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApplicationRequest(name string, type_ string) *CreateApplicationRequest {
	this := CreateApplicationRequest{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewCreateApplicationRequestWithDefaults instantiates a new CreateApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApplicationRequestWithDefaults() *CreateApplicationRequest {
	this := CreateApplicationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateApplicationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateApplicationRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateApplicationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateApplicationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateApplicationRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateApplicationRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateApplicationRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateApplicationRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value
func (o *CreateApplicationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateApplicationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateApplicationRequest) SetType(v string) {
	o.Type = v
}

// GetOidcClientMetadata returns the OidcClientMetadata field value if set, zero value otherwise.
func (o *CreateApplicationRequest) GetOidcClientMetadata() ListApplications200ResponseInnerOidcClientMetadata {
	if o == nil || IsNil(o.OidcClientMetadata) {
		var ret ListApplications200ResponseInnerOidcClientMetadata
		return ret
	}
	return *o.OidcClientMetadata
}

// GetOidcClientMetadataOk returns a tuple with the OidcClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationRequest) GetOidcClientMetadataOk() (*ListApplications200ResponseInnerOidcClientMetadata, bool) {
	if o == nil || IsNil(o.OidcClientMetadata) {
		return nil, false
	}
	return o.OidcClientMetadata, true
}

// HasOidcClientMetadata returns a boolean if a field has been set.
func (o *CreateApplicationRequest) HasOidcClientMetadata() bool {
	if o != nil && !IsNil(o.OidcClientMetadata) {
		return true
	}

	return false
}

// SetOidcClientMetadata gets a reference to the given ListApplications200ResponseInnerOidcClientMetadata and assigns it to the OidcClientMetadata field.
func (o *CreateApplicationRequest) SetOidcClientMetadata(v ListApplications200ResponseInnerOidcClientMetadata) {
	o.OidcClientMetadata = &v
}

// GetCustomClientMetadata returns the CustomClientMetadata field value if set, zero value otherwise.
func (o *CreateApplicationRequest) GetCustomClientMetadata() ListApplications200ResponseInnerCustomClientMetadata {
	if o == nil || IsNil(o.CustomClientMetadata) {
		var ret ListApplications200ResponseInnerCustomClientMetadata
		return ret
	}
	return *o.CustomClientMetadata
}

// GetCustomClientMetadataOk returns a tuple with the CustomClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationRequest) GetCustomClientMetadataOk() (*ListApplications200ResponseInnerCustomClientMetadata, bool) {
	if o == nil || IsNil(o.CustomClientMetadata) {
		return nil, false
	}
	return o.CustomClientMetadata, true
}

// HasCustomClientMetadata returns a boolean if a field has been set.
func (o *CreateApplicationRequest) HasCustomClientMetadata() bool {
	if o != nil && !IsNil(o.CustomClientMetadata) {
		return true
	}

	return false
}

// SetCustomClientMetadata gets a reference to the given ListApplications200ResponseInnerCustomClientMetadata and assigns it to the CustomClientMetadata field.
func (o *CreateApplicationRequest) SetCustomClientMetadata(v ListApplications200ResponseInnerCustomClientMetadata) {
	o.CustomClientMetadata = &v
}

// GetIsThirdParty returns the IsThirdParty field value if set, zero value otherwise.
func (o *CreateApplicationRequest) GetIsThirdParty() bool {
	if o == nil || IsNil(o.IsThirdParty) {
		var ret bool
		return ret
	}
	return *o.IsThirdParty
}

// GetIsThirdPartyOk returns a tuple with the IsThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationRequest) GetIsThirdPartyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsThirdParty) {
		return nil, false
	}
	return o.IsThirdParty, true
}

// HasIsThirdParty returns a boolean if a field has been set.
func (o *CreateApplicationRequest) HasIsThirdParty() bool {
	if o != nil && !IsNil(o.IsThirdParty) {
		return true
	}

	return false
}

// SetIsThirdParty gets a reference to the given bool and assigns it to the IsThirdParty field.
func (o *CreateApplicationRequest) SetIsThirdParty(v bool) {
	o.IsThirdParty = &v
}

// GetProtectedAppMetadata returns the ProtectedAppMetadata field value if set, zero value otherwise.
func (o *CreateApplicationRequest) GetProtectedAppMetadata() CreateApplicationRequestProtectedAppMetadata {
	if o == nil || IsNil(o.ProtectedAppMetadata) {
		var ret CreateApplicationRequestProtectedAppMetadata
		return ret
	}
	return *o.ProtectedAppMetadata
}

// GetProtectedAppMetadataOk returns a tuple with the ProtectedAppMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationRequest) GetProtectedAppMetadataOk() (*CreateApplicationRequestProtectedAppMetadata, bool) {
	if o == nil || IsNil(o.ProtectedAppMetadata) {
		return nil, false
	}
	return o.ProtectedAppMetadata, true
}

// HasProtectedAppMetadata returns a boolean if a field has been set.
func (o *CreateApplicationRequest) HasProtectedAppMetadata() bool {
	if o != nil && !IsNil(o.ProtectedAppMetadata) {
		return true
	}

	return false
}

// SetProtectedAppMetadata gets a reference to the given CreateApplicationRequestProtectedAppMetadata and assigns it to the ProtectedAppMetadata field.
func (o *CreateApplicationRequest) SetProtectedAppMetadata(v CreateApplicationRequestProtectedAppMetadata) {
	o.ProtectedAppMetadata = &v
}

func (o CreateApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApplicationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.OidcClientMetadata) {
		toSerialize["oidcClientMetadata"] = o.OidcClientMetadata
	}
	if !IsNil(o.CustomClientMetadata) {
		toSerialize["customClientMetadata"] = o.CustomClientMetadata
	}
	if !IsNil(o.IsThirdParty) {
		toSerialize["isThirdParty"] = o.IsThirdParty
	}
	if !IsNil(o.ProtectedAppMetadata) {
		toSerialize["protectedAppMetadata"] = o.ProtectedAppMetadata
	}
	return toSerialize, nil
}

func (o *CreateApplicationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateApplicationRequest := _CreateApplicationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateApplicationRequest)

	if err != nil {
		return err
	}

	*o = CreateApplicationRequest(varCreateApplicationRequest)

	return err
}

type NullableCreateApplicationRequest struct {
	value *CreateApplicationRequest
	isSet bool
}

func (v NullableCreateApplicationRequest) Get() *CreateApplicationRequest {
	return v.value
}

func (v *NullableCreateApplicationRequest) Set(val *CreateApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApplicationRequest(val *CreateApplicationRequest) *NullableCreateApplicationRequest {
	return &NullableCreateApplicationRequest{value: val, isSet: true}
}

func (v NullableCreateApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
