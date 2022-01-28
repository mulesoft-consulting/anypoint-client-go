/*
 * Identity Provider Management API
 *
 * Description of Identity Provider API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idp

import (
	"encoding/json"
)

// ServiceProvider1 struct for ServiceProvider1
type ServiceProvider1 struct {
	Urls *Urls4 `json:"urls,omitempty"`
}

// NewServiceProvider1 instantiates a new ServiceProvider1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProvider1() *ServiceProvider1 {
	this := ServiceProvider1{}
	return &this
}

// NewServiceProvider1WithDefaults instantiates a new ServiceProvider1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProvider1WithDefaults() *ServiceProvider1 {
	this := ServiceProvider1{}
	return &this
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *ServiceProvider1) GetUrls() Urls4 {
	if o == nil || o.Urls == nil {
		var ret Urls4
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider1) GetUrlsOk() (*Urls4, bool) {
	if o == nil || o.Urls == nil {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *ServiceProvider1) HasUrls() bool {
	if o != nil && o.Urls != nil {
		return true
	}

	return false
}

// SetUrls gets a reference to the given Urls4 and assigns it to the Urls field.
func (o *ServiceProvider1) SetUrls(v Urls4) {
	o.Urls = &v
}

func (o ServiceProvider1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Urls != nil {
		toSerialize["urls"] = o.Urls
	}
	return json.Marshal(toSerialize)
}

type NullableServiceProvider1 struct {
	value *ServiceProvider1
	isSet bool
}

func (v NullableServiceProvider1) Get() *ServiceProvider1 {
	return v.value
}

func (v *NullableServiceProvider1) Set(val *ServiceProvider1) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProvider1) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProvider1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProvider1(val *ServiceProvider1) *NullableServiceProvider1 {
	return &NullableServiceProvider1{value: val, isSet: true}
}

func (v NullableServiceProvider1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProvider1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

