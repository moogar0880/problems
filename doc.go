// Package problems provides an RFC-9457 (https://tools.ietf.org/html/rfc9457)
// and RFC 7807 (https://tools.ietf.org/html/rfc7807) compliant implementation
// of HTTP problem details. Which are defined as a means to carry
// machine-readable details of errors in an HTTP response to avoid the need to
// define new error response formats for HTTP APIs.
//
// The problem details specification was designed to allow for schema
// extensions. There are two possible ways to create problem extensions:
//
// 1. You can embed a problem in your extension problem type.
// 2. You can use the ExtendedProblem to leverage the existing types in this
// library.
//
// See the examples for references on how to use either of these extension
// mechanisms.
//
// Additionally, this library also ships with default http.HandlerFunc
// implementations which are capable of writing problems to a
// http.ResponseWriter in either of the two standard media formats, JSON and
// XML.
package problems
