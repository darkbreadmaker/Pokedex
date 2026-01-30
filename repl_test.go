package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "   Hello  World   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "   ",
			expected: nil,
		},
		{
			input: "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input: "hellO World",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected
		if len(actual) != len(expected) {
			t.Errorf("Incorrect length. -- Expected: %d -- Got: %d", len(expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Incorrect Word. -- Expected: %s -- Got: %s", expectedWord, word)
			}
		}
	}
}
