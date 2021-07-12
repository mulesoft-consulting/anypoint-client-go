/*
 * User API
 *
 * Description of the User API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package user

import (
	"encoding/json"
)

// UserId struct for UserId
type UserId struct {
	// The user Id
	Id string `json:"id"`
}

// NewUserId instantiates a new UserId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserId(id string) *UserId {
	this := UserId{}
	this.Id = id
	return &this
}

// NewUserIdWithDefaults instantiates a new UserId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdWithDefaults() *UserId {
	this := UserId{}
	return &this
}

// GetId returns the Id field value
func (o *UserId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserId) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserId) SetId(v string) {
	o.Id = v
}

func (o UserId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableUserId struct {
	value *UserId
	isSet bool
}

func (v NullableUserId) Get() *UserId {
	return v.value
}

func (v *NullableUserId) Set(val *UserId) {
	v.value = val
	v.isSet = true
}

func (v NullableUserId) IsSet() bool {
	return v.isSet
}

func (v *NullableUserId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserId(val *UserId) *NullableUserId {
	return &NullableUserId{value: val, isSet: true}
}

func (v NullableUserId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

