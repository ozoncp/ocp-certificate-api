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

// Validate checks the field values on Certificate with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Certificate) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return CertificateValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if m.GetUserId() <= 0 {
		return CertificateValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateValidationError{
				field:  "Created",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Link

	return nil
}

// CertificateValidationError is the validation error returned by
// Certificate.Validate if the designated constraints aren't met.
type CertificateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateValidationError) ErrorName() string { return "CertificateValidationError" }

// Error satisfies the builtin error interface
func (e CertificateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateValidationError{}

// Validate checks the field values on CreateCertificateV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateCertificateV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCertificate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCertificateV1RequestValidationError{
				field:  "Certificate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

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

// Validate checks the field values on DescribeCertificateV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeCertificateV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetCertificateId() <= 0 {
		return DescribeCertificateV1RequestValidationError{
			field:  "CertificateId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeCertificateV1RequestValidationError is the validation error returned
// by DescribeCertificateV1Request.Validate if the designated constraints
// aren't met.
type DescribeCertificateV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeCertificateV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeCertificateV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeCertificateV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeCertificateV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeCertificateV1RequestValidationError) ErrorName() string {
	return "DescribeCertificateV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeCertificateV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeCertificateV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeCertificateV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeCertificateV1RequestValidationError{}

// Validate checks the field values on DescribeCertificateV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeCertificateV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DescribeCertificateV1ResponseValidationError is the validation error
// returned by DescribeCertificateV1Response.Validate if the designated
// constraints aren't met.
type DescribeCertificateV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeCertificateV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeCertificateV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeCertificateV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeCertificateV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeCertificateV1ResponseValidationError) ErrorName() string {
	return "DescribeCertificateV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeCertificateV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeCertificateV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeCertificateV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeCertificateV1ResponseValidationError{}

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

	if v, ok := interface{}(m.GetCertificate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateCertificateV1RequestValidationError{
				field:  "Certificate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

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
