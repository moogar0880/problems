package problems

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// ProblemHandler returns a http.HandlerFunc which writes a provided problem
// to a http.ResponseWriter as JSON with the status code.
func ProblemHandler(p *Problem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ProblemMediaType)
		if p.Status != 0 {
			w.WriteHeader(p.Status)
		}
		_ = json.NewEncoder(w).Encode(p)
	}
}

// XMLProblemHandler returns a http.HandlerFunc which writes a provided problem
// to a http.ResponseWriter as XML with the status code.
func XMLProblemHandler(p *Problem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ProblemMediaTypeXML)
		if p.Status != 0 {
			w.WriteHeader(p.Status)
		}
		_ = xml.NewEncoder(w).Encode(struct {
			XMLName xml.Name `xml:"urn:ietf:rfc:7807 problem"`
			Problem
		}{
			Problem: *p,
		})
	}
}
