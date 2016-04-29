// Package problems provides an RFC 7807 (https://tools.ietf.org/html/rfc7807)
// compliant implementation of HTTP problem details. Which are defined as a
// means to carry machine-readable details of errors in an HTTP response to
// avoid the need to define new error response formats for HTTP APIs.
//
// The problem details specification was designed to allow for schema
// extensions. Because of this the exposed Problem interface only enforces the
// required Type and Title fields be set appropriately.
//
// Additionally, this library also ships with default http.HandlerFunc's capable
// of writing problems to http.ResponseWriter's in either of the two standard
// media formats JSON and XML.
package problems
