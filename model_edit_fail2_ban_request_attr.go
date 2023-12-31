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

// checks if the EditFail2BanRequestAttr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditFail2BanRequestAttr{}

// EditFail2BanRequestAttr array containing the fail2ban settings
type EditFail2BanRequestAttr struct {
	// the backlisted ips or hostnames separated by comma
	Backlist *string `json:"backlist,omitempty"`
	// the time an ip should be banned
	BanTime *float32 `json:"ban_time,omitempty"`
	// if the time of the ban should increase each time
	BanTimeIncrement *bool `json:"ban_time_increment,omitempty"`
	// the maximum numbe of wrong logins before a ip is banned
	MaxAttempts *float32 `json:"max_attempts,omitempty"`
	// the maximum time an ip should be banned
	MaxBanTime *float32 `json:"max_ban_time,omitempty"`
	// the networks mask to ban for ipv4
	NetbanIpv4 *float32 `json:"netban_ipv4,omitempty"`
	// the networks mask to ban for ipv6
	NetbanIpv6 *float32 `json:"netban_ipv6,omitempty"`
	// the maximum time in which a ip as to login with false credentials to be banned
	RetryWindow *float32 `json:"retry_window,omitempty"`
	// whitelisted ips or hostnames sepereated by comma
	Whitelist *string `json:"whitelist,omitempty"`
}

// NewEditFail2BanRequestAttr instantiates a new EditFail2BanRequestAttr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditFail2BanRequestAttr() *EditFail2BanRequestAttr {
	this := EditFail2BanRequestAttr{}
	return &this
}

// NewEditFail2BanRequestAttrWithDefaults instantiates a new EditFail2BanRequestAttr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditFail2BanRequestAttrWithDefaults() *EditFail2BanRequestAttr {
	this := EditFail2BanRequestAttr{}
	return &this
}

// GetBacklist returns the Backlist field value if set, zero value otherwise.
func (o *EditFail2BanRequestAttr) GetBacklist() string {
	if o == nil || IsNil(o.Backlist) {
		var ret string
		return ret
	}
	return *o.Backlist
}

// GetBacklistOk returns a tuple with the Backlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditFail2BanRequestAttr) GetBacklistOk() (*string, bool) {
	if o == nil || IsNil(o.Backlist) {
		return nil, false
	}
	return o.Backlist, true
}

// HasBacklist returns a boolean if a field has been set.
func (o *EditFail2BanRequestAttr) HasBacklist() bool {
	if o != nil && !IsNil(o.Backlist) {
		return true
	}

	return false
}

// SetBacklist gets a reference to the given string and assigns it to the Backlist field.
func (o *EditFail2BanRequestAttr) SetBacklist(v string) {
	o.Backlist = &v
}

// GetBanTime returns the BanTime field value if set, zero value otherwise.
func (o *EditFail2BanRequestAttr) GetBanTime() float32 {
	if o == nil || IsNil(o.BanTime) {
		var ret float32
		return ret
	}
	return *o.BanTime
}

// GetBanTimeOk returns a tuple with the BanTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditFail2BanRequestAttr) GetBanTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.BanTime) {
		return nil, false
	}
	return o.BanTime, true
}

// HasBanTime returns a boolean if a field has been set.
func (o *EditFail2BanRequestAttr) HasBanTime() bool {
	if o != nil && !IsNil(o.BanTime) {
		return true
	}

	return false
}

// SetBanTime gets a reference to the given float32 and assigns it to the BanTime field.
func (o *EditFail2BanRequestAttr) SetBanTime(v float32) {
	o.BanTime = &v
}

// GetBanTimeIncrement returns the BanTimeIncrement field value if set, zero value otherwise.
func (o *EditFail2BanRequestAttr) GetBanTimeIncrement() bool {
	if o == nil || IsNil(o.BanTimeIncrement) {
		var ret bool
		return ret
	}
	return *o.BanTimeIncrement
}

// GetBanTimeIncrementOk returns a tuple with the BanTimeIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditFail2BanRequestAttr) GetBanTimeIncrementOk() (*bool, bool) {
	if o == nil || IsNil(o.BanTimeIncrement) {
		return nil, false
	}
	return o.BanTimeIncrement, true
}

