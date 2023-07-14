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

// checks if the CreateAliasRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAliasRequest{}

// CreateAliasRequest struct for CreateAliasRequest
type CreateAliasRequest struct {
	// is alias active or not
	Active *bool `json:"active,omitempty"`
	// alias address, for catchall use \"@domain.tld\"
	Address *string `json:"address,omitempty"`
	// destination address, comma separated
	Goto *string `json:"goto,omitempty"`
	// learn as ham
	GotoHam *bool `json:"goto_ham,omitempty"`
	// silently ignore
	GotoNull *bool `json:"goto_null,omitempty"`
	// learn as spam
	GotoSpam *bool `json:"goto_spam,omitempty"`
	// toggle visibility as selectable sender in SOGo
	SogoVisible *bool `json:"sogo_visible,omitempty"`
}

// NewCreateAliasRequest instantiates a new CreateAliasRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAliasRequest() *CreateAliasRequest {
	this := CreateAliasRequest{}
	return &this
}

// NewCreateAliasRequestWithDefaults instantiates a new CreateAliasRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAliasRequestWithDefaults() *CreateAliasRequest {
	this := CreateAliasRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateAliasRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAliasRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateAliasRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CreateAliasRequest) SetActive(v bool) {
	o.Active = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CreateAliasRequest) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAliasRequest) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CreateAliasRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *CreateAliasRequest) SetAddress(v string) {
	o.Address = &v
}

// GetGoto returns the Goto field value if set, zero value otherwise.
func (o *CreateAliasRequest) GetGoto() string {
	if o == nil || IsNil(o.Goto) {
		var ret string
		return ret
	}
	return *o.Goto
}

// GetGotoOk returns a tuple with the Goto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAliasRequest) GetGotoOk() (*string, bool) {
	if o == nil || IsNil(o.Goto) {
		return nil, false
	}
	return o.Goto, true
}

// HasGoto returns a boolean if a field has been set.
func (o *CreateAliasRequest) HasGoto() bool {
	if o != nil && !IsNil(o.Goto) {
		return true
	}

	return false
}

// SetGoto gets a reference to the given string and assigns it to the Goto field.
func (o *CreateAliasRequest) SetGoto(v string) {
	o.Goto = &v
}

// GetGotoHam returns the GotoHam field value if set, zero value otherwise.
func (o *CreateAliasRequest) GetGotoHam() bool {
	if o == nil || IsNil(o.GotoHam) {
		var ret bool
		return ret
	}
	return *o.GotoHam
}

// GetGotoHamOk returns a tuple with the GotoHam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAliasRequest) GetGotoHamOk() (*bool, bool) {
	if o == nil || IsNil(o.GotoHam) {
		return nil, false
	}
	return o.GotoHam, true
}

// HasGotoHam returns a boolean if a field has been set.
func (o *CreateAliasRequest) HasGotoHam() bool {
	if o != nil && !IsNil(o.GotoHam) {
		return true
	}

	return false
}

// SetGotoHam gets a reference to the given bool and assigns it to the GotoHam field.
func (o *CreateAliasRequest) SetGotoHam(v bool) {
	o.GotoHam = &v
}

// GetGotoNull returns the GotoNull field value if set, zero value otherwise.
func (o *CreateAliasRequest) GetGotoNull() bool {
	if o == nil || IsNil(o.GotoNull) {
		var ret bool
		return ret
	}
	return *o.GotoNull
}

// GetGotoNullOk returns a tuple with the GotoNull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAliasRequest) GetGotoNullOk() (*bool, bool) {
	if o == nil || IsNil(o.GotoNull) {
		return nil, false
	}
	return o.GotoNull, true
}

// HasGotoNull returns a boolean if a field has been set.
func (o *CreateAliasRequest) HasGotoNull() bool {
	if o != nil && !IsNil(o.GotoNull) {
		return true
	}

	return false
}

// SetGotoNull gets a reference to the given bool and assigns it to the GotoNull field.
func (o *CreateAliasRequest) SetGotoNull(v bool) {
	o.GotoNull = &v
}

// GetGotoSpam returns the GotoSpam field value if set, zero value otherwise.
func (o *CreateAliasRequest) GetGotoSpam() bool {
	if o == nil || IsNil(o.GotoSpam) {
		var ret bool
		return ret
	}
	return *o.GotoSpam
}

// GetGotoSpamOk returns a tuple with the GotoSpam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAliasRequest) GetGotoSpamOk() (*bool, bool) {
	if o == nil || IsNil(o.GotoSpam) {
		return nil, false
	}
	return o.GotoSpam, true
}

// HasGotoSpam returns a boolean if a field has been set.
func (o *CreateAliasRequest) HasGotoSpam() bool {
	if o != nil && !IsNil(o.GotoSpam) {
		return true
	}

	return false
}

// SetGotoSpam gets a reference to the given bool and assigns it to the GotoSpam field.
func (o *CreateAliasRequest) SetGotoSpam(v bool) {
	o.GotoSpam = &v
}

// GetSogoVisible returns the SogoVisible field value if set, zero value otherwise.
func (o *CreateAliasRequest) GetSogoVisible() bool {
	if o == nil || IsNil(o.SogoVisible) {
		var ret bool
		return ret
	}
	return *o.SogoVisible
}

// GetSogoVisibleOk returns a tuple with the SogoVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAliasRequest) GetSogoVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.SogoVisible) {
		return nil, false
	}
	return o.SogoVisible, true
}

// HasSogoVisible returns a boolean if a field has been set.
func (o *CreateAliasRequest) HasSogoVisible() bool {
	if o != nil && !IsNil(o.SogoVisible) {
		return true
	}

	return false
}

// SetSogoVisible gets a reference to the given bool and assigns it to the SogoVisible field.
func (o *CreateAliasRequest) SetSogoVisible(v bool) {
	o.SogoVisible = &v
}

func (o CreateAliasRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAliasRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Goto) {
		toSerialize["goto"] = o.Goto
	}
	if !IsNil(o.GotoHam) {
		toSerialize["goto_ham"] = o.GotoHam
	}
	if !IsNil(o.GotoNull) {
		toSerialize["goto_null"] = o.GotoNull
	}
	if !IsNil(o.GotoSpam) {
		toSerialize["goto_spam"] = o.GotoSpam
	}
	if !IsNil(o.SogoVisible) {
		toSerialize["sogo_visible"] = o.SogoVisible
	}
	return toSerialize, nil
}

type NullableCreateAliasRequest struct {
	value *CreateAliasRequest
	isSet bool
}

func (v NullableCreateAliasRequest) Get() *CreateAliasRequest {
	return v.value
}

func (v *NullableCreateAliasRequest) Set(val *CreateAliasRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAliasRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAliasRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAliasRequest(val *CreateAliasRequest) *NullableCreateAliasRequest {
	return &NullableCreateAliasRequest{value: val, isSet: true}
}

func (v NullableCreateAliasRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAliasRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


