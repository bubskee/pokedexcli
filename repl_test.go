// repl_test.go
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
		//name: "basic trim",
		input:    "  hello   world  ",
		expected: []string{"hello", "world"},
	}, {
		//name: "trim with uppercase",
		input:    "  HeLlO WOrlD  ",
		expected: []string{"hello", "world"},
	},{
		input:    "  extra WORDS hello WORLD      ",
		expected: []string{"extra", "words","hello", "world"},
	}}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// check count of words, which is length
		actLength := len(actual)
		expLength := len(c.expected)
		if actLength != expLength {
			t.Errorf("wrong number of words, got %v; expected %v", actLength, expLength)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("wrong formatting, got %v; expected %v", word, expectedWord)
			}
		}
	}
}