// HasBanTimeIncrement returns a boolean if a field has been set.
func (o *EditFail2BanRequestAttr) HasBanTimeIncrement() bool {
	if o != nil && !IsNil(o.BanTimeIncrement) {
		return true
	}

	return false
}

// SetBanTimeIncrement gets a reference to the given bool and assigns it to the BanTimeIncrement field.
func (o *EditFail2BanRequestAttr) SetBanTimeIncrement(v bool) {
	o.BanTimeIncrement = &v
}

// GetMaxAttempts returns the MaxAttempts field value if set, zero value otherwise.
func (o *EditFail2BanRequestAttr) GetMaxAttempts() float32 {
	if o == nil || IsNil(o.MaxAttempts) {
		var ret float32
		return ret
	}
	return *o.MaxAttempts
}

// GetMaxAttemptsOk returns a tuple with the MaxAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditFail2BanRequestAttr) GetMaxAttemptsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxAttempts) {
		return nil, false
	}
	return o.MaxAttempts, true
}

// HasMaxAttempts returns a boolean if a field has been set.
func (o *EditFail2BanRequestAttr) HasMaxAttempts() bool {
	if o != nil && !IsNil(o.MaxAttempts) {
		return true
	}

	return false
}

// SetMaxAttempts gets a reference to the given float32 and assigns it to the MaxAttempts field.
func (o *EditFail2BanRequestAttr) SetMaxAttempts(v float32) {
	o.MaxAttempts = &v
}

// GetMaxBanTime returns the MaxBanTime field value if set, zero value otherwise.
func (o *EditFail2BanRequestAttr) GetMaxBanTime() float32 {
	if o == nil || IsNil(o.MaxBanTime) {
		var ret float32
		return ret
	}
	return *o.MaxBanTime
}

// GetMaxBanTimeOk returns a tuple with the MaxBanTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditFail2BanRequestAttr) GetMaxBanTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxBanTime) {
		return nil, false
	}
	return o.MaxBanTime, true
}

// HasMaxBanTime returns a boolean if a field has been set.
func (o *EditFail2BanRequestAttr) HasMaxBanTime() bool {
	if o != nil && !IsNil(o.MaxBanTime) {
		return true
	}

	return false
}

// SetMaxBanTime gets a reference to the given float32 and assigns it to the MaxBanTime field.
func (o *EditFail2BanRequestAttr) SetMaxBanTime(v float32) {
	o.MaxBanTime = &v
}

// GetNetbanIpv4 returns the NetbanIpv4 field value if set, zero value otherwise.
func (o *EditFail2BanRequestAttr) GetNetbanIpv4() float32 {
	if o == nil || IsNil(o.NetbanIpv4) {
		var ret float32
		return ret
	}
	return *o.NetbanIpv4
}

// GetNetbanIpv4Ok returns a tuple with the NetbanIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditFail2BanRequestAttr) GetNetbanIpv4Ok() (*float32, bool) {
	if o == nil || IsNil(o.NetbanIpv4) {
		return nil, false
	}
	return o.NetbanIpv4, true
}

// HasNetbanIpv4 returns a boolean if a field has been set.
func (o *EditFail2BanRequestAttr) HasNetbanIpv4() bool {
	if o != nil && !IsNil(o.NetbanIpv4) {
		return true
	}

	return false
}

// SetNetbanIpv4 gets a reference to the given float32 and assigns it to the NetbanIpv4 field.
func (o *EditFail2BanRequestAttr) SetNetbanIpv4(v float32) {
	o.NetbanIpv4 = &v
}

// GetNetbanIpv6 returns the NetbanIpv6 field value if set, zero value otherwise.
func (o *EditFail2BanRequestAttr) GetNetbanIpv6() float32 {
	if o == nil || IsNil(o.NetbanIpv6) {
		var ret float32
		return ret
	}
	return *o.NetbanIpv6
}

// GetNetbanIpv6Ok returns a tuple with the NetbanIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditFail2BanRequestAttr) GetNetbanIpv6Ok() (*float32, bool) {
	if o == nil || IsNil(o.NetbanIpv6) {
		return nil, false
	}
	return o.NetbanIpv6, true
}

// HasNetbanIpv6 returns a boolean if a field has been set.
func (o *EditFail2BanRequestAttr) HasNetbanIpv6() bool {
	if o != nil && !IsNil(o.NetbanIpv6) {
		return true
	}

	return false
}

