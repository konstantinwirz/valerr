package valerr

import (
	"fmt"
	"strings"
)

type (
	// ValidationError bundles a non-empty list of violations
	ValidationError interface {
		fmt.Stringer
		error
		// Violations returns a non empty list of violations
		Violations() []Violation
		// Append creates and returns a new instance of ValidationErrors
		// containing new violation and all previously existed ones
		Append(Violation) ValidationError
	}

	// Violation holds field name and reason why a validation rule is violated
	Violation interface {
		fmt.Stringer
		// Field returns the name of the field which violates some validation rules
		Field() string
		// Reason returns the reason why a rule has been violated
		Reason() string
	}
)

type valerr struct {
	violations []Violation
}

// NewValidationError creates and returns an instance of ValidationError
// with given violations
func NewValidationError(v Violation, vs ...Violation) ValidationError {
	return valerr{
		violations: append([]Violation{v}, vs...),
	}
}

func (e valerr) String() string {
	var vs []string
	for _, v := range e.violations {
		vs = append(vs, v.String())
	}
	return fmt.Sprintf("ValidationError[ %s ]", strings.Join(vs, ", "))
}

func (e valerr) Append(v Violation) ValidationError {
	return &valerr{
		violations: append(e.violations, v),
	}
}

func (e valerr) Error() string {
	var vs []string
	for _, v := range e.violations {
		vs = append(vs, fmt.Sprintf("{ '%s' : '%s' }", v.Field(), v.Reason()))
	}

	return fmt.Sprintf("[ %s ]", strings.Join(vs, ", "))
}

func (e valerr) Violations() []Violation {
	return e.violations
}

type violation struct {
	field, reason string
}

// NewViolation creates and returns a new violation
func NewViolation(field, reason string) Violation {
	return &violation{
		field:  field,
		reason: reason,
	}
}

func (v violation) Field() string {
	return v.field
}

func (v violation) Reason() string {
	return v.reason
}

func (v violation) String() string {
	return fmt.Sprintf("{ '%s' : '%s' }", v.field, v.reason)
}
