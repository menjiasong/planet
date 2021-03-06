// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pro.proto

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

// Validate checks the field values on ProductResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ProductResponse) Validate() error {
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

// ProductResponseValidationError is the validation error returned by
// ProductResponse.Validate if the designated constraints aren't met.
type ProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductResponseValidationError) ErrorName() string { return "ProductResponseValidationError" }

// Error satisfies the builtin error interface
func (e ProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductResponseValidationError{}

// Validate checks the field values on ProductListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProductListRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PageSize

	// no validation rules for PageNumber

	// no validation rules for OrderKey

	// no validation rules for OrderSort

	return nil
}

// ProductListRequestValidationError is the validation error returned by
// ProductListRequest.Validate if the designated constraints aren't met.
type ProductListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductListRequestValidationError) ErrorName() string {
	return "ProductListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProductListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductListRequestValidationError{}

// Validate checks the field values on ProductListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProductListResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for Message

	// no validation rules for Code

	// no validation rules for Error

	if v, ok := interface{}(m.GetDetails()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductListResponseValidationError{
				field:  "Details",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProductListResponseValidationError is the validation error returned by
// ProductListResponse.Validate if the designated constraints aren't met.
type ProductListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductListResponseValidationError) ErrorName() string {
	return "ProductListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProductListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductListResponseValidationError{}

// Validate checks the field values on ProductList with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ProductList) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProductListValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProductListValidationError is the validation error returned by
// ProductList.Validate if the designated constraints aren't met.
type ProductListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductListValidationError) ErrorName() string { return "ProductListValidationError" }

// Error satisfies the builtin error interface
func (e ProductListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductListValidationError{}

// Validate checks the field values on ProductOneRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ProductOneRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for DeletedAt

	// no validation rules for Code

	// no validation rules for Price

	return nil
}

// ProductOneRequestValidationError is the validation error returned by
// ProductOneRequest.Validate if the designated constraints aren't met.
type ProductOneRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductOneRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductOneRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductOneRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductOneRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductOneRequestValidationError) ErrorName() string {
	return "ProductOneRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProductOneRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductOneRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductOneRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductOneRequestValidationError{}

// Validate checks the field values on ProductIdRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ProductIdRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// ProductIdRequestValidationError is the validation error returned by
// ProductIdRequest.Validate if the designated constraints aren't met.
type ProductIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductIdRequestValidationError) ErrorName() string { return "ProductIdRequestValidationError" }

// Error satisfies the builtin error interface
func (e ProductIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductIdRequestValidationError{}

// Validate checks the field values on ProductDetailResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProductDetailResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for Message

	// no validation rules for Code

	// no validation rules for Error

	if v, ok := interface{}(m.GetDetails()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductDetailResponseValidationError{
				field:  "Details",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProductDetailResponseValidationError is the validation error returned by
// ProductDetailResponse.Validate if the designated constraints aren't met.
type ProductDetailResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductDetailResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductDetailResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductDetailResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductDetailResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductDetailResponseValidationError) ErrorName() string {
	return "ProductDetailResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProductDetailResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductDetailResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductDetailResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductDetailResponseValidationError{}
