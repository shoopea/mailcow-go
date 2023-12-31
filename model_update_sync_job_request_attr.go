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

// checks if the UpdateSyncJobRequestAttr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSyncJobRequestAttr{}

// UpdateSyncJobRequestAttr struct for UpdateSyncJobRequestAttr
type UpdateSyncJobRequestAttr struct {
	// Is sync job active
	Active *bool `json:"active,omitempty"`
	// Try to automap folders (\"Sent items\", \"Sent\" => \"Sent\" etc.)
	Automap *bool `json:"automap,omitempty"`
	// Custom parameters passed to imapsync command
	CustomParams *string `json:"custom_params,omitempty"`
	// Delete from source when completed
	Delete1 *bool `json:"delete1,omitempty"`
	// Delete messages on destination that are not on source
	Delete2 *bool `json:"delete2,omitempty"`
	// Delete duplicates on destination
	Delete2duplicates *bool `json:"delete2duplicates,omitempty"`
	// Encryption
	Enc1 *string `json:"enc1,omitempty"`
	// Exclude objects (regex)
	Exclude *string `json:"exclude,omitempty"`
	// Hostname
	Host1 *string `json:"host1,omitempty"`
	// Maximum age of messages in days that will be polled from remote (0 = ignore age)
	Maxage *float32 `json:"maxage,omitempty"`
	// Max. bytes per second (0 = unlimited)
	Maxbytespersecond *float32 `json:"maxbytespersecond,omitempty"`
	// Interval (min)
	MinsInterval *float32 `json:"mins_interval,omitempty"`
	// Password
	Password1 *string `json:"password1,omitempty"`
	// Port
	Port1 *string `json:"port1,omitempty"`
	// Skip duplicate messages across folders (first come, first serve)
	Skipcrossduplicates *bool `json:"skipcrossduplicates,omitempty"`
	// Sync into subfolder on destination (empty = do not use subfolder)
	Subfolder2 *string `json:"subfolder2,omitempty"`
	// Subscribe all folders
	Subscribeall *bool `json:"subscribeall,omitempty"`
	// Timeout for connection to remote host
	Timeout1 *float32 `json:"timeout1,omitempty"`
	// Timeout for connection to local host
	Timeout2 *float32 `json:"timeout2,omitempty"`
	// Username
	User1 *string `json:"user1,omitempty"`
}

// NewUpdateSyncJobRequestAttr instantiates a new UpdateSyncJobRequestAttr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSyncJobRequestAttr() *UpdateSyncJobRequestAttr {
	this := UpdateSyncJobRequestAttr{}
	return &this
}

// NewUpdateSyncJobRequestAttrWithDefaults instantiates a new UpdateSyncJobRequestAttr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSyncJobRequestAttrWithDefaults() *UpdateSyncJobRequestAttr {
	this := UpdateSyncJobRequestAttr{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateSyncJobRequestAttr) SetActive(v bool) {
	o.Active = &v
}

// GetAutomap returns the Automap field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetAutomap() bool {
	if o == nil || IsNil(o.Automap) {
		var ret bool
		return ret
	}
	return *o.Automap
}

// GetAutomapOk returns a tuple with the Automap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetAutomapOk() (*bool, bool) {
	if o == nil || IsNil(o.Automap) {
		return nil, false
	}
	return o.Automap, true
}

// HasAutomap returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasAutomap() bool {
	if o != nil && !IsNil(o.Automap) {
		return true
	}

	return false
}

// SetAutomap gets a reference to the given bool and assigns it to the Automap field.
func (o *UpdateSyncJobRequestAttr) SetAutomap(v bool) {
	o.Automap = &v
}

// GetCustomParams returns the CustomParams field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetCustomParams() string {
	if o == nil || IsNil(o.CustomParams) {
		var ret string
		return ret
	}
	return *o.CustomParams
}

// GetCustomParamsOk returns a tuple with the CustomParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetCustomParamsOk() (*string, bool) {
	if o == nil || IsNil(o.CustomParams) {
		return nil, false
	}
	return o.CustomParams, true
}

