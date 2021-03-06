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

// UserPutBody struct for UserPutBody
type UserPutBody struct {
	Username *string `json:"username,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	Email *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Password *string `json:"password,omitempty"`
	Properties *map[string]interface{} `json:"properties,omitempty"`
}

// NewUserPutBody instantiates a new UserPutBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPutBody() *UserPutBody {
	this := UserPutBody{}
	return &this
}

// NewUserPutBodyWithDefaults instantiates a new UserPutBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPutBodyWithDefaults() *UserPutBody {
	this := UserPutBody{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserPutBody) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPutBody) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserPutBody) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserPutBody) SetUsername(v string) {
	o.Username = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserPutBody) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPutBody) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserPutBody) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserPutBody) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserPutBody) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPutBody) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserPutBody) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserPutBody) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserPutBody) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPutBody) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserPutBody) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserPutBody) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UserPutBody) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPutBody) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UserPutBody) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UserPutBody) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserPutBody) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPutBody) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserPutBody) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserPutBody) SetPassword(v string) {
	o.Password = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *UserPutBody) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPutBody) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UserPutBody) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *UserPutBody) SetProperties(v map[string]interface{}) {
	o.Properties = &v
}

func (o UserPutBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableUserPutBody struct {
	value *UserPutBody
	isSet bool
}

func (v NullableUserPutBody) Get() *UserPutBody {
	return v.value
}

func (v *NullableUserPutBody) Set(val *UserPutBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPutBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPutBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPutBody(val *UserPutBody) *NullableUserPutBody {
	return &NullableUserPutBody{value: val, isSet: true}
}

func (v NullableUserPutBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPutBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


