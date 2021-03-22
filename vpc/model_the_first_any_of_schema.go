/*
 * VPC API
 *
 * Description of the VPC API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vpc

import (
	"encoding/json"
)

// TheFirstAnyOfSchema struct for TheFirstAnyOfSchema
type TheFirstAnyOfSchema struct {
	CIDR string `json:"CIDR"`
	NextHop string `json:"Next Hop"`
}

// NewTheFirstAnyOfSchema instantiates a new TheFirstAnyOfSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTheFirstAnyOfSchema(cIDR string, nextHop string) *TheFirstAnyOfSchema {
	this := TheFirstAnyOfSchema{}
	this.CIDR = cIDR
	this.NextHop = nextHop
	return &this
}

// NewTheFirstAnyOfSchemaWithDefaults instantiates a new TheFirstAnyOfSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTheFirstAnyOfSchemaWithDefaults() *TheFirstAnyOfSchema {
	this := TheFirstAnyOfSchema{}
	var cIDR string = ""
	this.CIDR = cIDR
	var nextHop string = ""
	this.NextHop = nextHop
	return &this
}

// GetCIDR returns the CIDR field value
func (o *TheFirstAnyOfSchema) GetCIDR() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CIDR
}

// GetCIDROk returns a tuple with the CIDR field value
// and a boolean to check if the value has been set.
func (o *TheFirstAnyOfSchema) GetCIDROk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CIDR, true
}

// SetCIDR sets field value
func (o *TheFirstAnyOfSchema) SetCIDR(v string) {
	o.CIDR = v
}

// GetNextHop returns the NextHop field value
func (o *TheFirstAnyOfSchema) GetNextHop() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextHop
}

// GetNextHopOk returns a tuple with the NextHop field value
// and a boolean to check if the value has been set.
func (o *TheFirstAnyOfSchema) GetNextHopOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NextHop, true
}

// SetNextHop sets field value
func (o *TheFirstAnyOfSchema) SetNextHop(v string) {
	o.NextHop = v
}

func (o TheFirstAnyOfSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["CIDR"] = o.CIDR
	}
	if true {
		toSerialize["Next Hop"] = o.NextHop
	}
	return json.Marshal(toSerialize)
}

type NullableTheFirstAnyOfSchema struct {
	value *TheFirstAnyOfSchema
	isSet bool
}

func (v NullableTheFirstAnyOfSchema) Get() *TheFirstAnyOfSchema {
	return v.value
}

func (v *NullableTheFirstAnyOfSchema) Set(val *TheFirstAnyOfSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableTheFirstAnyOfSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableTheFirstAnyOfSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTheFirstAnyOfSchema(val *TheFirstAnyOfSchema) *NullableTheFirstAnyOfSchema {
	return &NullableTheFirstAnyOfSchema{value: val, isSet: true}
}

func (v NullableTheFirstAnyOfSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTheFirstAnyOfSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


