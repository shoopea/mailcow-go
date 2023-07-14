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

// checks if the CreateBCCMapRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBCCMapRequest{}

// CreateBCCMapRequest struct for CreateBCCMapRequest
type CreateBCCMapRequest struct {
	// 1 for a active user account 0 for a disabled user account
	Active *float32 `json:"active,omitempty"`
	// the email address where all mails should be send to
	BccDest *string `json:"bcc_dest,omitempty"`
	// the domain which emails should be forwarded
	LocalDest *string `json:"local_dest,omitempty"`
	// the type of bcc map can be `sender` or `recipient`
	Type *string `json:"type,omitempty"`
}

// NewCreateBCCMapRequest instantiates a new CreateBCCMapRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBCCMapRequest() *CreateBCCMapRequest {
	this := CreateBCCMapRequest{}
	return &this
}

// NewCreateBCCMapRequestWithDefaults instantiates a new CreateBCCMapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBCCMapRequestWithDefaults() *CreateBCCMapRequest {
	this := CreateBCCMapRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateBCCMapRequest) GetActive() float32 {
	if o == nil || IsNil(o.Active) {
		var ret float32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBCCMapRequest) GetActiveOk() (*float32, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateBCCMapRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given float32 and assigns it to the Active field.
func (o *CreateBCCMapRequest) SetActive(v float32) {
	o.Active = &v
}

// GetBccDest returns the BccDest field value if set, zero value otherwise.
func (o *CreateBCCMapRequest) GetBccDest() string {
	if o == nil || IsNil(o.BccDest) {
		var ret string
		return ret
	}
	return *o.BccDest
}

// GetBccDestOk returns a tuple with the BccDest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBCCMapRequest) GetBccDestOk() (*string, bool) {
	if o == nil || IsNil(o.BccDest) {
		return nil, false
	}
	return o.BccDest, true
}

// HasBccDest returns a boolean if a field has been set.
func (o *CreateBCCMapRequest) HasBccDest() bool {
	if o != nil && !IsNil(o.BccDest) {
		return true
	}

	return false
}

// SetBccDest gets a reference to the given string and assigns it to the BccDest field.
func (o *CreateBCCMapRequest) SetBccDest(v string) {
	o.BccDest = &v
}

// GetLocalDest returns the LocalDest field value if set, zero value otherwise.
func (o *CreateBCCMapRequest) GetLocalDest() string {
	if o == nil || IsNil(o.LocalDest) {
		var ret string
		return ret
	}
	return *o.LocalDest
}

// GetLocalDestOk returns a tuple with the LocalDest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBCCMapRequest) GetLocalDestOk() (*string, bool) {
	if o == nil || IsNil(o.LocalDest) {
		return nil, false
	}
	return o.LocalDest, true
}

// HasLocalDest returns a boolean if a field has been set.
func (o *CreateBCCMapRequest) HasLocalDest() bool {
	if o != nil && !IsNil(o.LocalDest) {
		return true
	}

	return false
}

// SetLocalDest gets a reference to the given string and assigns it to the LocalDest field.
func (o *CreateBCCMapRequest) SetLocalDest(v string) {
	o.LocalDest = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateBCCMapRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBCCMapRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateBCCMapRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateBCCMapRequest) SetType(v string) {
	o.Type = &v
}

func (o CreateBCCMapRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBCCMapRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.BccDest) {
		toSerialize["bcc_dest"] = o.BccDest
	}
	if !IsNil(o.LocalDest) {
		toSerialize["local_dest"] = o.LocalDest
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCreateBCCMapRequest struct {
	value *CreateBCCMapRequest
	isSet bool
}

func (v NullableCreateBCCMapRequest) Get() *CreateBCCMapRequest {
	return v.value
}

func (v *NullableCreateBCCMapRequest) Set(val *CreateBCCMapRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBCCMapRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBCCMapRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBCCMapRequest(val *CreateBCCMapRequest) *NullableCreateBCCMapRequest {
	return &NullableCreateBCCMapRequest{value: val, isSet: true}
}

func (v NullableCreateBCCMapRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBCCMapRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


