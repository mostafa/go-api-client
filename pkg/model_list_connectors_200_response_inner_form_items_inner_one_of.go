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

// checks if the ListConnectors200ResponseInnerFormItemsInnerOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectors200ResponseInnerFormItemsInnerOneOf{}

// ListConnectors200ResponseInnerFormItemsInnerOneOf struct for ListConnectors200ResponseInnerFormItemsInnerOneOf
type ListConnectors200ResponseInnerFormItemsInnerOneOf struct {
	Type string `json:"type"`
	SelectItems []ListConnectors200ResponseInnerFormItemsInnerOneOfSelectItemsInner `json:"selectItems"`
	Key string `json:"key"`
	Label string `json:"label"`
	Placeholder *string `json:"placeholder,omitempty"`
	Required *bool `json:"required,omitempty"`
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	ShowConditions []ListConnectors200ResponseInnerFormItemsInnerOneOfShowConditionsInner `json:"showConditions,omitempty"`
	Description *string `json:"description,omitempty"`
	Tooltip *string `json:"tooltip,omitempty"`
	IsConfidential *bool `json:"isConfidential,omitempty"`
}

type _ListConnectors200ResponseInnerFormItemsInnerOneOf ListConnectors200ResponseInnerFormItemsInnerOneOf

// NewListConnectors200ResponseInnerFormItemsInnerOneOf instantiates a new ListConnectors200ResponseInnerFormItemsInnerOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectors200ResponseInnerFormItemsInnerOneOf(type_ string, selectItems []ListConnectors200ResponseInnerFormItemsInnerOneOfSelectItemsInner, key string, label string) *ListConnectors200ResponseInnerFormItemsInnerOneOf {
	this := ListConnectors200ResponseInnerFormItemsInnerOneOf{}
	this.Type = type_
	this.SelectItems = selectItems
	this.Key = key
	this.Label = label
	return &this
}

// NewListConnectors200ResponseInnerFormItemsInnerOneOfWithDefaults instantiates a new ListConnectors200ResponseInnerFormItemsInnerOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectors200ResponseInnerFormItemsInnerOneOfWithDefaults() *ListConnectors200ResponseInnerFormItemsInnerOneOf {
	this := ListConnectors200ResponseInnerFormItemsInnerOneOf{}
	return &this
}

// GetType returns the Type field value
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) SetType(v string) {
	o.Type = v
}

// GetSelectItems returns the SelectItems field value
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetSelectItems() []ListConnectors200ResponseInnerFormItemsInnerOneOfSelectItemsInner {
	if o == nil {
		var ret []ListConnectors200ResponseInnerFormItemsInnerOneOfSelectItemsInner
		return ret
	}

	return o.SelectItems
}

// GetSelectItemsOk returns a tuple with the SelectItems field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetSelectItemsOk() ([]ListConnectors200ResponseInnerFormItemsInnerOneOfSelectItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectItems, true
}

// SetSelectItems sets field value
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) SetSelectItems(v []ListConnectors200ResponseInnerFormItemsInnerOneOfSelectItemsInner) {
	o.SelectItems = v
}

// GetKey returns the Key field value
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) SetLabel(v string) {
	o.Label = v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) SetRequired(v bool) {
	o.Required = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetDefaultValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return &o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given interface{} and assigns it to the DefaultValue field.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

// GetShowConditions returns the ShowConditions field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetShowConditions() []ListConnectors200ResponseInnerFormItemsInnerOneOfShowConditionsInner {
	if o == nil || IsNil(o.ShowConditions) {
		var ret []ListConnectors200ResponseInnerFormItemsInnerOneOfShowConditionsInner
		return ret
	}
	return o.ShowConditions
}

// GetShowConditionsOk returns a tuple with the ShowConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetShowConditionsOk() ([]ListConnectors200ResponseInnerFormItemsInnerOneOfShowConditionsInner, bool) {
	if o == nil || IsNil(o.ShowConditions) {
		return nil, false
	}
	return o.ShowConditions, true
}

// HasShowConditions returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) HasShowConditions() bool {
	if o != nil && !IsNil(o.ShowConditions) {
		return true
	}

	return false
}

