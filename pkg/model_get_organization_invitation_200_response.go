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

// checks if the GetOrganizationInvitation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationInvitation200Response{}

// GetOrganizationInvitation200Response struct for GetOrganizationInvitation200Response
type GetOrganizationInvitation200Response struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
	InviterId NullableString `json:"inviterId"`
	Invitee string `json:"invitee"`
	AcceptedUserId NullableString `json:"acceptedUserId"`
	OrganizationId string `json:"organizationId"`
	Status string `json:"status"`
	CreatedAt float32 `json:"createdAt"`
	UpdatedAt float32 `json:"updatedAt"`
	ExpiresAt float32 `json:"expiresAt"`
	OrganizationRoles []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner `json:"organizationRoles"`
}

type _GetOrganizationInvitation200Response GetOrganizationInvitation200Response

// NewGetOrganizationInvitation200Response instantiates a new GetOrganizationInvitation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationInvitation200Response(tenantId string, id string, inviterId NullableString, invitee string, acceptedUserId NullableString, organizationId string, status string, createdAt float32, updatedAt float32, expiresAt float32, organizationRoles []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) *GetOrganizationInvitation200Response {
	this := GetOrganizationInvitation200Response{}
	this.TenantId = tenantId
	this.Id = id
	this.InviterId = inviterId
	this.Invitee = invitee
	this.AcceptedUserId = acceptedUserId
	this.OrganizationId = organizationId
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.ExpiresAt = expiresAt
	this.OrganizationRoles = organizationRoles
	return &this
}

// NewGetOrganizationInvitation200ResponseWithDefaults instantiates a new GetOrganizationInvitation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationInvitation200ResponseWithDefaults() *GetOrganizationInvitation200Response {
	this := GetOrganizationInvitation200Response{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *GetOrganizationInvitation200Response) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationInvitation200Response) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *GetOrganizationInvitation200Response) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *GetOrganizationInvitation200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationInvitation200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetOrganizationInvitation200Response) SetId(v string) {
	o.Id = v
}

// GetInviterId returns the InviterId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetOrganizationInvitation200Response) GetInviterId() string {
	if o == nil || o.InviterId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InviterId.Get()
}

// GetInviterIdOk returns a tuple with the InviterId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOrganizationInvitation200Response) GetInviterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InviterId.Get(), o.InviterId.IsSet()
}

// SetInviterId sets field value
func (o *GetOrganizationInvitation200Response) SetInviterId(v string) {
	o.InviterId.Set(&v)
}

// GetInvitee returns the Invitee field value
func (o *GetOrganizationInvitation200Response) GetInvitee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Invitee
}

// GetInviteeOk returns a tuple with the Invitee field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationInvitation200Response) GetInviteeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invitee, true
}

// SetInvitee sets field value
func (o *GetOrganizationInvitation200Response) SetInvitee(v string) {
	o.Invitee = v
}

// GetAcceptedUserId returns the AcceptedUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetOrganizationInvitation200Response) GetAcceptedUserId() string {
	if o == nil || o.AcceptedUserId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AcceptedUserId.Get()
}

// GetAcceptedUserIdOk returns a tuple with the AcceptedUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetOrganizationInvitation200Response) GetAcceptedUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptedUserId.Get(), o.AcceptedUserId.IsSet()
}

// SetAcceptedUserId sets field value
func (o *GetOrganizationInvitation200Response) SetAcceptedUserId(v string) {
	o.AcceptedUserId.Set(&v)
}

// GetOrganizationId returns the OrganizationId field value
func (o *GetOrganizationInvitation200Response) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationInvitation200Response) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *GetOrganizationInvitation200Response) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetStatus returns the Status field value
func (o *GetOrganizationInvitation200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationInvitation200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetOrganizationInvitation200Response) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetOrganizationInvitation200Response) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationInvitation200Response) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetOrganizationInvitation200Response) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetOrganizationInvitation200Response) GetUpdatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationInvitation200Response) GetUpdatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetOrganizationInvitation200Response) SetUpdatedAt(v float32) {
	o.UpdatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *GetOrganizationInvitation200Response) GetExpiresAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationInvitation200Response) GetExpiresAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *GetOrganizationInvitation200Response) SetExpiresAt(v float32) {
	o.ExpiresAt = v
}

// GetOrganizationRoles returns the OrganizationRoles field value
func (o *GetOrganizationInvitation200Response) GetOrganizationRoles() []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner {
	if o == nil {
		var ret []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner
		return ret
	}

	return o.OrganizationRoles
}

// GetOrganizationRolesOk returns a tuple with the OrganizationRoles field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationInvitation200Response) GetOrganizationRolesOk() ([]ListApplicationOrganizations200ResponseInnerOrganizationRolesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationRoles, true
}

// SetOrganizationRoles sets field value
func (o *GetOrganizationInvitation200Response) SetOrganizationRoles(v []ListApplicationOrganizations200ResponseInnerOrganizationRolesInner) {
	o.OrganizationRoles = v
}

func (o GetOrganizationInvitation200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationInvitation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	toSerialize["inviterId"] = o.InviterId.Get()
	toSerialize["invitee"] = o.Invitee
	toSerialize["acceptedUserId"] = o.AcceptedUserId.Get()
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["status"] = o.Status
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["expiresAt"] = o.ExpiresAt
	toSerialize["organizationRoles"] = o.OrganizationRoles
	return toSerialize, nil
}

func (o *GetOrganizationInvitation200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
		"inviterId",
		"invitee",
		"acceptedUserId",
		"organizationId",
		"status",
		"createdAt",
		"updatedAt",
		"expiresAt",
		"organizationRoles",
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

	varGetOrganizationInvitation200Response := _GetOrganizationInvitation200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetOrganizationInvitation200Response)

	if err != nil {
		return err
	}

	*o = GetOrganizationInvitation200Response(varGetOrganizationInvitation200Response)

	return err
}

type NullableGetOrganizationInvitation200Response struct {
	value *GetOrganizationInvitation200Response
	isSet bool
}

func (v NullableGetOrganizationInvitation200Response) Get() *GetOrganizationInvitation200Response {
	return v.value
}

func (v *NullableGetOrganizationInvitation200Response) Set(val *GetOrganizationInvitation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationInvitation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationInvitation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationInvitation200Response(val *GetOrganizationInvitation200Response) *NullableGetOrganizationInvitation200Response {
	return &NullableGetOrganizationInvitation200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationInvitation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationInvitation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
