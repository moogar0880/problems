package problems_test

import (
	"encoding/json"
	"os"

	"github.com/moogar0880/problems"
)

func ExampleNewStatusProblem() {
	notFound := problems.NewStatusProblem(404)
	b, _ := json.MarshalIndent(notFound, "", "  ")
	os.Stdout.Write(b)
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
	os.Stdout.Write(b)
	// Output: {
	//   "type": "about:blank",
	//   "title": "Not Found",
	//   "status": 404,
	//   "detail": "The item you've requested either does not exist or has been deleted."
	// }
}
