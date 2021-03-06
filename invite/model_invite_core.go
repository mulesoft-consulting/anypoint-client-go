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

// InviteCore struct for InviteCore
type InviteCore struct {
	Code *string `json:"code,omitempty"`
	InvitedAt *string `json:"invited_at,omitempty"`
	InvitedEmail *string `json:"invited_email,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewInviteCore instantiates a new InviteCore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteCore() *InviteCore {
	this := InviteCore{}
	return &this
}

// NewInviteCoreWithDefaults instantiates a new InviteCore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteCoreWithDefaults() *InviteCore {
	this := InviteCore{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InviteCore) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteCore) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InviteCore) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InviteCore) SetCode(v string) {
	o.Code = &v
}

// GetInvitedAt returns the InvitedAt field value if set, zero value otherwise.
func (o *InviteCore) GetInvitedAt() string {
	if o == nil || o.InvitedAt == nil {
		var ret string
		return ret
	}
	return *o.InvitedAt
}

// GetInvitedAtOk returns a tuple with the InvitedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteCore) GetInvitedAtOk() (*string, bool) {
	if o == nil || o.InvitedAt == nil {
		return nil, false
	}
	return o.InvitedAt, true
}

// HasInvitedAt returns a boolean if a field has been set.
func (o *InviteCore) HasInvitedAt() bool {
	if o != nil && o.InvitedAt != nil {
		return true
	}

	return false
}

// SetInvitedAt gets a reference to the given string and assigns it to the InvitedAt field.
func (o *InviteCore) SetInvitedAt(v string) {
	o.InvitedAt = &v
}

// GetInvitedEmail returns the InvitedEmail field value if set, zero value otherwise.
func (o *InviteCore) GetInvitedEmail() string {
	if o == nil || o.InvitedEmail == nil {
		var ret string
		return ret
	}
	return *o.InvitedEmail
}

// GetInvitedEmailOk returns a tuple with the InvitedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteCore) GetInvitedEmailOk() (*string, bool) {
	if o == nil || o.InvitedEmail == nil {
		return nil, false
	}
	return o.InvitedEmail, true
}

// HasInvitedEmail returns a boolean if a field has been set.
func (o *InviteCore) HasInvitedEmail() bool {
	if o != nil && o.InvitedEmail != nil {
		return true
	}

	return false
}

// SetInvitedEmail gets a reference to the given string and assigns it to the InvitedEmail field.
func (o *InviteCore) SetInvitedEmail(v string) {
	o.InvitedEmail = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *InviteCore) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteCore) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *InviteCore) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *InviteCore) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InviteCore) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteCore) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InviteCore) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InviteCore) SetStatus(v string) {
	o.Status = &v
}

func (o InviteCore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.InvitedAt != nil {
		toSerialize["invited_at"] = o.InvitedAt
	}
	if o.InvitedEmail != nil {
		toSerialize["invited_email"] = o.InvitedEmail
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInviteCore struct {
	value *InviteCore
	isSet bool
}

func (v NullableInviteCore) Get() *InviteCore {
	return v.value
}

func (v *NullableInviteCore) Set(val *InviteCore) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteCore) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteCore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteCore(val *InviteCore) *NullableInviteCore {
	return &NullableInviteCore{value: val, isSet: true}
}

func (v NullableInviteCore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteCore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


