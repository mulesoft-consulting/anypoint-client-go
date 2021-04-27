/*
 * Organization API
 *
 * Description of the Organization API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// Messaging An explanation about the purpose of this instance.
type Messaging struct {
	// An explanation about the purpose of this instance.
	Assigned int32 `json:"assigned"`
}

// NewMessaging instantiates a new Messaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessaging(assigned int32) *Messaging {
	this := Messaging{}
	this.Assigned = assigned
	return &this
}

// NewMessagingWithDefaults instantiates a new Messaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagingWithDefaults() *Messaging {
	this := Messaging{}
	var assigned int32 = 0
	this.Assigned = assigned
	return &this
}

// GetAssigned returns the Assigned field value
func (o *Messaging) GetAssigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value
// and a boolean to check if the value has been set.
func (o *Messaging) GetAssignedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Assigned, true
}

// SetAssigned sets field value
func (o *Messaging) SetAssigned(v int32) {
	o.Assigned = v
}

func (o Messaging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assigned"] = o.Assigned
	}
	return json.Marshal(toSerialize)
}

type NullableMessaging struct {
	value *Messaging
	isSet bool
}

func (v NullableMessaging) Get() *Messaging {
	return v.value
}

func (v *NullableMessaging) Set(val *Messaging) {
	v.value = val
	v.isSet = true
}

func (v NullableMessaging) IsSet() bool {
	return v.isSet
}

func (v *NullableMessaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessaging(val *Messaging) *NullableMessaging {
	return &NullableMessaging{value: val, isSet: true}
}

func (v NullableMessaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

