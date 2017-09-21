package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestServer() *httptest.Server {
	// set up httptest test server with our router
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router := getRouter()
		router.ServeHTTP(w, r)
	}))
}

func get(ts *httptest.Server, path string) (string, int, error) {
	// call the test server endpoint
	resp, err := http.Get(ts.URL + path)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	// read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}

	return string(body), resp.StatusCode, nil
}

func TestEcho(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	// set up test table (path, expected body, expected HTTP status)
	cases := []struct {
		path           string
		expectedBody   string
		expectedStatus int
	}{
		{"/api/v1/echo/hello!", "{\"message\":\"hello!\"}\n", 200},
		{"/api/v1/echo/hello", "I don't want to say that.", 400},
	}

	for _, c := range cases {
		body, status, err := get(ts, c.path)
		if err != nil {
			t.Error(err)
		}

		// assert
		if body != c.expectedBody {
			t.Errorf("body, want: %q got: %q", c.expectedBody, body)
		}
		if status != c.expectedStatus {
			t.Errorf("status, want: %d got: %d", c.expectedStatus, status)
		}
	}
}
