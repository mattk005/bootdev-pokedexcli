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

		{
			input:    "  BulbaSaur  PIKACHU  charmander  ",
			expected: []string{"bulbasaur", "pikachu", "charmander"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		lenAcutal := len(actual)
		lenExpected := len(c.expected)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of array does not match expected (%d : %d)", lenAcutal, lenExpected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s and %s do not match", word, expectedWord)
			}
		}
	}
}
