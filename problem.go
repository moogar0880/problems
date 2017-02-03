package problems

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	// ProblemMediaType is the default media type for a Problem response
	ProblemMediaType = "application/problem+json"

	// ProblemMediaTypeXML is the XML variant on the Problem Media type
	ProblemMediaTypeXML = "application/problem+xml"

	// DefaultURL is the default url to use for problem types
	DefaultURL = "about:blank"
)

// Problem is the interface describing an HTTP API problem. These "problem
// details" are designed to encompass a way to carry machine- readable details
// of errors in a HTTP response to avoid the need to define new error response
// formats for HTTP APIs.
type Problem interface {
	ProblemType() (*url.URL, error)
	ProblemTitle() string
	Error() string
}

// StatusProblem is the interface describing a problem with an associated
// Status code
type StatusProblem interface {
	Problem
	ProblemStatus() int
}

// ValidateProblem ensures that the provided Problem implementation meets the
// Problem description requirements. Which means that the Type is a valid uri,
// and that the Title be a non-empty string. Should the provided Problem be in
// violation of either of these requirements, an error is returned
func ValidateProblem(p Problem) error {
	_, err := p.ProblemType()
	if err != nil {
		return errors.New("Problem Type's must be valid URIs")
	}

	title := p.ProblemTitle()
	if title == "" {
		return errors.New("Problem Title's must be set")
	}
	return nil
}

// DefaultProblem is a default problem implementation. The Problem specification
// allows for arbitrary extensions to include new fields, in which case a new
// Problem implementation will likely be required
type DefaultProblem struct {
	// Type contains a URI that identifies the problem type. This URI will,
	// ideally, contain human-readable documentation for the problem when
	// dereferenced
	Type string `json:"type"`

	// Title is a short, human-readable summary of the problem type. This title
	// SHOULD NOT change from occurrence to occurrence of the problem, except for
	// purposes of localization
	Title string `json:"title"`

	// The HTTP status code for this occurrence of the problem
	Status int `json:"status,omitempty"`

	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`

	// A URI that identifies the specific occurrence of the problem. This URI
	// may or may not yield further information if dereferenced.
	Instance string `json:"instance,omitempty"`
}

// NewProblem returns a new instance of a DefaultProblem
func NewProblem() *DefaultProblem {
	problem := &DefaultProblem{Type: DefaultURL, Title: ""}
	return problem
}

// NewStatusProblem will generate a default problem for the provided HTTP status
// code. The Problem's Status field will be set to match the status argument,
// and the Title will be set to the default Go status text for that code.
func NewStatusProblem(status int) *DefaultProblem {
	problem := &DefaultProblem{Type: DefaultURL,
		Title:  http.StatusText(int(status)),
		Status: status}
	return problem
}

// NewDetailedProblem returns a new DefaultProblem with a Detail string set for
// a more detailed explanation of the problem being returned
func NewDetailedProblem(status int, details string) *DefaultProblem {
	problem := NewStatusProblem(status)
	problem.Detail = details
	return problem
}

// ProblemType returns the uri for the problem type being defined and an
// optional error
func (p *DefaultProblem) ProblemType() (*url.URL, error) {
	u, err := url.Parse(p.Type)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// ProblemTitle returns the unique title field for this Problem
func (p *DefaultProblem) ProblemTitle() string {
	return p.Title
}

// ProblemStatus allows the DefaultStatusProblem to implement the StatusProblem
// interface, returning the Status code for this problem
func (p *DefaultProblem) ProblemStatus() int {
	return p.Status
}

// Implement Error() so it can satisfy the default error Interface
// It returns a string in the form of:
//     <Title> (Status) - <Detail>
func (p *DefaultProblem) Error() string {
	return fmt.Sprintf("%s (%d) - %s", p.Title, p.Status, p.Detail)
}