// HasCustomParams returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasCustomParams() bool {
	if o != nil && !IsNil(o.CustomParams) {
		return true
	}

	return false
}

// SetCustomParams gets a reference to the given string and assigns it to the CustomParams field.
func (o *UpdateSyncJobRequestAttr) SetCustomParams(v string) {
	o.CustomParams = &v
}

// GetDelete1 returns the Delete1 field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetDelete1() bool {
	if o == nil || IsNil(o.Delete1) {
		var ret bool
		return ret
	}
	return *o.Delete1
}

// GetDelete1Ok returns a tuple with the Delete1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetDelete1Ok() (*bool, bool) {
	if o == nil || IsNil(o.Delete1) {
		return nil, false
	}
	return o.Delete1, true
}

// HasDelete1 returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasDelete1() bool {
	if o != nil && !IsNil(o.Delete1) {
		return true
	}

	return false
}

// SetDelete1 gets a reference to the given bool and assigns it to the Delete1 field.
func (o *UpdateSyncJobRequestAttr) SetDelete1(v bool) {
	o.Delete1 = &v
}

// GetDelete2 returns the Delete2 field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetDelete2() bool {
	if o == nil || IsNil(o.Delete2) {
		var ret bool
		return ret
	}
	return *o.Delete2
}

// GetDelete2Ok returns a tuple with the Delete2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetDelete2Ok() (*bool, bool) {
	if o == nil || IsNil(o.Delete2) {
		return nil, false
	}
	return o.Delete2, true
}

// HasDelete2 returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasDelete2() bool {
	if o != nil && !IsNil(o.Delete2) {
		return true
	}

	return false
}

// SetDelete2 gets a reference to the given bool and assigns it to the Delete2 field.
func (o *UpdateSyncJobRequestAttr) SetDelete2(v bool) {
	o.Delete2 = &v
}

// GetDelete2duplicates returns the Delete2duplicates field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetDelete2duplicates() bool {
	if o == nil || IsNil(o.Delete2duplicates) {
		var ret bool
		return ret
	}
	return *o.Delete2duplicates
}

// GetDelete2duplicatesOk returns a tuple with the Delete2duplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetDelete2duplicatesOk() (*bool, bool) {
	if o == nil || IsNil(o.Delete2duplicates) {
		return nil, false
	}
	return o.Delete2duplicates, true
}

// HasDelete2duplicates returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasDelete2duplicates() bool {
	if o != nil && !IsNil(o.Delete2duplicates) {
		return true
	}

	return false
}

// SetDelete2duplicates gets a reference to the given bool and assigns it to the Delete2duplicates field.
func (o *UpdateSyncJobRequestAttr) SetDelete2duplicates(v bool) {
	o.Delete2duplicates = &v
}

// GetEnc1 returns the Enc1 field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetEnc1() string {
	if o == nil || IsNil(o.Enc1) {
		var ret string
		return ret
	}
	return *o.Enc1
}

// GetEnc1Ok returns a tuple with the Enc1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetEnc1Ok() (*string, bool) {
	if o == nil || IsNil(o.Enc1) {
		return nil, false
	}
	return o.Enc1, true
}

// HasEnc1 returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasEnc1() bool {
	if o != nil && !IsNil(o.Enc1) {
		return true
	}

	return false
}

// SetEnc1 gets a reference to the given string and assigns it to the Enc1 field.
func (o *UpdateSyncJobRequestAttr) SetEnc1(v string) {
	o.Enc1 = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetExclude() string {
	if o == nil || IsNil(o.Exclude) {
		var ret string
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetExcludeOk() (*string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given string and assigns it to the Exclude field.
func (o *UpdateSyncJobRequestAttr) SetExclude(v string) {
	o.Exclude = &v
}

// GetHost1 returns the Host1 field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetHost1() string {
	if o == nil || IsNil(o.Host1) {
		var ret string
		return ret
	}
	return *o.Host1
}

// GetHost1Ok returns a tuple with the Host1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetHost1Ok() (*string, bool) {
	if o == nil || IsNil(o.Host1) {
		return nil, false
	}
	return o.Host1, true
}

// HasHost1 returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasHost1() bool {
	if o != nil && !IsNil(o.Host1) {
		return true
	}

	return false
}

// SetHost1 gets a reference to the given string and assigns it to the Host1 field.
func (o *UpdateSyncJobRequestAttr) SetHost1(v string) {
	o.Host1 = &v
}

// GetMaxage returns the Maxage field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetMaxage() float32 {
	if o == nil || IsNil(o.Maxage) {
		var ret float32
		return ret
	}
	return *o.Maxage
}

