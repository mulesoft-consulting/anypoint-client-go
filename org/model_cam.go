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

// Cam An explanation about the purpose of this instance.
type Cam struct {
	// An explanation about the purpose of this instance.
	Enabled bool `json:"enabled"`
}

// NewCam instantiates a new Cam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCam(enabled bool) *Cam {
	this := Cam{}
	this.Enabled = enabled
	return &this
}

// NewCamWithDefaults instantiates a new Cam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCamWithDefaults() *Cam {
	this := Cam{}
	var enabled bool = false
	this.Enabled = enabled
	return &this
}

// GetEnabled returns the Enabled field value
func (o *Cam) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Cam) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Cam) SetEnabled(v bool) {
	o.Enabled = v
}

func (o Cam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableCam struct {
	value *Cam
	isSet bool
}

func (v NullableCam) Get() *Cam {
	return v.value
}

func (v *NullableCam) Set(val *Cam) {
	v.value = val
	v.isSet = true
}

func (v NullableCam) IsSet() bool {
	return v.isSet
}

func (v *NullableCam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCam(val *Cam) *NullableCam {
	return &NullableCam{value: val, isSet: true}
}

func (v NullableCam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

