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

// checks if the CreateSyncJobRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSyncJobRequest{}

// CreateSyncJobRequest struct for CreateSyncJobRequest
type CreateSyncJobRequest struct {
	// your local mailcow mailbox
	Parameters *string `json:"parameters,omitempty"`
	// the smtp server where mails should be synced from
	Host1 *string `json:"host1,omitempty"`
	// the smtp port of the target mail server
	Port1 *string `json:"port1,omitempty"`
	// the password of the mailbox
	Password *string `json:"password,omitempty"`
	// the encryption method used to connect to the mailserver
	Enc1 *string `json:"enc1,omitempty"`
	// the interval in which messages should be syned
	MinsInternal *float32 `json:"mins_internal,omitempty"`
	// sync into subfolder on destination (empty = do not use subfolder)
	Subfolder2 *string `json:"subfolder2,omitempty"`
	// only sync messages up to this age in days
	Maxage *float32 `json:"maxage,omitempty"`
	// max speed transfer limit for the sync
	Maxbytespersecond *float32 `json:"maxbytespersecond,omitempty"`
	// timeout for connection to remote host
	Timeout1 *float32 `json:"timeout1,omitempty"`
	// timeout for connection to local host
	Timeout2 *float32 `json:"timeout2,omitempty"`
	// exclude objects (regex)
	Exclude *string `json:"exclude,omitempty"`
	// custom parameters
	CustomParams *string `json:"custom_params,omitempty"`
	// delete duplicates on destination (--delete2duplicates)
	Delete2duplicates *bool `json:"delete2duplicates,omitempty"`
	// delete from source when completed (--delete1)
	Delete1 *bool `json:"delete1,omitempty"`
	// delete messages on destination that are not on source (--delete2)
	Delete2 *bool `json:"delete2,omitempty"`
	// try to automap folders (\"Sent items\", \"Sent\" => \"Sent\" etc.) (--automap)
	Automap *bool `json:"automap,omitempty"`
	// skip duplicate messages across folders (first come, first serve) (--skipcrossduplicates)
	Skipcrossduplicates *bool `json:"skipcrossduplicates,omitempty"`
	// subscribe all folders (--subscribeall)
	Subscribeall *bool `json:"subscribeall,omitempty"`
	// enables or disables the sync job
	Active *bool `json:"active,omitempty"`
}

// NewCreateSyncJobRequest instantiates a new CreateSyncJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSyncJobRequest() *CreateSyncJobRequest {
	this := CreateSyncJobRequest{}
	return &this
}

// NewCreateSyncJobRequestWithDefaults instantiates a new CreateSyncJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSyncJobRequestWithDefaults() *CreateSyncJobRequest {
	this := CreateSyncJobRequest{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetParameters() string {
	if o == nil || IsNil(o.Parameters) {
		var ret string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetParametersOk() (*string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given string and assigns it to the Parameters field.
func (o *CreateSyncJobRequest) SetParameters(v string) {
	o.Parameters = &v
}

// GetHost1 returns the Host1 field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetHost1() string {
	if o == nil || IsNil(o.Host1) {
		var ret string
		return ret
	}
	return *o.Host1
}

// GetHost1Ok returns a tuple with the Host1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetHost1Ok() (*string, bool) {
	if o == nil || IsNil(o.Host1) {
		return nil, false
	}
	return o.Host1, true
}

// HasHost1 returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasHost1() bool {
	if o != nil && !IsNil(o.Host1) {
		return true
	}

	return false
}

// SetHost1 gets a reference to the given string and assigns it to the Host1 field.
func (o *CreateSyncJobRequest) SetHost1(v string) {
	o.Host1 = &v
}

// GetPort1 returns the Port1 field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetPort1() string {
	if o == nil || IsNil(o.Port1) {
		var ret string
		return ret
	}
	return *o.Port1
}

// GetPort1Ok returns a tuple with the Port1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetPort1Ok() (*string, bool) {
	if o == nil || IsNil(o.Port1) {
		return nil, false
	}
	return o.Port1, true
}

// HasPort1 returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasPort1() bool {
	if o != nil && !IsNil(o.Port1) {
		return true
	}

	return false
}

// SetPort1 gets a reference to the given string and assigns it to the Port1 field.
func (o *CreateSyncJobRequest) SetPort1(v string) {
	o.Port1 = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateSyncJobRequest) SetPassword(v string) {
	o.Password = &v
}

// GetEnc1 returns the Enc1 field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetEnc1() string {
	if o == nil || IsNil(o.Enc1) {
		var ret string
		return ret
	}
	return *o.Enc1
}

