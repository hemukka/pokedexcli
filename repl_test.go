package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
		{
			input:    "  Charmande Bulbasaur  PIKAchu",
			expected: []string{"charmande", "bulbasaur", "pikachu"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf(
				"lengths don't match:\nactual:   %v \nexpected: %v",
				actual,
				c.expected,
			)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf(
					`%v != %v
input:    %v
actual:   %v
expected: %v`,
					word,
					expectedWord,
					c.input,
					actual,
					c.expected,
				)
			}

		}
	}
}
