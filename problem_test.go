package problems

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		name    string
		problem *Problem
		expect  Problem
	}{
		{
			name:    "should support problem with only title and status",
			problem: NewStatusProblem(http.StatusNotFound),
			expect: Problem{
				Type:   DefaultURL,
				Title:  http.StatusText(http.StatusNotFound),
				Status: http.StatusNotFound,
			},
		},
		{
			name:    "should support problem with title, status, and details",
			problem: NewDetailedProblem(http.StatusNotFound, "couldn't find it"),
			expect: Problem{
				Type:   DefaultURL,
				Title:  http.StatusText(http.StatusNotFound),
				Status: http.StatusNotFound,
				Detail: "couldn't find it",
			},
		},
		{
			name: "should default title to match status code",
			problem: New().
				WithType("https://example.com").
				WithStatus(http.StatusBadRequest).
				WithDetail("Here are some details").
				WithInstance("https://example.com/errors/150"),
			expect: Problem{
				Type:     "https://example.com",
				Title:    "Bad Request",
				Status:   http.StatusBadRequest,
				Detail:   "Here are some details",
				Instance: "https://example.com/errors/150",
			},
		},
		{
			name: "should maintain custom title when setting status code",
			problem: New().
				WithType("https://example.com").
				WithTitle("This is an example").
				WithStatus(http.StatusBadRequest).
				WithDetail("Here are some details").
				WithInstance("https://example.com/errors/150"),
			expect: Problem{
				Type:     "https://example.com",
				Title:    "This is an example",
				Status:   http.StatusBadRequest,
				Detail:   "Here are some details",
				Instance: "https://example.com/errors/150",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, err := json.Marshal(test.problem)
			if err != nil {
				t.Errorf("Error marshalling problem data: %s", err)
			}

			var problem Problem
			if err = json.Unmarshal(data, &problem); err != nil {
				t.Errorf("Error unmarshalling problem data: %s", err)
			}

			if !reflect.DeepEqual(problem, test.expect) {
				t.Errorf("problems were not equal: wanted\n%#+v\n but got\n%#+v", test.expect, problem)
			}
		})
	}
}

// func TestDefaultProblem(t *testing.T) {
// 	problem := NewDetailedProblem(http.StatusUnauthorized, unAuthDetails)
//
// 	if _, err := problem.Validate(); err != nil {
// 		t.Errorf("problem is not valid")
// 	}
// }

//	func (cp *creditProblem) ProblemType() (*url.URL, error) {
//		u, err := url.Parse(cp.Type)
//		if err != nil {
//			return nil, err
//		}
//		return u, nil
//	}
//
//	func (cp *creditProblem) ProblemTitle() string {
//		return cp.Title
//	}
// var unAuthDetails = "you are unauthorized to access this resource"

// func TestCreditProblem(t *testing.T) {
// 	// problem := &creditProblem{
// 	// 	DefaultProblem: *NewDetailedProblem(http.StatusUnauthorized, unAuthDetails),
// 	// 	Balance:        30,
// 	// 	Accounts:       []string{"/account/12345", "/account/67890"},
// 	// }
//
// 	problem := NewExt[creditProblemExt]().
// 		WithStatus(http.StatusUnauthorized).
// 		WithDetail(unAuthDetails).
// 		WithExtension(creditProblemExt{
// 			Balance:  30,
// 			Accounts: []string{"/account/12345", "/account/67890"},
// 		})
//
// 	// typ, err := problem.ProblemType()
// 	// if err != nil {
// 	// 	t.Errorf("Unable to read problem type")
// 	// }
// 	// if typ != nil && typ.String() != problem.Type {
// 	// 	t.Errorf("Problem Types did not match")
// 	// }
// 	//
// 	// if problem.ProblemTitle() != problem.Title {
// 	// 	t.Errorf("Problem Titles did not match")
// 	// }
// 	//
// 	// err = ValidateProblem(problem)
// 	if _, err := problem.Validate(); err != nil {
// 		t.Errorf("problem is not valid: %s", err)
// 	}
// }
