// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/datastore/model/cluster.proto

package model

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

// Validate checks the field values on Cluster with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// ClusterMultiError, or nil if none found. Otherwise, only the first error is
// returned, if any.
func (m *Cluster) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := ClusterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Kubeconfig

	if m.GetUpdatedAt() <= 0 {
		err := ClusterValidationError{
			field:  "UpdatedAt",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ClusterMultiError(errors)
	}
	return nil
}

// ClusterMultiError is an error wrapping multiple validation errors returned
// by Cluster.Validate(true) if the designated constraints aren't met.
type ClusterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClusterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClusterMultiError) AllErrors() []error { return m }

// ClusterValidationError is the validation error returned by Cluster.Validate
// if the designated constraints aren't met.
type ClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterValidationError) ErrorName() string { return "ClusterValidationError" }

// Error satisfies the builtin error interface
func (e ClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterValidationError{}