// GetMaxageOk returns a tuple with the Maxage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetMaxageOk() (*float32, bool) {
	if o == nil || IsNil(o.Maxage) {
		return nil, false
	}
	return o.Maxage, true
}

// HasMaxage returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasMaxage() bool {
	if o != nil && !IsNil(o.Maxage) {
		return true
	}

	return false
}

// SetMaxage gets a reference to the given float32 and assigns it to the Maxage field.
func (o *UpdateSyncJobRequestAttr) SetMaxage(v float32) {
	o.Maxage = &v
}

// GetMaxbytespersecond returns the Maxbytespersecond field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetMaxbytespersecond() float32 {
	if o == nil || IsNil(o.Maxbytespersecond) {
		var ret float32
		return ret
	}
	return *o.Maxbytespersecond
}

// GetMaxbytespersecondOk returns a tuple with the Maxbytespersecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetMaxbytespersecondOk() (*float32, bool) {
	if o == nil || IsNil(o.Maxbytespersecond) {
		return nil, false
	}
	return o.Maxbytespersecond, true
}

// HasMaxbytespersecond returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasMaxbytespersecond() bool {
	if o != nil && !IsNil(o.Maxbytespersecond) {
		return true
	}

	return false
}

// SetMaxbytespersecond gets a reference to the given float32 and assigns it to the Maxbytespersecond field.
func (o *UpdateSyncJobRequestAttr) SetMaxbytespersecond(v float32) {
	o.Maxbytespersecond = &v
}

// GetMinsInterval returns the MinsInterval field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetMinsInterval() float32 {
	if o == nil || IsNil(o.MinsInterval) {
		var ret float32
		return ret
	}
	return *o.MinsInterval
}

// GetMinsIntervalOk returns a tuple with the MinsInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetMinsIntervalOk() (*float32, bool) {
	if o == nil || IsNil(o.MinsInterval) {
		return nil, false
	}
	return o.MinsInterval, true
}

// HasMinsInterval returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasMinsInterval() bool {
	if o != nil && !IsNil(o.MinsInterval) {
		return true
	}

	return false
}

// SetMinsInterval gets a reference to the given float32 and assigns it to the MinsInterval field.
func (o *UpdateSyncJobRequestAttr) SetMinsInterval(v float32) {
	o.MinsInterval = &v
}

// GetPassword1 returns the Password1 field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetPassword1() string {
	if o == nil || IsNil(o.Password1) {
		var ret string
		return ret
	}
	return *o.Password1
}

// GetPassword1Ok returns a tuple with the Password1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetPassword1Ok() (*string, bool) {
	if o == nil || IsNil(o.Password1) {
		return nil, false
	}
	return o.Password1, true
}

// HasPassword1 returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasPassword1() bool {
	if o != nil && !IsNil(o.Password1) {
		return true
	}

	return false
}

// SetPassword1 gets a reference to the given string and assigns it to the Password1 field.
func (o *UpdateSyncJobRequestAttr) SetPassword1(v string) {
	o.Password1 = &v
}

// GetPort1 returns the Port1 field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetPort1() string {
	if o == nil || IsNil(o.Port1) {
		var ret string
		return ret
	}
	return *o.Port1
}

// GetPort1Ok returns a tuple with the Port1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetPort1Ok() (*string, bool) {
	if o == nil || IsNil(o.Port1) {
		return nil, false
	}
	return o.Port1, true
}

// HasPort1 returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasPort1() bool {
	if o != nil && !IsNil(o.Port1) {
		return true
	}

	return false
}

