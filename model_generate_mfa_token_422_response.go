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

// GenerateMfaToken422Response struct for GenerateMfaToken422Response
type GenerateMfaToken422Response struct {
	StatusCode *int32 `json:"statusCode,omitempty"`
	Name *string `json:"name,omitempty"`
	Message *string `json:"message,omitempty"`
	Details *GenerateMfaToken422ResponseDetails `json:"details,omitempty"`
}

// NewGenerateMfaToken422Response instantiates a new GenerateMfaToken422Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateMfaToken422Response() *GenerateMfaToken422Response {
	this := GenerateMfaToken422Response{}
	return &this
}

// NewGenerateMfaToken422ResponseWithDefaults instantiates a new GenerateMfaToken422Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateMfaToken422ResponseWithDefaults() *GenerateMfaToken422Response {
	this := GenerateMfaToken422Response{}
	return &this
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *GenerateMfaToken422Response) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateMfaToken422Response) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *GenerateMfaToken422Response) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *GenerateMfaToken422Response) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GenerateMfaToken422Response) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateMfaToken422Response) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GenerateMfaToken422Response) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GenerateMfaToken422Response) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GenerateMfaToken422Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateMfaToken422Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GenerateMfaToken422Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GenerateMfaToken422Response) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GenerateMfaToken422Response) GetDetails() GenerateMfaToken422ResponseDetails {
	if o == nil || o.Details == nil {
		var ret GenerateMfaToken422ResponseDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateMfaToken422Response) GetDetailsOk() (*GenerateMfaToken422ResponseDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GenerateMfaToken422Response) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given GenerateMfaToken422ResponseDetails and assigns it to the Details field.
func (o *GenerateMfaToken422Response) SetDetails(v GenerateMfaToken422ResponseDetails) {
	o.Details = &v
}

func (o GenerateMfaToken422Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StatusCode != nil {
		toSerialize["statusCode"] = o.StatusCode
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateMfaToken422Response struct {
	value *GenerateMfaToken422Response
	isSet bool
}

func (v NullableGenerateMfaToken422Response) Get() *GenerateMfaToken422Response {
	return v.value
}

func (v *NullableGenerateMfaToken422Response) Set(val *GenerateMfaToken422Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateMfaToken422Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateMfaToken422Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateMfaToken422Response(val *GenerateMfaToken422Response) *NullableGenerateMfaToken422Response {
	return &NullableGenerateMfaToken422Response{value: val, isSet: true}
}

func (v NullableGenerateMfaToken422Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateMfaToken422Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


