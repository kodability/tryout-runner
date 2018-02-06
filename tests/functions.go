package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// makeGet makes a GET request
func makeGet(url string) (req *http.Request, write *httptest.ResponseRecorder, err error) {
	r, e := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	return r, w, e
}

// makePostJSON makes a POST request with JSON content-type
func makePostJSON(url string, body interface{}) (req *http.Request, write *httptest.ResponseRecorder, err error) {
	jsonValue, e := json.Marshal(body)
	if e != nil {
		return nil, nil, e
	}

	r, e := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	return r, w, e
}