// GetEnc1Ok returns a tuple with the Enc1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetEnc1Ok() (*string, bool) {
	if o == nil || IsNil(o.Enc1) {
		return nil, false
	}
	return o.Enc1, true
}

// HasEnc1 returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasEnc1() bool {
	if o != nil && !IsNil(o.Enc1) {
		return true
	}

	return false
}

// SetEnc1 gets a reference to the given string and assigns it to the Enc1 field.
func (o *CreateSyncJobRequest) SetEnc1(v string) {
	o.Enc1 = &v
}

// GetMinsInternal returns the MinsInternal field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetMinsInternal() float32 {
	if o == nil || IsNil(o.MinsInternal) {
		var ret float32
		return ret
	}
	return *o.MinsInternal
}

// GetMinsInternalOk returns a tuple with the MinsInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetMinsInternalOk() (*float32, bool) {
	if o == nil || IsNil(o.MinsInternal) {
		return nil, false
	}
	return o.MinsInternal, true
}

// HasMinsInternal returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasMinsInternal() bool {
	if o != nil && !IsNil(o.MinsInternal) {
		return true
	}

	return false
}

// SetMinsInternal gets a reference to the given float32 and assigns it to the MinsInternal field.
func (o *CreateSyncJobRequest) SetMinsInternal(v float32) {
	o.MinsInternal = &v
}

// GetSubfolder2 returns the Subfolder2 field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetSubfolder2() string {
	if o == nil || IsNil(o.Subfolder2) {
		var ret string
		return ret
	}
	return *o.Subfolder2
}

// GetSubfolder2Ok returns a tuple with the Subfolder2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetSubfolder2Ok() (*string, bool) {
	if o == nil || IsNil(o.Subfolder2) {
		return nil, false
	}
	return o.Subfolder2, true
}

// HasSubfolder2 returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasSubfolder2() bool {
	if o != nil && !IsNil(o.Subfolder2) {
		return true
	}

	return false
}

// SetSubfolder2 gets a reference to the given string and assigns it to the Subfolder2 field.
func (o *CreateSyncJobRequest) SetSubfolder2(v string) {
	o.Subfolder2 = &v
}

// GetMaxage returns the Maxage field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetMaxage() float32 {
	if o == nil || IsNil(o.Maxage) {
		var ret float32
		return ret
	}
	return *o.Maxage
}

// GetMaxageOk returns a tuple with the Maxage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetMaxageOk() (*float32, bool) {
	if o == nil || IsNil(o.Maxage) {
		return nil, false
	}
	return o.Maxage, true
}

// HasMaxage returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasMaxage() bool {
	if o != nil && !IsNil(o.Maxage) {
		return true
	}

	return false
}

// SetMaxage gets a reference to the given float32 and assigns it to the Maxage field.
func (o *CreateSyncJobRequest) SetMaxage(v float32) {
	o.Maxage = &v
}

// GetMaxbytespersecond returns the Maxbytespersecond field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetMaxbytespersecond() float32 {
	if o == nil || IsNil(o.Maxbytespersecond) {
		var ret float32
		return ret
	}
	return *o.Maxbytespersecond
}

// GetMaxbytespersecondOk returns a tuple with the Maxbytespersecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetMaxbytespersecondOk() (*float32, bool) {
	if o == nil || IsNil(o.Maxbytespersecond) {
		return nil, false
	}
	return o.Maxbytespersecond, true
}

// HasMaxbytespersecond returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasMaxbytespersecond() bool {
	if o != nil && !IsNil(o.Maxbytespersecond) {
		return true
	}

	return false
}

// SetMaxbytespersecond gets a reference to the given float32 and assigns it to the Maxbytespersecond field.
func (o *CreateSyncJobRequest) SetMaxbytespersecond(v float32) {
	o.Maxbytespersecond = &v
}

// GetTimeout1 returns the Timeout1 field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetTimeout1() float32 {
	if o == nil || IsNil(o.Timeout1) {
		var ret float32
		return ret
	}
	return *o.Timeout1
}

// GetTimeout1Ok returns a tuple with the Timeout1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetTimeout1Ok() (*float32, bool) {
	if o == nil || IsNil(o.Timeout1) {
		return nil, false
	}
	return o.Timeout1, true
}

// HasTimeout1 returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasTimeout1() bool {
	if o != nil && !IsNil(o.Timeout1) {
		return true
	}

	return false
}

// SetTimeout1 gets a reference to the given float32 and assigns it to the Timeout1 field.
func (o *CreateSyncJobRequest) SetTimeout1(v float32) {
	o.Timeout1 = &v
}

