// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-certificate-api/ocp-certificate-api.proto

package ocp_certificate_api

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate checks the field values on MultiCreateCertificatesV1Request with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *MultiCreateCertificatesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCertificates()) < 1 {
		return MultiCreateCertificatesV1RequestValidationError{
			field:  "Certificates",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetCertificates() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateCertificatesV1RequestValidationError{
					field:  fmt.Sprintf("Certificates[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateCertificatesV1RequestValidationError is the validation error
// returned by MultiCreateCertificatesV1Request.Validate if the designated
// constraints aren't met.
type MultiCreateCertificatesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateCertificatesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateCertificatesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateCertificatesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateCertificatesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateCertificatesV1RequestValidationError) ErrorName() string {
	return "MultiCreateCertificatesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateCertificatesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateCertificatesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateCertificatesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateCertificatesV1RequestValidationError{}

// Validate checks the field values on MultiCreateCertificatesV1Response with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *MultiCreateCertificatesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCertificateIds()) < 1 {
		return MultiCreateCertificatesV1ResponseValidationError{
			field:  "CertificateIds",
			reason: "value must contain at least 1 item(s)",
		}
	}

	return nil
}

// MultiCreateCertificatesV1ResponseValidationError is the validation error
// returned by MultiCreateCertificatesV1Response.Validate if the designated
// constraints aren't met.
type MultiCreateCertificatesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateCertificatesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateCertificatesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateCertificatesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateCertificatesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateCertificatesV1ResponseValidationError) ErrorName() string {
	return "MultiCreateCertificatesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateCertificatesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateCertificatesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateCertificatesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateCertificatesV1ResponseValidationError{}

// Validate checks the field values on CreateCertificateV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateCertificateV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return CreateCertificateV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCertificateV1RequestValidationError{
				field:  "Created",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Link

	return nil
}

// CreateCertificateV1RequestValidationError is the validation error returned
// by CreateCertificateV1Request.Validate if the designated constraints aren't met.
type CreateCertificateV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCertificateV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCertificateV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCertificateV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCertificateV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCertificateV1RequestValidationError) ErrorName() string {
	return "CreateCertificateV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCertificateV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCertificateV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCertificateV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCertificateV1RequestValidationError{}

// Validate checks the field values on CreateCertificateV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateCertificateV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CertificateId

	return nil
}

// CreateCertificateV1ResponseValidationError is the validation error returned
// by CreateCertificateV1Response.Validate if the designated constraints
// aren't met.
type CreateCertificateV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCertificateV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCertificateV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCertificateV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCertificateV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCertificateV1ResponseValidationError) ErrorName() string {
	return "CreateCertificateV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCertificateV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCertificateV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCertificateV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCertificateV1ResponseValidationError{}

// Validate checks the field values on GetCertificateV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetCertificateV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetCertificateId() <= 0 {
		return GetCertificateV1RequestValidationError{
			field:  "CertificateId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// GetCertificateV1RequestValidationError is the validation error returned by
// GetCertificateV1Request.Validate if the designated constraints aren't met.
type GetCertificateV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCertificateV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCertificateV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCertificateV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCertificateV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCertificateV1RequestValidationError) ErrorName() string {
	return "GetCertificateV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCertificateV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCertificateV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCertificateV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCertificateV1RequestValidationError{}

// Validate checks the field values on GetCertificateV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetCertificateV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return GetCertificateV1ResponseValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if m.GetUserId() <= 0 {
		return GetCertificateV1ResponseValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCertificateV1ResponseValidationError{
				field:  "Created",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Link

	return nil
}

// GetCertificateV1ResponseValidationError is the validation error returned by
// GetCertificateV1Response.Validate if the designated constraints aren't met.
type GetCertificateV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCertificateV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCertificateV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCertificateV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCertificateV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCertificateV1ResponseValidationError) ErrorName() string {
	return "GetCertificateV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCertificateV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCertificateV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCertificateV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCertificateV1ResponseValidationError{}

// Validate checks the field values on ListCertificateV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListCertificateV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimit() <= 0 {
		return ListCertificateV1RequestValidationError{
			field:  "Limit",
			reason: "value must be greater than 0",
		}
	}

	if m.GetOffset() <= 0 {
		return ListCertificateV1RequestValidationError{
			field:  "Offset",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// ListCertificateV1RequestValidationError is the validation error returned by
// ListCertificateV1Request.Validate if the designated constraints aren't met.
type ListCertificateV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCertificateV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCertificateV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCertificateV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCertificateV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCertificateV1RequestValidationError) ErrorName() string {
	return "ListCertificateV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCertificateV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCertificateV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCertificateV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCertificateV1RequestValidationError{}

// Validate checks the field values on ListCertificateV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListCertificateV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetCertificates() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCertificateV1ResponseValidationError{
					field:  fmt.Sprintf("Certificates[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListCertificateV1ResponseValidationError is the validation error returned by
// ListCertificateV1Response.Validate if the designated constraints aren't met.
type ListCertificateV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCertificateV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCertificateV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCertificateV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCertificateV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCertificateV1ResponseValidationError) ErrorName() string {
	return "ListCertificateV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListCertificateV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCertificateV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCertificateV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCertificateV1ResponseValidationError{}

// Validate checks the field values on UpdateCertificateV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateCertificateV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return UpdateCertificateV1RequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if m.GetUserId() <= 0 {
		return UpdateCertificateV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateCertificateV1RequestValidationError{
				field:  "Created",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Link

	return nil
}

// UpdateCertificateV1RequestValidationError is the validation error returned
// by UpdateCertificateV1Request.Validate if the designated constraints aren't met.
type UpdateCertificateV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCertificateV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCertificateV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCertificateV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCertificateV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCertificateV1RequestValidationError) ErrorName() string {
	return "UpdateCertificateV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCertificateV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCertificateV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCertificateV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCertificateV1RequestValidationError{}

// Validate checks the field values on UpdateCertificateV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateCertificateV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Updated

	return nil
}

// UpdateCertificateV1ResponseValidationError is the validation error returned
// by UpdateCertificateV1Response.Validate if the designated constraints
// aren't met.
type UpdateCertificateV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCertificateV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCertificateV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCertificateV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCertificateV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCertificateV1ResponseValidationError) ErrorName() string {
	return "UpdateCertificateV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCertificateV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCertificateV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCertificateV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCertificateV1ResponseValidationError{}

// Validate checks the field values on RemoveCertificateV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveCertificateV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetCertificateId() <= 0 {
		return RemoveCertificateV1RequestValidationError{
			field:  "CertificateId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveCertificateV1RequestValidationError is the validation error returned
// by RemoveCertificateV1Request.Validate if the designated constraints aren't met.
type RemoveCertificateV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveCertificateV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveCertificateV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveCertificateV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveCertificateV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveCertificateV1RequestValidationError) ErrorName() string {
	return "RemoveCertificateV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveCertificateV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveCertificateV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveCertificateV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveCertificateV1RequestValidationError{}

// Validate checks the field values on RemoveCertificateV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveCertificateV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Removed

	return nil
}

// RemoveCertificateV1ResponseValidationError is the validation error returned
// by RemoveCertificateV1Response.Validate if the designated constraints
// aren't met.
type RemoveCertificateV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveCertificateV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveCertificateV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveCertificateV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveCertificateV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveCertificateV1ResponseValidationError) ErrorName() string {
	return "RemoveCertificateV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveCertificateV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveCertificateV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveCertificateV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveCertificateV1ResponseValidationError{}
