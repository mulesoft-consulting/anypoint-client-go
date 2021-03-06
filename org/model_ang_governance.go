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

// AngGovernance An explanation about the purpose of this instance.
type AngGovernance struct {
	// An explanation about the purpose of this instance.
	Level int32 `json:"level"`
}

// NewAngGovernance instantiates a new AngGovernance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAngGovernance(level int32) *AngGovernance {
	this := AngGovernance{}
	this.Level = level
	return &this
}

// NewAngGovernanceWithDefaults instantiates a new AngGovernance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAngGovernanceWithDefaults() *AngGovernance {
	this := AngGovernance{}
	var level int32 = 0
	this.Level = level
	return &this
}

// GetLevel returns the Level field value
func (o *AngGovernance) GetLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *AngGovernance) GetLevelOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *AngGovernance) SetLevel(v int32) {
	o.Level = v
}

func (o AngGovernance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableAngGovernance struct {
	value *AngGovernance
	isSet bool
}

func (v NullableAngGovernance) Get() *AngGovernance {
	return v.value
}

func (v *NullableAngGovernance) Set(val *AngGovernance) {
	v.value = val
	v.isSet = true
}

func (v NullableAngGovernance) IsSet() bool {
	return v.isSet
}

func (v *NullableAngGovernance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAngGovernance(val *AngGovernance) *NullableAngGovernance {
	return &NullableAngGovernance{value: val, isSet: true}
}

func (v NullableAngGovernance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAngGovernance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


