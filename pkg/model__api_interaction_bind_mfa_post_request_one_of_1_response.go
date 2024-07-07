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

// checks if the ApiInteractionBindMfaPostRequestOneOf1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInteractionBindMfaPostRequestOneOf1Response{}

// ApiInteractionBindMfaPostRequestOneOf1Response struct for ApiInteractionBindMfaPostRequestOneOf1Response
type ApiInteractionBindMfaPostRequestOneOf1Response struct {
	ClientDataJSON string `json:"clientDataJSON"`
	AttestationObject string `json:"attestationObject"`
	AuthenticatorData *string `json:"authenticatorData,omitempty"`
	Transports []string `json:"transports,omitempty"`
	PublicKeyAlgorithm *float32 `json:"publicKeyAlgorithm,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
}

type _ApiInteractionBindMfaPostRequestOneOf1Response ApiInteractionBindMfaPostRequestOneOf1Response

// NewApiInteractionBindMfaPostRequestOneOf1Response instantiates a new ApiInteractionBindMfaPostRequestOneOf1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInteractionBindMfaPostRequestOneOf1Response(clientDataJSON string, attestationObject string) *ApiInteractionBindMfaPostRequestOneOf1Response {
	this := ApiInteractionBindMfaPostRequestOneOf1Response{}
	this.ClientDataJSON = clientDataJSON
	this.AttestationObject = attestationObject
	return &this
}

// NewApiInteractionBindMfaPostRequestOneOf1ResponseWithDefaults instantiates a new ApiInteractionBindMfaPostRequestOneOf1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInteractionBindMfaPostRequestOneOf1ResponseWithDefaults() *ApiInteractionBindMfaPostRequestOneOf1Response {
	this := ApiInteractionBindMfaPostRequestOneOf1Response{}
	return &this
}

// GetClientDataJSON returns the ClientDataJSON field value
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetClientDataJSON() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientDataJSON
}

// GetClientDataJSONOk returns a tuple with the ClientDataJSON field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetClientDataJSONOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientDataJSON, true
}

// SetClientDataJSON sets field value
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) SetClientDataJSON(v string) {
	o.ClientDataJSON = v
}

// GetAttestationObject returns the AttestationObject field value
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetAttestationObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttestationObject
}

// GetAttestationObjectOk returns a tuple with the AttestationObject field value
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetAttestationObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttestationObject, true
}

// SetAttestationObject sets field value
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) SetAttestationObject(v string) {
	o.AttestationObject = v
}

// GetAuthenticatorData returns the AuthenticatorData field value if set, zero value otherwise.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetAuthenticatorData() string {
	if o == nil || IsNil(o.AuthenticatorData) {
		var ret string
		return ret
	}
	return *o.AuthenticatorData
}

// GetAuthenticatorDataOk returns a tuple with the AuthenticatorData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetAuthenticatorDataOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticatorData) {
		return nil, false
	}
	return o.AuthenticatorData, true
}

// HasAuthenticatorData returns a boolean if a field has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) HasAuthenticatorData() bool {
	if o != nil && !IsNil(o.AuthenticatorData) {
		return true
	}

	return false
}

// SetAuthenticatorData gets a reference to the given string and assigns it to the AuthenticatorData field.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) SetAuthenticatorData(v string) {
	o.AuthenticatorData = &v
}

// GetTransports returns the Transports field value if set, zero value otherwise.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetTransports() []string {
	if o == nil || IsNil(o.Transports) {
		var ret []string
		return ret
	}
	return o.Transports
}

// GetTransportsOk returns a tuple with the Transports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetTransportsOk() ([]string, bool) {
	if o == nil || IsNil(o.Transports) {
		return nil, false
	}
	return o.Transports, true
}

// HasTransports returns a boolean if a field has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) HasTransports() bool {
	if o != nil && !IsNil(o.Transports) {
		return true
	}

	return false
}

// SetTransports gets a reference to the given []string and assigns it to the Transports field.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) SetTransports(v []string) {
	o.Transports = v
}

// GetPublicKeyAlgorithm returns the PublicKeyAlgorithm field value if set, zero value otherwise.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetPublicKeyAlgorithm() float32 {
	if o == nil || IsNil(o.PublicKeyAlgorithm) {
		var ret float32
		return ret
	}
	return *o.PublicKeyAlgorithm
}

// GetPublicKeyAlgorithmOk returns a tuple with the PublicKeyAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetPublicKeyAlgorithmOk() (*float32, bool) {
	if o == nil || IsNil(o.PublicKeyAlgorithm) {
		return nil, false
	}
	return o.PublicKeyAlgorithm, true
}

// HasPublicKeyAlgorithm returns a boolean if a field has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) HasPublicKeyAlgorithm() bool {
	if o != nil && !IsNil(o.PublicKeyAlgorithm) {
		return true
	}

	return false
}

// SetPublicKeyAlgorithm gets a reference to the given float32 and assigns it to the PublicKeyAlgorithm field.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) SetPublicKeyAlgorithm(v float32) {
	o.PublicKeyAlgorithm = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *ApiInteractionBindMfaPostRequestOneOf1Response) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o ApiInteractionBindMfaPostRequestOneOf1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInteractionBindMfaPostRequestOneOf1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientDataJSON"] = o.ClientDataJSON
	toSerialize["attestationObject"] = o.AttestationObject
	if !IsNil(o.AuthenticatorData) {
		toSerialize["authenticatorData"] = o.AuthenticatorData
	}
	if !IsNil(o.Transports) {
		toSerialize["transports"] = o.Transports
	}
	if !IsNil(o.PublicKeyAlgorithm) {
		toSerialize["publicKeyAlgorithm"] = o.PublicKeyAlgorithm
	}
	if !IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	return toSerialize, nil
}

func (o *ApiInteractionBindMfaPostRequestOneOf1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientDataJSON",
		"attestationObject",
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

	varApiInteractionBindMfaPostRequestOneOf1Response := _ApiInteractionBindMfaPostRequestOneOf1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiInteractionBindMfaPostRequestOneOf1Response)

	if err != nil {
		return err
	}

	*o = ApiInteractionBindMfaPostRequestOneOf1Response(varApiInteractionBindMfaPostRequestOneOf1Response)

	return err
}

type NullableApiInteractionBindMfaPostRequestOneOf1Response struct {
	value *ApiInteractionBindMfaPostRequestOneOf1Response
	isSet bool
}

func (v NullableApiInteractionBindMfaPostRequestOneOf1Response) Get() *ApiInteractionBindMfaPostRequestOneOf1Response {
	return v.value
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf1Response) Set(val *ApiInteractionBindMfaPostRequestOneOf1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInteractionBindMfaPostRequestOneOf1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInteractionBindMfaPostRequestOneOf1Response(val *ApiInteractionBindMfaPostRequestOneOf1Response) *NullableApiInteractionBindMfaPostRequestOneOf1Response {
	return &NullableApiInteractionBindMfaPostRequestOneOf1Response{value: val, isSet: true}
}

func (v NullableApiInteractionBindMfaPostRequestOneOf1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInteractionBindMfaPostRequestOneOf1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

