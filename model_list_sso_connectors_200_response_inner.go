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

// checks if the ListSsoConnectors200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSsoConnectors200ResponseInner{}

// ListSsoConnectors200ResponseInner struct for ListSsoConnectors200ResponseInner
type ListSsoConnectors200ResponseInner struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	ProviderName string `json:"providerName"`
	ConnectorName string `json:"connectorName"`
	// arbitrary
	Config map[string]interface{} `json:"config"`
	Domains []string `json:"domains"`
	Branding ListOrganizationJitSsoConnectors200ResponseInnerBranding `json:"branding"`
	SyncProfile bool `json:"syncProfile"`
	CreatedAt float32 `json:"createdAt"`
	Name string `json:"name"`
	ProviderType string `json:"providerType"`
	ProviderLogo string `json:"providerLogo"`
	ProviderLogoDark string `json:"providerLogoDark"`
	ProviderConfig map[string]interface{} `json:"providerConfig,omitempty"`
}

type _ListSsoConnectors200ResponseInner ListSsoConnectors200ResponseInner

// NewListSsoConnectors200ResponseInner instantiates a new ListSsoConnectors200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSsoConnectors200ResponseInner(tenantId string, id string, providerName string, connectorName string, config map[string]interface{}, domains []string, branding ListOrganizationJitSsoConnectors200ResponseInnerBranding, syncProfile bool, createdAt float32, name string, providerType string, providerLogo string, providerLogoDark string) *ListSsoConnectors200ResponseInner {
	this := ListSsoConnectors200ResponseInner{}
	this.TenantId = tenantId
	this.Id = id
	this.ProviderName = providerName
	this.ConnectorName = connectorName
	this.Config = config
	this.Domains = domains
	this.Branding = branding
	this.SyncProfile = syncProfile
	this.CreatedAt = createdAt
	this.Name = name
	this.ProviderType = providerType
	this.ProviderLogo = providerLogo
	this.ProviderLogoDark = providerLogoDark
	return &this
}

// NewListSsoConnectors200ResponseInnerWithDefaults instantiates a new ListSsoConnectors200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSsoConnectors200ResponseInnerWithDefaults() *ListSsoConnectors200ResponseInner {
	this := ListSsoConnectors200ResponseInner{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListSsoConnectors200ResponseInner) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListSsoConnectors200ResponseInner) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *ListSsoConnectors200ResponseInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListSsoConnectors200ResponseInner) SetId(v string) {
	o.Id = v
}

// GetProviderName returns the ProviderName field value
func (o *ListSsoConnectors200ResponseInner) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *ListSsoConnectors200ResponseInner) SetProviderName(v string) {
	o.ProviderName = v
}

// GetConnectorName returns the ConnectorName field value
func (o *ListSsoConnectors200ResponseInner) GetConnectorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorName
}

// GetConnectorNameOk returns a tuple with the ConnectorName field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetConnectorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorName, true
}

// SetConnectorName sets field value
func (o *ListSsoConnectors200ResponseInner) SetConnectorName(v string) {
	o.ConnectorName = v
}

// GetConfig returns the Config field value
func (o *ListSsoConnectors200ResponseInner) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// SetConfig sets field value
func (o *ListSsoConnectors200ResponseInner) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetDomains returns the Domains field value
func (o *ListSsoConnectors200ResponseInner) GetDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetDomainsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domains, true
}

// SetDomains sets field value
func (o *ListSsoConnectors200ResponseInner) SetDomains(v []string) {
	o.Domains = v
}

// GetBranding returns the Branding field value
func (o *ListSsoConnectors200ResponseInner) GetBranding() ListOrganizationJitSsoConnectors200ResponseInnerBranding {
	if o == nil {
		var ret ListOrganizationJitSsoConnectors200ResponseInnerBranding
		return ret
	}

	return o.Branding
}

// GetBrandingOk returns a tuple with the Branding field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetBrandingOk() (*ListOrganizationJitSsoConnectors200ResponseInnerBranding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branding, true
}

// SetBranding sets field value
func (o *ListSsoConnectors200ResponseInner) SetBranding(v ListOrganizationJitSsoConnectors200ResponseInnerBranding) {
	o.Branding = v
}

// GetSyncProfile returns the SyncProfile field value
func (o *ListSsoConnectors200ResponseInner) GetSyncProfile() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SyncProfile
}

// GetSyncProfileOk returns a tuple with the SyncProfile field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetSyncProfileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncProfile, true
}

