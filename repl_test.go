package main

// small unit test for the repl (input-loop)

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "HEllo   wOrld",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello  ",
			expected: []string{"hello"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Different lengths: '%v' vs '%v", actual, c.expected)
			// fmt.Println("Error: different length of strings")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			diff := cmp.Diff(expectedWord, word)
			if diff != "" {
				t.Fatalf("%s", diff)
			}

		}
	}
}
