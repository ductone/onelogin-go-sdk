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

// RemoveRoleUsersRequest struct for RemoveRoleUsersRequest
type RemoveRoleUsersRequest struct {
	UserId []int32 `json:"user_id,omitempty"`
}

// NewRemoveRoleUsersRequest instantiates a new RemoveRoleUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveRoleUsersRequest() *RemoveRoleUsersRequest {
	this := RemoveRoleUsersRequest{}
	return &this
}

// NewRemoveRoleUsersRequestWithDefaults instantiates a new RemoveRoleUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveRoleUsersRequestWithDefaults() *RemoveRoleUsersRequest {
	this := RemoveRoleUsersRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *RemoveRoleUsersRequest) GetUserId() []int32 {
	if o == nil || o.UserId == nil {
		var ret []int32
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveRoleUsersRequest) GetUserIdOk() ([]int32, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *RemoveRoleUsersRequest) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given []int32 and assigns it to the UserId field.
func (o *RemoveRoleUsersRequest) SetUserId(v []int32) {
	o.UserId = v
}

func (o RemoveRoleUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveRoleUsersRequest struct {
	value *RemoveRoleUsersRequest
	isSet bool
}

func (v NullableRemoveRoleUsersRequest) Get() *RemoveRoleUsersRequest {
	return v.value
}

func (v *NullableRemoveRoleUsersRequest) Set(val *RemoveRoleUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveRoleUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveRoleUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveRoleUsersRequest(val *RemoveRoleUsersRequest) *NullableRemoveRoleUsersRequest {
	return &NullableRemoveRoleUsersRequest{value: val, isSet: true}
}

func (v NullableRemoveRoleUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveRoleUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

