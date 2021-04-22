// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bas.proto

package pb

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
)

// Validate checks the field values on DlxConsumerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DlxConsumerRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Header

	// no validation rules for Body

	return nil
}

// DlxConsumerRequestValidationError is the validation error returned by
// DlxConsumerRequest.Validate if the designated constraints aren't met.
type DlxConsumerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DlxConsumerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DlxConsumerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DlxConsumerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DlxConsumerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DlxConsumerRequestValidationError) ErrorName() string {
	return "DlxConsumerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DlxConsumerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDlxConsumerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DlxConsumerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DlxConsumerRequestValidationError{}

// Validate checks the field values on BaseResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *BaseResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// BaseResponseValidationError is the validation error returned by
// BaseResponse.Validate if the designated constraints aren't met.
type BaseResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BaseResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BaseResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BaseResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BaseResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BaseResponseValidationError) ErrorName() string { return "BaseResponseValidationError" }

// Error satisfies the builtin error interface
func (e BaseResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBaseResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BaseResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BaseResponseValidationError{}
