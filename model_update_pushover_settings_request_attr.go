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

// checks if the UpdatePushoverSettingsRequestAttr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePushoverSettingsRequestAttr{}

// UpdatePushoverSettingsRequestAttr struct for UpdatePushoverSettingsRequestAttr
type UpdatePushoverSettingsRequestAttr struct {
	// Enables pushover 1 disable pushover 0
	Active *float32 `json:"active,omitempty"`
	// Send the Push with High priority
	EvaluateXPrio *float32 `json:"evaluate_x_prio,omitempty"`
	// Pushover key
	Key *string `json:"key,omitempty"`
	// Only send push for prio mails
	OnlyXPrio *float32 `json:"only_x_prio,omitempty"`
	// Set notification sound
	Sound *string `json:"sound,omitempty"`
	// Only send push for emails from these senders
	Senders *string `json:"senders,omitempty"`
	// Regex to match senders for which a push will be send
	SendersRegex *string `json:"senders_regex,omitempty"`
	// Custom push noficiation text
	Text *string `json:"text,omitempty"`
	// Push title
	Title *string `json:"title,omitempty"`
	// Pushover token
	Token *string `json:"token,omitempty"`
}

// NewUpdatePushoverSettingsRequestAttr instantiates a new UpdatePushoverSettingsRequestAttr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePushoverSettingsRequestAttr() *UpdatePushoverSettingsRequestAttr {
	this := UpdatePushoverSettingsRequestAttr{}
	return &this
}

// NewUpdatePushoverSettingsRequestAttrWithDefaults instantiates a new UpdatePushoverSettingsRequestAttr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePushoverSettingsRequestAttrWithDefaults() *UpdatePushoverSettingsRequestAttr {
	this := UpdatePushoverSettingsRequestAttr{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdatePushoverSettingsRequestAttr) GetActive() float32 {
	if o == nil || IsNil(o.Active) {
		var ret float32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePushoverSettingsRequestAttr) GetActiveOk() (*float32, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdatePushoverSettingsRequestAttr) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given float32 and assigns it to the Active field.
func (o *UpdatePushoverSettingsRequestAttr) SetActive(v float32) {
	o.Active = &v
}

// GetEvaluateXPrio returns the EvaluateXPrio field value if set, zero value otherwise.
func (o *UpdatePushoverSettingsRequestAttr) GetEvaluateXPrio() float32 {
	if o == nil || IsNil(o.EvaluateXPrio) {
		var ret float32
		return ret
	}
	return *o.EvaluateXPrio
}

// GetEvaluateXPrioOk returns a tuple with the EvaluateXPrio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePushoverSettingsRequestAttr) GetEvaluateXPrioOk() (*float32, bool) {
	if o == nil || IsNil(o.EvaluateXPrio) {
		return nil, false
	}
	return o.EvaluateXPrio, true
}

// HasEvaluateXPrio returns a boolean if a field has been set.
func (o *UpdatePushoverSettingsRequestAttr) HasEvaluateXPrio() bool {
	if o != nil && !IsNil(o.EvaluateXPrio) {
		return true
	}

	return false
}

// SetEvaluateXPrio gets a reference to the given float32 and assigns it to the EvaluateXPrio field.
func (o *UpdatePushoverSettingsRequestAttr) SetEvaluateXPrio(v float32) {
	o.EvaluateXPrio = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdatePushoverSettingsRequestAttr) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePushoverSettingsRequestAttr) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdatePushoverSettingsRequestAttr) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdatePushoverSettingsRequestAttr) SetKey(v string) {
	o.Key = &v
}

// GetOnlyXPrio returns the OnlyXPrio field value if set, zero value otherwise.
func (o *UpdatePushoverSettingsRequestAttr) GetOnlyXPrio() float32 {
	if o == nil || IsNil(o.OnlyXPrio) {
		var ret float32
		return ret
	}
	return *o.OnlyXPrio
}

// GetOnlyXPrioOk returns a tuple with the OnlyXPrio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePushoverSettingsRequestAttr) GetOnlyXPrioOk() (*float32, bool) {
	if o == nil || IsNil(o.OnlyXPrio) {
		return nil, false
	}
	return o.OnlyXPrio, true
}

// HasOnlyXPrio returns a boolean if a field has been set.
func (o *UpdatePushoverSettingsRequestAttr) HasOnlyXPrio() bool {
	if o != nil && !IsNil(o.OnlyXPrio) {
		return true
	}

	return false
}

// SetOnlyXPrio gets a reference to the given float32 and assigns it to the OnlyXPrio field.
func (o *UpdatePushoverSettingsRequestAttr) SetOnlyXPrio(v float32) {
	o.OnlyXPrio = &v
}

// GetSound returns the Sound field value if set, zero value otherwise.
func (o *UpdatePushoverSettingsRequestAttr) GetSound() string {
	if o == nil || IsNil(o.Sound) {
		var ret string
		return ret
	}
	return *o.Sound
}

