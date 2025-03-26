package problems_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/moogar0880/problems"
)

func ExampleNewStatusProblem() {
	notFound := problems.NewStatusProblem(404)
	b, _ := json.MarshalIndent(notFound, "", "  ")
	fmt.Println(string(b))
	// Output: {
	//   "type": "about:blank",
	//   "title": "Not Found",
	//   "status": 404
	// }
}

func ExampleNewStatusProblem_detailed() {
	notFound := problems.NewStatusProblem(404)
	notFound.Detail = "The item you've requested either does not exist or has been deleted."
	b, _ := json.MarshalIndent(notFound, "", "  ")
	fmt.Println(string(b))
	// Output: {
	//   "type": "about:blank",
	//   "title": "Not Found",
	//   "status": 404,
	//   "detail": "The item you've requested either does not exist or has been deleted."
	// }
}

func ExampleFromError() {
	err := func() error {
		// Some fallible function.
		return errors.New("something bad happened")
	}()
	internalServerError := problems.FromError(err).WithStatus(http.StatusInternalServerError)
	b, _ := json.MarshalIndent(internalServerError, "", "  ")
	fmt.Println(string(b))
	// Output: {
	//   "type": "about:blank",
	//   "title": "Internal Server Error",
	//   "status": 500,
	//   "detail": "something bad happened"
	// }
}

func ExampleExtendedProblem() {
	type CreditProblemExt struct {
		Balance  float64  `json:"balance"`
		Accounts []string `json:"accounts"`
	}
	problem := problems.NewExt[CreditProblemExt]().
		WithStatus(http.StatusForbidden).
		WithDetail("You do not have sufficient funds to complete this transaction.").
		WithExtension(CreditProblemExt{
			Balance:  30,
			Accounts: []string{"/account/12345", "/account/67890"},
		})
	b, _ := json.MarshalIndent(problem, "", "  ")
	fmt.Println(string(b))
	// Output: {
	//   "type": "about:blank",
	//   "title": "Forbidden",
	//   "status": 403,
	//   "detail": "You do not have sufficient funds to complete this transaction.",
	//   "extensions": {
	//     "balance": 30,
	//     "accounts": [
	//       "/account/12345",
	//       "/account/67890"
	//     ]
	//   }
	// }
}

func ExampleExtendedProblem_embedding() {
	type CreditProblem struct {
		problems.Problem

		Balance  float64  `json:"balance"`
		Accounts []string `json:"accounts"`
	}
	problem := &CreditProblem{
		Problem: *problems.New().
			WithStatus(http.StatusForbidden).
			WithDetail("You do not have sufficient funds to complete this transaction."),
		Balance:  30,
		Accounts: []string{"/account/12345", "/account/67890"},
	}

	b, _ := json.MarshalIndent(problem, "", "  ")
	fmt.Println(string(b))
	// Output: {
	//   "type": "about:blank",
	//   "title": "Forbidden",
	//   "status": 403,
	//   "detail": "You do not have sufficient funds to complete this transaction.",
	//   "balance": 30,
	//   "accounts": [
	//     "/account/12345",
	//     "/account/67890"
	//   ]
	// }
}
