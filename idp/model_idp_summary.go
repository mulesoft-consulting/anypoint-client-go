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

// IdpSummary struct for IdpSummary
type IdpSummary struct {
	ProviderId *string `json:"provider_id,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *ModelType `json:"type,omitempty"`
}

// NewIdpSummary instantiates a new IdpSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpSummary() *IdpSummary {
	this := IdpSummary{}
	return &this
}

// NewIdpSummaryWithDefaults instantiates a new IdpSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpSummaryWithDefaults() *IdpSummary {
	this := IdpSummary{}
	return &this
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *IdpSummary) GetProviderId() string {
	if o == nil || o.ProviderId == nil {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpSummary) GetProviderIdOk() (*string, bool) {
	if o == nil || o.ProviderId == nil {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *IdpSummary) HasProviderId() bool {
	if o != nil && o.ProviderId != nil {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *IdpSummary) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *IdpSummary) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpSummary) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *IdpSummary) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *IdpSummary) SetOrgId(v string) {
	o.OrgId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdpSummary) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpSummary) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdpSummary) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdpSummary) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdpSummary) GetType() ModelType {
	if o == nil || o.Type == nil {
		var ret ModelType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpSummary) GetTypeOk() (*ModelType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdpSummary) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ModelType and assigns it to the Type field.
func (o *IdpSummary) SetType(v ModelType) {
	o.Type = &v
}

func (o IdpSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProviderId != nil {
		toSerialize["provider_id"] = o.ProviderId
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIdpSummary struct {
	value *IdpSummary
	isSet bool
}

func (v NullableIdpSummary) Get() *IdpSummary {
	return v.value
}

func (v *NullableIdpSummary) Set(val *IdpSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpSummary(val *IdpSummary) *NullableIdpSummary {
	return &NullableIdpSummary{value: val, isSet: true}
}

func (v NullableIdpSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


