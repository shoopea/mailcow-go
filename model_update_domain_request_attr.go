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

// checks if the UpdateDomainRequestAttr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDomainRequestAttr{}

// UpdateDomainRequestAttr struct for UpdateDomainRequestAttr
type UpdateDomainRequestAttr struct {
	// is domain active or not
	Active *bool `json:"active,omitempty"`
	// limit count of aliases associated with this domain
	Aliases *float32 `json:"aliases,omitempty"`
	// relay domain or not
	Backupmx *bool `json:"backupmx,omitempty"`
	// predefined mailbox quota in `add mailbox` form
	Defquota *float32 `json:"defquota,omitempty"`
	// Description of domain
	Description *string `json:"description,omitempty"`
	// is domain global address list active or not, it enables shared contacts accross domain in SOGo webmail
	Gal *bool `json:"gal,omitempty"`
	// limit count of mailboxes associated with this domain
	Mailboxes *float32 `json:"mailboxes,omitempty"`
	// maximum quota per mailbox
	Maxquota *float32 `json:"maxquota,omitempty"`
	// maximum quota for this domain (for all mailboxes in sum)
	Quota *float32 `json:"quota,omitempty"`
	// if not, them you have to create \"dummy\" mailbox for each address to relay
	RelayAllRecipients *bool `json:"relay_all_recipients,omitempty"`
	// Relay non-existing mailboxes only. Existing mailboxes will be delivered locally.
	RelayUnknownOnly *bool `json:"relay_unknown_only,omitempty"`
	// id of relayhost
	Relayhost *float32 `json:"relayhost,omitempty"`
	RlFrame *string `json:"rl_frame,omitempty"`
	// rate limit value
	RlValue *float32 `json:"rl_value,omitempty"`
	// tags for this Domain
	Tags []string `json:"tags,omitempty"`
}

// NewUpdateDomainRequestAttr instantiates a new UpdateDomainRequestAttr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDomainRequestAttr() *UpdateDomainRequestAttr {
	this := UpdateDomainRequestAttr{}
	return &this
}

// NewUpdateDomainRequestAttrWithDefaults instantiates a new UpdateDomainRequestAttr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDomainRequestAttrWithDefaults() *UpdateDomainRequestAttr {
	this := UpdateDomainRequestAttr{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateDomainRequestAttr) SetActive(v bool) {
	o.Active = &v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetAliases() float32 {
	if o == nil || IsNil(o.Aliases) {
		var ret float32
		return ret
	}
	return *o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetAliasesOk() (*float32, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given float32 and assigns it to the Aliases field.
func (o *UpdateDomainRequestAttr) SetAliases(v float32) {
	o.Aliases = &v
}

// GetBackupmx returns the Backupmx field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetBackupmx() bool {
	if o == nil || IsNil(o.Backupmx) {
		var ret bool
		return ret
	}
	return *o.Backupmx
}

// GetBackupmxOk returns a tuple with the Backupmx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetBackupmxOk() (*bool, bool) {
	if o == nil || IsNil(o.Backupmx) {
		return nil, false
	}
	return o.Backupmx, true
}

// HasBackupmx returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasBackupmx() bool {
	if o != nil && !IsNil(o.Backupmx) {
		return true
	}

	return false
}

// SetBackupmx gets a reference to the given bool and assigns it to the Backupmx field.
func (o *UpdateDomainRequestAttr) SetBackupmx(v bool) {
	o.Backupmx = &v
}

// GetDefquota returns the Defquota field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetDefquota() float32 {
	if o == nil || IsNil(o.Defquota) {
		var ret float32
		return ret
	}
	return *o.Defquota
}

// GetDefquotaOk returns a tuple with the Defquota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetDefquotaOk() (*float32, bool) {
	if o == nil || IsNil(o.Defquota) {
		return nil, false
	}
	return o.Defquota, true
}

// HasDefquota returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasDefquota() bool {
	if o != nil && !IsNil(o.Defquota) {
		return true
	}

	return false
}

// SetDefquota gets a reference to the given float32 and assigns it to the Defquota field.
func (o *UpdateDomainRequestAttr) SetDefquota(v float32) {
	o.Defquota = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateDomainRequestAttr) SetDescription(v string) {
	o.Description = &v
}

