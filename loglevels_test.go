package main

import (
	"reflect"
	"testing"
)

func TestCountLogLevels(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{
			name: "multiple log levels",
			input: []string{
				"INFO user logged in",
				"WARN disk 95% full",
				"INFO request completed",
				"ERROR application unavailable",
			},
			expected: map[string]int{
				"INFO":  2,
				"WARN":  1,
				"ERROR": 1,
			},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: map[string]int{},
		},
		{
			name: "single log line",
			input: []string{
				"INFO user logged in",
			},
			expected: map[string]int{
				"INFO": 1,
			},
		},
		{
			name: "duplicate log levels",
			input: []string{
				"INFO user logged in",
				"INFO request completed",
				"INFO connection closed",
			},
			expected: map[string]int{
				"INFO": 3,
			},
		},
		{
			name: "line containing only a level",
			input: []string{
				"INFO",
			},
			expected: map[string]int{
				"INFO": 1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountLogLevels(test.input)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("got %v, want %v", result, test.expected)
			}
		})
	}
}
