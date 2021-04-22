// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: usr.proto

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

// Validate checks the field values on UserResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for Message

	// no validation rules for Code

	// no validation rules for Error

	// no validation rules for Details

	return nil
}

// UserResponseValidationError is the validation error returned by
// UserResponse.Validate if the designated constraints aren't met.
type UserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserResponseValidationError) ErrorName() string { return "UserResponseValidationError" }

// Error satisfies the builtin error interface
func (e UserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserResponseValidationError{}

// Validate checks the field values on UserListRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UserListRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PageSize

	// no validation rules for PageNumber

	// no validation rules for OrderKey

	// no validation rules for OrderSort

	return nil
}

// UserListRequestValidationError is the validation error returned by
// UserListRequest.Validate if the designated constraints aren't met.
type UserListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListRequestValidationError) ErrorName() string { return "UserListRequestValidationError" }

// Error satisfies the builtin error interface
func (e UserListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListRequestValidationError{}

// Validate checks the field values on UserListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UserListResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for Message

	// no validation rules for Code

	// no validation rules for Error

	if v, ok := interface{}(m.GetDetails()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserListResponseValidationError{
				field:  "Details",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UserListResponseValidationError is the validation error returned by
// UserListResponse.Validate if the designated constraints aren't met.
type UserListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListResponseValidationError) ErrorName() string { return "UserListResponseValidationError" }

// Error satisfies the builtin error interface
func (e UserListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListResponseValidationError{}

// Validate checks the field values on UserList with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UserList) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserListValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// UserListValidationError is the validation error returned by
// UserList.Validate if the designated constraints aren't met.
type UserListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListValidationError) ErrorName() string { return "UserListValidationError" }

// Error satisfies the builtin error interface
func (e UserListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListValidationError{}

// Validate checks the field values on UserOneRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserOneRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Password

	return nil
}

// UserOneRequestValidationError is the validation error returned by
// UserOneRequest.Validate if the designated constraints aren't met.
type UserOneRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserOneRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserOneRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserOneRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserOneRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserOneRequestValidationError) ErrorName() string { return "UserOneRequestValidationError" }

// Error satisfies the builtin error interface
func (e UserOneRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserOneRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserOneRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserOneRequestValidationError{}

// Validate checks the field values on UserIdRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserIdRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// UserIdRequestValidationError is the validation error returned by
// UserIdRequest.Validate if the designated constraints aren't met.
type UserIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserIdRequestValidationError) ErrorName() string { return "UserIdRequestValidationError" }

// Error satisfies the builtin error interface
func (e UserIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserIdRequestValidationError{}

// Validate checks the field values on UserDetailResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UserDetailResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for Message

	// no validation rules for Code

	// no validation rules for Error

	if v, ok := interface{}(m.GetDetails()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDetailResponseValidationError{
				field:  "Details",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UserDetailResponseValidationError is the validation error returned by
// UserDetailResponse.Validate if the designated constraints aren't met.
type UserDetailResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserDetailResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserDetailResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserDetailResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserDetailResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserDetailResponseValidationError) ErrorName() string {
	return "UserDetailResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UserDetailResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserDetailResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserDetailResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserDetailResponseValidationError{}
