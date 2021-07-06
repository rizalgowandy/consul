// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/common/dynamic_forward_proxy/v3/dns_cache.proto

package envoy_extensions_common_dynamic_forward_proxy_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"

	v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}

	_ = v3.Cluster_DnsLookupFamily(0)
)

// define the regex for a UUID once up-front
var _dns_cache_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on DnsCacheConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DnsCacheConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return DnsCacheConfigValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	if _, ok := v3.Cluster_DnsLookupFamily_name[int32(m.GetDnsLookupFamily())]; !ok {
		return DnsCacheConfigValidationError{
			field:  "DnsLookupFamily",
			reason: "value must be one of the defined enum values",
		}
	}

	if d := m.GetDnsRefreshRate(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return DnsCacheConfigValidationError{
				field:  "DnsRefreshRate",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gte := time.Duration(0*time.Second + 1000000*time.Nanosecond)

		if dur < gte {
			return DnsCacheConfigValidationError{
				field:  "DnsRefreshRate",
				reason: "value must be greater than or equal to 1ms",
			}
		}

	}

	if d := m.GetHostTtl(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return DnsCacheConfigValidationError{
				field:  "HostTtl",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return DnsCacheConfigValidationError{
				field:  "HostTtl",
				reason: "value must be greater than 0s",
			}
		}

	}

	if wrapper := m.GetMaxHosts(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			return DnsCacheConfigValidationError{
				field:  "MaxHosts",
				reason: "value must be greater than 0",
			}
		}

	}

	if v, ok := interface{}(m.GetDnsFailureRefreshRate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheConfigValidationError{
				field:  "DnsFailureRefreshRate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DnsCacheConfigValidationError is the validation error returned by
// DnsCacheConfig.Validate if the designated constraints aren't met.
type DnsCacheConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsCacheConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsCacheConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsCacheConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsCacheConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsCacheConfigValidationError) ErrorName() string { return "DnsCacheConfigValidationError" }

// Error satisfies the builtin error interface
func (e DnsCacheConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsCacheConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsCacheConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsCacheConfigValidationError{}