// GetGal returns the Gal field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetGal() bool {
	if o == nil || IsNil(o.Gal) {
		var ret bool
		return ret
	}
	return *o.Gal
}

// GetGalOk returns a tuple with the Gal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetGalOk() (*bool, bool) {
	if o == nil || IsNil(o.Gal) {
		return nil, false
	}
	return o.Gal, true
}

// HasGal returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasGal() bool {
	if o != nil && !IsNil(o.Gal) {
		return true
	}

	return false
}

// SetGal gets a reference to the given bool and assigns it to the Gal field.
func (o *UpdateDomainRequestAttr) SetGal(v bool) {
	o.Gal = &v
}

// GetMailboxes returns the Mailboxes field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetMailboxes() float32 {
	if o == nil || IsNil(o.Mailboxes) {
		var ret float32
		return ret
	}
	return *o.Mailboxes
}

// GetMailboxesOk returns a tuple with the Mailboxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetMailboxesOk() (*float32, bool) {
	if o == nil || IsNil(o.Mailboxes) {
		return nil, false
	}
	return o.Mailboxes, true
}

// HasMailboxes returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasMailboxes() bool {
	if o != nil && !IsNil(o.Mailboxes) {
		return true
	}

	return false
}

// SetMailboxes gets a reference to the given float32 and assigns it to the Mailboxes field.
func (o *UpdateDomainRequestAttr) SetMailboxes(v float32) {
	o.Mailboxes = &v
}

// GetMaxquota returns the Maxquota field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetMaxquota() float32 {
	if o == nil || IsNil(o.Maxquota) {
		var ret float32
		return ret
	}
	return *o.Maxquota
}

// GetMaxquotaOk returns a tuple with the Maxquota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetMaxquotaOk() (*float32, bool) {
	if o == nil || IsNil(o.Maxquota) {
		return nil, false
	}
	return o.Maxquota, true
}

// HasMaxquota returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasMaxquota() bool {
	if o != nil && !IsNil(o.Maxquota) {
		return true
	}

	return false
}

// SetMaxquota gets a reference to the given float32 and assigns it to the Maxquota field.
func (o *UpdateDomainRequestAttr) SetMaxquota(v float32) {
	o.Maxquota = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetQuota() float32 {
	if o == nil || IsNil(o.Quota) {
		var ret float32
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetQuotaOk() (*float32, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given float32 and assigns it to the Quota field.
func (o *UpdateDomainRequestAttr) SetQuota(v float32) {
	o.Quota = &v
}

// GetRelayAllRecipients returns the RelayAllRecipients field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetRelayAllRecipients() bool {
	if o == nil || IsNil(o.RelayAllRecipients) {
		var ret bool
		return ret
	}
	return *o.RelayAllRecipients
}

// GetRelayAllRecipientsOk returns a tuple with the RelayAllRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetRelayAllRecipientsOk() (*bool, bool) {
	if o == nil || IsNil(o.RelayAllRecipients) {
		return nil, false
	}
	return o.RelayAllRecipients, true
}

// HasRelayAllRecipients returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasRelayAllRecipients() bool {
	if o != nil && !IsNil(o.RelayAllRecipients) {
		return true
	}

	return false
}

// SetRelayAllRecipients gets a reference to the given bool and assigns it to the RelayAllRecipients field.
func (o *UpdateDomainRequestAttr) SetRelayAllRecipients(v bool) {
	o.RelayAllRecipients = &v
}

// GetRelayUnknownOnly returns the RelayUnknownOnly field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetRelayUnknownOnly() bool {
	if o == nil || IsNil(o.RelayUnknownOnly) {
		var ret bool
		return ret
	}
	return *o.RelayUnknownOnly
}

// GetRelayUnknownOnlyOk returns a tuple with the RelayUnknownOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetRelayUnknownOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.RelayUnknownOnly) {
		return nil, false
	}
	return o.RelayUnknownOnly, true
}

// HasRelayUnknownOnly returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasRelayUnknownOnly() bool {
	if o != nil && !IsNil(o.RelayUnknownOnly) {
		return true
	}

	return false
}

// SetRelayUnknownOnly gets a reference to the given bool and assigns it to the RelayUnknownOnly field.
func (o *UpdateDomainRequestAttr) SetRelayUnknownOnly(v bool) {
	o.RelayUnknownOnly = &v
}

