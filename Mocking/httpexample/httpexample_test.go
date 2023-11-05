package httpexample_test

import (
	"errors"
	"livrogolang/Mocking/httpexample"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMakeHTTPCall(t *testing.T) {
	testTable := []struct {
		name             string
		server           *httptest.Server
		expectedResponse *httpexample.Response
		expectedError    error
	}{
		{
			name: "happy-path",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"id": 55, "name": "foo", "description": "bar"}`))
			})),
			expectedResponse: &httpexample.Response{
				ID:          55,
				Name:        "foo",
				Description: "bar",
			},
			expectedError: nil,
		},
		{
			name: "error-bad-request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
			})),
			expectedResponse: nil,
			expectedError:    httpexample.ErrBadStatusCode,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()
			response, err := httpexample.MakeHttpCall(tc.server.URL)
			if !errors.Is(err, tc.expectedError) {
				t.Fatalf("expected error %v, got %v", tc.expectedError, err)
			}
			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Fatalf("expected response %v, got %v", tc.expectedResponse, response)
			}
		})
	}
}
