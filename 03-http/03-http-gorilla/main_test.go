package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEcho(t *testing.T) {
	// set up httptest test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router := getRouter()
		router.ServeHTTP(w, r)
	}))
	defer ts.Close()

	// call the test server endpoint /api/v1/echo/hello!
	resp, err := http.Get(ts.URL + "/api/v1/echo/hello!")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// read the response
	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// assert
	expected := []byte("hello!")
	if bytes.Compare(actual, expected) != 0 {
		t.Fatalf("want: %s got: %s", expected, actual)
	}
}
