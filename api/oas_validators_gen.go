// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *User) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Group == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "group",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *UserCreateRequest) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Group == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "group",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *UserList) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Items == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.Items {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "items",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}