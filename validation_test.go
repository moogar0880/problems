package problems

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestExtendedProblem_Validate(t *testing.T) {
	tests := []struct {
		name    string
		problem ExtendedProblem[creditProblemExt]
	}{
		{
			name: "should fail to parse invalid problem type",
			problem: ExtendedProblem[creditProblemExt]{
				Problem: Problem{
					Type:  "::/",
					Title: http.StatusText(http.StatusBadRequest),
				},
				Extensions: creditProblemExt{
					Balance:  30,
					Accounts: []string{"/account/12345", "/account/67890"},
				},
			},
		},
		{
			name: "should fail to validate an empty problem title",
			problem: ExtendedProblem[creditProblemExt]{
				Problem: Problem{
					Type: DefaultURL,
				},
				Extensions: creditProblemExt{
					Balance:  30,
					Accounts: []string{"/account/12345", "/account/67890"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.problem.Type, func(t *testing.T) {
			_, err := test.problem.Validate()
			if err == nil {
				t.Errorf("problem is not valid but passed validation")
			}

			if len(err.Error()) == 0 {
				t.Errorf("problem is invalid and no error message was present")
			}
		})
	}
}

func TestProblem_Validate(t *testing.T) {
	tests := []struct {
		name    string
		problem Problem
	}{
		{
			name: "should fail to parse invalid problem type",
			problem: Problem{
				Type:  "::/",
				Title: http.StatusText(http.StatusBadRequest),
			},
		},
		{
			name: "should fail to validate an empty problem title",
			problem: Problem{
				Type: DefaultURL,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.problem.Type, func(t *testing.T) {
			_, err := test.problem.Validate()
			if err == nil {
				t.Errorf("problem is not valid but passed validation")
			}

			if len(err.Error()) == 0 {
				t.Errorf("problem is invalid and no error message was present")
			}
		})
	}
}

func TestValidExtendedProblem_MarshalJSON(t *testing.T) {
	problem := NewExt[creditProblemExt]().
		WithStatus(http.StatusUnauthorized).
		WithDetail(unAuthDetails).
		WithExtension(creditProblemExt{
			Balance:  30,
			Accounts: []string{"/account/12345", "/account/67890"},
		})

	valid, err := problem.Validate()
	if err != nil {
		t.Errorf("problem is not valid: %s", err)
	}

	expect, err := json.Marshal(problem)
	if err != nil {
		t.Errorf("failed to marshal extended problem as json: %s", err)
	}

	validated, err := json.Marshal(valid)
	if err != nil {
		t.Errorf("failed to marshal valid extended problem as json: %s", err)
	}

	if !bytes.Equal(expect, validated) {
		t.Errorf("extended problem does not match validated")
	}
}

func TestValidProblem_MarshalJSON(t *testing.T) {
	problem := New().
		WithStatus(http.StatusUnauthorized).
		WithDetail(unAuthDetails)

	valid, err := problem.Validate()
	if err != nil {
		t.Errorf("problem is not valid: %s", err)
	}

	expect, err := json.Marshal(problem)
	if err != nil {
		t.Errorf("failed to marshal problem as json: %s", err)
	}

	validated, err := json.Marshal(valid)
	if err != nil {
		t.Errorf("failed to marshal valid problem as json: %s", err)
	}

	if !bytes.Equal(expect, validated) {
		t.Errorf("problem does not match validated")
	}
}
