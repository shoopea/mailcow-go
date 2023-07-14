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

// checks if the CreateTimeLimitedAliasRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTimeLimitedAliasRequest{}

// CreateTimeLimitedAliasRequest struct for CreateTimeLimitedAliasRequest
type CreateTimeLimitedAliasRequest struct {
	// the mailbox an alias should be created for
	Username *string `json:"username,omitempty"`
	// the domain
	Domain *string `json:"domain,omitempty"`
}

// NewCreateTimeLimitedAliasRequest instantiates a new CreateTimeLimitedAliasRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTimeLimitedAliasRequest() *CreateTimeLimitedAliasRequest {
	this := CreateTimeLimitedAliasRequest{}
	return &this
}

// NewCreateTimeLimitedAliasRequestWithDefaults instantiates a new CreateTimeLimitedAliasRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTimeLimitedAliasRequestWithDefaults() *CreateTimeLimitedAliasRequest {
	this := CreateTimeLimitedAliasRequest{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateTimeLimitedAliasRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTimeLimitedAliasRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateTimeLimitedAliasRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateTimeLimitedAliasRequest) SetUsername(v string) {
	o.Username = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CreateTimeLimitedAliasRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTimeLimitedAliasRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CreateTimeLimitedAliasRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *CreateTimeLimitedAliasRequest) SetDomain(v string) {
	o.Domain = &v
}

func (o CreateTimeLimitedAliasRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTimeLimitedAliasRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	return toSerialize, nil
}

type NullableCreateTimeLimitedAliasRequest struct {
	value *CreateTimeLimitedAliasRequest
	isSet bool
}

func (v NullableCreateTimeLimitedAliasRequest) Get() *CreateTimeLimitedAliasRequest {
	return v.value
}

func (v *NullableCreateTimeLimitedAliasRequest) Set(val *CreateTimeLimitedAliasRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTimeLimitedAliasRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTimeLimitedAliasRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTimeLimitedAliasRequest(val *CreateTimeLimitedAliasRequest) *NullableCreateTimeLimitedAliasRequest {
	return &NullableCreateTimeLimitedAliasRequest{value: val, isSet: true}
}

func (v NullableCreateTimeLimitedAliasRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTimeLimitedAliasRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

