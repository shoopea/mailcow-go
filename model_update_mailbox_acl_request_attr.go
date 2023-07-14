/*
mailcow API

mailcow is complete e-mailing solution with advanced antispam, antivirus, nice UI and API.  In order to use this API you have to create a API key and add your IP address to the whitelist of allowed IPs this can be done by logging into the Mailcow UI using your admin account, then go to Configuration > Access > Edit administrator details > API. There you will find a collapsed API menu.  There are two types of API keys   - The read only key can only be used for all get endpoints   - The read write key can be used for all endpoints

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mailcow

import (
	"encoding/json"
)

// checks if the UpdateMailboxACLRequestAttr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMailboxACLRequestAttr{}

// UpdateMailboxACLRequestAttr struct for UpdateMailboxACLRequestAttr
type UpdateMailboxACLRequestAttr struct {
	// contains a list of active user acls
	UserAcl map[string]interface{} `json:"user_acl,omitempty"`
}

// NewUpdateMailboxACLRequestAttr instantiates a new UpdateMailboxACLRequestAttr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMailboxACLRequestAttr() *UpdateMailboxACLRequestAttr {
	this := UpdateMailboxACLRequestAttr{}
	return &this
}

// NewUpdateMailboxACLRequestAttrWithDefaults instantiates a new UpdateMailboxACLRequestAttr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMailboxACLRequestAttrWithDefaults() *UpdateMailboxACLRequestAttr {
	this := UpdateMailboxACLRequestAttr{}
	return &this
}

// GetUserAcl returns the UserAcl field value if set, zero value otherwise.
func (o *UpdateMailboxACLRequestAttr) GetUserAcl() map[string]interface{} {
	if o == nil || IsNil(o.UserAcl) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserAcl
}

// GetUserAclOk returns a tuple with the UserAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailboxACLRequestAttr) GetUserAclOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserAcl) {
		return map[string]interface{}{}, false
	}
	return o.UserAcl, true
}

// HasUserAcl returns a boolean if a field has been set.
func (o *UpdateMailboxACLRequestAttr) HasUserAcl() bool {
	if o != nil && !IsNil(o.UserAcl) {
		return true
	}

	return false
}

// SetUserAcl gets a reference to the given map[string]interface{} and assigns it to the UserAcl field.
func (o *UpdateMailboxACLRequestAttr) SetUserAcl(v map[string]interface{}) {
	o.UserAcl = v
}

func (o UpdateMailboxACLRequestAttr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMailboxACLRequestAttr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserAcl) {
		toSerialize["user_acl"] = o.UserAcl
	}
	return toSerialize, nil
}

type NullableUpdateMailboxACLRequestAttr struct {
	value *UpdateMailboxACLRequestAttr
	isSet bool
}

func (v NullableUpdateMailboxACLRequestAttr) Get() *UpdateMailboxACLRequestAttr {
	return v.value
}

func (v *NullableUpdateMailboxACLRequestAttr) Set(val *UpdateMailboxACLRequestAttr) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMailboxACLRequestAttr) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMailboxACLRequestAttr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMailboxACLRequestAttr(val *UpdateMailboxACLRequestAttr) *NullableUpdateMailboxACLRequestAttr {
	return &NullableUpdateMailboxACLRequestAttr{value: val, isSet: true}
}

func (v NullableUpdateMailboxACLRequestAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMailboxACLRequestAttr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

