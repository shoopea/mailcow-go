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

// checks if the UpdateMailboxRequestAttr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMailboxRequestAttr{}

// UpdateMailboxRequestAttr struct for UpdateMailboxRequestAttr
type UpdateMailboxRequestAttr struct {
	// is mailbox active or not
	Active *bool `json:"active,omitempty"`
	// force user to change password on next login
	ForcePwUpdate *bool `json:"force_pw_update,omitempty"`
	// Full name of the mailbox user
	Name *string `json:"name,omitempty"`
	// new mailbox password for confirmation
	Password2 *string `json:"password2,omitempty"`
	// new mailbox password
	Password *string `json:"password,omitempty"`
	// mailbox quota
	Quota *float32 `json:"quota,omitempty"`
	// list of allowed send from addresses
	SenderAcl map[string]interface{} `json:"sender_acl,omitempty"`
	// is access to SOGo webmail active or not
	SogoAccess *bool `json:"sogo_access,omitempty"`
}

// NewUpdateMailboxRequestAttr instantiates a new UpdateMailboxRequestAttr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMailboxRequestAttr() *UpdateMailboxRequestAttr {
	this := UpdateMailboxRequestAttr{}
	return &this
}

// NewUpdateMailboxRequestAttrWithDefaults instantiates a new UpdateMailboxRequestAttr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMailboxRequestAttrWithDefaults() *UpdateMailboxRequestAttr {
	this := UpdateMailboxRequestAttr{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateMailboxRequestAttr) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailboxRequestAttr) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateMailboxRequestAttr) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateMailboxRequestAttr) SetActive(v bool) {
	o.Active = &v
}

// GetForcePwUpdate returns the ForcePwUpdate field value if set, zero value otherwise.
func (o *UpdateMailboxRequestAttr) GetForcePwUpdate() bool {
	if o == nil || IsNil(o.ForcePwUpdate) {
		var ret bool
		return ret
	}
	return *o.ForcePwUpdate
}

// GetForcePwUpdateOk returns a tuple with the ForcePwUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailboxRequestAttr) GetForcePwUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.ForcePwUpdate) {
		return nil, false
	}
	return o.ForcePwUpdate, true
}

// HasForcePwUpdate returns a boolean if a field has been set.
func (o *UpdateMailboxRequestAttr) HasForcePwUpdate() bool {
	if o != nil && !IsNil(o.ForcePwUpdate) {
		return true
	}

	return false
}

// SetForcePwUpdate gets a reference to the given bool and assigns it to the ForcePwUpdate field.
func (o *UpdateMailboxRequestAttr) SetForcePwUpdate(v bool) {
	o.ForcePwUpdate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateMailboxRequestAttr) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailboxRequestAttr) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateMailboxRequestAttr) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateMailboxRequestAttr) SetName(v string) {
	o.Name = &v
}

// GetPassword2 returns the Password2 field value if set, zero value otherwise.
func (o *UpdateMailboxRequestAttr) GetPassword2() string {
	if o == nil || IsNil(o.Password2) {
		var ret string
		return ret
	}
	return *o.Password2
}

// GetPassword2Ok returns a tuple with the Password2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailboxRequestAttr) GetPassword2Ok() (*string, bool) {
	if o == nil || IsNil(o.Password2) {
		return nil, false
	}
	return o.Password2, true
}

// HasPassword2 returns a boolean if a field has been set.
func (o *UpdateMailboxRequestAttr) HasPassword2() bool {
	if o != nil && !IsNil(o.Password2) {
		return true
	}

	return false
}

// SetPassword2 gets a reference to the given string and assigns it to the Password2 field.
func (o *UpdateMailboxRequestAttr) SetPassword2(v string) {
	o.Password2 = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateMailboxRequestAttr) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailboxRequestAttr) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateMailboxRequestAttr) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateMailboxRequestAttr) SetPassword(v string) {
	o.Password = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *UpdateMailboxRequestAttr) GetQuota() float32 {
	if o == nil || IsNil(o.Quota) {
		var ret float32
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailboxRequestAttr) GetQuotaOk() (*float32, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *UpdateMailboxRequestAttr) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given float32 and assigns it to the Quota field.
func (o *UpdateMailboxRequestAttr) SetQuota(v float32) {
	o.Quota = &v
}

// GetSenderAcl returns the SenderAcl field value if set, zero value otherwise.
func (o *UpdateMailboxRequestAttr) GetSenderAcl() map[string]interface{} {
	if o == nil || IsNil(o.SenderAcl) {
		var ret map[string]interface{}
		return ret
	}
	return o.SenderAcl
}

// GetSenderAclOk returns a tuple with the SenderAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailboxRequestAttr) GetSenderAclOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SenderAcl) {
		return map[string]interface{}{}, false
	}
	return o.SenderAcl, true
}

// HasSenderAcl returns a boolean if a field has been set.
func (o *UpdateMailboxRequestAttr) HasSenderAcl() bool {
	if o != nil && !IsNil(o.SenderAcl) {
		return true
	}

	return false
}

// SetSenderAcl gets a reference to the given map[string]interface{} and assigns it to the SenderAcl field.
func (o *UpdateMailboxRequestAttr) SetSenderAcl(v map[string]interface{}) {
	o.SenderAcl = v
}

// GetSogoAccess returns the SogoAccess field value if set, zero value otherwise.
func (o *UpdateMailboxRequestAttr) GetSogoAccess() bool {
	if o == nil || IsNil(o.SogoAccess) {
		var ret bool
		return ret
	}
	return *o.SogoAccess
}

// GetSogoAccessOk returns a tuple with the SogoAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailboxRequestAttr) GetSogoAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.SogoAccess) {
		return nil, false
	}
	return o.SogoAccess, true
}

// HasSogoAccess returns a boolean if a field has been set.
func (o *UpdateMailboxRequestAttr) HasSogoAccess() bool {
	if o != nil && !IsNil(o.SogoAccess) {
		return true
	}

	return false
}

// SetSogoAccess gets a reference to the given bool and assigns it to the SogoAccess field.
func (o *UpdateMailboxRequestAttr) SetSogoAccess(v bool) {
	o.SogoAccess = &v
}

func (o UpdateMailboxRequestAttr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMailboxRequestAttr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ForcePwUpdate) {
		toSerialize["force_pw_update"] = o.ForcePwUpdate
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Password2) {
		toSerialize["password2"] = o.Password2
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !IsNil(o.SenderAcl) {
		toSerialize["sender_acl"] = o.SenderAcl
	}
	if !IsNil(o.SogoAccess) {
		toSerialize["sogo_access"] = o.SogoAccess
	}
	return toSerialize, nil
}

type NullableUpdateMailboxRequestAttr struct {
	value *UpdateMailboxRequestAttr
	isSet bool
}

func (v NullableUpdateMailboxRequestAttr) Get() *UpdateMailboxRequestAttr {
	return v.value
}

func (v *NullableUpdateMailboxRequestAttr) Set(val *UpdateMailboxRequestAttr) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMailboxRequestAttr) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMailboxRequestAttr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMailboxRequestAttr(val *UpdateMailboxRequestAttr) *NullableUpdateMailboxRequestAttr {
	return &NullableUpdateMailboxRequestAttr{value: val, isSet: true}
}

func (v NullableUpdateMailboxRequestAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMailboxRequestAttr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