// SetPort1 gets a reference to the given string and assigns it to the Port1 field.
func (o *UpdateSyncJobRequestAttr) SetPort1(v string) {
	o.Port1 = &v
}

// GetSkipcrossduplicates returns the Skipcrossduplicates field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetSkipcrossduplicates() bool {
	if o == nil || IsNil(o.Skipcrossduplicates) {
		var ret bool
		return ret
	}
	return *o.Skipcrossduplicates
}

// GetSkipcrossduplicatesOk returns a tuple with the Skipcrossduplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetSkipcrossduplicatesOk() (*bool, bool) {
	if o == nil || IsNil(o.Skipcrossduplicates) {
		return nil, false
	}
	return o.Skipcrossduplicates, true
}

// HasSkipcrossduplicates returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasSkipcrossduplicates() bool {
	if o != nil && !IsNil(o.Skipcrossduplicates) {
		return true
	}

	return false
}

// SetSkipcrossduplicates gets a reference to the given bool and assigns it to the Skipcrossduplicates field.
func (o *UpdateSyncJobRequestAttr) SetSkipcrossduplicates(v bool) {
	o.Skipcrossduplicates = &v
}

// GetSubfolder2 returns the Subfolder2 field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetSubfolder2() string {
	if o == nil || IsNil(o.Subfolder2) {
		var ret string
		return ret
	}
	return *o.Subfolder2
}

// GetSubfolder2Ok returns a tuple with the Subfolder2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetSubfolder2Ok() (*string, bool) {
	if o == nil || IsNil(o.Subfolder2) {
		return nil, false
	}
	return o.Subfolder2, true
}

// HasSubfolder2 returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasSubfolder2() bool {
	if o != nil && !IsNil(o.Subfolder2) {
		return true
	}

	return false
}

// SetSubfolder2 gets a reference to the given string and assigns it to the Subfolder2 field.
func (o *UpdateSyncJobRequestAttr) SetSubfolder2(v string) {
	o.Subfolder2 = &v
}

// GetSubscribeall returns the Subscribeall field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetSubscribeall() bool {
	if o == nil || IsNil(o.Subscribeall) {
		var ret bool
		return ret
	}
	return *o.Subscribeall
}

// GetSubscribeallOk returns a tuple with the Subscribeall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetSubscribeallOk() (*bool, bool) {
	if o == nil || IsNil(o.Subscribeall) {
		return nil, false
	}
	return o.Subscribeall, true
}

// HasSubscribeall returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasSubscribeall() bool {
	if o != nil && !IsNil(o.Subscribeall) {
		return true
	}

	return false
}

// SetSubscribeall gets a reference to the given bool and assigns it to the Subscribeall field.
func (o *UpdateSyncJobRequestAttr) SetSubscribeall(v bool) {
	o.Subscribeall = &v
}

// GetTimeout1 returns the Timeout1 field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetTimeout1() float32 {
	if o == nil || IsNil(o.Timeout1) {
		var ret float32
		return ret
	}
	return *o.Timeout1
}

// GetTimeout1Ok returns a tuple with the Timeout1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetTimeout1Ok() (*float32, bool) {
	if o == nil || IsNil(o.Timeout1) {
		return nil, false
	}
	return o.Timeout1, true
}

// HasTimeout1 returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasTimeout1() bool {
	if o != nil && !IsNil(o.Timeout1) {
		return true
	}

	return false
}

// SetTimeout1 gets a reference to the given float32 and assigns it to the Timeout1 field.
func (o *UpdateSyncJobRequestAttr) SetTimeout1(v float32) {
	o.Timeout1 = &v
}

// GetTimeout2 returns the Timeout2 field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetTimeout2() float32 {
	if o == nil || IsNil(o.Timeout2) {
		var ret float32
		return ret
	}
	return *o.Timeout2
}

// GetTimeout2Ok returns a tuple with the Timeout2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetTimeout2Ok() (*float32, bool) {
	if o == nil || IsNil(o.Timeout2) {
		return nil, false
	}
	return o.Timeout2, true
}

