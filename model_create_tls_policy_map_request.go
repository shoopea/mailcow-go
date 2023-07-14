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

// checks if the CreateTLSPolicyMapRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTLSPolicyMapRequest{}

// CreateTLSPolicyMapRequest struct for CreateTLSPolicyMapRequest
type CreateTLSPolicyMapRequest struct {
	// custom parameters you find out more about them [here](http://www.postfix.org/postconf.5.html#smtp_tls_policy_maps)
	Parameters *string `json:"parameters,omitempty"`
	// 1 for a active user account 0 for a disabled user account
	Active *float32 `json:"active,omitempty"`
	// the target domain or email address
	Dest *string `json:"dest,omitempty"`
	// the policy
	Policy *string `json:"policy,omitempty"`
}

// NewCreateTLSPolicyMapRequest instantiates a new CreateTLSPolicyMapRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTLSPolicyMapRequest() *CreateTLSPolicyMapRequest {
	this := CreateTLSPolicyMapRequest{}
	return &this
}

// NewCreateTLSPolicyMapRequestWithDefaults instantiates a new CreateTLSPolicyMapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTLSPolicyMapRequestWithDefaults() *CreateTLSPolicyMapRequest {
	this := CreateTLSPolicyMapRequest{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *CreateTLSPolicyMapRequest) GetParameters() string {
	if o == nil || IsNil(o.Parameters) {
		var ret string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTLSPolicyMapRequest) GetParametersOk() (*string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CreateTLSPolicyMapRequest) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given string and assigns it to the Parameters field.
func (o *CreateTLSPolicyMapRequest) SetParameters(v string) {
	o.Parameters = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateTLSPolicyMapRequest) GetActive() float32 {
	if o == nil || IsNil(o.Active) {
		var ret float32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTLSPolicyMapRequest) GetActiveOk() (*float32, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateTLSPolicyMapRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given float32 and assigns it to the Active field.
func (o *CreateTLSPolicyMapRequest) SetActive(v float32) {
	o.Active = &v
}

// GetDest returns the Dest field value if set, zero value otherwise.
func (o *CreateTLSPolicyMapRequest) GetDest() string {
	if o == nil || IsNil(o.Dest) {
		var ret string
		return ret
	}
	return *o.Dest
}

// GetDestOk returns a tuple with the Dest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTLSPolicyMapRequest) GetDestOk() (*string, bool) {
	if o == nil || IsNil(o.Dest) {
		return nil, false
	}
	return o.Dest, true
}

// HasDest returns a boolean if a field has been set.
func (o *CreateTLSPolicyMapRequest) HasDest() bool {
	if o != nil && !IsNil(o.Dest) {
		return true
	}

	return false
}

// SetDest gets a reference to the given string and assigns it to the Dest field.
func (o *CreateTLSPolicyMapRequest) SetDest(v string) {
	o.Dest = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *CreateTLSPolicyMapRequest) GetPolicy() string {
	if o == nil || IsNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTLSPolicyMapRequest) GetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *CreateTLSPolicyMapRequest) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *CreateTLSPolicyMapRequest) SetPolicy(v string) {
	o.Policy = &v
}

func (o CreateTLSPolicyMapRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTLSPolicyMapRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Dest) {
		toSerialize["dest"] = o.Dest
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	return toSerialize, nil
}

type NullableCreateTLSPolicyMapRequest struct {
	value *CreateTLSPolicyMapRequest
	isSet bool
}

func (v NullableCreateTLSPolicyMapRequest) Get() *CreateTLSPolicyMapRequest {
	return v.value
}

func (v *NullableCreateTLSPolicyMapRequest) Set(val *CreateTLSPolicyMapRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTLSPolicyMapRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTLSPolicyMapRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTLSPolicyMapRequest(val *CreateTLSPolicyMapRequest) *NullableCreateTLSPolicyMapRequest {
	return &NullableCreateTLSPolicyMapRequest{value: val, isSet: true}
}

func (v NullableCreateTLSPolicyMapRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTLSPolicyMapRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