// GetTimeout2 returns the Timeout2 field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetTimeout2() float32 {
	if o == nil || IsNil(o.Timeout2) {
		var ret float32
		return ret
	}
	return *o.Timeout2
}

// GetTimeout2Ok returns a tuple with the Timeout2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetTimeout2Ok() (*float32, bool) {
	if o == nil || IsNil(o.Timeout2) {
		return nil, false
	}
	return o.Timeout2, true
}

// HasTimeout2 returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasTimeout2() bool {
	if o != nil && !IsNil(o.Timeout2) {
		return true
	}

	return false
}

// SetTimeout2 gets a reference to the given float32 and assigns it to the Timeout2 field.
func (o *CreateSyncJobRequest) SetTimeout2(v float32) {
	o.Timeout2 = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetExclude() string {
	if o == nil || IsNil(o.Exclude) {
		var ret string
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetExcludeOk() (*string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given string and assigns it to the Exclude field.
func (o *CreateSyncJobRequest) SetExclude(v string) {
	o.Exclude = &v
}

// GetCustomParams returns the CustomParams field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetCustomParams() string {
	if o == nil || IsNil(o.CustomParams) {
		var ret string
		return ret
	}
	return *o.CustomParams
}

// GetCustomParamsOk returns a tuple with the CustomParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetCustomParamsOk() (*string, bool) {
	if o == nil || IsNil(o.CustomParams) {
		return nil, false
	}
	return o.CustomParams, true
}

// HasCustomParams returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasCustomParams() bool {
	if o != nil && !IsNil(o.CustomParams) {
		return true
	}

	return false
}

// SetCustomParams gets a reference to the given string and assigns it to the CustomParams field.
func (o *CreateSyncJobRequest) SetCustomParams(v string) {
	o.CustomParams = &v
}

// GetDelete2duplicates returns the Delete2duplicates field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetDelete2duplicates() bool {
	if o == nil || IsNil(o.Delete2duplicates) {
		var ret bool
		return ret
	}
	return *o.Delete2duplicates
}

// GetDelete2duplicatesOk returns a tuple with the Delete2duplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetDelete2duplicatesOk() (*bool, bool) {
	if o == nil || IsNil(o.Delete2duplicates) {
		return nil, false
	}
	return o.Delete2duplicates, true
}

// HasDelete2duplicates returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasDelete2duplicates() bool {
	if o != nil && !IsNil(o.Delete2duplicates) {
		return true
	}

	return false
}

// SetDelete2duplicates gets a reference to the given bool and assigns it to the Delete2duplicates field.
func (o *CreateSyncJobRequest) SetDelete2duplicates(v bool) {
	o.Delete2duplicates = &v
}

// GetDelete1 returns the Delete1 field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetDelete1() bool {
	if o == nil || IsNil(o.Delete1) {
		var ret bool
		return ret
	}
	return *o.Delete1
}

// GetDelete1Ok returns a tuple with the Delete1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetDelete1Ok() (*bool, bool) {
	if o == nil || IsNil(o.Delete1) {
		return nil, false
	}
	return o.Delete1, true
}

// HasDelete1 returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasDelete1() bool {
	if o != nil && !IsNil(o.Delete1) {
		return true
	}

	return false
}

// SetDelete1 gets a reference to the given bool and assigns it to the Delete1 field.
func (o *CreateSyncJobRequest) SetDelete1(v bool) {
	o.Delete1 = &v
}

// GetDelete2 returns the Delete2 field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetDelete2() bool {
	if o == nil || IsNil(o.Delete2) {
		var ret bool
		return ret
	}
	return *o.Delete2
}

// GetDelete2Ok returns a tuple with the Delete2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetDelete2Ok() (*bool, bool) {
	if o == nil || IsNil(o.Delete2) {
		return nil, false
	}
	return o.Delete2, true
}

// HasDelete2 returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasDelete2() bool {
	if o != nil && !IsNil(o.Delete2) {
		return true
	}

	return false
}

// SetDelete2 gets a reference to the given bool and assigns it to the Delete2 field.
func (o *CreateSyncJobRequest) SetDelete2(v bool) {
	o.Delete2 = &v
}

// GetAutomap returns the Automap field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetAutomap() bool {
	if o == nil || IsNil(o.Automap) {
		var ret bool
		return ret
	}
	return *o.Automap
}

// GetAutomapOk returns a tuple with the Automap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetAutomapOk() (*bool, bool) {
	if o == nil || IsNil(o.Automap) {
		return nil, false
	}
	return o.Automap, true
}