// SetSyncProfile sets field value
func (o *ListSsoConnectors200ResponseInner) SetSyncProfile(v bool) {
	o.SyncProfile = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ListSsoConnectors200ResponseInner) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ListSsoConnectors200ResponseInner) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetName returns the Name field value
func (o *ListSsoConnectors200ResponseInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListSsoConnectors200ResponseInner) SetName(v string) {
	o.Name = v
}

// GetProviderType returns the ProviderType field value
func (o *ListSsoConnectors200ResponseInner) GetProviderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *ListSsoConnectors200ResponseInner) SetProviderType(v string) {
	o.ProviderType = v
}

// GetProviderLogo returns the ProviderLogo field value
func (o *ListSsoConnectors200ResponseInner) GetProviderLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderLogo
}

// GetProviderLogoOk returns a tuple with the ProviderLogo field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetProviderLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderLogo, true
}

// SetProviderLogo sets field value
func (o *ListSsoConnectors200ResponseInner) SetProviderLogo(v string) {
	o.ProviderLogo = v
}

// GetProviderLogoDark returns the ProviderLogoDark field value
func (o *ListSsoConnectors200ResponseInner) GetProviderLogoDark() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderLogoDark
}

// GetProviderLogoDarkOk returns a tuple with the ProviderLogoDark field value
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetProviderLogoDarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderLogoDark, true
}

// SetProviderLogoDark sets field value
func (o *ListSsoConnectors200ResponseInner) SetProviderLogoDark(v string) {
	o.ProviderLogoDark = v
}

// GetProviderConfig returns the ProviderConfig field value if set, zero value otherwise.
func (o *ListSsoConnectors200ResponseInner) GetProviderConfig() map[string]interface{} {
	if o == nil || IsNil(o.ProviderConfig) {
		var ret map[string]interface{}
		return ret
	}
	return o.ProviderConfig
}

// GetProviderConfigOk returns a tuple with the ProviderConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSsoConnectors200ResponseInner) GetProviderConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ProviderConfig) {
		return map[string]interface{}{}, false
	}
	return o.ProviderConfig, true
}

// HasProviderConfig returns a boolean if a field has been set.
func (o *ListSsoConnectors200ResponseInner) HasProviderConfig() bool {
	if o != nil && !IsNil(o.ProviderConfig) {
		return true
	}

	return false
}

// SetProviderConfig gets a reference to the given map[string]interface{} and assigns it to the ProviderConfig field.
func (o *ListSsoConnectors200ResponseInner) SetProviderConfig(v map[string]interface{}) {
	o.ProviderConfig = v
}

func (o ListSsoConnectors200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSsoConnectors200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["providerName"] = o.ProviderName
	toSerialize["connectorName"] = o.ConnectorName
	toSerialize["config"] = o.Config
	toSerialize["domains"] = o.Domains
	toSerialize["branding"] = o.Branding
	toSerialize["syncProfile"] = o.SyncProfile
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["name"] = o.Name
	toSerialize["providerType"] = o.ProviderType
	toSerialize["providerLogo"] = o.ProviderLogo
	toSerialize["providerLogoDark"] = o.ProviderLogoDark
	if !IsNil(o.ProviderConfig) {
		toSerialize["providerConfig"] = o.ProviderConfig
	}
	return toSerialize, nil
}

func (o *ListSsoConnectors200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"providerName",
		"connectorName",
		"config",
		"domains",
		"branding",
		"syncProfile",
		"createdAt",
		"name",
		"providerType",
		"providerLogo",
		"providerLogoDark",
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

	varListSsoConnectors200ResponseInner := _ListSsoConnectors200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListSsoConnectors200ResponseInner)

	if err != nil {
		return err
	}

	*o = ListSsoConnectors200ResponseInner(varListSsoConnectors200ResponseInner)

	return err
}

type NullableListSsoConnectors200ResponseInner struct {
	value *ListSsoConnectors200ResponseInner
	isSet bool
}

func (v NullableListSsoConnectors200ResponseInner) Get() *ListSsoConnectors200ResponseInner {
	return v.value
}

func (v *NullableListSsoConnectors200ResponseInner) Set(val *ListSsoConnectors200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListSsoConnectors200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListSsoConnectors200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSsoConnectors200ResponseInner(val *ListSsoConnectors200ResponseInner) *NullableListSsoConnectors200ResponseInner {
	return &NullableListSsoConnectors200ResponseInner{value: val, isSet: true}
}

func (v NullableListSsoConnectors200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSsoConnectors200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
