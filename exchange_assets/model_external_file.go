/*
 * Exchange Assets
 *
 * Description of the Exchange Assets API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package exchange_assets

import (
	"encoding/json"
)

// ExternalFile struct for ExternalFile
type ExternalFile struct {
	Url *string `json:"url,omitempty"`
}

// NewExternalFile instantiates a new ExternalFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalFile() *ExternalFile {
	this := ExternalFile{}
	return &this
}

// NewExternalFileWithDefaults instantiates a new ExternalFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalFileWithDefaults() *ExternalFile {
	this := ExternalFile{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ExternalFile) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalFile) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ExternalFile) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ExternalFile) SetUrl(v string) {
	o.Url = &v
}

func (o ExternalFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableExternalFile struct {
	value *ExternalFile
	isSet bool
}

func (v NullableExternalFile) Get() *ExternalFile {
	return v.value
}

func (v *NullableExternalFile) Set(val *ExternalFile) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalFile) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalFile(val *ExternalFile) *NullableExternalFile {
	return &NullableExternalFile{value: val, isSet: true}
}

func (v NullableExternalFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

