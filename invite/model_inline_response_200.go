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

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	Data *[]Invite `json:"data,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewInlineResponse200 instantiates a new InlineResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200WithDefaults() *InlineResponse200 {
	this := InlineResponse200{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse200) GetData() []Invite {
	if o == nil || o.Data == nil {
		var ret []Invite
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetDataOk() (*[]Invite, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse200) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Invite and assigns it to the Data field.
func (o *InlineResponse200) SetData(v []Invite) {
	o.Data = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200) SetTotal(v int32) {
	o.Total = &v
}

func (o InlineResponse200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200 struct {
	value *InlineResponse200
	isSet bool
}

func (v NullableInlineResponse200) Get() *InlineResponse200 {
	return v.value
}

func (v *NullableInlineResponse200) Set(val *InlineResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200(val *InlineResponse200) *NullableInlineResponse200 {
	return &NullableInlineResponse200{value: val, isSet: true}
}

func (v NullableInlineResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


