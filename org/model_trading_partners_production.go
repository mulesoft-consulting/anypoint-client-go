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

// TradingPartnersProduction An explanation about the purpose of this instance.
type TradingPartnersProduction struct {
	// An explanation about the purpose of this instance.
	Assigned int32 `json:"assigned"`
}

// NewTradingPartnersProduction instantiates a new TradingPartnersProduction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingPartnersProduction(assigned int32) *TradingPartnersProduction {
	this := TradingPartnersProduction{}
	this.Assigned = assigned
	return &this
}

// NewTradingPartnersProductionWithDefaults instantiates a new TradingPartnersProduction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingPartnersProductionWithDefaults() *TradingPartnersProduction {
	this := TradingPartnersProduction{}
	var assigned int32 = 0
	this.Assigned = assigned
	return &this
}

// GetAssigned returns the Assigned field value
func (o *TradingPartnersProduction) GetAssigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value
// and a boolean to check if the value has been set.
func (o *TradingPartnersProduction) GetAssignedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Assigned, true
}

// SetAssigned sets field value
func (o *TradingPartnersProduction) SetAssigned(v int32) {
	o.Assigned = v
}

func (o TradingPartnersProduction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assigned"] = o.Assigned
	}
	return json.Marshal(toSerialize)
}

type NullableTradingPartnersProduction struct {
	value *TradingPartnersProduction
	isSet bool
}

func (v NullableTradingPartnersProduction) Get() *TradingPartnersProduction {
	return v.value
}

func (v *NullableTradingPartnersProduction) Set(val *TradingPartnersProduction) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingPartnersProduction) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingPartnersProduction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingPartnersProduction(val *TradingPartnersProduction) *NullableTradingPartnersProduction {
	return &NullableTradingPartnersProduction{value: val, isSet: true}
}

func (v NullableTradingPartnersProduction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingPartnersProduction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


