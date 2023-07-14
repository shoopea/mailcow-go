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

// checks if the CreateAppPasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAppPasswordRequest{}

// CreateAppPasswordRequest struct for CreateAppPasswordRequest
type CreateAppPasswordRequest struct {
	// is alias active or not
	Active *bool `json:"active,omitempty"`
	// mailbox for which the app password should be created
	Username *string `json:"username,omitempty"`
	// name of your app password
	AppName *string `json:"app_name,omitempty"`
	// your app password
	AppPasswd *string `json:"app_passwd,omitempty"`
	// your app password
	AppPasswd2 *string `json:"app_passwd2,omitempty"`
}

// NewCreateAppPasswordRequest instantiates a new CreateAppPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAppPasswordRequest() *CreateAppPasswordRequest {
	this := CreateAppPasswordRequest{}
	return &this
}

// NewCreateAppPasswordRequestWithDefaults instantiates a new CreateAppPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAppPasswordRequestWithDefaults() *CreateAppPasswordRequest {
	this := CreateAppPasswordRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateAppPasswordRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppPasswordRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateAppPasswordRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CreateAppPasswordRequest) SetActive(v bool) {
	o.Active = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateAppPasswordRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppPasswordRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateAppPasswordRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateAppPasswordRequest) SetUsername(v string) {
	o.Username = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *CreateAppPasswordRequest) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppPasswordRequest) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *CreateAppPasswordRequest) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *CreateAppPasswordRequest) SetAppName(v string) {
	o.AppName = &v
}

// GetAppPasswd returns the AppPasswd field value if set, zero value otherwise.
func (o *CreateAppPasswordRequest) GetAppPasswd() string {
	if o == nil || IsNil(o.AppPasswd) {
		var ret string
		return ret
	}
	return *o.AppPasswd
}

// GetAppPasswdOk returns a tuple with the AppPasswd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppPasswordRequest) GetAppPasswdOk() (*string, bool) {
	if o == nil || IsNil(o.AppPasswd) {
		return nil, false
	}
	return o.AppPasswd, true
}

// HasAppPasswd returns a boolean if a field has been set.
func (o *CreateAppPasswordRequest) HasAppPasswd() bool {
	if o != nil && !IsNil(o.AppPasswd) {
		return true
	}

	return false
}

// SetAppPasswd gets a reference to the given string and assigns it to the AppPasswd field.
func (o *CreateAppPasswordRequest) SetAppPasswd(v string) {
	o.AppPasswd = &v
}

// GetAppPasswd2 returns the AppPasswd2 field value if set, zero value otherwise.
func (o *CreateAppPasswordRequest) GetAppPasswd2() string {
	if o == nil || IsNil(o.AppPasswd2) {
		var ret string
		return ret
	}
	return *o.AppPasswd2
}

// GetAppPasswd2Ok returns a tuple with the AppPasswd2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppPasswordRequest) GetAppPasswd2Ok() (*string, bool) {
	if o == nil || IsNil(o.AppPasswd2) {
		return nil, false
	}
	return o.AppPasswd2, true
}

// HasAppPasswd2 returns a boolean if a field has been set.
func (o *CreateAppPasswordRequest) HasAppPasswd2() bool {
	if o != nil && !IsNil(o.AppPasswd2) {
		return true
	}

	return false
}

// SetAppPasswd2 gets a reference to the given string and assigns it to the AppPasswd2 field.
func (o *CreateAppPasswordRequest) SetAppPasswd2(v string) {
	o.AppPasswd2 = &v
}

func (o CreateAppPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAppPasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.AppPasswd) {
		toSerialize["app_passwd"] = o.AppPasswd
	}
	if !IsNil(o.AppPasswd2) {
		toSerialize["app_passwd2"] = o.AppPasswd2
	}
	return toSerialize, nil
}

type NullableCreateAppPasswordRequest struct {
	value *CreateAppPasswordRequest
	isSet bool
}

func (v NullableCreateAppPasswordRequest) Get() *CreateAppPasswordRequest {
	return v.value
}

func (v *NullableCreateAppPasswordRequest) Set(val *CreateAppPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAppPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAppPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAppPasswordRequest(val *CreateAppPasswordRequest) *NullableCreateAppPasswordRequest {
	return &NullableCreateAppPasswordRequest{value: val, isSet: true}
}

func (v NullableCreateAppPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAppPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

