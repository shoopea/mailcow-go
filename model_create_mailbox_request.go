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

// checks if the CreateMailboxRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMailboxRequest{}

// CreateMailboxRequest struct for CreateMailboxRequest
type CreateMailboxRequest struct {
	// is mailbox active or not
	Active *bool `json:"active,omitempty"`
	// domain name
	Domain *string `json:"domain,omitempty"`
	// left part of email address
	LocalPart *string `json:"local_part,omitempty"`
	// Full name of the mailbox user
	Name *string `json:"name,omitempty"`
	// mailbox password for confirmation
	Password2 *string `json:"password2,omitempty"`
	// mailbox password
	Password *string `json:"password,omitempty"`
	// mailbox quota
	Quota *float32 `json:"quota,omitempty"`
	// forces the user to update its password on first login
	ForcePwUpdate *bool `json:"force_pw_update,omitempty"`
	// force inbound email tls encryption
	TlsEnforceIn *bool `json:"tls_enforce_in,omitempty"`
	// force oubound tmail tls encryption
	TlsEnforceOut *bool `json:"tls_enforce_out,omitempty"`
}

// NewCreateMailboxRequest instantiates a new CreateMailboxRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMailboxRequest() *CreateMailboxRequest {
	this := CreateMailboxRequest{}
	return &this
}

// NewCreateMailboxRequestWithDefaults instantiates a new CreateMailboxRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMailboxRequestWithDefaults() *CreateMailboxRequest {
	this := CreateMailboxRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateMailboxRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMailboxRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateMailboxRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CreateMailboxRequest) SetActive(v bool) {
	o.Active = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CreateMailboxRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMailboxRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CreateMailboxRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *CreateMailboxRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetLocalPart returns the LocalPart field value if set, zero value otherwise.
func (o *CreateMailboxRequest) GetLocalPart() string {
	if o == nil || IsNil(o.LocalPart) {
		var ret string
		return ret
	}
	return *o.LocalPart
}

// GetLocalPartOk returns a tuple with the LocalPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMailboxRequest) GetLocalPartOk() (*string, bool) {
	if o == nil || IsNil(o.LocalPart) {
		return nil, false
	}
	return o.LocalPart, true
}

// HasLocalPart returns a boolean if a field has been set.
func (o *CreateMailboxRequest) HasLocalPart() bool {
	if o != nil && !IsNil(o.LocalPart) {
		return true
	}

	return false
}

// SetLocalPart gets a reference to the given string and assigns it to the LocalPart field.
func (o *CreateMailboxRequest) SetLocalPart(v string) {
	o.LocalPart = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateMailboxRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMailboxRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateMailboxRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateMailboxRequest) SetName(v string) {
	o.Name = &v
}

// GetPassword2 returns the Password2 field value if set, zero value otherwise.
func (o *CreateMailboxRequest) GetPassword2() string {
	if o == nil || IsNil(o.Password2) {
		var ret string
		return ret
	}
	return *o.Password2
}

// GetPassword2Ok returns a tuple with the Password2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMailboxRequest) GetPassword2Ok() (*string, bool) {
	if o == nil || IsNil(o.Password2) {
		return nil, false
	}
	return o.Password2, true
}

// HasPassword2 returns a boolean if a field has been set.
func (o *CreateMailboxRequest) HasPassword2() bool {
	if o != nil && !IsNil(o.Password2) {
		return true
	}

	return false
}

// SetPassword2 gets a reference to the given string and assigns it to the Password2 field.
func (o *CreateMailboxRequest) SetPassword2(v string) {
	o.Password2 = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateMailboxRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMailboxRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateMailboxRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateMailboxRequest) SetPassword(v string) {
	o.Password = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *CreateMailboxRequest) GetQuota() float32 {
	if o == nil || IsNil(o.Quota) {
		var ret float32
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMailboxRequest) GetQuotaOk() (*float32, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *CreateMailboxRequest) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given float32 and assigns it to the Quota field.
func (o *CreateMailboxRequest) SetQuota(v float32) {
	o.Quota = &v
}

// GetForcePwUpdate returns the ForcePwUpdate field value if set, zero value otherwise.
func (o *CreateMailboxRequest) GetForcePwUpdate() bool {
	if o == nil || IsNil(o.ForcePwUpdate) {
		var ret bool
		return ret
	}
	return *o.ForcePwUpdate
}

// GetForcePwUpdateOk returns a tuple with the ForcePwUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMailboxRequest) GetForcePwUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.ForcePwUpdate) {
		return nil, false
	}
	return o.ForcePwUpdate, true
}

