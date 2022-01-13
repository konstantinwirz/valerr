package valerr

import (
	"reflect"
	"testing"
)

func TestNewValidationError(t *testing.T) {
	tests := []struct {
		name  string
		first Violation
		rest  []Violation
		want  ValidationError
	}{
		{
			name:  "one argument",
			first: NewViolation("email", "invalid"),
			rest:  nil,
			want:  valerr{violations: []Violation{NewViolation("email", "invalid")}},
		},
		{
			name:  "two arguments",
			first: NewViolation("email", "invalid"),
			rest:  []Violation{NewViolation("password", "insecure")},
			want:  valerr{violations: []Violation{NewViolation("email", "invalid"), NewViolation("password", "insecure")}},
		},
		{
			name:  "three arguments",
			first: NewViolation("email", "invalid"),
			rest:  []Violation{NewViolation("password", "insecure"), NewViolation("name", "empty")},
			want: valerr{
				violations: []Violation{
					NewViolation("email", "invalid"),
					NewViolation("password", "insecure"),
					NewViolation("name", "empty")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewValidationError(tt.first, tt.rest...)
			compare(t, got, tt.want)
		})
	}
}

func TestAppend(t *testing.T) {
	tests := []struct {
		name   string
		actual ValidationError
		toAdd  Violation
		want   ValidationError
	}{
		{
			name:   "add to one violation",
			actual: NewValidationError(NewViolation("email", "empty")),
			toAdd:  NewViolation("password", "insecure"),
			want: NewValidationError(
				NewViolation("email", "empty"),
				NewViolation("password", "insecure"),
			),
		},
		{
			name:   "add to multiple violations",
			actual: NewValidationError(NewViolation("email", "empty"), NewViolation("password", "insecure")),
			toAdd:  NewViolation("name", "empty"),
			want: NewValidationError(
				NewViolation("email", "empty"),
				NewViolation("password", "insecure"),
				NewViolation("name", "empty"),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.actual.Append(tt.toAdd)
			compare(t, got, tt.want)
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name string
		err  ValidationError
		want string
	}{
		{
			name: "one violation",
			err:  NewValidationError(NewViolation("email", "empty")),
			want: "[ { 'email' : 'empty' } ]",
		},
		{
			name: "multiple violations",
			err:  NewValidationError(NewViolation("email", "empty"), NewViolation("password", "insecure")),
			want: "[ { 'email' : 'empty' }, { 'password' : 'insecure' } ]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.Error()
			if got != tt.want {
				t.Errorf("expected %q, got %q", tt.want, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		err  ValidationError
		want string
	}{
		{
			name: "one violation",
			err:  NewValidationError(NewViolation("email", "empty")),
			want: "ValidationError[ { 'email' : 'empty' } ]",
		},
		{
			name: "multiple violations",
			err:  NewValidationError(NewViolation("email", "empty"), NewViolation("password", "insecure")),
			want: "ValidationError[ { 'email' : 'empty' }, { 'password' : 'insecure' } ]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.String()
			if got != tt.want {
				t.Errorf("expected %q, got %q", tt.want, got)
			}
		})
	}
}

func hasViolation(err ValidationError, v Violation) bool {
	for _, vv := range err.Violations() {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

func compare(t *testing.T, got, want ValidationError) {
	gotLen := len(got.Violations())
	wantLen := len(want.Violations())
	if gotLen != wantLen {
		t.Errorf("want %d violations, got %d", wantLen, gotLen)
	}

	for _, v := range got.Violations() {
		if !hasViolation(want, v) {
			t.Errorf("want %v to be part of violations", v)
		}
	}
}