// HasAutomap returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasAutomap() bool {
	if o != nil && !IsNil(o.Automap) {
		return true
	}

	return false
}

// SetAutomap gets a reference to the given bool and assigns it to the Automap field.
func (o *CreateSyncJobRequest) SetAutomap(v bool) {
	o.Automap = &v
}

// GetSkipcrossduplicates returns the Skipcrossduplicates field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetSkipcrossduplicates() bool {
	if o == nil || IsNil(o.Skipcrossduplicates) {
		var ret bool
		return ret
	}
	return *o.Skipcrossduplicates
}

// GetSkipcrossduplicatesOk returns a tuple with the Skipcrossduplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetSkipcrossduplicatesOk() (*bool, bool) {
	if o == nil || IsNil(o.Skipcrossduplicates) {
		return nil, false
	}
	return o.Skipcrossduplicates, true
}

// HasSkipcrossduplicates returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasSkipcrossduplicates() bool {
	if o != nil && !IsNil(o.Skipcrossduplicates) {
		return true
	}

	return false
}

// SetSkipcrossduplicates gets a reference to the given bool and assigns it to the Skipcrossduplicates field.
func (o *CreateSyncJobRequest) SetSkipcrossduplicates(v bool) {
	o.Skipcrossduplicates = &v
}

// GetSubscribeall returns the Subscribeall field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetSubscribeall() bool {
	if o == nil || IsNil(o.Subscribeall) {
		var ret bool
		return ret
	}
	return *o.Subscribeall
}

// GetSubscribeallOk returns a tuple with the Subscribeall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetSubscribeallOk() (*bool, bool) {
	if o == nil || IsNil(o.Subscribeall) {
		return nil, false
	}
	return o.Subscribeall, true
}

// HasSubscribeall returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasSubscribeall() bool {
	if o != nil && !IsNil(o.Subscribeall) {
		return true
	}

	return false
}

// SetSubscribeall gets a reference to the given bool and assigns it to the Subscribeall field.
func (o *CreateSyncJobRequest) SetSubscribeall(v bool) {
	o.Subscribeall = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateSyncJobRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSyncJobRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateSyncJobRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CreateSyncJobRequest) SetActive(v bool) {
	o.Active = &v
}

func (o CreateSyncJobRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSyncJobRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Host1) {
		toSerialize["host1"] = o.Host1
	}
	if !IsNil(o.Port1) {
		toSerialize["port1"] = o.Port1
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Enc1) {
		toSerialize["enc1"] = o.Enc1
	}
	if !IsNil(o.MinsInternal) {
		toSerialize["mins_internal"] = o.MinsInternal
	}
	if !IsNil(o.Subfolder2) {
		toSerialize["subfolder2"] = o.Subfolder2
	}
	if !IsNil(o.Maxage) {
		toSerialize["maxage"] = o.Maxage
	}
	if !IsNil(o.Maxbytespersecond) {
		toSerialize["maxbytespersecond"] = o.Maxbytespersecond
	}
	if !IsNil(o.Timeout1) {
		toSerialize["timeout1"] = o.Timeout1
	}
	if !IsNil(o.Timeout2) {
		toSerialize["timeout2"] = o.Timeout2
	}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !IsNil(o.CustomParams) {
		toSerialize["custom_params"] = o.CustomParams
	}
	if !IsNil(o.Delete2duplicates) {
		toSerialize["delete2duplicates"] = o.Delete2duplicates
	}
	if !IsNil(o.Delete1) {
		toSerialize["delete1"] = o.Delete1
	}
	if !IsNil(o.Delete2) {
		toSerialize["delete2"] = o.Delete2
	}
	if !IsNil(o.Automap) {
		toSerialize["automap"] = o.Automap
	}
	if !IsNil(o.Skipcrossduplicates) {
		toSerialize["skipcrossduplicates"] = o.Skipcrossduplicates
	}
	if !IsNil(o.Subscribeall) {
		toSerialize["subscribeall"] = o.Subscribeall
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableCreateSyncJobRequest struct {
	value *CreateSyncJobRequest
	isSet bool
}

func (v NullableCreateSyncJobRequest) Get() *CreateSyncJobRequest {
	return v.value
}

func (v *NullableCreateSyncJobRequest) Set(val *CreateSyncJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSyncJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSyncJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSyncJobRequest(val *CreateSyncJobRequest) *NullableCreateSyncJobRequest {
	return &NullableCreateSyncJobRequest{value: val, isSet: true}
}

func (v NullableCreateSyncJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSyncJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


