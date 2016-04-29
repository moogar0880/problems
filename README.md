# problems
Problems is an RFC-7807 compliant library for describing HTTP errors, written
purely in Go. For more information see [RFC-7807](https://tools.ietf.org/html/rfc7807).

## Usage
The problems library exposes an assortment of interfaces, structs, and functions
for defining and using HTTP Problem detail resources.

### Predefined Errors
You can define basic Problem details up front by using the `NewStatusProblem`
function

```go
import (
  "github.com/moogar0880/problems"
)

var (
  // The NotFound problem will be built with an appropriate status code and
  // informative title set. Additional information can be provided in the Detail
  // field of the generated struct
  NotFound = problems.NewStatusProblem(404)
)
```

### Detailed Errors
New errors can also be created a head of time, or on the fly like so:

```go
import (
  "github.com/moogar0880/problems"
)

func NoSuchUser() *problems.DefaultProblem {
  nosuch := problems.NewStatusProblem(404)
  nosuch.Detail = "Sorry, that user does not exist"
  return nosuch
}
```

### Expanded Errors
The specification for these HTTP problems was designed to allow for arbitrary
expansion of the problem resources. This can be accomplished through this
library by implementing the `Problem` interface:

```go
import (
  "github.com/moogar0880/problems"
)

type CreditProblem struct {
  problems.DefaultProblem

  Balance float64
  Accounts []string
}

func (cp *CreditProblem) ProblemType() (*url.URL, error) {
	u, err := url.Parse(cp.Type)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (cp *CreditProblem) ProblemTitle() string {
	return cp.Title
}
```

## Serving Problems
Additionally, RFC-7807 defines two new media types for problem resources,
`application/problem+json"` and `application/problem+xml`. This library defines
those media types as the constants `ProblemMediaType` and
`ProblemMediaTypeXML`.

In order to facilitate serving problem definitions, this library exposes two
`http.HandlerFunc` implementations which accept a problem, and return a
functioning `HandlerFunc` that will server that error.

```go
package main

import (
  "net/http"

  "github.com/moogar0880/problems"
)

var (
	Unauthorized = problems.NewStatusError(401)
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/secrets", problems.ProblemHandler(Unauthorized))

	server := http.Server{Handler: mux, Addr: ":80"}
	server.ListenAndServe()
}
```
