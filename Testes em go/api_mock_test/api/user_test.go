package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type MockHTTP struct {
	response *http.Response
	err      error
}

func (mock *MockHTTP) Do(req *http.Request) (*http.Response, error) {
	return mock.response, mock.err
}

func TestGetRandomUser(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	//NewRecorder implements the interface that satisfy the http.ResponseWriter
	w := httptest.NewRecorder()

	res := http.Response{
		Body: io.NopCloser(strings.NewReader("")),
	}

	mock := MockHTTP{
		response: &res,
	}

	service := Service{
		HTTPClient: &mock,
	}

	service.GetRandomUser(w, r)

	got := w.Result()

	if !reflect.DeepEqual(http.StatusOK, got.StatusCode) {
		t.Errorf("wanted: %d, got: %d", http.StatusOK, got.StatusCode)
	}
}
