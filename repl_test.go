package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  --- this is a --- normal   phrase",
			expected: []string{"---", "this", "is", "a", "---", "normal", "phrase"},
		},
		{
			input:    "CAPITAL case WORDS",
			expected: []string{"capital", "case", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual length not matching expected length: %d - %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("words does not match: %s - %s", actual, expected)
			}
		}
	}
}
