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

// MqAdvancedFeatures An explanation about the purpose of this instance.
type MqAdvancedFeatures struct {
	// An explanation about the purpose of this instance.
	Enabled bool `json:"enabled"`
}

// NewMqAdvancedFeatures instantiates a new MqAdvancedFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMqAdvancedFeatures(enabled bool) *MqAdvancedFeatures {
	this := MqAdvancedFeatures{}
	this.Enabled = enabled
	return &this
}

// NewMqAdvancedFeaturesWithDefaults instantiates a new MqAdvancedFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMqAdvancedFeaturesWithDefaults() *MqAdvancedFeatures {
	this := MqAdvancedFeatures{}
	var enabled bool = false
	this.Enabled = enabled
	return &this
}

// GetEnabled returns the Enabled field value
func (o *MqAdvancedFeatures) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *MqAdvancedFeatures) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *MqAdvancedFeatures) SetEnabled(v bool) {
	o.Enabled = v
}

func (o MqAdvancedFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableMqAdvancedFeatures struct {
	value *MqAdvancedFeatures
	isSet bool
}

func (v NullableMqAdvancedFeatures) Get() *MqAdvancedFeatures {
	return v.value
}

func (v *NullableMqAdvancedFeatures) Set(val *MqAdvancedFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableMqAdvancedFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableMqAdvancedFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMqAdvancedFeatures(val *MqAdvancedFeatures) *NullableMqAdvancedFeatures {
	return &NullableMqAdvancedFeatures{value: val, isSet: true}
}

func (v NullableMqAdvancedFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMqAdvancedFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


