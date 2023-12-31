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

// checks if the CreateAlias200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAlias200Response{}

// CreateAlias200Response struct for CreateAlias200Response
type CreateAlias200Response struct {
	// contains request object
	Log []interface{} `json:"log,omitempty"`
	Msg []interface{} `json:"msg,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewCreateAlias200Response instantiates a new CreateAlias200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAlias200Response() *CreateAlias200Response {
	this := CreateAlias200Response{}
	return &this
}

// NewCreateAlias200ResponseWithDefaults instantiates a new CreateAlias200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAlias200ResponseWithDefaults() *CreateAlias200Response {
	this := CreateAlias200Response{}
	return &this
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *CreateAlias200Response) GetLog() []interface{} {
	if o == nil || IsNil(o.Log) {
		var ret []interface{}
		return ret
	}
	return o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlias200Response) GetLogOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Log) {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *CreateAlias200Response) HasLog() bool {
	if o != nil && !IsNil(o.Log) {
		return true
	}

	return false
}

// SetLog gets a reference to the given []interface{} and assigns it to the Log field.
func (o *CreateAlias200Response) SetLog(v []interface{}) {
	o.Log = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateAlias200Response) GetMsg() []interface{} {
	if o == nil || IsNil(o.Msg) {
		var ret []interface{}
		return ret
	}
	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlias200Response) GetMsgOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CreateAlias200Response) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given []interface{} and assigns it to the Msg field.
func (o *CreateAlias200Response) SetMsg(v []interface{}) {
	o.Msg = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateAlias200Response) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlias200Response) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateAlias200Response) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateAlias200Response) SetType(v string) {
	o.Type = &v
}

func (o CreateAlias200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAlias200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Log) {
		toSerialize["log"] = o.Log
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCreateAlias200Response struct {
	value *CreateAlias200Response
	isSet bool
}

func (v NullableCreateAlias200Response) Get() *CreateAlias200Response {
	return v.value
}

func (v *NullableCreateAlias200Response) Set(val *CreateAlias200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAlias200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAlias200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAlias200Response(val *CreateAlias200Response) *NullableCreateAlias200Response {
	return &NullableCreateAlias200Response{value: val, isSet: true}
}

func (v NullableCreateAlias200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAlias200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


