package problems

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testServer(funcs ...http.HandlerFunc) *httptest.Server {
	mux := http.NewServeMux()

	for _, f := range funcs {
		mux.HandleFunc("/", f)
	}
	ts := httptest.NewServer(mux)
	return ts
}

func getResponse(uri string, server *httptest.Server) (*http.Response, error) {
	res, err := http.Get(server.URL + uri)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func TestJSONProblems(t *testing.T) {
	notFound := NewDetailedProblem(http.StatusNotFound, "That thing doesn't exist.")

	server := testServer(ProblemHandler(notFound))
	defer server.Close()

	w, err := getResponse("/", server)
	if err != nil {
		t.Error(err)
	}

	if w.StatusCode != notFound.Status {
		t.Errorf("Expected HTTP status code to be %d, got %d", notFound.Status, w.StatusCode)
	}

	var response Problem
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Error(err)
	}

	if response.Status != notFound.Status {
		t.Errorf("Expected response Status to be %d, but got %d", notFound.Status, response.Status)
	}

	if response.Title != notFound.Title {
		t.Errorf("Expected response Title to be %q, but got %q", notFound.Title, response.Title)
	}

	if response.Detail != notFound.Detail {
		t.Errorf("Expected response Detail to be %q, but got %q", notFound.Detail, response.Detail)
	}
}

func TestXMLProblems(t *testing.T) {
	notFound := NewDetailedProblem(http.StatusNotFound, "That thing doesn't exist.")

	server := testServer(XMLProblemHandler(notFound))
	defer server.Close()

	w, err := getResponse("/", server)
	if err != nil {
		t.Error(err)
	}

	if w.StatusCode != notFound.Status {
		t.Errorf("Expected HTTP status code to be %d, got %d", notFound.Status, w.StatusCode)
	}

	var response Problem
	err = xml.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Error(err)
	}

	if response.Status != notFound.Status {
		t.Errorf("Expected response Status to be %d, but got %d", notFound.Status, response.Status)
	}

	if response.Title != notFound.Title {
		t.Errorf("Expected response Title to be %q, but got %q", notFound.Title, response.Title)
	}

	if response.Detail != notFound.Detail {
		t.Errorf("Expected response Detail to be %q, but got %q", notFound.Detail, response.Detail)
	}
}
