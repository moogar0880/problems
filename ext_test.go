package problems

import (
	"encoding/json"
	"net/http"
	"testing"
)

type creditProblemExt struct {
	Balance  float64  `json:"balance"`
	Accounts []string `json:"accounts"`
}

type creditProblem struct {
	Problem

	Balance  float64  `json:"balance"`
	Accounts []string `json:"accounts"`
}

var unAuthDetails = "you are unauthorized to access this resource"

func TestExtend(t *testing.T) {
	problem := Extend[creditProblemExt](
		New().
			WithStatus(http.StatusUnauthorized).
			WithDetail(unAuthDetails),
		creditProblemExt{
			Balance:  30,
			Accounts: []string{"/account/12345", "/account/67890"},
		},
	)

	if _, err := problem.Validate(); err != nil {
		t.Errorf("problem is not valid: %s", err)
	}

	data, err := json.Marshal(problem)
	if err != nil {
		t.Errorf("failed to marshal extended problem as json: %s", err)
	}

	expect := `{"type":"about:blank","title":"Unauthorized","status":401,"detail":"you are unauthorized to access this resource","extensions":{"balance":30,"accounts":["/account/12345","/account/67890"]}}`

	if string(data) != expect {
		t.Errorf("extended problem does not match expectation:\ngot\n%s\nwant\n%s", string(data), expect)
	}
}

func TestExtensionViaEmbedding(t *testing.T) {
	problem := &creditProblem{
		Problem: *New().
			WithStatus(http.StatusUnauthorized).
			WithDetail(unAuthDetails),
		Balance:  30,
		Accounts: []string{"/account/12345", "/account/67890"},
	}

	if _, err := problem.Validate(); err != nil {
		t.Errorf("problem is not valid: %s", err)
	}

	data, err := json.Marshal(problem)
	if err != nil {
		t.Errorf("failed to marshal extended problem as json: %s", err)
	}

	expect := `{"type":"about:blank","title":"Unauthorized","status":401,"detail":"you are unauthorized to access this resource","balance":30,"accounts":["/account/12345","/account/67890"]}`

	if string(data) != expect {
		t.Errorf("extended problem does not match expectation:\ngot\n%s\nwant\n%s", string(data), expect)
	}
}

func TestExtension(t *testing.T) {
	problem := NewExt[creditProblemExt]().
		WithType("https://example.com").
		WithTitle("This is a custom title").
		WithStatus(http.StatusBadRequest).
		WithDetail("Here are some details").
		WithInstance("https://example.com/errors/150")

	if _, err := problem.Validate(); err != nil {
		t.Errorf("extended problem is not valid but should be: %s", err)
	}
}
