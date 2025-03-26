package problems

import (
	"encoding/json"
	"errors"
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
		{
			name: "should format detail message",
			problem: New().
				WithType("https://example.com").
				WithStatus(http.StatusBadRequest).
				WithDetailf("%q is not a valid integer", "foo").
				WithInstance("https://example.com/errors/150"),
			expect: Problem{
				Type:     "https://example.com",
				Title:    "Bad Request",
				Status:   http.StatusBadRequest,
				Detail:   `"foo" is not a valid integer`,
				Instance: "https://example.com/errors/150",
			},
		},
		{
			name: "should use error for detail message",
			problem: FromError(errors.New("an error occurred")).
				WithType("https://example.com").
				WithStatus(http.StatusBadRequest).
				WithInstance("https://example.com/errors/150"),
			expect: Problem{
				Type:     "https://example.com",
				Title:    "Bad Request",
				Status:   http.StatusBadRequest,
				Detail:   "an error occurred",
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

			if len(problem.Error()) == 0 {
				t.Error("error message should be set but was empty")
			}
		})
	}
}
