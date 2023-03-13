package cqs

import "strings"

var _ error = &MultiError{}

const (
	multiErrorPrefix    = "multi error: "
	multiErrorSeparator = "; "
)

// NewMultiError is a constructor.
func NewMultiError() *MultiError {
	return &MultiError{
		errors: make([]error, 0),
	}
}

// MultiError is self described.
type MultiError struct {
	errors []error
}

// Error implements the Error interface.
func (e MultiError) Error() string {
	if len(e.errors) == 0 {
		return ""
	}

	errorMsgs := make([]string, len(e.errors))
	for i, e := range e.errors {
		errorMsgs[i] = e.Error()
	}

	return multiErrorPrefix + strings.Join(errorMsgs, multiErrorSeparator)
}

// ErrResult returns nil if not errors have been added.
func (e MultiError) ErrResult() error {
	if len(e.errors) == 0 {
		return nil
	}

	return &e
}

// Add adds a new error.
func (e *MultiError) Add(err error) {
	e.errors = append(e.errors, err)
}
