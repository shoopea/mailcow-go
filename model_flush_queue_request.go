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

// checks if the FlushQueueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlushQueueRequest{}

// FlushQueueRequest struct for FlushQueueRequest
type FlushQueueRequest struct {
	// use flush to flush the mail queue
	Action *string `json:"action,omitempty"`
}

// NewFlushQueueRequest instantiates a new FlushQueueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlushQueueRequest() *FlushQueueRequest {
	this := FlushQueueRequest{}
	return &this
}

// NewFlushQueueRequestWithDefaults instantiates a new FlushQueueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlushQueueRequestWithDefaults() *FlushQueueRequest {
	this := FlushQueueRequest{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *FlushQueueRequest) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlushQueueRequest) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *FlushQueueRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *FlushQueueRequest) SetAction(v string) {
	o.Action = &v
}

func (o FlushQueueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlushQueueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableFlushQueueRequest struct {
	value *FlushQueueRequest
	isSet bool
}

func (v NullableFlushQueueRequest) Get() *FlushQueueRequest {
	return v.value
}

func (v *NullableFlushQueueRequest) Set(val *FlushQueueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlushQueueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlushQueueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlushQueueRequest(val *FlushQueueRequest) *NullableFlushQueueRequest {
	return &NullableFlushQueueRequest{value: val, isSet: true}
}

func (v NullableFlushQueueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlushQueueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


