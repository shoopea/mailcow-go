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

// checks if the DeleteMailsInQuarantineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteMailsInQuarantineRequest{}

// DeleteMailsInQuarantineRequest struct for DeleteMailsInQuarantineRequest
type DeleteMailsInQuarantineRequest struct {
	// contains list of emails you want to delete
	Items map[string]interface{} `json:"items,omitempty"`
}

// NewDeleteMailsInQuarantineRequest instantiates a new DeleteMailsInQuarantineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteMailsInQuarantineRequest() *DeleteMailsInQuarantineRequest {
	this := DeleteMailsInQuarantineRequest{}
	return &this
}

// NewDeleteMailsInQuarantineRequestWithDefaults instantiates a new DeleteMailsInQuarantineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteMailsInQuarantineRequestWithDefaults() *DeleteMailsInQuarantineRequest {
	this := DeleteMailsInQuarantineRequest{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DeleteMailsInQuarantineRequest) GetItems() map[string]interface{} {
	if o == nil || IsNil(o.Items) {
		var ret map[string]interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMailsInQuarantineRequest) GetItemsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Items) {
		return map[string]interface{}{}, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DeleteMailsInQuarantineRequest) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given map[string]interface{} and assigns it to the Items field.
func (o *DeleteMailsInQuarantineRequest) SetItems(v map[string]interface{}) {
	o.Items = v
}

func (o DeleteMailsInQuarantineRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteMailsInQuarantineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableDeleteMailsInQuarantineRequest struct {
	value *DeleteMailsInQuarantineRequest
	isSet bool
}

func (v NullableDeleteMailsInQuarantineRequest) Get() *DeleteMailsInQuarantineRequest {
	return v.value
}

func (v *NullableDeleteMailsInQuarantineRequest) Set(val *DeleteMailsInQuarantineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteMailsInQuarantineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteMailsInQuarantineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteMailsInQuarantineRequest(val *DeleteMailsInQuarantineRequest) *NullableDeleteMailsInQuarantineRequest {
	return &NullableDeleteMailsInQuarantineRequest{value: val, isSet: true}
}

func (v NullableDeleteMailsInQuarantineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteMailsInQuarantineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


