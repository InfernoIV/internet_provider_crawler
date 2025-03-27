package main

import (
	"reflect"
	"testing"
)

func Test_get_urls_from_html(t *testing.T) {
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
			name:     "absolute and relative URLs 2",
			inputURL: "https://github.com",
			inputBody: `
<html>
	<body>
		<a href="/path/two">
			<span>Boot.dev</span>
		</a>
		<a href="https://otherone.com/path/three">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://github.com/path/two", "https://otherone.com/path/three"},
		},
		{
			name:     "absolute and relative URLs 3",
			inputURL: "https://google.com",
			inputBody: `
<html>
	<body>
		<a href="/search?q=boot.dev">
			<span>Boot.dev</span>
		</a>
		<a href="https://reddit.com/r/learnprogramming/comments/1i51n07/do_bootdev_is_worth_paying_the_subscription/">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://google.com/search", "https://reddit.com/r/learnprogramming/comments/1i51n07/do_bootdev_is_worth_paying_the_subscription/"},
		},

		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := get_urls_from_html(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
