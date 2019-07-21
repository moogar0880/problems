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
		_ = json.NewEncoder(w).Encode(p)
	}
}

// XMLProblemHandler returns an http.HandlerFunc which writes a provided problem
// to an http.ResponseWriter as XML
func XMLProblemHandler(p Problem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ProblemMediaTypeXML)
		_ = xml.NewEncoder(w).Encode(p)
	}
}

// StatusProblemHandler returns an http.HandlerFunc which writes a provided
// problem to an http.ResponseWriter as JSON with the status code
func StatusProblemHandler(p StatusProblem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ProblemMediaType)
		if p.ProblemStatus() != 0 {
			w.WriteHeader(p.ProblemStatus())
		}
		_ = json.NewEncoder(w).Encode(p)
	}
}

// XMLStatusProblemHandler returns an http.HandlerFunc which writes a provided
// problem to an http.ResponseWriter as XML with the status code
func XMLStatusProblemHandler(p StatusProblem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ProblemMediaTypeXML)
		if p.ProblemStatus() != 0 {
			w.WriteHeader(p.ProblemStatus())
		}
		_ = xml.NewEncoder(w).Encode(p)
	}
}