// SetNetbanIpv6 gets a reference to the given float32 and assigns it to the NetbanIpv6 field.
func (o *EditFail2BanRequestAttr) SetNetbanIpv6(v float32) {
	o.NetbanIpv6 = &v
}

// GetRetryWindow returns the RetryWindow field value if set, zero value otherwise.
func (o *EditFail2BanRequestAttr) GetRetryWindow() float32 {
	if o == nil || IsNil(o.RetryWindow) {
		var ret float32
		return ret
	}
	return *o.RetryWindow
}

// GetRetryWindowOk returns a tuple with the RetryWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditFail2BanRequestAttr) GetRetryWindowOk() (*float32, bool) {
	if o == nil || IsNil(o.RetryWindow) {
		return nil, false
	}
	return o.RetryWindow, true
}

// HasRetryWindow returns a boolean if a field has been set.
func (o *EditFail2BanRequestAttr) HasRetryWindow() bool {
	if o != nil && !IsNil(o.RetryWindow) {
		return true
	}

	return false
}

// SetRetryWindow gets a reference to the given float32 and assigns it to the RetryWindow field.
func (o *EditFail2BanRequestAttr) SetRetryWindow(v float32) {
	o.RetryWindow = &v
}

// GetWhitelist returns the Whitelist field value if set, zero value otherwise.
func (o *EditFail2BanRequestAttr) GetWhitelist() string {
	if o == nil || IsNil(o.Whitelist) {
		var ret string
		return ret
	}
	return *o.Whitelist
}

// GetWhitelistOk returns a tuple with the Whitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditFail2BanRequestAttr) GetWhitelistOk() (*string, bool) {
	if o == nil || IsNil(o.Whitelist) {
		return nil, false
	}
	return o.Whitelist, true
}

// HasWhitelist returns a boolean if a field has been set.
func (o *EditFail2BanRequestAttr) HasWhitelist() bool {
	if o != nil && !IsNil(o.Whitelist) {
		return true
	}

	return false
}

// SetWhitelist gets a reference to the given string and assigns it to the Whitelist field.
func (o *EditFail2BanRequestAttr) SetWhitelist(v string) {
	o.Whitelist = &v
}

func (o EditFail2BanRequestAttr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditFail2BanRequestAttr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Backlist) {
		toSerialize["backlist"] = o.Backlist
	}
	if !IsNil(o.BanTime) {
		toSerialize["ban_time"] = o.BanTime
	}
	if !IsNil(o.BanTimeIncrement) {
		toSerialize["ban_time_increment"] = o.BanTimeIncrement
	}
	if !IsNil(o.MaxAttempts) {
		toSerialize["max_attempts"] = o.MaxAttempts
	}
	if !IsNil(o.MaxBanTime) {
		toSerialize["max_ban_time"] = o.MaxBanTime
	}
	if !IsNil(o.NetbanIpv4) {
		toSerialize["netban_ipv4"] = o.NetbanIpv4
	}
	if !IsNil(o.NetbanIpv6) {
		toSerialize["netban_ipv6"] = o.NetbanIpv6
	}
	if !IsNil(o.RetryWindow) {
		toSerialize["retry_window"] = o.RetryWindow
	}
	if !IsNil(o.Whitelist) {
		toSerialize["whitelist"] = o.Whitelist
	}
	return toSerialize, nil
}

type NullableEditFail2BanRequestAttr struct {
	value *EditFail2BanRequestAttr
	isSet bool
}

func (v NullableEditFail2BanRequestAttr) Get() *EditFail2BanRequestAttr {
	return v.value
}

func (v *NullableEditFail2BanRequestAttr) Set(val *EditFail2BanRequestAttr) {
	v.value = val
	v.isSet = true
}

func (v NullableEditFail2BanRequestAttr) IsSet() bool {
	return v.isSet
}

func (v *NullableEditFail2BanRequestAttr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditFail2BanRequestAttr(val *EditFail2BanRequestAttr) *NullableEditFail2BanRequestAttr {
	return &NullableEditFail2BanRequestAttr{value: val, isSet: true}
}

func (v NullableEditFail2BanRequestAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditFail2BanRequestAttr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


