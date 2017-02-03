package problems

import (
	"errors"
	"net/http"
	"net/url"
	"testing"
)

func TestDefaultProblem(t *testing.T) {
	problem := &DefaultProblem{Type: DefaultURL,
		Status: 401,
		Title:  http.StatusText(http.StatusUnauthorized),
		Detail: "You dun did somethin' wrong.",
	}

	typ, err := problem.ProblemType()
	if err != nil {
		t.Errorf("Unable to read problem type")
	}
	if typ.String() != problem.Type {
		t.Errorf("Problem Type's did not match")
	}

	if problem.ProblemTitle() != problem.Title {
		t.Errorf("Problem Title's did not match")
	}

	err = ValidateProblem(problem)
	if err != nil {
		t.Errorf("problem is not valid")
	}
}

type badProblemType struct{}

func (p badProblemType) ProblemType() (*url.URL, error) {
	return nil, errors.New("I'm a bad problem type")
}

func (p badProblemType) ProblemTitle() string {
	return "something valid"
}

func (p badProblemType) Error() string {
	return ""
}

type badProblemTitle struct{}

func (p badProblemTitle) ProblemType() (*url.URL, error) {
	return &url.URL{}, nil
}

func (p badProblemTitle) ProblemTitle() string {
	return ""
}

func (p badProblemTitle) Error() string {
	return ""
}

func TestValidateProblem(t *testing.T) {
	var err error
	err = ValidateProblem(badProblemType{})
	if err == nil {
		t.Error("Only valid URI's should be allowed as problem types")
	}

	err = ValidateProblem(badProblemTitle{})
	if err == nil {
		t.Errorf("Empty strings should not be allowed as problem titles")
	}

	badURI := "::/"
	err = ValidateProblem(&DefaultProblem{Type: badURI})
	if err == nil {
		t.Errorf("%q was allowed as a valid URI", badURI)
	}
}

type creditProblem struct {
	DefaultProblem

	Balance  float64  `json:"balance"`
	Accounts []string `json:"accounts"`
}

func (cp *creditProblem) ProblemType() (*url.URL, error) {
	u, err := url.Parse(cp.Type)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (cp *creditProblem) ProblemTitle() string {
	return cp.Title
}

func TestCreditProblem(t *testing.T) {
	problem := &creditProblem{
		DefaultProblem: DefaultProblem{Type: DefaultURL,
			Status: 401,
			Title:  http.StatusText(http.StatusUnauthorized),
			Detail: "You dun did somethin' wrong.",
		},
		Balance:  30,
		Accounts: []string{"/account/12345", "/account/67890"},
	}

	typ, err := problem.ProblemType()
	if err != nil {
		t.Errorf("Unable to read problem type")
	}
	if typ.String() != problem.Type {
		t.Errorf("Problem Type's did not match")
	}

	if problem.ProblemTitle() != problem.Title {
		t.Errorf("Problem Title's did not match")
	}

	err = ValidateProblem(problem)
	if err != nil {
		t.Errorf("problem is not valid")
	}
}

func TestErrorInterface(t *testing.T) {
	var i interface{} = &DefaultProblem{Type: DefaultURL,
		Status: 401,
		Title:  http.StatusText(http.StatusUnauthorized),
		Detail: "Detail Message.",
	}
	p, ok := i.(error)
	if !ok {
		t.Errorf("DefaultProblem does not implement error interface")
	}

	expected := http.StatusText(http.StatusUnauthorized) + " (401) - Detail Message."
	if p.Error() != expected {
		t.Errorf("Error message was not what we were expecting. Expected '%s', got '%s'", expected, p.Error())
	}
}