// SetShowConditions gets a reference to the given []ListConnectors200ResponseInnerFormItemsInnerOneOfShowConditionsInner and assigns it to the ShowConditions field.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) SetShowConditions(v []ListConnectors200ResponseInnerFormItemsInnerOneOfShowConditionsInner) {
	o.ShowConditions = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) SetDescription(v string) {
	o.Description = &v
}

// GetTooltip returns the Tooltip field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetTooltip() string {
	if o == nil || IsNil(o.Tooltip) {
		var ret string
		return ret
	}
	return *o.Tooltip
}

// GetTooltipOk returns a tuple with the Tooltip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetTooltipOk() (*string, bool) {
	if o == nil || IsNil(o.Tooltip) {
		return nil, false
	}
	return o.Tooltip, true
}

// HasTooltip returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) HasTooltip() bool {
	if o != nil && !IsNil(o.Tooltip) {
		return true
	}

	return false
}

// SetTooltip gets a reference to the given string and assigns it to the Tooltip field.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) SetTooltip(v string) {
	o.Tooltip = &v
}

// GetIsConfidential returns the IsConfidential field value if set, zero value otherwise.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetIsConfidential() bool {
	if o == nil || IsNil(o.IsConfidential) {
		var ret bool
		return ret
	}
	return *o.IsConfidential
}

// GetIsConfidentialOk returns a tuple with the IsConfidential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) GetIsConfidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.IsConfidential) {
		return nil, false
	}
	return o.IsConfidential, true
}

// HasIsConfidential returns a boolean if a field has been set.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) HasIsConfidential() bool {
	if o != nil && !IsNil(o.IsConfidential) {
		return true
	}

	return false
}

// SetIsConfidential gets a reference to the given bool and assigns it to the IsConfidential field.
func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) SetIsConfidential(v bool) {
	o.IsConfidential = &v
}

func (o ListConnectors200ResponseInnerFormItemsInnerOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectors200ResponseInnerFormItemsInnerOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["selectItems"] = o.SelectItems
	toSerialize["key"] = o.Key
	toSerialize["label"] = o.Label
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.ShowConditions) {
		toSerialize["showConditions"] = o.ShowConditions
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tooltip) {
		toSerialize["tooltip"] = o.Tooltip
	}
	if !IsNil(o.IsConfidential) {
		toSerialize["isConfidential"] = o.IsConfidential
	}
	return toSerialize, nil
}

func (o *ListConnectors200ResponseInnerFormItemsInnerOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"selectItems",
		"key",
		"label",
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

	varListConnectors200ResponseInnerFormItemsInnerOneOf := _ListConnectors200ResponseInnerFormItemsInnerOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListConnectors200ResponseInnerFormItemsInnerOneOf)

	if err != nil {
		return err
	}

	*o = ListConnectors200ResponseInnerFormItemsInnerOneOf(varListConnectors200ResponseInnerFormItemsInnerOneOf)

	return err
}

type NullableListConnectors200ResponseInnerFormItemsInnerOneOf struct {
	value *ListConnectors200ResponseInnerFormItemsInnerOneOf
	isSet bool
}

func (v NullableListConnectors200ResponseInnerFormItemsInnerOneOf) Get() *ListConnectors200ResponseInnerFormItemsInnerOneOf {
	return v.value
}

func (v *NullableListConnectors200ResponseInnerFormItemsInnerOneOf) Set(val *ListConnectors200ResponseInnerFormItemsInnerOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectors200ResponseInnerFormItemsInnerOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectors200ResponseInnerFormItemsInnerOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectors200ResponseInnerFormItemsInnerOneOf(val *ListConnectors200ResponseInnerFormItemsInnerOneOf) *NullableListConnectors200ResponseInnerFormItemsInnerOneOf {
	return &NullableListConnectors200ResponseInnerFormItemsInnerOneOf{value: val, isSet: true}
}

func (v NullableListConnectors200ResponseInnerFormItemsInnerOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectors200ResponseInnerFormItemsInnerOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

