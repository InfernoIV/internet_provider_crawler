package main

import "testing"

func Test_normalize_url(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "https, remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "https://blog.boot.dev/path",
		},
		{
			name:     "https, extra path",
			inputURL: "https://blog.boot.dev/path/",
			expected: "https://blog.boot.dev/path/",
		},
		{
			name:     "http, extra path",
			inputURL: "http://blog.boot.dev/path/",
			expected: "http://blog.boot.dev/path/",
		},
		{
			name:     "http, remove scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "http://blog.boot.dev/path",
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalize_url(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
