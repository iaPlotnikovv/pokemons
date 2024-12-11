package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "POKE MON",
			expected: []string{"poke", "mon"},
		},
		{
			input:    "Hi my name is Ilia",
			expected: []string{"hi", "my", "name", "is", "ilia"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("UNMATCH DETECTED!  %v test length VS %v", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("UNMATCH DETECTED! test %v VS %v", word, expectedWord)
				continue
			}

		}

	}

}
