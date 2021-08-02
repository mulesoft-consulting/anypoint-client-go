/*
 * Team Members API
 *
 * Description of the Team Members API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package team_members

import (
	"encoding/json"
)

// TeamMemberCollection struct for TeamMemberCollection
type TeamMemberCollection struct {
	Data *[]TeamMember `json:"data,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewTeamMemberCollection instantiates a new TeamMemberCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamMemberCollection() *TeamMemberCollection {
	this := TeamMemberCollection{}
	return &this
}

// NewTeamMemberCollectionWithDefaults instantiates a new TeamMemberCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamMemberCollectionWithDefaults() *TeamMemberCollection {
	this := TeamMemberCollection{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TeamMemberCollection) GetData() []TeamMember {
	if o == nil || o.Data == nil {
		var ret []TeamMember
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamMemberCollection) GetDataOk() (*[]TeamMember, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TeamMemberCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []TeamMember and assigns it to the Data field.
func (o *TeamMemberCollection) SetData(v []TeamMember) {
	o.Data = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TeamMemberCollection) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamMemberCollection) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TeamMemberCollection) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *TeamMemberCollection) SetTotal(v int32) {
	o.Total = &v
}

func (o TeamMemberCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableTeamMemberCollection struct {
	value *TeamMemberCollection
	isSet bool
}

func (v NullableTeamMemberCollection) Get() *TeamMemberCollection {
	return v.value
}

func (v *NullableTeamMemberCollection) Set(val *TeamMemberCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamMemberCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamMemberCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamMemberCollection(val *TeamMemberCollection) *NullableTeamMemberCollection {
	return &NullableTeamMemberCollection{value: val, isSet: true}
}

func (v NullableTeamMemberCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamMemberCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


