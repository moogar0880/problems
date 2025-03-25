package problems

import "fmt"

const errPrefix = "problems"

// ErrTitleMustBeSet is the error returned from a call to ValidateProblem if
// the problem is validated without a title.
var ErrTitleMustBeSet = fmt.Errorf("%s: problem title must be set", errPrefix)

// ErrInvalidProblemType is the error type returned if a problems type is not a
// valid URI when it is validated. The inner Err will contain the error
// returned from attempting to parse the invalid URI.
type ErrInvalidProblemType struct {
	Err   error
	Value string
}

// NewErrInvalidProblemType returns a new ErrInvalidProblemType instance which
// wraps the provided error.
func NewErrInvalidProblemType(value string, e error) error {
	return &ErrInvalidProblemType{
		Err:   e,
		Value: value,
	}
}

func (e *ErrInvalidProblemType) Error() string {
	return fmt.Sprintf("%s: problem type must be a valid uri: %s", errPrefix, e.Err)
}