// GetSoundOk returns a tuple with the Sound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePushoverSettingsRequestAttr) GetSoundOk() (*string, bool) {
	if o == nil || IsNil(o.Sound) {
		return nil, false
	}
	return o.Sound, true
}

// HasSound returns a boolean if a field has been set.
func (o *UpdatePushoverSettingsRequestAttr) HasSound() bool {
	if o != nil && !IsNil(o.Sound) {
		return true
	}

	return false
}

// SetSound gets a reference to the given string and assigns it to the Sound field.
func (o *UpdatePushoverSettingsRequestAttr) SetSound(v string) {
	o.Sound = &v
}

// GetSenders returns the Senders field value if set, zero value otherwise.
func (o *UpdatePushoverSettingsRequestAttr) GetSenders() string {
	if o == nil || IsNil(o.Senders) {
		var ret string
		return ret
	}
	return *o.Senders
}

// GetSendersOk returns a tuple with the Senders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePushoverSettingsRequestAttr) GetSendersOk() (*string, bool) {
	if o == nil || IsNil(o.Senders) {
		return nil, false
	}
	return o.Senders, true
}

// HasSenders returns a boolean if a field has been set.
func (o *UpdatePushoverSettingsRequestAttr) HasSenders() bool {
	if o != nil && !IsNil(o.Senders) {
		return true
	}

	return false
}

// SetSenders gets a reference to the given string and assigns it to the Senders field.
func (o *UpdatePushoverSettingsRequestAttr) SetSenders(v string) {
	o.Senders = &v
}

// GetSendersRegex returns the SendersRegex field value if set, zero value otherwise.
func (o *UpdatePushoverSettingsRequestAttr) GetSendersRegex() string {
	if o == nil || IsNil(o.SendersRegex) {
		var ret string
		return ret
	}
	return *o.SendersRegex
}

// GetSendersRegexOk returns a tuple with the SendersRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePushoverSettingsRequestAttr) GetSendersRegexOk() (*string, bool) {
	if o == nil || IsNil(o.SendersRegex) {
		return nil, false
	}
	return o.SendersRegex, true
}

// HasSendersRegex returns a boolean if a field has been set.
func (o *UpdatePushoverSettingsRequestAttr) HasSendersRegex() bool {
	if o != nil && !IsNil(o.SendersRegex) {
		return true
	}

	return false
}

// SetSendersRegex gets a reference to the given string and assigns it to the SendersRegex field.
func (o *UpdatePushoverSettingsRequestAttr) SetSendersRegex(v string) {
	o.SendersRegex = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *UpdatePushoverSettingsRequestAttr) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePushoverSettingsRequestAttr) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *UpdatePushoverSettingsRequestAttr) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *UpdatePushoverSettingsRequestAttr) SetText(v string) {
	o.Text = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdatePushoverSettingsRequestAttr) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePushoverSettingsRequestAttr) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdatePushoverSettingsRequestAttr) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdatePushoverSettingsRequestAttr) SetTitle(v string) {
	o.Title = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UpdatePushoverSettingsRequestAttr) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePushoverSettingsRequestAttr) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UpdatePushoverSettingsRequestAttr) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UpdatePushoverSettingsRequestAttr) SetToken(v string) {
	o.Token = &v
}

func (o UpdatePushoverSettingsRequestAttr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePushoverSettingsRequestAttr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.EvaluateXPrio) {
		toSerialize["evaluate_x_prio"] = o.EvaluateXPrio
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.OnlyXPrio) {
		toSerialize["only_x_prio"] = o.OnlyXPrio
	}
	if !IsNil(o.Sound) {
		toSerialize["sound"] = o.Sound
	}
	if !IsNil(o.Senders) {
		toSerialize["senders"] = o.Senders
	}
	if !IsNil(o.SendersRegex) {
		toSerialize["senders_regex"] = o.SendersRegex
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableUpdatePushoverSettingsRequestAttr struct {
	value *UpdatePushoverSettingsRequestAttr
	isSet bool
}

func (v NullableUpdatePushoverSettingsRequestAttr) Get() *UpdatePushoverSettingsRequestAttr {
	return v.value
}

func (v *NullableUpdatePushoverSettingsRequestAttr) Set(val *UpdatePushoverSettingsRequestAttr) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePushoverSettingsRequestAttr) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePushoverSettingsRequestAttr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePushoverSettingsRequestAttr(val *UpdatePushoverSettingsRequestAttr) *NullableUpdatePushoverSettingsRequestAttr {
	return &NullableUpdatePushoverSettingsRequestAttr{value: val, isSet: true}
}

func (v NullableUpdatePushoverSettingsRequestAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePushoverSettingsRequestAttr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


