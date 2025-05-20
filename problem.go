package problems

import (
	"fmt"
	"net/http"
)

const (
	// ProblemMediaType is the default media type for a Problem response
	ProblemMediaType = "application/problem+json"

	// ProblemMediaTypeXML is the XML variant on the Problem Media type
	ProblemMediaTypeXML = "application/problem+xml"

	// DefaultURL is the default url to use for problem types
	DefaultURL = "about:blank"
)

// A Problem defines all the standard problem detail fields as defined by
// RFC-9457 and can easily be serialized to either JSON or XML.
//
// To add extensions to a Problem definition, see ExtendedProblem or consider
// embedding a Problem in your extension struct.
type Problem struct {
	// Type contains a URI that identifies the problem type. This URI will,
	// ideally, contain human-readable documentation for the problem when
	// de-referenced.
	Type string `json:"type" xml:"type"`

	// Title is a short, human-readable summary of the problem type. This title
	// SHOULD NOT change from occurrence to occurrence of the problem, except
	// for purposes of localization.
	Title string `json:"title" xml:"title"`

	// The HTTP status code for this occurrence of the problem.
	Status int `json:"status,omitempty" xml:"status,omitempty"`

	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`

	// A URI that identifies the specific occurrence of the problem. This URI
	// may or may not yield further information if de-referenced.
	Instance string `json:"instance,omitempty" xml:"instance,omitempty"`
}

// New returns a new Problem instance with the type field set to DefaultURL.
func New() *Problem {
	return &Problem{Type: DefaultURL}
}

// FromError returns a new Problem instance which contains the string version
// of the provided error as the details of the problem.
func FromError(err error) *Problem {
	return New().WithError(err)
}

// NewStatusProblem will generate a default problem for the provided HTTP status
// code. The Problem's Status field will be set to match the status argument,
// and the Title will be set to the default Go status text for that code.
func NewStatusProblem(status int) *Problem {
	return New().WithTitle(http.StatusText(status)).WithStatus(status)
}

// NewDetailedProblem returns a new Problem with a Detail string set for
// a more detailed explanation of the problem being returned.
func NewDetailedProblem(status int, details string) *Problem {
	return NewStatusProblem(status).WithDetail(details)
}

// WithType sets the type field to the provided string.
func (p *Problem) WithType(typ string) *Problem {
	p.Type = typ
	return p
}

// WithTitle sets the title field to the provided string.
func (p *Problem) WithTitle(title string) *Problem {
	p.Title = title
	return p
}

// WithStatus sets the status field to the provided int.
//
// If no title is set then this call will also set the title to the return
// value of http.StatusText for the provided status code.
func (p *Problem) WithStatus(status int) *Problem {
	p.Status = status
	if p.Title == "" {
		p.Title = http.StatusText(status)
	}

	return p
}

// WithDetail sets the detail message to the provided string.
func (p *Problem) WithDetail(detail string) *Problem {
	p.Detail = detail
	return p
}

// WithDetailf behaves identically to WithDetail, but allows consumers to
// provide a format string and arguments which will be formatted internally.
func (p *Problem) WithDetailf(format string, args ...interface{}) *Problem {
	p.Detail = fmt.Sprintf(format, args...)
	return p
}

// WithError sets the detail message to the provided error.
func (p *Problem) WithError(err error) *Problem {
	p.Detail = err.Error()
	return p
}

// WithInstance sets the instance uri to the provided string.
func (p *Problem) WithInstance(instance string) *Problem {
	p.Instance = instance
	return p
}

// Error implements the error interface and allows a Problem to be used as a
// native error.
func (p *Problem) Error() string {
	return fmt.Sprintf("%s (%d) - %s", p.Title, p.Status, p.Detail)
}
