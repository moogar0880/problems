# Problems

Problems is an RFC-7807 and RFC-9457 compliant library for describing HTTP errors.
For more information see [RFC-9457](https://tools.ietf.org/html/rfc9457), and it's predecessor [RFC-7807](https://tools.ietf.org/html/rfc7807).

[![Build Status](https://travis-ci.org/moogar0880/problems.svg?branch=master)](https://travis-ci.org/moogar0880/problems)
[![Go Report Card](https://goreportcard.com/badge/github.com/moogar0880/problems)](https://goreportcard.com/report/github.com/moogar0880/problems)
[![GoDoc](https://godoc.org/github.com/moogar0880/problems?status.svg)](https://godoc.org/github.com/moogar0880/problems)

## Usage

The problems library exposes an assortment of types to aid HTTP service authors
in defining and using HTTP Problem detail resources.

### Predefined Errors

You can define basic Problem details up front by using the `NewStatusProblem`
function

```go
package main

import "github.com/moogar0880/problems"

var (
  // The NotFound problem will be built with an appropriate status code and
  // informative title set. Additional information can be provided in the Detail
  // field of the generated struct
  NotFound = problems.NewStatusProblem(404)
)
```

Which, when served over HTTP as JSON will look like the following:

```json
{
   "type": "about:blank",
   "title": "Not Found",
   "status": 404
}
```

### Detailed Errors

New errors can also be created a head of time, or on the fly like so:

```go
package main

import "github.com/moogar0880/problems"

func NoSuchUser() *problems.Problem {
    nosuch := problems.NewStatusProblem(404)
    nosuch.Detail = "Sorry, that user does not exist."
    return nosuch
}
```

Which, when served over HTTP as JSON will look like the following:

```json
{
   "type": "about:blank",
   "title": "Not Found",
   "status": 404,
   "detail": "Sorry, that user does not exist."
}
```

### Extended Errors

The specification for these HTTP problems was designed to allow for arbitrary
expansion of the problem resources. This can be accomplished through this
library by either embedding the `Problem` struct in your extension type:

```go
package main

import "github.com/moogar0880/problems"

type CreditProblem struct {
    problems.Problem

    Balance  float64  `json:"balance"`
    Accounts []string `json:"accounts"`
}
```

Which, when served over HTTP as JSON will look like the following:

```json
{
   "type": "about:blank",
   "title": "Unauthorized",
   "status": 401,
   "balance": 30,
   "accounts": ["/account/12345", "/account/67890"]
}
```

Or by using the `problems.ExtendedProblem` type:

```go
package main

import (
    "net/http"

    "github.com/moogar0880/problems"
)

type CreditProblemExt struct {
    Balance  float64  `json:"balance"`
    Accounts []string `json:"accounts"`
}

func main() {
    problems.NewExt[CreditProblemExt]().
        WithStatus(http.StatusForbidden).
        WithDetail("Your account does not have sufficient funds to complete this transaction").
        WithExtension(CreditProblemExt{
            Balance:  30,
            Accounts: []string{"/account/12345", "/account/67890"},
        })
}
```

Which, when served over HTTP as JSON will look like the following:

```json
{
   "type": "about:blank",
   "title": "Unauthorized",
   "status": 401, 
   "extensions": {
       "balance": 30,
       "accounts": ["/account/12345", "/account/67890"]    
   }
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

var Unauthorized = problems.NewStatusProblem(401)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/secrets", problems.ProblemHandler(Unauthorized))

    server := http.Server{Handler: mux, Addr: ":8080"}
    server.ListenAndServe()
}
```
