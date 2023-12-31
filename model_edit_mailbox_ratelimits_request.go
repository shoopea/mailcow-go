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

// checks if the EditMailboxRatelimitsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditMailboxRatelimitsRequest{}

// EditMailboxRatelimitsRequest struct for EditMailboxRatelimitsRequest
type EditMailboxRatelimitsRequest struct {
	Attr *EditMailboxRatelimitsRequestAttr `json:"attr,omitempty"`
	// contains list of mailboxes you want to edit the ratelimit of
	Items map[string]interface{} `json:"items,omitempty"`
}

// NewEditMailboxRatelimitsRequest instantiates a new EditMailboxRatelimitsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditMailboxRatelimitsRequest() *EditMailboxRatelimitsRequest {
	this := EditMailboxRatelimitsRequest{}
	return &this
}

// NewEditMailboxRatelimitsRequestWithDefaults instantiates a new EditMailboxRatelimitsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditMailboxRatelimitsRequestWithDefaults() *EditMailboxRatelimitsRequest {
	this := EditMailboxRatelimitsRequest{}
	return &this
}

// GetAttr returns the Attr field value if set, zero value otherwise.
func (o *EditMailboxRatelimitsRequest) GetAttr() EditMailboxRatelimitsRequestAttr {
	if o == nil || IsNil(o.Attr) {
		var ret EditMailboxRatelimitsRequestAttr
		return ret
	}
	return *o.Attr
}

// GetAttrOk returns a tuple with the Attr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMailboxRatelimitsRequest) GetAttrOk() (*EditMailboxRatelimitsRequestAttr, bool) {
	if o == nil || IsNil(o.Attr) {
		return nil, false
	}
	return o.Attr, true
}

// HasAttr returns a boolean if a field has been set.
func (o *EditMailboxRatelimitsRequest) HasAttr() bool {
	if o != nil && !IsNil(o.Attr) {
		return true
	}

	return false
}

// SetAttr gets a reference to the given EditMailboxRatelimitsRequestAttr and assigns it to the Attr field.
func (o *EditMailboxRatelimitsRequest) SetAttr(v EditMailboxRatelimitsRequestAttr) {
	o.Attr = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *EditMailboxRatelimitsRequest) GetItems() map[string]interface{} {
	if o == nil || IsNil(o.Items) {
		var ret map[string]interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMailboxRatelimitsRequest) GetItemsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Items) {
		return map[string]interface{}{}, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *EditMailboxRatelimitsRequest) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given map[string]interface{} and assigns it to the Items field.
func (o *EditMailboxRatelimitsRequest) SetItems(v map[string]interface{}) {
	o.Items = v
}

func (o EditMailboxRatelimitsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditMailboxRatelimitsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attr) {
		toSerialize["attr"] = o.Attr
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableEditMailboxRatelimitsRequest struct {
	value *EditMailboxRatelimitsRequest
	isSet bool
}

func (v NullableEditMailboxRatelimitsRequest) Get() *EditMailboxRatelimitsRequest {
	return v.value
}

func (v *NullableEditMailboxRatelimitsRequest) Set(val *EditMailboxRatelimitsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditMailboxRatelimitsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditMailboxRatelimitsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditMailboxRatelimitsRequest(val *EditMailboxRatelimitsRequest) *NullableEditMailboxRatelimitsRequest {
	return &NullableEditMailboxRatelimitsRequest{value: val, isSet: true}
}

func (v NullableEditMailboxRatelimitsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditMailboxRatelimitsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


