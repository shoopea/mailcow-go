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

// checks if the CreateDomainPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDomainPolicyRequest{}

// CreateDomainPolicyRequest struct for CreateDomainPolicyRequest
type CreateDomainPolicyRequest struct {
	// domain name to which policy is associated to
	Domain *string `json:"domain,omitempty"`
	// exact address or use wildcard to match whole domain
	ObjectFrom *string `json:"object_from,omitempty"`
	ObjectList *string `json:"object_list,omitempty"`
}

// NewCreateDomainPolicyRequest instantiates a new CreateDomainPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDomainPolicyRequest() *CreateDomainPolicyRequest {
	this := CreateDomainPolicyRequest{}
	return &this
}

// NewCreateDomainPolicyRequestWithDefaults instantiates a new CreateDomainPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDomainPolicyRequestWithDefaults() *CreateDomainPolicyRequest {
	this := CreateDomainPolicyRequest{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CreateDomainPolicyRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomainPolicyRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CreateDomainPolicyRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *CreateDomainPolicyRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetObjectFrom returns the ObjectFrom field value if set, zero value otherwise.
func (o *CreateDomainPolicyRequest) GetObjectFrom() string {
	if o == nil || IsNil(o.ObjectFrom) {
		var ret string
		return ret
	}
	return *o.ObjectFrom
}

// GetObjectFromOk returns a tuple with the ObjectFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomainPolicyRequest) GetObjectFromOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectFrom) {
		return nil, false
	}
	return o.ObjectFrom, true
}

// HasObjectFrom returns a boolean if a field has been set.
func (o *CreateDomainPolicyRequest) HasObjectFrom() bool {
	if o != nil && !IsNil(o.ObjectFrom) {
		return true
	}

	return false
}

// SetObjectFrom gets a reference to the given string and assigns it to the ObjectFrom field.
func (o *CreateDomainPolicyRequest) SetObjectFrom(v string) {
	o.ObjectFrom = &v
}

// GetObjectList returns the ObjectList field value if set, zero value otherwise.
func (o *CreateDomainPolicyRequest) GetObjectList() string {
	if o == nil || IsNil(o.ObjectList) {
		var ret string
		return ret
	}
	return *o.ObjectList
}

// GetObjectListOk returns a tuple with the ObjectList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomainPolicyRequest) GetObjectListOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectList) {
		return nil, false
	}
	return o.ObjectList, true
}

// HasObjectList returns a boolean if a field has been set.
func (o *CreateDomainPolicyRequest) HasObjectList() bool {
	if o != nil && !IsNil(o.ObjectList) {
		return true
	}

	return false
}

// SetObjectList gets a reference to the given string and assigns it to the ObjectList field.
func (o *CreateDomainPolicyRequest) SetObjectList(v string) {
	o.ObjectList = &v
}

func (o CreateDomainPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDomainPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.ObjectFrom) {
		toSerialize["object_from"] = o.ObjectFrom
	}
	if !IsNil(o.ObjectList) {
		toSerialize["object_list"] = o.ObjectList
	}
	return toSerialize, nil
}

type NullableCreateDomainPolicyRequest struct {
	value *CreateDomainPolicyRequest
	isSet bool
}

func (v NullableCreateDomainPolicyRequest) Get() *CreateDomainPolicyRequest {
	return v.value
}

func (v *NullableCreateDomainPolicyRequest) Set(val *CreateDomainPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDomainPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDomainPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDomainPolicyRequest(val *CreateDomainPolicyRequest) *NullableCreateDomainPolicyRequest {
	return &NullableCreateDomainPolicyRequest{value: val, isSet: true}
}

func (v NullableCreateDomainPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDomainPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


