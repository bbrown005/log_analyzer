package main

import (
	"reflect"
	"testing"
)

func TestReadLogFile(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		expected      []string
		expectedError bool
	}{
		{
			name: "reads multiple lines",
			path: "testdata/simple.log",
			expected: []string{
				"INFO user logged in",
				"ERROR database unavailable",
				"WARN cache miss",
			},
			expectedError: false,
		},
		{
			name: "reads one line",
			path: "testdata/one_line.log",
			expected: []string{
				"INFO server started",
			},
			expectedError: false,
		},
		{
			name:          "reads empty file",
			path:          "testdata/empty.log",
			expected:      []string{},
			expectedError: false,
		},
		{
			name:          "missing file",
			path:          "testdata/missing.log",
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ReadLogFile(test.path)

			if test.expectedError && err == nil {
				t.Errorf("ReadLogFile(%q) expected error but got none", test.path)
			}

			if !test.expectedError && err != nil {
				t.Errorf("ReadLogFile(%q) expected no error but got %v", test.path, err)
			}

			if !test.expectedError && !reflect.DeepEqual(result, test.expected) {
				t.Errorf("got %v, want %v", result, test.expected)
			}
		})
	}
}
