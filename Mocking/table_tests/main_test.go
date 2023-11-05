package main

import "testing"

func TestTableTestq(t *testing.T) {
	testTable := []struct {
		name string
		// input, output
	}{
		{name: "test1"},
		// inoput, output
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			// call function or method under test
			// ....
			// check response
		})
	}
}
