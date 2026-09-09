package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world  AGain",
			expected: []string{"hello", "world", "again"},
		},
		{
			input:    "THiS WoRks JuSt FinE !!!",
			expected: []string{"this", "works", "just", "fine", "!!!"},
		},
	}
	for count, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("Error: Not the expected lenght in case #%d", count+1)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s is not %s !!! Test #%d failed", word, expectedWord, count+1)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}

}

//func Test
