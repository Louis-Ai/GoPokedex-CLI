package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    " help ",
			expected: []string{"help"},
		},
		{
			input:    " help text",
			expected: []string{"help", "text"},
		},
		{
			input:    " hEllO World",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanUserInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lenghts do not match, want: '%v', got: '%v'", c.expected, actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match, want: '%v', got: '%v'", expectedWord, word)
			}
		}
	}
}