// HasForcePwUpdate returns a boolean if a field has been set.
func (o *CreateMailboxRequest) HasForcePwUpdate() bool {
	if o != nil && !IsNil(o.ForcePwUpdate) {
		return true
	}

	return false
}

// SetForcePwUpdate gets a reference to the given bool and assigns it to the ForcePwUpdate field.
func (o *CreateMailboxRequest) SetForcePwUpdate(v bool) {
	o.ForcePwUpdate = &v
}

// GetTlsEnforceIn returns the TlsEnforceIn field value if set, zero value otherwise.
func (o *CreateMailboxRequest) GetTlsEnforceIn() bool {
	if o == nil || IsNil(o.TlsEnforceIn) {
		var ret bool
		return ret
	}
	return *o.TlsEnforceIn
}

// GetTlsEnforceInOk returns a tuple with the TlsEnforceIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMailboxRequest) GetTlsEnforceInOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsEnforceIn) {
		return nil, false
	}
	return o.TlsEnforceIn, true
}

// HasTlsEnforceIn returns a boolean if a field has been set.
func (o *CreateMailboxRequest) HasTlsEnforceIn() bool {
	if o != nil && !IsNil(o.TlsEnforceIn) {
		return true
	}

	return false
}

// SetTlsEnforceIn gets a reference to the given bool and assigns it to the TlsEnforceIn field.
func (o *CreateMailboxRequest) SetTlsEnforceIn(v bool) {
	o.TlsEnforceIn = &v
}

// GetTlsEnforceOut returns the TlsEnforceOut field value if set, zero value otherwise.
func (o *CreateMailboxRequest) GetTlsEnforceOut() bool {
	if o == nil || IsNil(o.TlsEnforceOut) {
		var ret bool
		return ret
	}
	return *o.TlsEnforceOut
}

// GetTlsEnforceOutOk returns a tuple with the TlsEnforceOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMailboxRequest) GetTlsEnforceOutOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsEnforceOut) {
		return nil, false
	}
	return o.TlsEnforceOut, true
}

// HasTlsEnforceOut returns a boolean if a field has been set.
func (o *CreateMailboxRequest) HasTlsEnforceOut() bool {
	if o != nil && !IsNil(o.TlsEnforceOut) {
		return true
	}

	return false
}

// SetTlsEnforceOut gets a reference to the given bool and assigns it to the TlsEnforceOut field.
func (o *CreateMailboxRequest) SetTlsEnforceOut(v bool) {
	o.TlsEnforceOut = &v
}

func (o CreateMailboxRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMailboxRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.LocalPart) {
		toSerialize["local_part"] = o.LocalPart
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
	if !IsNil(o.ForcePwUpdate) {
		toSerialize["force_pw_update"] = o.ForcePwUpdate
	}
	if !IsNil(o.TlsEnforceIn) {
		toSerialize["tls_enforce_in"] = o.TlsEnforceIn
	}
	if !IsNil(o.TlsEnforceOut) {
		toSerialize["tls_enforce_out"] = o.TlsEnforceOut
	}
	return toSerialize, nil
}

type NullableCreateMailboxRequest struct {
	value *CreateMailboxRequest
	isSet bool
}

func (v NullableCreateMailboxRequest) Get() *CreateMailboxRequest {
	return v.value
}

func (v *NullableCreateMailboxRequest) Set(val *CreateMailboxRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMailboxRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMailboxRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMailboxRequest(val *CreateMailboxRequest) *NullableCreateMailboxRequest {
	return &NullableCreateMailboxRequest{value: val, isSet: true}
}

func (v NullableCreateMailboxRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMailboxRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


