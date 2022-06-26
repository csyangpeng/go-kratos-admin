// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/center/admin/v1/center_admin.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on LoginReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReqMultiError, or nil
// if none found.
func (m *LoginReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUsername()) < 4 {
		err := LoginReqValidationError{
			field:  "Username",
			reason: "value length must be at least 4 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 5 {
		err := LoginReqValidationError{
			field:  "Password",
			reason: "value length must be at least 5 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LoginReqMultiError(errors)
	}

	return nil
}

// LoginReqMultiError is an error wrapping multiple validation errors returned
// by LoginReq.ValidateAll() if the designated constraints aren't met.
type LoginReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReqMultiError) AllErrors() []error { return m }

// LoginReqValidationError is the validation error returned by
// LoginReq.Validate if the designated constraints aren't met.
type LoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReqValidationError) ErrorName() string { return "LoginReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReqValidationError{}

// Validate checks the field values on LoginReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReplyMultiError, or
// nil if none found.
func (m *LoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return LoginReplyMultiError(errors)
	}

	return nil
}

// LoginReplyMultiError is an error wrapping multiple validation errors
// returned by LoginReply.ValidateAll() if the designated constraints aren't met.
type LoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReplyMultiError) AllErrors() []error { return m }

// LoginReplyValidationError is the validation error returned by
// LoginReply.Validate if the designated constraints aren't met.
type LoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReplyValidationError) ErrorName() string { return "LoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReplyValidationError{}

// Validate checks the field values on LogoutReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LogoutReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogoutReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LogoutReqMultiError, or nil
// if none found.
func (m *LogoutReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LogoutReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LogoutReqMultiError(errors)
	}

	return nil
}

// LogoutReqMultiError is an error wrapping multiple validation errors returned
// by LogoutReq.ValidateAll() if the designated constraints aren't met.
type LogoutReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogoutReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogoutReqMultiError) AllErrors() []error { return m }

// LogoutReqValidationError is the validation error returned by
// LogoutReq.Validate if the designated constraints aren't met.
type LogoutReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogoutReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogoutReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogoutReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogoutReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogoutReqValidationError) ErrorName() string { return "LogoutReqValidationError" }

// Error satisfies the builtin error interface
func (e LogoutReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogoutReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogoutReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogoutReqValidationError{}

// Validate checks the field values on LogoutReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LogoutReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogoutReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LogoutReplyMultiError, or
// nil if none found.
func (m *LogoutReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LogoutReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LogoutReplyMultiError(errors)
	}

	return nil
}

// LogoutReplyMultiError is an error wrapping multiple validation errors
// returned by LogoutReply.ValidateAll() if the designated constraints aren't met.
type LogoutReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogoutReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogoutReplyMultiError) AllErrors() []error { return m }

// LogoutReplyValidationError is the validation error returned by
// LogoutReply.Validate if the designated constraints aren't met.
type LogoutReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogoutReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogoutReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogoutReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogoutReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogoutReplyValidationError) ErrorName() string { return "LogoutReplyValidationError" }

// Error satisfies the builtin error interface
func (e LogoutReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogoutReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogoutReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogoutReplyValidationError{}

// Validate checks the field values on GetUserReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserReqMultiError, or
// nil if none found.
func (m *GetUserReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetUserReqMultiError(errors)
	}

	return nil
}

// GetUserReqMultiError is an error wrapping multiple validation errors
// returned by GetUserReq.ValidateAll() if the designated constraints aren't met.
type GetUserReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserReqMultiError) AllErrors() []error { return m }

// GetUserReqValidationError is the validation error returned by
// GetUserReq.Validate if the designated constraints aren't met.
type GetUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReqValidationError) ErrorName() string { return "GetUserReqValidationError" }

// Error satisfies the builtin error interface
func (e GetUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReqValidationError{}

// Validate checks the field values on GetUserReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserReplyMultiError, or
// nil if none found.
func (m *GetUserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetUsername()) < 4 {
		err := GetUserReplyValidationError{
			field:  "Username",
			reason: "value length must be at least 4 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetUserReplyMultiError(errors)
	}

	return nil
}

// GetUserReplyMultiError is an error wrapping multiple validation errors
// returned by GetUserReply.ValidateAll() if the designated constraints aren't met.
type GetUserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserReplyMultiError) AllErrors() []error { return m }

// GetUserReplyValidationError is the validation error returned by
// GetUserReply.Validate if the designated constraints aren't met.
type GetUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReplyValidationError) ErrorName() string { return "GetUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReplyValidationError{}

// Validate checks the field values on ListUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListUserReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListUserReqMultiError, or
// nil if none found.
func (m *ListUserReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageIndex

	// no validation rules for PageSize

	if len(errors) > 0 {
		return ListUserReqMultiError(errors)
	}

	return nil
}

// ListUserReqMultiError is an error wrapping multiple validation errors
// returned by ListUserReq.ValidateAll() if the designated constraints aren't met.
type ListUserReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserReqMultiError) AllErrors() []error { return m }

// ListUserReqValidationError is the validation error returned by
// ListUserReq.Validate if the designated constraints aren't met.
type ListUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserReqValidationError) ErrorName() string { return "ListUserReqValidationError" }

// Error satisfies the builtin error interface
func (e ListUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserReqValidationError{}

// Validate checks the field values on ListUserReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListUserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListUserReplyMultiError, or
// nil if none found.
func (m *ListUserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUserReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUserReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUserReplyValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListUserReplyMultiError(errors)
	}

	return nil
}

// ListUserReplyMultiError is an error wrapping multiple validation errors
// returned by ListUserReply.ValidateAll() if the designated constraints
// aren't met.
type ListUserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserReplyMultiError) AllErrors() []error { return m }

// ListUserReplyValidationError is the validation error returned by
// ListUserReply.Validate if the designated constraints aren't met.
type ListUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserReplyValidationError) ErrorName() string { return "ListUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserReplyValidationError{}

// Validate checks the field values on ListUserReply_User with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListUserReply_User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserReply_User with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserReply_UserMultiError, or nil if none found.
func (m *ListUserReply_User) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserReply_User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Username

	if len(errors) > 0 {
		return ListUserReply_UserMultiError(errors)
	}

	return nil
}

// ListUserReply_UserMultiError is an error wrapping multiple validation errors
// returned by ListUserReply_User.ValidateAll() if the designated constraints
// aren't met.
type ListUserReply_UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserReply_UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserReply_UserMultiError) AllErrors() []error { return m }

// ListUserReply_UserValidationError is the validation error returned by
// ListUserReply_User.Validate if the designated constraints aren't met.
type ListUserReply_UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserReply_UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserReply_UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserReply_UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserReply_UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserReply_UserValidationError) ErrorName() string {
	return "ListUserReply_UserValidationError"
}

// Error satisfies the builtin error interface
func (e ListUserReply_UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserReply_User.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserReply_UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserReply_UserValidationError{}