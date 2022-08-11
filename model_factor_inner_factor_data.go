/*
OneLogin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FactorInnerFactorData Array of factor specific properties. For example, the token or totp code for OneLogin Protect.
type FactorInnerFactorData struct {
	// The token which can be used to verify the factor registration.
	VerificationToken *string `json:"verification_token,omitempty"`
	// OTP Url that can be leveraged for any authenticator that supports the key uri format.
	TotpUrl *string `json:"totp_url,omitempty"`
}

// NewFactorInnerFactorData instantiates a new FactorInnerFactorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFactorInnerFactorData() *FactorInnerFactorData {
	this := FactorInnerFactorData{}
	return &this
}

// NewFactorInnerFactorDataWithDefaults instantiates a new FactorInnerFactorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFactorInnerFactorDataWithDefaults() *FactorInnerFactorData {
	this := FactorInnerFactorData{}
	return &this
}

// GetVerificationToken returns the VerificationToken field value if set, zero value otherwise.
func (o *FactorInnerFactorData) GetVerificationToken() string {
	if o == nil || o.VerificationToken == nil {
		var ret string
		return ret
	}
	return *o.VerificationToken
}

// GetVerificationTokenOk returns a tuple with the VerificationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FactorInnerFactorData) GetVerificationTokenOk() (*string, bool) {
	if o == nil || o.VerificationToken == nil {
		return nil, false
	}
	return o.VerificationToken, true
}

// HasVerificationToken returns a boolean if a field has been set.
func (o *FactorInnerFactorData) HasVerificationToken() bool {
	if o != nil && o.VerificationToken != nil {
		return true
	}

	return false
}

// SetVerificationToken gets a reference to the given string and assigns it to the VerificationToken field.
func (o *FactorInnerFactorData) SetVerificationToken(v string) {
	o.VerificationToken = &v
}

// GetTotpUrl returns the TotpUrl field value if set, zero value otherwise.
func (o *FactorInnerFactorData) GetTotpUrl() string {
	if o == nil || o.TotpUrl == nil {
		var ret string
		return ret
	}
	return *o.TotpUrl
}

// GetTotpUrlOk returns a tuple with the TotpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FactorInnerFactorData) GetTotpUrlOk() (*string, bool) {
	if o == nil || o.TotpUrl == nil {
		return nil, false
	}
	return o.TotpUrl, true
}

// HasTotpUrl returns a boolean if a field has been set.
func (o *FactorInnerFactorData) HasTotpUrl() bool {
	if o != nil && o.TotpUrl != nil {
		return true
	}

	return false
}

// SetTotpUrl gets a reference to the given string and assigns it to the TotpUrl field.
func (o *FactorInnerFactorData) SetTotpUrl(v string) {
	o.TotpUrl = &v
}

func (o FactorInnerFactorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VerificationToken != nil {
		toSerialize["verification_token"] = o.VerificationToken
	}
	if o.TotpUrl != nil {
		toSerialize["totp_url"] = o.TotpUrl
	}
	return json.Marshal(toSerialize)
}

type NullableFactorInnerFactorData struct {
	value *FactorInnerFactorData
	isSet bool
}

func (v NullableFactorInnerFactorData) Get() *FactorInnerFactorData {
	return v.value
}

func (v *NullableFactorInnerFactorData) Set(val *FactorInnerFactorData) {
	v.value = val
	v.isSet = true
}

func (v NullableFactorInnerFactorData) IsSet() bool {
	return v.isSet
}

func (v *NullableFactorInnerFactorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFactorInnerFactorData(val *FactorInnerFactorData) *NullableFactorInnerFactorData {
	return &NullableFactorInnerFactorData{value: val, isSet: true}
}

func (v NullableFactorInnerFactorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFactorInnerFactorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


