package main

import (
	"slices"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/one">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a>
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: nil,
		},
		{
			name:     "invalid href URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a href=":\\invalidURL">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: nil,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i+1, tc.name, err)
				return
			}
			if !slices.Equal(actual, tc.expected) {
				t.Errorf(`
				Test %v - %s FAIL:
					expected URL: %v,
					actual:       %v`, i+1, tc.name, tc.expected, actual)
			}
		})
	}
}
