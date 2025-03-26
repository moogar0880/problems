package problems

import (
	"encoding/json"
	"net/url"
)

// A ValidProblem is a sealed variant of Problem which is guaranteed to have
// valid fields.
//
// Instances of ValidProblem can be created by using the Validate method from
// the Problem type.
type ValidProblem struct {
	// Type contains a URI that identifies the problem type. This URI will,
	// ideally, contain human-readable documentation for the problem when
	// de-referenced.
	typ *url.URL

	// Title is a short, human-readable summary of the problem type. This title
	// SHOULD NOT change from occurrence to occurrence of the problem, except
	// for purposes of localization.
	title string

	// The HTTP status code for this occurrence of the problem.
	status int

	// A human-readable explanation specific to this occurrence of the problem.
	detail string

	// A URI that identifies the specific occurrence of the problem. This URI
	// may or may not yield further information if de-referenced.
	instance string
}

// Validate validates the content of the Problem instance. If the Problem is
// invalid, as defined by RFC-9457, then an error explaining the validation
// error is returned. Otherwise, a sealed ValidProblem instance is returned.
//
// See the documentation for ErrTitleMustBeSet and ErrInvalidProblemType for
// more information on the validation errors returned by this method.
func (p *Problem) Validate() (*ValidProblem, error) {
	typ, err := validate(p.Type, p.Title)
	if err != nil {
		return nil, err
	}

	return &ValidProblem{
		typ:      typ,
		title:    p.Title,
		status:   p.Status,
		detail:   p.Detail,
		instance: p.Instance,
	}, nil
}

// IntoProblem allows you to convert from a ValidProblem back into a Problem.
func (p *ValidProblem) IntoProblem() *Problem {
	return &Problem{
		Type:     p.typ.String(),
		Title:    p.title,
		Status:   p.status,
		Detail:   p.detail,
		Instance: p.instance,
	}
}

// MarshalJSON implements the json.Marshaler interface and ensures that a
// ValidProblem is properly serialized into JSON.
func (p *ValidProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.IntoProblem())
}

// A ValidExtendedProblem is a sealed variant of ExtendedProblem which is
// guaranteed to contain valid fields.
//
// Instances of ValidExtendedProblem can be created by using the Validate
// method on the ExtendedProblem type.
type ValidExtendedProblem[T any] struct {
	ValidProblem

	// Extensions allows for Problem type definitions to extend the standard
	// problem details object with additional members that are specific to that
	// problem type.
	extensions T
}

// Validate validates the content of the ExtendedProblem instance. If the
// ExtendedProblem is invalid, as defined by RFC-9457, then an error explaining
// the validation error is returned. Otherwise, a sealed ValidProblem instance
// is returned.
//
// See the documentation for ErrTitleMustBeSet and ErrInvalidProblemType for
// more information on the validation errors returned by this method.
func (p *ExtendedProblem[T]) Validate() (*ValidExtendedProblem[T], error) {
	typ, err := validate(p.Type, p.Title)
	if err != nil {
		return nil, err
	}

	return &ValidExtendedProblem[T]{
		ValidProblem: ValidProblem{
			typ:      typ,
			title:    p.Title,
			status:   p.Status,
			detail:   p.Detail,
			instance: p.Instance,
		},
		extensions: p.Extensions,
	}, nil
}

// IntoProblem allows you to convert from a ValidExtendedProblem back into a
// Problem.
func (p *ValidExtendedProblem[T]) IntoProblem() *Problem {
	return &Problem{
		Type:     p.typ.String(),
		Title:    p.title,
		Status:   p.status,
		Detail:   p.detail,
		Instance: p.instance,
	}
}

// IntoExtendedProblem allows you to convert from a ValidExtendedProblem back
// into an ExtendedProblem.
func (p *ValidExtendedProblem[T]) IntoExtendedProblem() *ExtendedProblem[T] {
	return &ExtendedProblem[T]{
		Problem:    *p.IntoProblem(),
		Extensions: p.extensions,
	}
}

// MarshalJSON implements the json.Marshaler interface and ensures that a
// ValidExtendedProblem is properly serialized into JSON.
func (p *ValidExtendedProblem[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.IntoExtendedProblem())
}

func validate(typ, title string) (*url.URL, error) {
	if len(title) == 0 {
		return nil, ErrTitleMustBeSet
	}

	typURL, err := url.Parse(typ)
	if err != nil {
		return nil, NewErrInvalidProblemType(typ, err)
	}

	return typURL, nil
}
