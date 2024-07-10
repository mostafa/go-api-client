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

// checks if the ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner{}

// ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner struct for ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner
type ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner struct {
	Domain string `json:"domain"`
	Status string `json:"status"`
	ErrorMessage NullableString `json:"errorMessage"`
	DnsRecords []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner `json:"dnsRecords"`
	CloudflareData NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData `json:"cloudflareData"`
}

type _ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner

// NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner instantiates a new ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner(domain string, status string, errorMessage NullableString, dnsRecords []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner, cloudflareData NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner {
	this := ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner{}
	this.Domain = domain
	this.Status = status
	this.ErrorMessage = errorMessage
	this.DnsRecords = dnsRecords
	this.CloudflareData = cloudflareData
	return &this
}

// NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerWithDefaults instantiates a new ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerWithDefaults() *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner {
	this := ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner{}
	return &this
}

// GetDomain returns the Domain field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetDomain(v string) {
	o.Domain = v
}

// GetStatus returns the Status field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetStatus(v string) {
	o.Status = v
}

// GetErrorMessage returns the ErrorMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// SetErrorMessage sets field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// GetDnsRecords returns the DnsRecords field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetDnsRecords() []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner {
	if o == nil {
		var ret []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner
		return ret
	}

	return o.DnsRecords
}

// GetDnsRecordsOk returns a tuple with the DnsRecords field value
// and a boolean to check if the value has been set.
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetDnsRecordsOk() ([]ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsRecords, true
}

// SetDnsRecords sets field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetDnsRecords(v []ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerDnsRecordsInner) {
	o.DnsRecords = v
}

// GetCloudflareData returns the CloudflareData field value
// If the value is explicit nil, the zero value for ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData will be returned
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetCloudflareData() ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData {
	if o == nil || o.CloudflareData.Get() == nil {
		var ret ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData
		return ret
	}

	return *o.CloudflareData.Get()
}

// GetCloudflareDataOk returns a tuple with the CloudflareData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) GetCloudflareDataOk() (*ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudflareData.Get(), o.CloudflareData.IsSet()
}

// SetCloudflareData sets field value
func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) SetCloudflareData(v ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInnerCloudflareData) {
	o.CloudflareData.Set(&v)
}

func (o ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	toSerialize["status"] = o.Status
	toSerialize["errorMessage"] = o.ErrorMessage.Get()
	toSerialize["dnsRecords"] = o.DnsRecords
	toSerialize["cloudflareData"] = o.CloudflareData.Get()
	return toSerialize, nil
}

func (o *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"status",
		"errorMessage",
		"dnsRecords",
		"cloudflareData",
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

	varListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner := _ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner)

	if err != nil {
		return err
	}

	*o = ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner(varListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner)

	return err
}

type NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner struct {
	value *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner
	isSet bool
}

func (v NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) Get() *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner {
	return v.value
}

func (v *NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) Set(val *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner(val *ListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) *NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner {
	return &NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner{value: val, isSet: true}
}

func (v NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplications200ResponseInnerProtectedAppMetadataCustomDomainsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
