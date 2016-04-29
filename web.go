package problems

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// ProblemHandler returns an http.HandlerFunc which writes a provided problem
// to an http.ResponseWriter as JSON
func ProblemHandler(p Problem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ProblemMediaType)
		json.NewEncoder(w).Encode(p)
	}
}

// XMLProblemHandler returns an http.HandlerFunc which writes a provided problem
// to an http.ResponseWriter as XML
func XMLProblemHandler(p Problem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ProblemMediaTypeXML)
		xml.NewEncoder(w).Encode(p)
	}
}
