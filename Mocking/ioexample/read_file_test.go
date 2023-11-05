package ioexample_test

import (
	"errors"
	"io"
	"livrogolang/Mocking/ioexample"
	"reflect"
	"testing"
)

type mockReadCloser struct {
	expected_data []byte
	expecterr     error
}

// allow test table to fill out expectedData and expectedErr filed
func (mrc *mockReadCloser) Read(p []byte) (n int, err error) {
	copy(p, mrc.expected_data)
	return 0, mrc.expecterr
}

// do nothing on close
func (mrc *mockReadCloser) Close() error { return nil }

// TestReadFileContent tests the ReadFileContent function
func TestReadFileContent(t *testing.T) {
	errRead := errors.New("uh oh")
	testCases := []struct {
		name         string
		readCloser   io.ReadCloser
		numBytes     int
		expectedData []byte
		expectedErr  error
	}{
		{
			name: "happy-path",
			readCloser: &mockReadCloser{
				expected_data: []byte(`hello`),
				expecterr:     nil,
			},
			numBytes:     5,
			expectedData: []byte("hello"),
			expectedErr:  nil,
		},
		{
			name: "error-from-reader",
			readCloser: &mockReadCloser{
				expected_data: nil,
				expecterr:     errRead,
			},
			numBytes:     0,
			expectedData: nil,
			expectedErr:  errRead,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// call ReadFileContent with mockReadCloser
			data, err := ioexample.ReadFileContent(tc.readCloser, tc.numBytes)
			if !reflect.DeepEqual(data, tc.expectedData) {
				t.Errorf("expected %v, got %v", tc.expectedData, data)
			}
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
