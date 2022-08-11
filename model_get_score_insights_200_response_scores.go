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

// GetScoreInsights200ResponseScores struct for GetScoreInsights200ResponseScores
type GetScoreInsights200ResponseScores struct {
	Minimal *int32 `json:"minimal,omitempty"`
	Low *int32 `json:"low,omitempty"`
	Medium *int32 `json:"medium,omitempty"`
	High *int32 `json:"high,omitempty"`
	VeryHigh *int32 `json:"very_high,omitempty"`
}

// NewGetScoreInsights200ResponseScores instantiates a new GetScoreInsights200ResponseScores object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetScoreInsights200ResponseScores() *GetScoreInsights200ResponseScores {
	this := GetScoreInsights200ResponseScores{}
	return &this
}

// NewGetScoreInsights200ResponseScoresWithDefaults instantiates a new GetScoreInsights200ResponseScores object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetScoreInsights200ResponseScoresWithDefaults() *GetScoreInsights200ResponseScores {
	this := GetScoreInsights200ResponseScores{}
	return &this
}

// GetMinimal returns the Minimal field value if set, zero value otherwise.
func (o *GetScoreInsights200ResponseScores) GetMinimal() int32 {
	if o == nil || o.Minimal == nil {
		var ret int32
		return ret
	}
	return *o.Minimal
}

// GetMinimalOk returns a tuple with the Minimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScoreInsights200ResponseScores) GetMinimalOk() (*int32, bool) {
	if o == nil || o.Minimal == nil {
		return nil, false
	}
	return o.Minimal, true
}

// HasMinimal returns a boolean if a field has been set.
func (o *GetScoreInsights200ResponseScores) HasMinimal() bool {
	if o != nil && o.Minimal != nil {
		return true
	}

	return false
}

// SetMinimal gets a reference to the given int32 and assigns it to the Minimal field.
func (o *GetScoreInsights200ResponseScores) SetMinimal(v int32) {
	o.Minimal = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *GetScoreInsights200ResponseScores) GetLow() int32 {
	if o == nil || o.Low == nil {
		var ret int32
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScoreInsights200ResponseScores) GetLowOk() (*int32, bool) {
	if o == nil || o.Low == nil {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *GetScoreInsights200ResponseScores) HasLow() bool {
	if o != nil && o.Low != nil {
		return true
	}

	return false
}

// SetLow gets a reference to the given int32 and assigns it to the Low field.
func (o *GetScoreInsights200ResponseScores) SetLow(v int32) {
	o.Low = &v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *GetScoreInsights200ResponseScores) GetMedium() int32 {
	if o == nil || o.Medium == nil {
		var ret int32
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScoreInsights200ResponseScores) GetMediumOk() (*int32, bool) {
	if o == nil || o.Medium == nil {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *GetScoreInsights200ResponseScores) HasMedium() bool {
	if o != nil && o.Medium != nil {
		return true
	}

	return false
}

// SetMedium gets a reference to the given int32 and assigns it to the Medium field.
func (o *GetScoreInsights200ResponseScores) SetMedium(v int32) {
	o.Medium = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *GetScoreInsights200ResponseScores) GetHigh() int32 {
	if o == nil || o.High == nil {
		var ret int32
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScoreInsights200ResponseScores) GetHighOk() (*int32, bool) {
	if o == nil || o.High == nil {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *GetScoreInsights200ResponseScores) HasHigh() bool {
	if o != nil && o.High != nil {
		return true
	}

	return false
}

// SetHigh gets a reference to the given int32 and assigns it to the High field.
func (o *GetScoreInsights200ResponseScores) SetHigh(v int32) {
	o.High = &v
}

// GetVeryHigh returns the VeryHigh field value if set, zero value otherwise.
func (o *GetScoreInsights200ResponseScores) GetVeryHigh() int32 {
	if o == nil || o.VeryHigh == nil {
		var ret int32
		return ret
	}
	return *o.VeryHigh
}

// GetVeryHighOk returns a tuple with the VeryHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScoreInsights200ResponseScores) GetVeryHighOk() (*int32, bool) {
	if o == nil || o.VeryHigh == nil {
		return nil, false
	}
	return o.VeryHigh, true
}

// HasVeryHigh returns a boolean if a field has been set.
func (o *GetScoreInsights200ResponseScores) HasVeryHigh() bool {
	if o != nil && o.VeryHigh != nil {
		return true
	}

	return false
}

// SetVeryHigh gets a reference to the given int32 and assigns it to the VeryHigh field.
func (o *GetScoreInsights200ResponseScores) SetVeryHigh(v int32) {
	o.VeryHigh = &v
}

func (o GetScoreInsights200ResponseScores) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Minimal != nil {
		toSerialize["minimal"] = o.Minimal
	}
	if o.Low != nil {
		toSerialize["low"] = o.Low
	}
	if o.Medium != nil {
		toSerialize["medium"] = o.Medium
	}
	if o.High != nil {
		toSerialize["high"] = o.High
	}
	if o.VeryHigh != nil {
		toSerialize["very_high"] = o.VeryHigh
	}
	return json.Marshal(toSerialize)
}

type NullableGetScoreInsights200ResponseScores struct {
	value *GetScoreInsights200ResponseScores
	isSet bool
}

func (v NullableGetScoreInsights200ResponseScores) Get() *GetScoreInsights200ResponseScores {
	return v.value
}

func (v *NullableGetScoreInsights200ResponseScores) Set(val *GetScoreInsights200ResponseScores) {
	v.value = val
	v.isSet = true
}

func (v NullableGetScoreInsights200ResponseScores) IsSet() bool {
	return v.isSet
}

func (v *NullableGetScoreInsights200ResponseScores) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetScoreInsights200ResponseScores(val *GetScoreInsights200ResponseScores) *NullableGetScoreInsights200ResponseScores {
	return &NullableGetScoreInsights200ResponseScores{value: val, isSet: true}
}

func (v NullableGetScoreInsights200ResponseScores) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetScoreInsights200ResponseScores) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