// HasTimeout2 returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasTimeout2() bool {
	if o != nil && !IsNil(o.Timeout2) {
		return true
	}

	return false
}

// SetTimeout2 gets a reference to the given float32 and assigns it to the Timeout2 field.
func (o *UpdateSyncJobRequestAttr) SetTimeout2(v float32) {
	o.Timeout2 = &v
}

// GetUser1 returns the User1 field value if set, zero value otherwise.
func (o *UpdateSyncJobRequestAttr) GetUser1() string {
	if o == nil || IsNil(o.User1) {
		var ret string
		return ret
	}
	return *o.User1
}

// GetUser1Ok returns a tuple with the User1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSyncJobRequestAttr) GetUser1Ok() (*string, bool) {
	if o == nil || IsNil(o.User1) {
		return nil, false
	}
	return o.User1, true
}

// HasUser1 returns a boolean if a field has been set.
func (o *UpdateSyncJobRequestAttr) HasUser1() bool {
	if o != nil && !IsNil(o.User1) {
		return true
	}

	return false
}

// SetUser1 gets a reference to the given string and assigns it to the User1 field.
func (o *UpdateSyncJobRequestAttr) SetUser1(v string) {
	o.User1 = &v
}

func (o UpdateSyncJobRequestAttr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSyncJobRequestAttr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Automap) {
		toSerialize["automap"] = o.Automap
	}
	if !IsNil(o.CustomParams) {
		toSerialize["custom_params"] = o.CustomParams
	}
	if !IsNil(o.Delete1) {
		toSerialize["delete1"] = o.Delete1
	}
	if !IsNil(o.Delete2) {
		toSerialize["delete2"] = o.Delete2
	}
	if !IsNil(o.Delete2duplicates) {
		toSerialize["delete2duplicates"] = o.Delete2duplicates
	}
	if !IsNil(o.Enc1) {
		toSerialize["enc1"] = o.Enc1
	}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !IsNil(o.Host1) {
		toSerialize["host1"] = o.Host1
	}
	if !IsNil(o.Maxage) {
		toSerialize["maxage"] = o.Maxage
	}
	if !IsNil(o.Maxbytespersecond) {
		toSerialize["maxbytespersecond"] = o.Maxbytespersecond
	}
	if !IsNil(o.MinsInterval) {
		toSerialize["mins_interval"] = o.MinsInterval
	}
	if !IsNil(o.Password1) {
		toSerialize["password1"] = o.Password1
	}
	if !IsNil(o.Port1) {
		toSerialize["port1"] = o.Port1
	}
	if !IsNil(o.Skipcrossduplicates) {
		toSerialize["skipcrossduplicates"] = o.Skipcrossduplicates
	}
	if !IsNil(o.Subfolder2) {
		toSerialize["subfolder2"] = o.Subfolder2
	}
	if !IsNil(o.Subscribeall) {
		toSerialize["subscribeall"] = o.Subscribeall
	}
	if !IsNil(o.Timeout1) {
		toSerialize["timeout1"] = o.Timeout1
	}
	if !IsNil(o.Timeout2) {
		toSerialize["timeout2"] = o.Timeout2
	}
	if !IsNil(o.User1) {
		toSerialize["user1"] = o.User1
	}
	return toSerialize, nil
}

type NullableUpdateSyncJobRequestAttr struct {
	value *UpdateSyncJobRequestAttr
	isSet bool
}

func (v NullableUpdateSyncJobRequestAttr) Get() *UpdateSyncJobRequestAttr {
	return v.value
}

func (v *NullableUpdateSyncJobRequestAttr) Set(val *UpdateSyncJobRequestAttr) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSyncJobRequestAttr) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSyncJobRequestAttr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSyncJobRequestAttr(val *UpdateSyncJobRequestAttr) *NullableUpdateSyncJobRequestAttr {
	return &NullableUpdateSyncJobRequestAttr{value: val, isSet: true}
}

func (v NullableUpdateSyncJobRequestAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSyncJobRequestAttr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