// GetRelayhost returns the Relayhost field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetRelayhost() float32 {
	if o == nil || IsNil(o.Relayhost) {
		var ret float32
		return ret
	}
	return *o.Relayhost
}

// GetRelayhostOk returns a tuple with the Relayhost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetRelayhostOk() (*float32, bool) {
	if o == nil || IsNil(o.Relayhost) {
		return nil, false
	}
	return o.Relayhost, true
}

// HasRelayhost returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasRelayhost() bool {
	if o != nil && !IsNil(o.Relayhost) {
		return true
	}

	return false
}

// SetRelayhost gets a reference to the given float32 and assigns it to the Relayhost field.
func (o *UpdateDomainRequestAttr) SetRelayhost(v float32) {
	o.Relayhost = &v
}

// GetRlFrame returns the RlFrame field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetRlFrame() string {
	if o == nil || IsNil(o.RlFrame) {
		var ret string
		return ret
	}
	return *o.RlFrame
}

// GetRlFrameOk returns a tuple with the RlFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetRlFrameOk() (*string, bool) {
	if o == nil || IsNil(o.RlFrame) {
		return nil, false
	}
	return o.RlFrame, true
}

// HasRlFrame returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasRlFrame() bool {
	if o != nil && !IsNil(o.RlFrame) {
		return true
	}

	return false
}

// SetRlFrame gets a reference to the given string and assigns it to the RlFrame field.
func (o *UpdateDomainRequestAttr) SetRlFrame(v string) {
	o.RlFrame = &v
}

// GetRlValue returns the RlValue field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetRlValue() float32 {
	if o == nil || IsNil(o.RlValue) {
		var ret float32
		return ret
	}
	return *o.RlValue
}

// GetRlValueOk returns a tuple with the RlValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetRlValueOk() (*float32, bool) {
	if o == nil || IsNil(o.RlValue) {
		return nil, false
	}
	return o.RlValue, true
}

// HasRlValue returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasRlValue() bool {
	if o != nil && !IsNil(o.RlValue) {
		return true
	}

	return false
}

// SetRlValue gets a reference to the given float32 and assigns it to the RlValue field.
func (o *UpdateDomainRequestAttr) SetRlValue(v float32) {
	o.RlValue = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateDomainRequestAttr) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequestAttr) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateDomainRequestAttr) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateDomainRequestAttr) SetTags(v []string) {
	o.Tags = v
}

func (o UpdateDomainRequestAttr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDomainRequestAttr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !IsNil(o.Backupmx) {
		toSerialize["backupmx"] = o.Backupmx
	}
	if !IsNil(o.Defquota) {
		toSerialize["defquota"] = o.Defquota
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Gal) {
		toSerialize["gal"] = o.Gal
	}
	if !IsNil(o.Mailboxes) {
		toSerialize["mailboxes"] = o.Mailboxes
	}
	if !IsNil(o.Maxquota) {
		toSerialize["maxquota"] = o.Maxquota
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !IsNil(o.RelayAllRecipients) {
		toSerialize["relay_all_recipients"] = o.RelayAllRecipients
	}
	if !IsNil(o.RelayUnknownOnly) {
		toSerialize["relay_unknown_only"] = o.RelayUnknownOnly
	}
	if !IsNil(o.Relayhost) {
		toSerialize["relayhost"] = o.Relayhost
	}
	if !IsNil(o.RlFrame) {
		toSerialize["rl_frame"] = o.RlFrame
	}
	if !IsNil(o.RlValue) {
		toSerialize["rl_value"] = o.RlValue
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableUpdateDomainRequestAttr struct {
	value *UpdateDomainRequestAttr
	isSet bool
}

func (v NullableUpdateDomainRequestAttr) Get() *UpdateDomainRequestAttr {
	return v.value
}

func (v *NullableUpdateDomainRequestAttr) Set(val *UpdateDomainRequestAttr) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDomainRequestAttr) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDomainRequestAttr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDomainRequestAttr(val *UpdateDomainRequestAttr) *NullableUpdateDomainRequestAttr {
	return &NullableUpdateDomainRequestAttr{value: val, isSet: true}
}

func (v NullableUpdateDomainRequestAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDomainRequestAttr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


