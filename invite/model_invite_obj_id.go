/*
 * Invite API
 *
 * Description of the Invite API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invite

import (
	"encoding/json"
)

// InviteObjId struct for InviteObjId
type InviteObjId struct {
	Id string `json:"id"`
}

// NewInviteObjId instantiates a new InviteObjId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteObjId(id string) *InviteObjId {
	this := InviteObjId{}
	this.Id = id
	return &this
}

// NewInviteObjIdWithDefaults instantiates a new InviteObjId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteObjIdWithDefaults() *InviteObjId {
	this := InviteObjId{}
	return &this
}

// GetId returns the Id field value
func (o *InviteObjId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InviteObjId) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InviteObjId) SetId(v string) {
	o.Id = v
}

func (o InviteObjId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInviteObjId struct {
	value *InviteObjId
	isSet bool
}

func (v NullableInviteObjId) Get() *InviteObjId {
	return v.value
}

func (v *NullableInviteObjId) Set(val *InviteObjId) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteObjId) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteObjId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteObjId(val *InviteObjId) *NullableInviteObjId {
	return &NullableInviteObjId{value: val, isSet: true}
}

func (v NullableInviteObjId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteObjId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

