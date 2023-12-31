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

// checks if the DeleteQueueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteQueueRequest{}

// DeleteQueueRequest struct for DeleteQueueRequest
type DeleteQueueRequest struct {
	// use super_delete to delete the mail queue
	Action *string `json:"action,omitempty"`
}

// NewDeleteQueueRequest instantiates a new DeleteQueueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteQueueRequest() *DeleteQueueRequest {
	this := DeleteQueueRequest{}
	return &this
}

// NewDeleteQueueRequestWithDefaults instantiates a new DeleteQueueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteQueueRequestWithDefaults() *DeleteQueueRequest {
	this := DeleteQueueRequest{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *DeleteQueueRequest) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteQueueRequest) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *DeleteQueueRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *DeleteQueueRequest) SetAction(v string) {
	o.Action = &v
}

func (o DeleteQueueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteQueueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableDeleteQueueRequest struct {
	value *DeleteQueueRequest
	isSet bool
}

func (v NullableDeleteQueueRequest) Get() *DeleteQueueRequest {
	return v.value
}

func (v *NullableDeleteQueueRequest) Set(val *DeleteQueueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteQueueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteQueueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteQueueRequest(val *DeleteQueueRequest) *NullableDeleteQueueRequest {
	return &NullableDeleteQueueRequest{value: val, isSet: true}
}

func (v NullableDeleteQueueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteQueueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


