package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input	string
		expected	[]string
	}{
		{
			input:	"hello world",
			expected: []string{"hello", "world"}
		},
		{
			input: "testnowhitespace"
			expected: []string{"testnowhitespace"}
		},
		{
			input: "Hello World"
			expected: []string{"hello", "world"}
		},
		{
			input: "testing leading whitespace    "
			expected: []string{"testing", "leading", "whitespace"}
		}
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual and expected do not match, len(actual)=%v, len(expected)=%v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word and expectedWord do not match: word=%s expected=%s", word, expectedWord)
			}
		}
	}
}