package problems

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// An ExtendedProblem extends the Problem type with a new field, Extensions,
// of type T.
type ExtendedProblem[T any] struct {
	Problem

	// Extensions allows for Problem type definitions to extend the standard
	// problem details object with additional members that are specific to that
	// problem type.
	Extensions T `json:"extensions,omitempty" xml:"extensions,omitempty"`
}

// NewExt returns a new ExtendedProblem with all the same default values
// as applied by a call to New.
func NewExt[T any]() *ExtendedProblem[T] {
	return &ExtendedProblem[T]{
		Problem: *New(),
	}
}

// ExtFromError returns a new ExtendedProblem instance which contains the
// string version of the provided error as the details of the problem.
func ExtFromError[T any](err error) *ExtendedProblem[T] {
	return NewExt[T]().WithError(err)
}

// Extend allows you to convert a standard Problem instance to an
// ExtendedProblem with the provided extension data.
func Extend[T any](p *Problem, ext T) *ExtendedProblem[T] {
	return &ExtendedProblem[T]{
		Problem:    *p,
		Extensions: ext,
	}
}

// WithType sets the type field to the provided string.
func (p *ExtendedProblem[T]) WithType(typ string) *ExtendedProblem[T] {
	p.Type = typ
	return p
}

// WithTitle sets the title field to the provided string.
func (p *ExtendedProblem[T]) WithTitle(title string) *ExtendedProblem[T] {
	p.Title = title
	return p
}

// WithStatus sets the status field to the provided int.
//
// If no title is set then this call will also set the title to the return
// value of http.StatusText for the provided status code.
func (p *ExtendedProblem[T]) WithStatus(status int) *ExtendedProblem[T] {
	p.Status = status
	if p.Title == "" {
		p.Title = http.StatusText(status)
	}
	return p
}

// WithDetail sets the detail message to the provided string.
func (p *ExtendedProblem[T]) WithDetail(detail string) *ExtendedProblem[T] {
	p.Detail = detail
	return p
}

// WithDetailf behaves identically to WithDetail, but allows consumers to
// provide a format string and arguments which will be formatted internally.
func (p *ExtendedProblem[T]) WithDetailf(format string, args ...interface{}) *ExtendedProblem[T] {
	p.Detail = fmt.Sprintf(format, args...)
	return p
}

// WithError sets the detail message to the provided error.
func (p *ExtendedProblem[T]) WithError(err error) *ExtendedProblem[T] {
	p.Detail = err.Error()
	return p
}

// WithInstance sets the instance uri to the provided string.
func (p *ExtendedProblem[T]) WithInstance(instance string) *ExtendedProblem[T] {
	p.Instance = instance
	return p
}

// WithExtension sets the extensions value to the provided extension of type T.
func (p *ExtendedProblem[T]) WithExtension(ext T) *ExtendedProblem[T] {
	p.Extensions = ext
	return p
}

// Error implements the error interface and allows a Problem to be used as a
// native error.
func (p *ExtendedProblem[T]) Error() string {
	ext, _ := json.Marshal(p.Extensions)
	return fmt.Sprintf("%s (%d) - %s - %s", p.Title, p.Status, p.Detail, string(ext))
}
