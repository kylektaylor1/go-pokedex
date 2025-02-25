package main

import (
	"fmt"
	"testing"
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
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "my name is kyle",
			expected: []string{"my", "name", "is", "kyle"},
		},
		{
			input:    "New world OP",
			expected: []string{"new", "world", "op"},
		},
		{
			input:    "coffEe iS dope",
			expected: []string{"coffee", "is", "dope"},
		},
	}

	for _, c := range cases {
		fmt.Printf("c: %v\n", c)
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words, got %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word: %s does not match expected: %s", word, expectedWord)
			}
		}
	}
}
